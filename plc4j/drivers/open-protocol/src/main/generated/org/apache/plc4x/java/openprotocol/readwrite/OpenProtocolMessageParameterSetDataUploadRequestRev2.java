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
package org.apache.plc4x.java.openprotocol.readwrite;

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

public class OpenProtocolMessageParameterSetDataUploadRequestRev2
    extends OpenProtocolMessageParameterSetDataUploadRequest implements Message {

  // Accessors for discriminator values.
  public Integer getRevision() {
    return (int) 2;
  }

  // Properties.
  protected final int parameterSetId;

  public OpenProtocolMessageParameterSetDataUploadRequestRev2(
      Integer midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      int parameterSetId) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.parameterSetId = parameterSetId;
  }

  public int getParameterSetId() {
    return parameterSetId;
  }

  @Override
  protected void serializeOpenProtocolMessageParameterSetDataUploadRequestChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenProtocolMessageParameterSetDataUploadRequestRev2");

    // Simple Field (parameterSetId)
    writeSimpleField(
        "parameterSetId",
        parameterSetId,
        writeUnsignedInt(writeBuffer, 24),
        WithOption.WithEncoding("ASCII"));

    writeBuffer.popContext("OpenProtocolMessageParameterSetDataUploadRequestRev2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageParameterSetDataUploadRequestRev2 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (parameterSetId)
    lengthInBits += 24;

    return lengthInBits;
  }

  public static OpenProtocolMessageParameterSetDataUploadRequestBuilder
      staticParseOpenProtocolMessageParameterSetDataUploadRequestBuilder(
          ReadBuffer readBuffer, Integer revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageParameterSetDataUploadRequestRev2");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int parameterSetId =
        readSimpleField(
            "parameterSetId", readUnsignedInt(readBuffer, 24), WithOption.WithEncoding("ASCII"));

    readBuffer.closeContext("OpenProtocolMessageParameterSetDataUploadRequestRev2");
    // Create the instance
    return new OpenProtocolMessageParameterSetDataUploadRequestRev2BuilderImpl(parameterSetId);
  }

  public static class OpenProtocolMessageParameterSetDataUploadRequestRev2BuilderImpl
      implements OpenProtocolMessageParameterSetDataUploadRequest
          .OpenProtocolMessageParameterSetDataUploadRequestBuilder {
    private final int parameterSetId;

    public OpenProtocolMessageParameterSetDataUploadRequestRev2BuilderImpl(int parameterSetId) {
      this.parameterSetId = parameterSetId;
    }

    public OpenProtocolMessageParameterSetDataUploadRequestRev2 build(
        Integer midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageParameterSetDataUploadRequestRev2
          openProtocolMessageParameterSetDataUploadRequestRev2 =
              new OpenProtocolMessageParameterSetDataUploadRequestRev2(
                  midRevision,
                  noAckFlag,
                  targetStationId,
                  targetSpindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber,
                  parameterSetId);
      return openProtocolMessageParameterSetDataUploadRequestRev2;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageParameterSetDataUploadRequestRev2)) {
      return false;
    }
    OpenProtocolMessageParameterSetDataUploadRequestRev2 that =
        (OpenProtocolMessageParameterSetDataUploadRequestRev2) o;
    return (getParameterSetId() == that.getParameterSetId()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getParameterSetId());
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
