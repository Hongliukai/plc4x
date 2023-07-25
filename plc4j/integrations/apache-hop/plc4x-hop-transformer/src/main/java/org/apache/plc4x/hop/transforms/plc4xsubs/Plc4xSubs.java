/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.apache.plc4x.hop.transforms.plc4xsubs;

import java.time.Duration;
import java.util.ArrayList;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentLinkedQueue;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.ReentrantLock;
import java.util.function.Consumer;
import java.util.logging.Level;
import java.util.logging.Logger;
import org.apache.commons.codec.binary.BinaryCodec;
import org.apache.commons.lang3.StringUtils;
import org.apache.commons.lang3.tuple.ImmutablePair;
import org.apache.commons.lang3.tuple.Pair;
import org.apache.hop.core.CheckResult;
import org.apache.hop.core.Const;
import org.apache.hop.core.ICheckResult;
import org.apache.hop.core.RowMetaAndData;
import org.apache.hop.core.exception.HopException;
import org.apache.hop.core.exception.HopPluginException;
import org.apache.hop.core.exception.HopValueException;
import org.apache.hop.core.logging.LogLevel;
import org.apache.hop.core.row.IRowMeta;
import org.apache.hop.core.row.IValueMeta;
import org.apache.hop.core.row.RowDataUtil;
import org.apache.hop.core.row.RowMeta;
import org.apache.hop.core.row.value.ValueMetaDate;
import org.apache.hop.core.row.value.ValueMetaFactory;
import org.apache.hop.core.util.StringUtil;
import org.apache.hop.core.util.Utils;
import org.apache.hop.i18n.BaseMessages;
import org.apache.hop.metadata.api.IHopMetadataProvider;
import org.apache.hop.pipeline.Pipeline;
import org.apache.hop.pipeline.PipelineMeta;
import org.apache.hop.pipeline.transform.BaseTransform;
import org.apache.hop.pipeline.transform.ITransform;
import org.apache.hop.pipeline.transform.TransformMeta;
import org.apache.plc4x.hop.metadata.Plc4xConnection;
import org.apache.plc4x.hop.metadata.util.Plc4xLookup;
import org.apache.plc4x.hop.transforms.util.Plc4xGeneratorField;
import org.apache.plc4x.hop.transforms.util.Plc4xPlcTag;
import org.apache.plc4x.hop.transforms.util.Plc4xPlcSubscriptionTag;
import org.apache.plc4x.hop.metadata.util.Plc4xWrapperConnection;
import org.apache.plc4x.hop.transforms.plc4xinput.Plc4xRead;
import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.api.listener.ConnectionStateListener;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcSubscriptionEvent;
import org.apache.plc4x.java.api.messages.PlcSubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionResponse;
import org.apache.plc4x.java.api.model.PlcConsumerRegistration;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcValueType;
import static org.apache.plc4x.java.api.types.PlcValueType.BYTE;
import static org.apache.plc4x.java.api.types.PlcValueType.DATE;
import static org.apache.plc4x.java.api.types.PlcValueType.DATE_AND_TIME;
import static org.apache.plc4x.java.api.types.PlcValueType.DWORD;
import static org.apache.plc4x.java.api.types.PlcValueType.LDATE;
import static org.apache.plc4x.java.api.types.PlcValueType.LDATE_AND_TIME;
import static org.apache.plc4x.java.api.types.PlcValueType.LTIME;
import static org.apache.plc4x.java.api.types.PlcValueType.LWORD;
import static org.apache.plc4x.java.api.types.PlcValueType.STRING;
import static org.apache.plc4x.java.api.types.PlcValueType.TIME;
import static org.apache.plc4x.java.api.types.PlcValueType.WORD;
import org.apache.plc4x.java.s7.events.S7CyclicEvent;
import org.apache.plc4x.java.s7.events.S7Event;
import org.apache.plc4x.java.s7.readwrite.protocol.S7HPlcConnection;
import org.apache.plc4x.java.s7.readwrite.tag.S7SubscriptionTag;
import org.apache.plc4x.java.spi.utils.hex.Hex;
import org.openide.util.Lookup;

