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

public class OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1
    extends OpenProtocolMessageLockAtBatchDoneUnsubscribe implements Message {

  // Accessors for discriminator values.
  public Integer getRevision() {
    return (int) 1;
  }

  public OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1(
      Integer midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
  }

  @Override
  protected void serializeOpenProtocolMessageLockAtBatchDoneUnsubscribeChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1");

    writeBuffer.popContext("OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    return lengthInBits;
  }

  public static OpenProtocolMessageLockAtBatchDoneUnsubscribeBuilder
      staticParseOpenProtocolMessageLockAtBatchDoneUnsubscribeBuilder(
          ReadBuffer readBuffer, Integer revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    readBuffer.closeContext("OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1");
    // Create the instance
    return new OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1BuilderImpl();
  }

  public static class OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1BuilderImpl
      implements OpenProtocolMessageLockAtBatchDoneUnsubscribe
          .OpenProtocolMessageLockAtBatchDoneUnsubscribeBuilder {

    public OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1BuilderImpl() {}

    public OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1 build(
        Integer midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1
          openProtocolMessageLockAtBatchDoneUnsubscribeRev1 =
              new OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1(
                  midRevision,
                  noAckFlag,
                  targetStationId,
                  targetSpindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber);
      return openProtocolMessageLockAtBatchDoneUnsubscribeRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1)) {
      return false;
    }
    OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1 that =
        (OpenProtocolMessageLockAtBatchDoneUnsubscribeRev1) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
