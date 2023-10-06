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

public class IdentifyReplyCommandMaximumLevels extends IdentifyReplyCommand implements Message {

  // Accessors for discriminator values.
  public Attribute getAttribute() {
    return Attribute.MaximumLevels;
  }

  // Properties.
  protected final byte[] maximumLevels;

  // Arguments.
  protected final Byte numBytes;

  public IdentifyReplyCommandMaximumLevels(byte[] maximumLevels, Byte numBytes) {
    super(numBytes);
    this.maximumLevels = maximumLevels;
    this.numBytes = numBytes;
  }

  public byte[] getMaximumLevels() {
    return maximumLevels;
  }

  @Override
  protected void serializeIdentifyReplyCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("IdentifyReplyCommandMaximumLevels");

    // Array Field (maximumLevels)
    writeByteArrayField("maximumLevels", maximumLevels, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("IdentifyReplyCommandMaximumLevels");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    IdentifyReplyCommandMaximumLevels _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Array field
    if (maximumLevels != null) {
      lengthInBits += 8 * maximumLevels.length;
    }

    return lengthInBits;
  }

  public static IdentifyReplyCommandBuilder staticParseIdentifyReplyCommandBuilder(
      ReadBuffer readBuffer, Attribute attribute, Byte numBytes) throws ParseException {
    readBuffer.pullContext("IdentifyReplyCommandMaximumLevels");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte[] maximumLevels = readBuffer.readByteArray("maximumLevels", Math.toIntExact(numBytes));

    readBuffer.closeContext("IdentifyReplyCommandMaximumLevels");
    // Create the instance
    return new IdentifyReplyCommandMaximumLevelsBuilderImpl(maximumLevels, numBytes);
  }

  public static class IdentifyReplyCommandMaximumLevelsBuilderImpl
      implements IdentifyReplyCommand.IdentifyReplyCommandBuilder {
    private final byte[] maximumLevels;
    private final Byte numBytes;

    public IdentifyReplyCommandMaximumLevelsBuilderImpl(byte[] maximumLevels, Byte numBytes) {
      this.maximumLevels = maximumLevels;
      this.numBytes = numBytes;
    }

    public IdentifyReplyCommandMaximumLevels build(Byte numBytes) {

      IdentifyReplyCommandMaximumLevels identifyReplyCommandMaximumLevels =
          new IdentifyReplyCommandMaximumLevels(maximumLevels, numBytes);
      return identifyReplyCommandMaximumLevels;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IdentifyReplyCommandMaximumLevels)) {
      return false;
    }
    IdentifyReplyCommandMaximumLevels that = (IdentifyReplyCommandMaximumLevels) o;
    return (getMaximumLevels() == that.getMaximumLevels()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMaximumLevels());
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
