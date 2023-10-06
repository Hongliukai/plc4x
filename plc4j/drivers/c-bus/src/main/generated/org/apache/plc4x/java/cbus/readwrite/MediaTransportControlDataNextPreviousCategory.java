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

public class MediaTransportControlDataNextPreviousCategory extends MediaTransportControlData
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte operation;

  public MediaTransportControlDataNextPreviousCategory(
      MediaTransportControlCommandTypeContainer commandTypeContainer,
      byte mediaLinkGroup,
      byte operation) {
    super(commandTypeContainer, mediaLinkGroup);
    this.operation = operation;
  }

  public byte getOperation() {
    return operation;
  }

  public boolean getIsSetThePreviousCategory() {
    return (boolean) ((getOperation()) == (0x00));
  }

  public boolean getIsSetTheNextCategory() {
    return (boolean) ((getOperation()) != (0x00));
  }

  @Override
  protected void serializeMediaTransportControlDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MediaTransportControlDataNextPreviousCategory");

    // Simple Field (operation)
    writeSimpleField("operation", operation, writeByte(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isSetThePreviousCategory = getIsSetThePreviousCategory();
    writeBuffer.writeVirtual("isSetThePreviousCategory", isSetThePreviousCategory);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isSetTheNextCategory = getIsSetTheNextCategory();
    writeBuffer.writeVirtual("isSetTheNextCategory", isSetTheNextCategory);

    writeBuffer.popContext("MediaTransportControlDataNextPreviousCategory");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MediaTransportControlDataNextPreviousCategory _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (operation)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static MediaTransportControlDataBuilder staticParseMediaTransportControlDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("MediaTransportControlDataNextPreviousCategory");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte operation = readSimpleField("operation", readByte(readBuffer, 8));
    boolean isSetThePreviousCategory =
        readVirtualField("isSetThePreviousCategory", boolean.class, (operation) == (0x00));
    boolean isSetTheNextCategory =
        readVirtualField("isSetTheNextCategory", boolean.class, (operation) != (0x00));

    readBuffer.closeContext("MediaTransportControlDataNextPreviousCategory");
    // Create the instance
    return new MediaTransportControlDataNextPreviousCategoryBuilderImpl(operation);
  }

  public static class MediaTransportControlDataNextPreviousCategoryBuilderImpl
      implements MediaTransportControlData.MediaTransportControlDataBuilder {
    private final byte operation;

    public MediaTransportControlDataNextPreviousCategoryBuilderImpl(byte operation) {
      this.operation = operation;
    }

    public MediaTransportControlDataNextPreviousCategory build(
        MediaTransportControlCommandTypeContainer commandTypeContainer, byte mediaLinkGroup) {
      MediaTransportControlDataNextPreviousCategory mediaTransportControlDataNextPreviousCategory =
          new MediaTransportControlDataNextPreviousCategory(
              commandTypeContainer, mediaLinkGroup, operation);
      return mediaTransportControlDataNextPreviousCategory;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MediaTransportControlDataNextPreviousCategory)) {
      return false;
    }
    MediaTransportControlDataNextPreviousCategory that =
        (MediaTransportControlDataNextPreviousCategory) o;
    return (getOperation() == that.getOperation()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getOperation());
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
