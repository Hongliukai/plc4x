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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class QualifierOfParameterOfMeasuredValues implements Message {

  // Properties.
  protected final boolean parameterInOperation;
  protected final boolean localParameterChange;
  protected final byte kindOfParameter;

  public QualifierOfParameterOfMeasuredValues(
      boolean parameterInOperation, boolean localParameterChange, byte kindOfParameter) {
    super();
    this.parameterInOperation = parameterInOperation;
    this.localParameterChange = localParameterChange;
    this.kindOfParameter = kindOfParameter;
  }

  public boolean getParameterInOperation() {
    return parameterInOperation;
  }

  public boolean getLocalParameterChange() {
    return localParameterChange;
  }

  public byte getKindOfParameter() {
    return kindOfParameter;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("QualifierOfParameterOfMeasuredValues");

    // Simple Field (parameterInOperation)
    writeSimpleField(
        "parameterInOperation",
        parameterInOperation,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (localParameterChange)
    writeSimpleField(
        "localParameterChange",
        localParameterChange,
        writeBoolean(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (kindOfParameter)
    writeSimpleField(
        "kindOfParameter",
        kindOfParameter,
        writeUnsignedByte(writeBuffer, 6),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("QualifierOfParameterOfMeasuredValues");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    QualifierOfParameterOfMeasuredValues _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (parameterInOperation)
    lengthInBits += 1;

    // Simple field (localParameterChange)
    lengthInBits += 1;

    // Simple field (kindOfParameter)
    lengthInBits += 6;

    return lengthInBits;
  }

  public static QualifierOfParameterOfMeasuredValues staticParse(
      ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static QualifierOfParameterOfMeasuredValues staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("QualifierOfParameterOfMeasuredValues");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean parameterInOperation =
        readSimpleField(
            "parameterInOperation",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    boolean localParameterChange =
        readSimpleField(
            "localParameterChange",
            readBoolean(readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    byte kindOfParameter =
        readSimpleField(
            "kindOfParameter",
            readUnsignedByte(readBuffer, 6),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("QualifierOfParameterOfMeasuredValues");
    // Create the instance
    QualifierOfParameterOfMeasuredValues _qualifierOfParameterOfMeasuredValues;
    _qualifierOfParameterOfMeasuredValues =
        new QualifierOfParameterOfMeasuredValues(
            parameterInOperation, localParameterChange, kindOfParameter);
    return _qualifierOfParameterOfMeasuredValues;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof QualifierOfParameterOfMeasuredValues)) {
      return false;
    }
    QualifierOfParameterOfMeasuredValues that = (QualifierOfParameterOfMeasuredValues) o;
    return (getParameterInOperation() == that.getParameterInOperation())
        && (getLocalParameterChange() == that.getLocalParameterChange())
        && (getKindOfParameter() == that.getKindOfParameter())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getParameterInOperation(), getLocalParameterChange(), getKindOfParameter());
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