/**
 *
 *
 */
public class Plc4xSubs extends BaseTransform<Plc4xSubsMeta, Plc4xSubsData> implements ConnectionStateListener {

    private static final Class<?> PKG = Plc4xSubs.class; // Needed by Translator

    private Plc4xConnection connmeta = null;
    private Plc4xWrapperConnection connwrapper = null;
    private PlcSubscriptionRequest subsRequest = null;
    private PlcSubscriptionResponse subsResponse = null;

    private PlcConsumerRegistration csmr;

    private int maxwait = 0;
    private static final ReentrantLock lock = new ReentrantLock();

    private Plc4xLookup lookup = Plc4xLookup.getDefault();
    private Lookup.Template template = null;
    private Lookup.Result<Plc4xWrapperConnection> result = null;

    private ConcurrentLinkedQueue<PlcSubscriptionEvent> events = new ConcurrentLinkedQueue();
    private boolean stopBundle = false;
    private boolean isConnected = false;

    private String jobId = null;

    private BinaryCodec binarycodec = new BinaryCodec();

    private static final String DUMMY = "dummy";
    private static final String DATA = "DATA_";

    private int counter = 0;
    private String strData;
    private PlcValueType valuetype = null;

    private Map<String, String> tags_index = new HashMap();
    private Map<String, Pair<String, S7SubscriptionTag>> plctags = new HashMap();

    public Plc4xSubs(TransformMeta transformMeta, Plc4xSubsMeta meta, Plc4xSubsData data, int copyNr, PipelineMeta pipelineMeta,
            Pipeline pipeline) {
        super(transformMeta, meta, data, copyNr, pipelineMeta, pipeline);
    }

