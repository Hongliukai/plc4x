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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetLogDataLogDataEntryIntegerValue extends BACnetLogDataLogDataEntry
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagSignedInteger integerValue;

  public BACnetLogDataLogDataEntryIntegerValue(
      BACnetTagHeader peekedTagHeader, BACnetContextTagSignedInteger integerValue) {
    super(peekedTagHeader);
    this.integerValue = integerValue;
  }

  public BACnetContextTagSignedInteger getIntegerValue() {
    return integerValue;
  }

  @Override
  protected void serializeBACnetLogDataLogDataEntryChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetLogDataLogDataEntryIntegerValue");

    // Simple Field (integerValue)
    writeSimpleField("integerValue", integerValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetLogDataLogDataEntryIntegerValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetLogDataLogDataEntryIntegerValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (integerValue)
    lengthInBits += integerValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLogDataLogDataEntryBuilder staticParseBACnetLogDataLogDataEntryBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetLogDataLogDataEntryIntegerValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagSignedInteger integerValue =
        readSimpleField(
            "integerValue",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagSignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (4),
                            (BACnetDataType) (BACnetDataType.SIGNED_INTEGER)),
                readBuffer));

    readBuffer.closeContext("BACnetLogDataLogDataEntryIntegerValue");
    // Create the instance
    return new BACnetLogDataLogDataEntryIntegerValueBuilderImpl(integerValue);
  }

  public static class BACnetLogDataLogDataEntryIntegerValueBuilderImpl
      implements BACnetLogDataLogDataEntry.BACnetLogDataLogDataEntryBuilder {
    private final BACnetContextTagSignedInteger integerValue;

    public BACnetLogDataLogDataEntryIntegerValueBuilderImpl(
        BACnetContextTagSignedInteger integerValue) {
      this.integerValue = integerValue;
    }

    public BACnetLogDataLogDataEntryIntegerValue build(BACnetTagHeader peekedTagHeader) {
      BACnetLogDataLogDataEntryIntegerValue bACnetLogDataLogDataEntryIntegerValue =
          new BACnetLogDataLogDataEntryIntegerValue(peekedTagHeader, integerValue);
      return bACnetLogDataLogDataEntryIntegerValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogDataLogDataEntryIntegerValue)) {
      return false;
    }
    BACnetLogDataLogDataEntryIntegerValue that = (BACnetLogDataLogDataEntryIntegerValue) o;
    return (getIntegerValue() == that.getIntegerValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getIntegerValue());
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
