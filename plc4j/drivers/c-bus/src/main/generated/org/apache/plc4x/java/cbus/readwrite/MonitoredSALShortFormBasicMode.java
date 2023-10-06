/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.cbus.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class MonitoredSALShortFormBasicMode extends MonitoredSAL implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte counts;
  protected final Short bridgeCount;
  protected final Short networkNumber;
  protected final Byte noCounts;
  protected final ApplicationIdContainer application;
  protected final SALData salData;

  // Arguments.
  protected final CBusOptions cBusOptions;

  public MonitoredSALShortFormBasicMode(
      byte salType,
      byte counts,
      Short bridgeCount,
      Short networkNumber,
      Byte noCounts,
      ApplicationIdContainer application,
      SALData salData,
      CBusOptions cBusOptions) {
    super(salType, cBusOptions);
    this.counts = counts;
    this.bridgeCount = bridgeCount;
    this.networkNumber = networkNumber;
    this.noCounts = noCounts;
    this.application = application;
    this.salData = salData;
    this.cBusOptions = cBusOptions;
  }

  public byte getCounts() {
    return counts;
  }

  public Short getBridgeCount() {
    return bridgeCount;
  }

  public Short getNetworkNumber() {
    return networkNumber;
  }

  public Byte getNoCounts() {
    return noCounts;
  }

  public ApplicationIdContainer getApplication() {
    return application;
  }

  public SALData getSalData() {
    return salData;
  }

  @Override
  protected void serializeMonitoredSALChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MonitoredSALShortFormBasicMode");

    // Optional Field (bridgeCount) (Can be skipped, if the value is null)
    writeOptionalField(
        "bridgeCount", bridgeCount, writeUnsignedShort(writeBuffer, 8), (getCounts()) != (0x00));

    // Optional Field (networkNumber) (Can be skipped, if the value is null)
    writeOptionalField(
        "networkNumber",
        networkNumber,
        writeUnsignedShort(writeBuffer, 8),
        (getCounts()) != (0x00));

    // Optional Field (noCounts) (Can be skipped, if the value is null)
    writeOptionalField("noCounts", noCounts, writeByte(writeBuffer, 8), (getCounts()) == (0x00));

    // Simple Field (application)
    writeSimpleEnumField(
        "application",
        "ApplicationIdContainer",
        application,
        new DataWriterEnumDefault<>(
            ApplicationIdContainer::getValue,
            ApplicationIdContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Optional Field (salData) (Can be skipped, if the value is null)
    writeOptionalField("salData", salData, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("MonitoredSALShortFormBasicMode");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MonitoredSALShortFormBasicMode _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Optional Field (bridgeCount)
    if (bridgeCount != null) {
      lengthInBits += 8;
    }

    // Optional Field (networkNumber)
    if (networkNumber != null) {
      lengthInBits += 8;
    }

    // Optional Field (noCounts)
    if (noCounts != null) {
      lengthInBits += 8;
    }

    // Simple field (application)
    lengthInBits += 8;

    // Optional Field (salData)
    if (salData != null) {
      lengthInBits += salData.getLengthInBits();
    }

    return lengthInBits;
  }

  public static MonitoredSALBuilder staticParseMonitoredSALBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("MonitoredSALShortFormBasicMode");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte counts = readPeekField("counts", readByte(readBuffer, 8));

    Short bridgeCount =
        readOptionalField("bridgeCount", readUnsignedShort(readBuffer, 8), (counts) != (0x00));

    Short networkNumber =
        readOptionalField("networkNumber", readUnsignedShort(readBuffer, 8), (counts) != (0x00));

    Byte noCounts = readOptionalField("noCounts", readByte(readBuffer, 8), (counts) == (0x00));

    ApplicationIdContainer application =
        readEnumField(
            "application",
            "ApplicationIdContainer",
            new DataReaderEnumDefault<>(
                ApplicationIdContainer::enumForValue, readUnsignedShort(readBuffer, 8)));

    SALData salData =
        readOptionalField(
            "salData",
            new DataReaderComplexDefault<>(
                () ->
                    SALData.staticParse(
                        readBuffer, (ApplicationId) (application.getApplicationId())),
                readBuffer));

    readBuffer.closeContext("MonitoredSALShortFormBasicMode");
    // Create the instance
    return new MonitoredSALShortFormBasicModeBuilderImpl(
        counts, bridgeCount, networkNumber, noCounts, application, salData, cBusOptions);
  }

  public static class MonitoredSALShortFormBasicModeBuilderImpl
      implements MonitoredSAL.MonitoredSALBuilder {
    private final byte counts;
    private final Short bridgeCount;
    private final Short networkNumber;
    private final Byte noCounts;
    private final ApplicationIdContainer application;
    private final SALData salData;
    private final CBusOptions cBusOptions;

    public MonitoredSALShortFormBasicModeBuilderImpl(
        byte counts,
        Short bridgeCount,
        Short networkNumber,
        Byte noCounts,
        ApplicationIdContainer application,
        SALData salData,
        CBusOptions cBusOptions) {
      this.counts = counts;
      this.bridgeCount = bridgeCount;
      this.networkNumber = networkNumber;
      this.noCounts = noCounts;
      this.application = application;
      this.salData = salData;
      this.cBusOptions = cBusOptions;
    }

    public MonitoredSALShortFormBasicMode build(byte salType, CBusOptions cBusOptions) {
      MonitoredSALShortFormBasicMode monitoredSALShortFormBasicMode =
          new MonitoredSALShortFormBasicMode(
              salType,
              counts,
              bridgeCount,
              networkNumber,
              noCounts,
              application,
              salData,
              cBusOptions);
      return monitoredSALShortFormBasicMode;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MonitoredSALShortFormBasicMode)) {
      return false;
    }
    MonitoredSALShortFormBasicMode that = (MonitoredSALShortFormBasicMode) o;
    return (getCounts() == that.getCounts())
        && (getBridgeCount() == that.getBridgeCount())
        && (getNetworkNumber() == that.getNetworkNumber())
        && (getNoCounts() == that.getNoCounts())
        && (getApplication() == that.getApplication())
        && (getSalData() == that.getSalData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getCounts(),
        getBridgeCount(),
        getNetworkNumber(),
        getNoCounts(),
        getApplication(),
        getSalData());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
