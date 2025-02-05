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

public abstract class Ethernet_FramePayload implements Message {

  // Abstract accessors for discriminator values.
  public abstract Integer getPacketType();

  public Ethernet_FramePayload() {
    super();
  }

  protected abstract void serializeEthernet_FramePayloadChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("Ethernet_FramePayload");

    // Discriminator Field (packetType) (Used as input to a switch field)
    writeDiscriminatorField("packetType", getPacketType(), writeUnsignedInt(writeBuffer, 16));

    // Switch field (Serialize the sub-type)
    serializeEthernet_FramePayloadChild(writeBuffer);

    writeBuffer.popContext("Ethernet_FramePayload");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    Ethernet_FramePayload _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (packetType)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static Ethernet_FramePayload staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static Ethernet_FramePayload staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("Ethernet_FramePayload");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int packetType = readDiscriminatorField("packetType", readUnsignedInt(readBuffer, 16));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    Ethernet_FramePayloadBuilder builder = null;
    if (EvaluationHelper.equals(packetType, (int) 0x0800)) {
      builder = Ethernet_FramePayload_IPv4.staticParseEthernet_FramePayloadBuilder(readBuffer);
    } else if (EvaluationHelper.equals(packetType, (int) 0x8100)) {
      builder =
          Ethernet_FramePayload_VirtualLan.staticParseEthernet_FramePayloadBuilder(readBuffer);
    } else if (EvaluationHelper.equals(packetType, (int) 0x8892)) {
      builder = Ethernet_FramePayload_PnDcp.staticParseEthernet_FramePayloadBuilder(readBuffer);
    } else if (EvaluationHelper.equals(packetType, (int) 0x88cc)) {
      builder = Ethernet_FramePayload_LLDP.staticParseEthernet_FramePayloadBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "packetType="
              + packetType
              + "]");
    }

    readBuffer.closeContext("Ethernet_FramePayload");
    // Create the instance
    Ethernet_FramePayload _ethernet_FramePayload = builder.build();
    return _ethernet_FramePayload;
  }

  public interface Ethernet_FramePayloadBuilder {
    Ethernet_FramePayload build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Ethernet_FramePayload)) {
      return false;
    }
    Ethernet_FramePayload that = (Ethernet_FramePayload) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
