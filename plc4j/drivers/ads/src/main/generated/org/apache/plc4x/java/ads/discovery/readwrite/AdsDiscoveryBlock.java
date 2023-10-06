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
package org.apache.plc4x.java.ads.discovery.readwrite;

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

public abstract class AdsDiscoveryBlock implements Message {

  // Abstract accessors for discriminator values.
  public abstract AdsDiscoveryBlockType getBlockType();

  public AdsDiscoveryBlock() {
    super();
  }

  protected abstract void serializeAdsDiscoveryBlockChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AdsDiscoveryBlock");

    // Discriminator Field (blockType) (Used as input to a switch field)
    writeDiscriminatorEnumField(
        "blockType",
        "AdsDiscoveryBlockType",
        getBlockType(),
        new DataWriterEnumDefault<>(
            AdsDiscoveryBlockType::getValue,
            AdsDiscoveryBlockType::name,
            writeUnsignedInt(writeBuffer, 16)));

    // Switch field (Serialize the sub-type)
    serializeAdsDiscoveryBlockChild(writeBuffer);

    writeBuffer.popContext("AdsDiscoveryBlock");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AdsDiscoveryBlock _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (blockType)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static AdsDiscoveryBlock staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static AdsDiscoveryBlock staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AdsDiscoveryBlock");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    AdsDiscoveryBlockType blockType =
        readDiscriminatorField(
            "blockType",
            new DataReaderEnumDefault<>(
                AdsDiscoveryBlockType::enumForValue, readUnsignedInt(readBuffer, 16)));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    AdsDiscoveryBlockBuilder builder = null;
    if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.STATUS)) {
      builder = AdsDiscoveryBlockStatus.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.PASSWORD)) {
      builder = AdsDiscoveryBlockPassword.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.VERSION)) {
      builder = AdsDiscoveryBlockVersion.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.OS_DATA)) {
      builder = AdsDiscoveryBlockOsData.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.HOST_NAME)) {
      builder = AdsDiscoveryBlockHostName.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.AMS_NET_ID)) {
      builder = AdsDiscoveryBlockAmsNetId.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.ROUTE_NAME)) {
      builder = AdsDiscoveryBlockRouteName.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.USER_NAME)) {
      builder = AdsDiscoveryBlockUserName.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    } else if (EvaluationHelper.equals(blockType, AdsDiscoveryBlockType.FINGERPRINT)) {
      builder = AdsDiscoveryBlockFingerprint.staticParseAdsDiscoveryBlockBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "blockType="
              + blockType
              + "]");
    }

    readBuffer.closeContext("AdsDiscoveryBlock");
    // Create the instance
    AdsDiscoveryBlock _adsDiscoveryBlock = builder.build();
    return _adsDiscoveryBlock;
  }

  public interface AdsDiscoveryBlockBuilder {
    AdsDiscoveryBlock build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsDiscoveryBlock)) {
      return false;
    }
    AdsDiscoveryBlock that = (AdsDiscoveryBlock) o;
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
