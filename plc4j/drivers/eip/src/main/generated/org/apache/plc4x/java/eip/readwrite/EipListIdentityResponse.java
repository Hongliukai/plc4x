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
package org.apache.plc4x.java.eip.readwrite;

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

public class EipListIdentityResponse extends EipPacket implements Message {

  // Accessors for discriminator values.
  public Integer getCommand() {
    return (int) 0x0063;
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  public Integer getPacketLength() {
    return 0;
  }

  // Properties.
  protected final List<CommandSpecificDataItem> items;

  public EipListIdentityResponse(
      long sessionHandle,
      long status,
      byte[] senderContext,
      long options,
      List<CommandSpecificDataItem> items) {
    super(sessionHandle, status, senderContext, options);
    this.items = items;
  }

  public List<CommandSpecificDataItem> getItems() {
    return items;
  }

  @Override
  protected void serializeEipPacketChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("EipListIdentityResponse");

    // Implicit Field (itemCount) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int itemCount = (int) (COUNT(getItems()));
    writeImplicitField("itemCount", itemCount, writeUnsignedInt(writeBuffer, 16));

    // Array Field (items)
    writeComplexTypeArrayField("items", items, writeBuffer);

    writeBuffer.popContext("EipListIdentityResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    EipListIdentityResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (itemCount)
    lengthInBits += 16;

    // Array field
    if (items != null) {
      int i = 0;
      for (CommandSpecificDataItem element : items) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= items.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static EipPacketBuilder staticParseEipPacketBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("EipListIdentityResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int itemCount = readImplicitField("itemCount", readUnsignedInt(readBuffer, 16));

    List<CommandSpecificDataItem> items =
        readCountArrayField(
            "items",
            new DataReaderComplexDefault<>(
                () -> CommandSpecificDataItem.staticParse(readBuffer), readBuffer),
            itemCount);

    readBuffer.closeContext("EipListIdentityResponse");
    // Create the instance
    return new EipListIdentityResponseBuilderImpl(items);
  }

  public static class EipListIdentityResponseBuilderImpl implements EipPacket.EipPacketBuilder {
    private final List<CommandSpecificDataItem> items;

    public EipListIdentityResponseBuilderImpl(List<CommandSpecificDataItem> items) {
      this.items = items;
    }

    public EipListIdentityResponse build(
        long sessionHandle, long status, byte[] senderContext, long options) {
      EipListIdentityResponse eipListIdentityResponse =
          new EipListIdentityResponse(sessionHandle, status, senderContext, options, items);
      return eipListIdentityResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EipListIdentityResponse)) {
      return false;
    }
    EipListIdentityResponse that = (EipListIdentityResponse) o;
    return (getItems() == that.getItems()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getItems());
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
