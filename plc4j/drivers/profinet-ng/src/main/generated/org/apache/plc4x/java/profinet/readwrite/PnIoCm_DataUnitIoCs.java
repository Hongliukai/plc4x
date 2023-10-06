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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnIoCm_DataUnitIoCs implements Message {

  // Properties.
  protected final boolean dataState;
  protected final byte instance;
  protected final boolean extension;

  // Reserved Fields
  private Byte reservedField0;

  public PnIoCm_DataUnitIoCs(boolean dataState, byte instance, boolean extension) {
    super();
    this.dataState = dataState;
    this.instance = instance;
    this.extension = extension;
  }

  public boolean getDataState() {
    return dataState;
  }

  public byte getInstance() {
    return instance;
  }

  public boolean getExtension() {
    return extension;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnIoCm_DataUnitIoCs");

    // Simple Field (dataState)
    writeSimpleField(
        "dataState",
        dataState,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (instance)
    writeSimpleField(
        "instance",
        instance,
        writeUnsignedByte(writeBuffer, 2),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0x00,
        writeUnsignedByte(writeBuffer, 4),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (extension)
    writeSimpleField(
        "extension",
        extension,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCm_DataUnitIoCs");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    PnIoCm_DataUnitIoCs _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (dataState)
    lengthInBits += 1;

    // Simple field (instance)
    lengthInBits += 2;

    // Reserved Field (reserved)
    lengthInBits += 4;

    // Simple field (extension)
    lengthInBits += 1;

    return lengthInBits;
  }

  public static PnIoCm_DataUnitIoCs staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static PnIoCm_DataUnitIoCs staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("PnIoCm_DataUnitIoCs");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean dataState =
        readSimpleField(
            "dataState", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    byte instance =
        readSimpleField(
            "instance",
            readUnsignedByte(readBuffer, 2),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Byte reservedField0 =
        readReservedField(
            "reserved",
            readUnsignedByte(readBuffer, 4),
            (byte) 0x00,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    boolean extension =
        readSimpleField(
            "extension", readBoolean(readBuffer), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_DataUnitIoCs");
    // Create the instance
    PnIoCm_DataUnitIoCs _pnIoCm_DataUnitIoCs;
    _pnIoCm_DataUnitIoCs = new PnIoCm_DataUnitIoCs(dataState, instance, extension);
    _pnIoCm_DataUnitIoCs.reservedField0 = reservedField0;
    return _pnIoCm_DataUnitIoCs;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_DataUnitIoCs)) {
      return false;
    }
    PnIoCm_DataUnitIoCs that = (PnIoCm_DataUnitIoCs) o;
    return (getDataState() == that.getDataState())
        && (getInstance() == that.getInstance())
        && (getExtension() == that.getExtension())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getDataState(), getInstance(), getExtension());
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