    /*
  * Including Date and Time field for every row 
  *
  * @param meta Meta data from user dialog
  * @param remarks Error registers
  * @param origin transform instance name
     */
    public static final RowMetaAndData buildRow(
            Plc4xSubsMeta meta, List<ICheckResult> remarks, String origin) throws HopPluginException {
        IRowMeta rowMeta = new RowMeta();
        Object[] rowData = RowDataUtil.allocateRowData(meta.getFields().size() + 2);
        int index = 0;

        if (!Utils.isEmpty(meta.getRowTimeField())) {
            rowMeta.addValueMeta(new ValueMetaDate(meta.getRowTimeField()));
            rowData[index++] = null;
        }

        if (!Utils.isEmpty(meta.getLastTimeField())) {
            rowMeta.addValueMeta(new ValueMetaDate(meta.getLastTimeField()));
            rowData[index++] = null;
        }

        for (Plc4xGeneratorField field : meta.getFields()) {
            int typeString = ValueMetaFactory.getIdForValueMeta(field.getType());
            if (StringUtils.isNotEmpty(field.getType())) {

                IValueMeta valueMeta
                        = ValueMetaFactory.createValueMeta(field.getName(), typeString); // build a
                // value!
                valueMeta.setLength(field.getLength());
                valueMeta.setPrecision(field.getPrecision());
                valueMeta.setConversionMask(field.getFormat());
                valueMeta.setCurrencySymbol(field.getCurrency());
                valueMeta.setGroupingSymbol(field.getGroup());
                valueMeta.setDecimalSymbol(field.getDecimal());
                valueMeta.setOrigin(origin);

                IValueMeta stringMeta = ValueMetaFactory.cloneValueMeta(valueMeta, IValueMeta.TYPE_STRING);

                if (field.isSetEmptyString()) {
                    // Set empty string
                    rowData[index] = StringUtil.EMPTY_STRING;
                } else {
                    String stringValue = field.getValue();

                    // If the value is empty: consider it to be NULL.
                    if (Utils.isEmpty(stringValue)) {
                        rowData[index] = null;

                        if (valueMeta.getType() == IValueMeta.TYPE_NONE) {
                            String message
                                    = BaseMessages.getString(
                                            PKG,
                                            "Plc4x.Read.Meta.CheckResult.SpecifyTypeError",
                                            valueMeta.getName(),
                                            stringValue);
                            remarks.add(new CheckResult(ICheckResult.TYPE_RESULT_ERROR, message, null));
                        }
                    } else {
                        // Convert the data from String to the specified type ...
                        //
                        try {
                            System.out.println("stringValue: " + stringValue);
                            rowData[index] = valueMeta.convertData(stringMeta, stringValue);
                        } catch (HopValueException e) {
                            switch (valueMeta.getType()) {
                                case IValueMeta.TYPE_NUMBER:
                                    String message
                                            = BaseMessages.getString(
                                                    PKG,
                                                    "Plc4x.Read.Meta.BuildRow.Error.Parsing.Number",
                                                    valueMeta.getName(),
                                                    stringValue,
                                                    e.toString());
                                    remarks.add(new CheckResult(ICheckResult.TYPE_RESULT_ERROR, message, null));
                                    break;

                                case IValueMeta.TYPE_DATE:
                                    message
                                            = BaseMessages.getString(
                                                    PKG,
                                                    "Plc4x.Read.Meta.BuildRow.Error.Parsing.Date",
                                                    valueMeta.getName(),
                                                    stringValue,
                                                    e.toString());
                                    remarks.add(new CheckResult(ICheckResult.TYPE_RESULT_ERROR, message, null));
                                    break;

                                case IValueMeta.TYPE_INTEGER:
                                    message
                                            = BaseMessages.getString(
                                                    PKG,
                                                    "Plc4x.Read.Meta.BuildRow.Error.Parsing.Integer",
                                                    valueMeta.getName(),
                                                    stringValue,
                                                    e.toString());
                                    remarks.add(new CheckResult(ICheckResult.TYPE_RESULT_ERROR, message, null));
                                    break;

                                case IValueMeta.TYPE_BIGNUMBER:
                                    message
                                            = BaseMessages.getString(
                                                    PKG,
                                                    "Plc4x.Read.Meta.BuildRow.Error.Parsing.BigNumber",
                                                    valueMeta.getName(),
                                                    stringValue,
                                                    e.toString());
                                    remarks.add(new CheckResult(ICheckResult.TYPE_RESULT_ERROR, message, null));
                                    break;

                                case IValueMeta.TYPE_TIMESTAMP:
                                    message
                                            = BaseMessages.getString(
                                                    PKG,
                                                    "Plc4x.Read.Meta.BuildRow.Error.Parsing.Timestamp",
                                                    valueMeta.getName(),
                                                    stringValue,
                                                    e.toString());
                                    remarks.add(new CheckResult(ICheckResult.TYPE_RESULT_ERROR, message, null));
                                    break;

                                default:
                                    // Boolean and binary don't throw errors normally, so it's probably an unspecified
                                    // error problem...
                                    message
                                            = BaseMessages.getString(
                                                    PKG,
                                                    "Plc4x.Read.Meta.CheckResult.SpecifyTypeError",
                                                    valueMeta.getName(),
                                                    stringValue);
                                    remarks.add(new CheckResult(ICheckResult.TYPE_RESULT_ERROR, message, null));
                                    break;
                            }
                        }
                    }
                }

                // Now add value to the row!
                // This is in fact a copy from the fields row, but now with data.
                rowMeta.addValueMeta(valueMeta);
                index++;
            }
        }

        return new RowMetaAndData(rowMeta, rowData);
    }

