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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7VarPayloadDataItem implements Message {

  // Properties.
  protected final DataTransportErrorCode returnCode;
  protected final DataTransportSize transportSize;
  protected final byte[] data;

  public S7VarPayloadDataItem(
      DataTransportErrorCode returnCode, DataTransportSize transportSize, byte[] data) {
    super();
    this.returnCode = returnCode;
    this.transportSize = transportSize;
    this.data = data;
  }

  public DataTransportErrorCode getReturnCode() {
    return returnCode;
  }

  public DataTransportSize getTransportSize() {
    return transportSize;
  }

  public byte[] getData() {
    return data;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7VarPayloadDataItem");

    // Simple Field (returnCode)
    writeSimpleEnumField(
        "returnCode",
        "DataTransportErrorCode",
        returnCode,
        new DataWriterEnumDefault<>(
            DataTransportErrorCode::getValue,
            DataTransportErrorCode::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (transportSize)
    writeSimpleEnumField(
        "transportSize",
        "DataTransportSize",
        transportSize,
        new DataWriterEnumDefault<>(
            DataTransportSize::getValue,
            DataTransportSize::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Implicit Field (dataLength) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int dataLength =
        (int)
            ((COUNT(getData()))
                * ((((((getTransportSize()) == (DataTransportSize.BIT)))
                    ? 1
                    : (((getTransportSize().getSizeInBits()) ? 8 : 1))))));
    writeImplicitField("dataLength", dataLength, writeUnsignedInt(writeBuffer, 16));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    // Padding Field (padding)
    writePaddingField(
        "padding",
        (int) ((((!(_lastItem))) ? ((COUNT(data)) % (2)) : 0)),
        (short) 0x00,
        writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("S7VarPayloadDataItem");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    S7VarPayloadDataItem _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (returnCode)
    lengthInBits += 8;

    // Simple field (transportSize)
    lengthInBits += 8;

    // Implicit Field (dataLength)
    lengthInBits += 16;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    // Padding Field (padding)
    int _timesPadding = (int) ((((!(_lastItem))) ? ((COUNT(data)) % (2)) : 0));
    while (_timesPadding-- > 0) {
      lengthInBits += 8;
    }

    return lengthInBits;
  }

  public static S7VarPayloadDataItem staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static S7VarPayloadDataItem staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("S7VarPayloadDataItem");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    DataTransportErrorCode returnCode =
        readEnumField(
            "returnCode",
            "DataTransportErrorCode",
            new DataReaderEnumDefault<>(
                DataTransportErrorCode::enumForValue, readUnsignedShort(readBuffer, 8)));

    DataTransportSize transportSize =
        readEnumField(
            "transportSize",
            "DataTransportSize",
            new DataReaderEnumDefault<>(
                DataTransportSize::enumForValue, readUnsignedShort(readBuffer, 8)));

    int dataLength = readImplicitField("dataLength", readUnsignedInt(readBuffer, 16));

    byte[] data =
        readBuffer.readByteArray(
            "data",
            Math.toIntExact(
                ((transportSize.getSizeInBits()) ? CEIL((dataLength) / (8.0)) : dataLength)));

    readPaddingField(
        readUnsignedShort(readBuffer, 8), (int) ((((!(_lastItem))) ? ((COUNT(data)) % (2)) : 0)));

    readBuffer.closeContext("S7VarPayloadDataItem");
    // Create the instance
    S7VarPayloadDataItem _s7VarPayloadDataItem;
    _s7VarPayloadDataItem = new S7VarPayloadDataItem(returnCode, transportSize, data);
    return _s7VarPayloadDataItem;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7VarPayloadDataItem)) {
      return false;
    }
    S7VarPayloadDataItem that = (S7VarPayloadDataItem) o;
    return (getReturnCode() == that.getReturnCode())
        && (getTransportSize() == that.getTransportSize())
        && (getData() == that.getData())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getReturnCode(), getTransportSize(), getData());
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
