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

public class IODReadResponseHeader extends PnIoCm_Block implements Message {

  // Accessors for discriminator values.
  public PnIoCm_BlockType getBlockType() {
    return PnIoCm_BlockType.IOD_READ_RES_HEADER;
  }

  // Constant values.
  public static final Integer PADFIELD = 0x0000;

  // Properties.
  protected final short blockVersionHigh;
  protected final short blockVersionLow;
  protected final int sequenceNumber;
  protected final Uuid arUuid;
  protected final long api;
  protected final int slotNumber;
  protected final int subSlotNumber;
  protected final int index;
  protected final long recordDataLength;
  protected final int additionalValue1;
  protected final int additionalValue2;

  public IODReadResponseHeader(
      short blockVersionHigh,
      short blockVersionLow,
      int sequenceNumber,
      Uuid arUuid,
      long api,
      int slotNumber,
      int subSlotNumber,
      int index,
      long recordDataLength,
      int additionalValue1,
      int additionalValue2) {
    super();
    this.blockVersionHigh = blockVersionHigh;
    this.blockVersionLow = blockVersionLow;
    this.sequenceNumber = sequenceNumber;
    this.arUuid = arUuid;
    this.api = api;
    this.slotNumber = slotNumber;
    this.subSlotNumber = subSlotNumber;
    this.index = index;
    this.recordDataLength = recordDataLength;
    this.additionalValue1 = additionalValue1;
    this.additionalValue2 = additionalValue2;
  }

  public short getBlockVersionHigh() {
    return blockVersionHigh;
  }

  public short getBlockVersionLow() {
    return blockVersionLow;
  }

  public int getSequenceNumber() {
    return sequenceNumber;
  }

  public Uuid getArUuid() {
    return arUuid;
  }

  public long getApi() {
    return api;
  }

  public int getSlotNumber() {
    return slotNumber;
  }

  public int getSubSlotNumber() {
    return subSlotNumber;
  }

  public int getIndex() {
    return index;
  }

  public long getRecordDataLength() {
    return recordDataLength;
  }

  public int getAdditionalValue1() {
    return additionalValue1;
  }

  public int getAdditionalValue2() {
    return additionalValue2;
  }

  public int getPadField() {
    return PADFIELD;
  }

  @Override
  protected void serializePnIoCm_BlockChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("IODReadResponseHeader");