    /* 
  * 1. Block the other instances by means of a lock.  
  * 2. Try to locate an existing connection.
  * 3. If it doesn't exist, it tries to take control of the routine to 
  *    create an instance of PlcConnection and his wrapper.
  * 4. Register the connection wrapper for global access.
  * 5. If the connection to the PLC is made, then it creates the query 
  *    and executes it.
  *
     */
    @Override
    public boolean processRow() throws HopException {
        Object[] r = getRow(); // Get row from input rowset & set row busy!
        setLogLevel(LogLevel.DEBUG);

        if ((!meta.isNeverEnding() && data.rowsWritten >= data.rowLimit) && !isStopped()) {
            setOutputDone(); // signal end to receiver(s)
            return false;
        }

        if (first) {
            plctags.clear();
            //This performs a minimal check on the user item.
            //It guarantees that the rates are within those managed by Plc4x.
            meta.getFields().forEach((f) -> {
                plctags.put(f.getName(), new ImmutablePair(f.getName(), S7SubscriptionTag.of(f.getItem())));
            });
            first = false;
        }

        // If we do not have the jobId, it is because we have not completed 
        // the CYC subscription.
        if ((null == jobId) && (null == subsRequest)) {
            RegisterPlcTags();
            GetSubscriptions();
            RegisterCYCHandler();
        }

        while (events.size() == 0) {
            try {
                Thread.sleep(100);
                if (stopBundle) {
                    setOutputDone(); // signal end to receiver(s)
                    return false;
                }
            } catch (InterruptedException ex) {
                break;
            }
        }

        S7CyclicEvent s7event = (S7CyclicEvent) events.poll();
        //System.out.println("Ejecuto el consumidor en el cliente..." + s7event.getTagNames());

        r = data.outputRowMeta.cloneRow(data.outputRowData);

        //s7event.getTimestamp() Takes the time sent by the PLC
        data.prevDate = data.rowDate;
        data.rowDate = Date.from(s7event.getTimestamp());
        int index = 0;
        if (!Utils.isEmpty(meta.getRowTimeField())) {
            r[index++] = data.rowDate;
        }
        if (!Utils.isEmpty(meta.getLastTimeField())) {
            r[index++] = data.prevDate;
        }

        counter = 0;
        for (Plc4xGeneratorField field : meta.getFields()) {
            valuetype = ((S7SubscriptionTag) plctags.get(field.getName()).getRight()).getS7Tags()[0].getPlcValueType();
            if (null != valuetype) {
                System.out.println("Tipo enviado desde el PLC: " + valuetype.name());
            } else {
                System.out.println("El tipo de NULL.");
            }

            if (field.getType().equalsIgnoreCase("Avro Record")) {
                throw new HopException("'Avro Record' type is not supported");
            } else if (field.getType().equalsIgnoreCase("BigNumber")) {
                throw new HopException("'BigNumber' type is not supported");
            } else if (field.getType().equalsIgnoreCase("Binary")) {
                switch (valuetype) {
                    case BYTE:
                        r[index++] = binarycodec.toByteArray(
                                Long.toBinaryString(Long.valueOf(s7event.getByte(field.getName()))));
                        break;
                    case WORD:
                        r[index++] = binarycodec.toByteArray(
                                Long.toBinaryString(Long.valueOf(s7event.getShort(field.getName()))));
                        break;
                    case DWORD:
                        r[index++] = binarycodec.toByteArray(
                                Integer.toBinaryString(s7event.getInteger(field.getName())));
                        break;
                    case LWORD:
                        r[index++] = binarycodec.toByteArray(
                                Long.toBinaryString(s7event.getLong(field.getName())));
                        break;
                    default:
                        throw new HopException("'DATE, LDATE, DATE_AND_TIME, LDATE_AND_TIME' type is not supported");
                }

            } else if (field.getType().equalsIgnoreCase("Boolean")) {
                switch (valuetype) {
                    case BOOL:
                        r[index++] = s7event.getBoolean(plctags.get(field.getName()).getKey());
                        break;
                    default:
                }
            } else if (field.getType().equalsIgnoreCase("Date")) {
                switch (valuetype) {
                    case DATE:;
                        break;
                    case LDATE:;
                        break;
                    case DATE_AND_TIME:;
                        break;
                    case LDATE_AND_TIME:;
                        break;
                    default:
                        throw new HopException("'DATE, LDATE, DATE_AND_TIME, LDATE_AND_TIME' type is not supported");
                }
            } else if (field.getType().equalsIgnoreCase("Graph")) {
                throw new HopException("'Graph' type is not supported");
            } else if (field.getType().equalsIgnoreCase("Integer")) {
                switch (valuetype) {
                    case BYTE:
                        r[index++] = Long.valueOf(s7event.getByte(field.getName()));
                        break;
                    case WORD:
                        r[index++] = Long.valueOf(s7event.getShort(field.getName()));
                        break;
                    case DWORD:
                        r[index++] = Long.valueOf(s7event.getInteger(field.getName()));
                        break;
                    case LWORD:
                        r[index++] = s7event.getLong(field.getName());
                        break;
                    case INT:
                        r[index++] = Long.valueOf(s7event.getShort(field.getName()));
                        break;
                    case UINT:
                        r[index++] = Long.valueOf((s7event.getShort(field.getName())) & 0xFFFF);
                        break;
                    case SINT:
                        r[index++] = Long.valueOf(s7event.getByte(field.getName()));
                        break;
                    case USINT:
                        r[index++] = Long.valueOf(s7event.getByte(field.getName()) & 0xFF);
                        break;
                    case DINT:
                        r[index++] = Long.valueOf(s7event.getInteger(field.getName()));
                        break;
                    case UDINT:
                        r[index++] = Long.valueOf(s7event.getInteger(field.getName()) & 0xFFFF);
                        break;
                    case LINT:
                        r[index++] = s7event.getLong(field.getName());
                        break;
                    case ULINT:
                        r[index++] = (s7event.getLong(field.getName())) & 0xFFFFFFFFL;
                        break;
                    default:
                        throw new HopException("Tag type is not supported. Check tag definition.");
                }
            } else if (field.getType().equalsIgnoreCase("Internet Address")) {
                throw new HopException("'Internet Address' type is not supported");
            } else if (field.getType().equalsIgnoreCase("JSON")) {
                throw new HopException("'JSON' type is not supported");
            } else if (field.getType().equalsIgnoreCase("Number")) {
                switch (valuetype) {
                    case REAL:
                        System.out.println("Hex: " + Hex.dump((byte[]) s7event.getMap().get(field.getName())));
                        r[index++] = Double.valueOf(s7event.getFloat(field.getName()));
                        break;
                    case LREAL:
                        r[index++] = s7event.getDouble(field.getName());
                        break;
                    default:
                        throw new HopException("Tag type is not supported. Check tag definition.");
                }

            } else if (field.getType().equalsIgnoreCase("String")) {
                switch (valuetype) {
                    case STRING:;
                    case DWORD:
                        r[index++] = String.valueOf(s7event.getFloat(field.getName()));
                        break;
                    default:
                        throw new HopException("'STRING' type is not supported");
                }
            } else if (field.getType().equalsIgnoreCase("Timestamp")) {
                switch (valuetype) {
                    case TIME:;
                        r[index++] = s7event.getTime(plctags.get(field.getName()).getKey());
                        break;
                    case LTIME:;
                        break;
                    case TIME_OF_DAY:;
                        r[index++] = s7event.getTime(plctags.get(field.getName()).getKey());
                        break;
                    case LTIME_OF_DAY:;
                        break;
                    default:
                        throw new HopException("'LTIME, LTIME_OF_DAY' type is not supported");
                }
            }
        }

        putRow(data.outputRowMeta, r); // return your data
        data.rowsWritten++;
        return true;
    }

