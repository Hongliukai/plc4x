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

public class BACnetConstructedDataAutoSlaveDiscovery extends BACnetConstructedData
    implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return null;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.AUTO_SLAVE_DISCOVERY;
  }

  // Properties.
  protected final BACnetApplicationTagBoolean autoSlaveDiscovery;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataAutoSlaveDiscovery(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetApplicationTagBoolean autoSlaveDiscovery,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.autoSlaveDiscovery = autoSlaveDiscovery;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetApplicationTagBoolean getAutoSlaveDiscovery() {
    return autoSlaveDiscovery;
  }

  public BACnetApplicationTagBoolean getActualValue() {
    return (BACnetApplicationTagBoolean) (getAutoSlaveDiscovery());
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConstructedDataAutoSlaveDiscovery");

    // Simple Field (autoSlaveDiscovery)
    writeSimpleField(
        "autoSlaveDiscovery", autoSlaveDiscovery, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetApplicationTagBoolean actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetConstructedDataAutoSlaveDiscovery");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataAutoSlaveDiscovery _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (autoSlaveDiscovery)
    lengthInBits += autoSlaveDiscovery.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataAutoSlaveDiscovery");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagBoolean autoSlaveDiscovery =
        readSimpleField(
            "autoSlaveDiscovery",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagBoolean) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));
    BACnetApplicationTagBoolean actualValue =
        readVirtualField("actualValue", BACnetApplicationTagBoolean.class, autoSlaveDiscovery);

    readBuffer.closeContext("BACnetConstructedDataAutoSlaveDiscovery");
    // Create the instance
    return new BACnetConstructedDataAutoSlaveDiscoveryBuilderImpl(
        autoSlaveDiscovery, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataAutoSlaveDiscoveryBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final BACnetApplicationTagBoolean autoSlaveDiscovery;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataAutoSlaveDiscoveryBuilderImpl(
        BACnetApplicationTagBoolean autoSlaveDiscovery,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.autoSlaveDiscovery = autoSlaveDiscovery;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataAutoSlaveDiscovery build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataAutoSlaveDiscovery bACnetConstructedDataAutoSlaveDiscovery =
          new BACnetConstructedDataAutoSlaveDiscovery(
              openingTag,
              peekedTagHeader,
              closingTag,
              autoSlaveDiscovery,
              tagNumber,
              arrayIndexArgument);
      return bACnetConstructedDataAutoSlaveDiscovery;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataAutoSlaveDiscovery)) {
      return false;
    }
    BACnetConstructedDataAutoSlaveDiscovery that = (BACnetConstructedDataAutoSlaveDiscovery) o;
    return (getAutoSlaveDiscovery() == that.getAutoSlaveDiscovery()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAutoSlaveDiscovery());
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