    // Implicit Field (blockLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int blockLength =
        (int)
            ((((getIndex()) < (0x8000))
                ? (getLengthInBytes()) - (((4) + (getRecordDataLength())))
                : (getLengthInBytes()) - (4)));
    writeImplicitField(
        "blockLength",
        blockLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionHigh)
    writeSimpleField(
        "blockVersionHigh",
        blockVersionHigh,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (blockVersionLow)
    writeSimpleField(
        "blockVersionLow",
        blockVersionLow,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (sequenceNumber)
    writeSimpleField(
        "sequenceNumber",
        sequenceNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (arUuid)
    writeSimpleField(
        "arUuid",
        arUuid,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (api)
    writeSimpleField(
        "api",
        api,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (slotNumber)
    writeSimpleField(
        "slotNumber",
        slotNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (subSlotNumber)
    writeSimpleField(
        "subSlotNumber",
        subSlotNumber,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (padField)
    writeConstField(
        "padField",
        PADFIELD,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (index)
    writeSimpleField(
        "index",
        index,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (recordDataLength)
    writeSimpleField(
        "recordDataLength",
        recordDataLength,
        writeUnsignedLong(writeBuffer, 32),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (additionalValue1)
    writeSimpleField(
        "additionalValue1",
        additionalValue1,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (additionalValue2)
    writeSimpleField(
        "additionalValue2",
        additionalValue2,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Padding Field (padding)
    writePaddingField(
        "padding",
        (int) (20),
        (short) 0x00,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("IODReadResponseHeader");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    IODReadResponseHeader _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (blockLength)
    lengthInBits += 16;

    // Simple field (blockVersionHigh)
    lengthInBits += 8;

    // Simple field (blockVersionLow)
    lengthInBits += 8;

    // Simple field (sequenceNumber)
    lengthInBits += 16;

    // Simple field (arUuid)
    lengthInBits += arUuid.getLengthInBits();

    // Simple field (api)
    lengthInBits += 32;

    // Simple field (slotNumber)
    lengthInBits += 16;

    // Simple field (subSlotNumber)
    lengthInBits += 16;

    // Const Field (padField)
    lengthInBits += 16;

    // Simple field (index)
    lengthInBits += 16;

    // Simple field (recordDataLength)
    lengthInBits += 32;

    // Simple field (additionalValue1)
    lengthInBits += 16;

    // Simple field (additionalValue2)
    lengthInBits += 16;

    // Padding Field (padding)
    int _timesPadding = (int) (20);
    while (_timesPadding-- > 0) {
      lengthInBits += 8;
    }

    return lengthInBits;
  }

  public static PnIoCm_BlockBuilder staticParsePnIoCm_BlockBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("IODReadResponseHeader");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockLength =
        readImplicitField(
            "blockLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionHigh =
        readSimpleField(
            "blockVersionHigh",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short blockVersionLow =
        readSimpleField(
            "blockVersionLow",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int sequenceNumber =
        readSimpleField(
            "sequenceNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    Uuid arUuid =
        readSimpleField(
            "arUuid",
            new DataReaderComplexDefault<>(() -> Uuid.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    long api =
        readSimpleField(
            "api",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int slotNumber =
        readSimpleField(
            "slotNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int subSlotNumber =
        readSimpleField(
            "subSlotNumber",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int padField =
        readConstField(
            "padField",
            readUnsignedInt(readBuffer, 16),
            IODReadResponseHeader.PADFIELD,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int index =
        readSimpleField(
            "index",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    long recordDataLength =
        readSimpleField(
            "recordDataLength",
            readUnsignedLong(readBuffer, 32),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int additionalValue1 =
        readSimpleField(
            "additionalValue1",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int additionalValue2 =
        readSimpleField(
            "additionalValue2",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readPaddingField(
        readUnsignedShort(readBuffer, 8),
        (int) (20),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("IODReadResponseHeader");
    // Create the instance
    return new IODReadResponseHeaderBuilderImpl(
        blockVersionHigh,
        blockVersionLow,
        sequenceNumber,
        arUuid,
        api,
        slotNumber,
        subSlotNumber,
        index,
        recordDataLength,
        additionalValue1,
        additionalValue2);
  }

  public static class IODReadResponseHeaderBuilderImpl implements PnIoCm_Block.PnIoCm_BlockBuilder {
    private final short blockVersionHigh;
    private final short blockVersionLow;
    private final int sequenceNumber;
    private final Uuid arUuid;
    private final long api;
    private final int slotNumber;
    private final int subSlotNumber;
    private final int index;
    private final long recordDataLength;
    private final int additionalValue1;
    private final int additionalValue2;

    public IODReadResponseHeaderBuilderImpl(
        short blockVersionHigh,
        short blockVersionLow,
        int sequenceNumber,
        Uuid arUuid,
        long api,
        int slotNumber,
        int subSlotNumber,
        int index,
        long recordDataLength,
        int additionalValue1,
        int additionalValue2) {
      this.blockVersionHigh = blockVersionHigh;
      this.blockVersionLow = blockVersionLow;
      this.sequenceNumber = sequenceNumber;
      this.arUuid = arUuid;
      this.api = api;
      this.slotNumber = slotNumber;
      this.subSlotNumber = subSlotNumber;
      this.index = index;
      this.recordDataLength = recordDataLength;
      this.additionalValue1 = additionalValue1;
      this.additionalValue2 = additionalValue2;
    }

    public IODReadResponseHeader build() {
      IODReadResponseHeader iODReadResponseHeader =
          new IODReadResponseHeader(
              blockVersionHigh,
              blockVersionLow,
              sequenceNumber,
              arUuid,
              api,
              slotNumber,
              subSlotNumber,
              index,
              recordDataLength,
              additionalValue1,
              additionalValue2);
      return iODReadResponseHeader;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IODReadResponseHeader)) {
      return false;
    }
    IODReadResponseHeader that = (IODReadResponseHeader) o;
    return (getBlockVersionHigh() == that.getBlockVersionHigh())
        && (getBlockVersionLow() == that.getBlockVersionLow())
        && (getSequenceNumber() == that.getSequenceNumber())
        && (getArUuid() == that.getArUuid())
        && (getApi() == that.getApi())
        && (getSlotNumber() == that.getSlotNumber())
        && (getSubSlotNumber() == that.getSubSlotNumber())
        && (getIndex() == that.getIndex())
        && (getRecordDataLength() == that.getRecordDataLength())
        && (getAdditionalValue1() == that.getAdditionalValue1())
        && (getAdditionalValue2() == that.getAdditionalValue2())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getBlockVersionHigh(),
        getBlockVersionLow(),
        getSequenceNumber(),
        getArUuid(),
        getApi(),
        getSlotNumber(),
        getSubSlotNumber(),
        getIndex(),
        getRecordDataLength(),
        getAdditionalValue1(),
        getAdditionalValue2());
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