    @Override
    public boolean init() {
        try {
            if (super.init()) {
                // Determine the number of rows to generate...
                data.rowLimit = Const.toLong(resolve(meta.getRowLimit()), -1L);
                data.rowsWritten = 0L;
                data.delay = Const.toLong(resolve(meta.getIntervalInMs()), -1L);

                if (data.rowLimit < 0L) { // Unable to parse
                    logError(BaseMessages.getString(PKG, "Plc4x.Read.Meta.Wrong.RowLimit.Number"));
                    return false; // fail
                }

                // Create a row (constants) with all the values in it...
                List<ICheckResult> remarks = new ArrayList<>(); // stores the errors...
                RowMetaAndData outputRow = buildRow(meta, remarks, getTransformName());
                if (!remarks.isEmpty()) {
                    for (int i = 0; i < remarks.size(); i++) {
                        CheckResult cr = (CheckResult) remarks.get(i);
                        logError(cr.getText());
                    }

                    return false;
                }

                data.outputRowData = outputRow.getData();
                data.outputRowMeta = outputRow.getRowMeta();

                jobId = null;
                subsRequest = null;
                subsResponse = null;

                getPlcConnection();

                final S7HPlcConnection conn = (S7HPlcConnection) connwrapper.getConnection();
                conn.addEventListener(this);

                return true;
            }
            return false;
        } catch (Exception ex) {
            setErrors(1L);
            logError("Error initializing transform", ex);
            return false;
        }
    }

    /*
    * Here, must perform the cleaning of any resource, main of the connection to 
    * the associated PLC.
     */
    @Override
    public void cleanup() {
        super.cleanup();
        super.cleanup();
        logBasic("Cleanup. Release connection.");
        if (connwrapper != null) {
            connwrapper.release();
            if (connwrapper.refCnt() <= 0) {
                lookup.remove(connwrapper);
            }
            connwrapper = null;
            subsRequest = null;
        }
    }


    /*
    * Here, must perform the cleaning of any resource. 
    * 1. Check if we have reference to wrapper.
    * 2. Release de reference to object.
    * 3. The lastone remove the global reference to connection wrapper.
    * 4. Clear local references.
     */
    @Override
    public void dispose() {

        UnRegisterCYCHandler();

        if (connwrapper.getConnection() instanceof S7HPlcConnection) {
            final S7HPlcConnection conn = (S7HPlcConnection) connwrapper.getConnection();
            conn.removeEventListener(this);
        }

        super.dispose();
        if (connwrapper != null) {
            logBasic("Dispose. Release connection: " + connwrapper.refCnt());
            connwrapper.release();
            if (connwrapper.refCnt() <= 0) {
                lookup.remove(connwrapper);
            }
            connwrapper = null;
            subsRequest = null;
        }
    }

    @Override
    public void stopRunning() throws HopException {
        super.stopRunning();
        stopBundle = true;
    }

    private void getPlcConnection() {
        lock.lock(); //(01)
        try {

            IHopMetadataProvider metaprovider = getMetadataProvider();
            connmeta = metaprovider.getSerializer(Plc4xConnection.class).load(meta.getConnection());

            if (connwrapper == null) {
                template = new Lookup.Template<>(Plc4xWrapperConnection.class, meta.getConnection(), null);
                result = lookup.lookup(template);
                if (!result.allItems().isEmpty()) {
                    logBasic("Using connection: " + meta.getConnection());
                    connwrapper = (Plc4xWrapperConnection) result.allInstances().toArray()[0];
                    if (connwrapper != null) {
                        connwrapper.retain();
                    }
                }
            };

            if (connmeta == null) {
                logError(
                        BaseMessages.getString(
                                PKG,
                                "Plc4x.Read.Meta.Log.SetMetadata",
                                meta.getConnection()));
            }

            if ((connmeta != null) && (connwrapper == null)) {
                subsRequest = null;
                try {
                    final PlcConnection conn = new DefaultPlcDriverManager().getConnection(connmeta.getUrl()); //(03)
                    Thread.sleep(200);
                    if (conn.isConnected()) {
                        logBasic("Create new connection with url : " + connmeta.getUrl());
                        connwrapper = new Plc4xWrapperConnection(conn, meta.getConnection());
                        lookup.add(connwrapper);
                    }

                } catch (Exception ex) {
                    setErrors(1L);
                    logError("Unable to create connection to PLC. " + ex.getMessage());
                }
            }

        } catch (HopException ex) {
            logError(ex.getMessage());
        } finally {
            lock.unlock();
        }
    }

    /*
    * Registers the tags for the cyclical subscription.
    * In the first processing of the rows, a check of the tags is 
    * carried out in order that they are well formed, generating an exception 
    * if they are not.
     */
    public boolean RegisterPlcTags() {
        if ((connmeta != null) && (connwrapper != null)) {
            if (connwrapper.getConnection().isConnected()) {
                if (subsRequest == null) {
                    final PlcSubscriptionRequest.Builder builder = connwrapper.getConnection().subscriptionRequestBuilder(); //(05)
                    for (Plc4xGeneratorField field : meta.getFields()) {
                        builder.addEventTagAddress(field.getName(), field.getItem().toString());
                    }

                    subsRequest = builder.build();

                }
            } else {
                setErrors(1L);
                logError("PLC is not connected.");
                setOutputDone();
                return false;
            }

        } else {
            setErrors(1L);
            logError("PLC connection don't exist.");
            setOutputDone();
            return false;
        }
        return true;
    }

    /*
    * This method performs the subscription to the events within the PLC.
    * 1. Take the waiting time as a reference to monitor the request.
    * 2. Captures the jobId associated with the subscription 
    *    which is assigned by the PLC.
     */
    public void GetSubscriptions() {
        jobId = null;
        try {
            if (null != subsRequest) {
                maxwait = Integer.parseInt(meta.getMaxwaitInMs());
                maxwait = (maxwait < 1000) ? 1000 : maxwait;

                subsResponse = subsRequest.execute().get(maxwait * 2, TimeUnit.MILLISECONDS);

                //Only for Siemens S7
//                subsResponse.getTagNames().forEach(s -> logDebug("Capture JobID: " + s));
                //Only one JobId per subscription will be obtained.
                jobId = (String) subsResponse.getTagNames().toArray()[0];
            }
        } catch (Exception ex) {
            logError(ex.toString());
            jobId = null;
            subsResponse = null;
        }
    }

    /*
    * Subscribe to incoming events by transferring them to
    * the local event buffer.
     */
    public void RegisterCYCHandler() {
        if (null != subsResponse) {
            csmr = subsResponse
                    .getSubscriptionHandle(jobId)
                    .register(evt -> {
                        events.add(evt);
                    });
        }
    }

    /*
    * Unsubscribes to the event generator from the driver and requests
    * to cancel the sending of PLC events for the specified job.
     */
    public void UnRegisterCYCHandler() {
        if (null != csmr) {
            csmr.unregister();

            PlcSubscriptionResponse unregister_response = null;

            final PlcSubscriptionRequest.Builder unregister_builder = connwrapper.getConnection().subscriptionRequestBuilder();
            unregister_builder.addEventTagAddress("CANCELTAG", "CANCEL:" + jobId);

            final PlcSubscriptionRequest unregister_request = unregister_builder.build();

            try {
                unregister_response = unregister_request.execute().get(500, TimeUnit.MILLISECONDS);
            } catch (Exception ex) {
                logError(ex.getMessage());
            }

            logBasic("UnRegisterCYCHandles: " + unregister_response.getResponseCode(jobId));
        }
    }

    /*
    * When the driver makes a connection this method is called. 
    * In the case of the S7 driver, only the connection of the embedded channel
    * is served.
     */
    @Override
    public void connected() {
        if (connwrapper.getConnection().isConnected()) {
            RegisterPlcTags();
            GetSubscriptions();
            RegisterCYCHandler();
        }
    }

    /*
    * This method is called from the driver to indicate that it has 
    * disconnected from the PLC.
    * If the PLC is disconnected, it automatically deletes all 
    * the status associated with this connection.
     */
    @Override
    public void disconnected() {
        logError("Driver disconnected!");
        isConnected = false;
        csmr.unregister();
        csmr = null;
        subsRequest = null;
    }

}
