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

public class LightingDataLabel extends LightingData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte group;
  protected final LightingLabelOptions labelOptions;
  protected final Language language;
  protected final byte[] data;

  public LightingDataLabel(
      LightingCommandTypeContainer commandTypeContainer,
      byte group,
      LightingLabelOptions labelOptions,
      Language language,
      byte[] data) {
    super(commandTypeContainer);
    this.group = group;
    this.labelOptions = labelOptions;
    this.language = language;
    this.data = data;
  }

  public byte getGroup() {
    return group;
  }

  public LightingLabelOptions getLabelOptions() {
    return labelOptions;
  }

  public Language getLanguage() {
    return language;
  }

  public byte[] getData() {
    return data;
  }

  @Override
  protected void serializeLightingDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("LightingDataLabel");

    // Simple Field (group)
    writeSimpleField("group", group, writeByte(writeBuffer, 8));

    // Simple Field (labelOptions)
    writeSimpleField("labelOptions", labelOptions, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (language) (Can be skipped, if the value is null)
    writeOptionalEnumField(
        "language",
        "Language",
        language,
        new DataWriterEnumDefault<>(
            Language::getValue, Language::name, writeUnsignedShort(writeBuffer, 8)),
        (getLabelOptions().getLabelType()) != (LightingLabelType.LOAD_DYNAMIC_ICON));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("LightingDataLabel");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    LightingDataLabel _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (group)
    lengthInBits += 8;

    // Simple field (labelOptions)
    lengthInBits += labelOptions.getLengthInBits();

    // Optional Field (language)
    if (language != null) {
      lengthInBits += 8;
    }

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static LightingDataBuilder staticParseLightingDataBuilder(
      ReadBuffer readBuffer, LightingCommandTypeContainer commandTypeContainer)
      throws ParseException {
    readBuffer.pullContext("LightingDataLabel");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte group = readSimpleField("group", readByte(readBuffer, 8));

    LightingLabelOptions labelOptions =
        readSimpleField(
            "labelOptions",
            new DataReaderComplexDefault<>(
                () -> LightingLabelOptions.staticParse(readBuffer), readBuffer));

    Language language =
        readOptionalField(
            "language",
            new DataReaderEnumDefault<>(Language::enumForValue, readUnsignedShort(readBuffer, 8)),
            (labelOptions.getLabelType()) != (LightingLabelType.LOAD_DYNAMIC_ICON));

    byte[] data =
        readBuffer.readByteArray(
            "data",
            Math.toIntExact(
                ((commandTypeContainer.getNumBytes())
                    - ((((((labelOptions.getLabelType()) != (LightingLabelType.LOAD_DYNAMIC_ICON)))
                        ? (3)
                        : (2)))))));

    readBuffer.closeContext("LightingDataLabel");
    // Create the instance
    return new LightingDataLabelBuilderImpl(group, labelOptions, language, data);
  }

  public static class LightingDataLabelBuilderImpl implements LightingData.LightingDataBuilder {
    private final byte group;
    private final LightingLabelOptions labelOptions;
    private final Language language;
    private final byte[] data;

    public LightingDataLabelBuilderImpl(
        byte group, LightingLabelOptions labelOptions, Language language, byte[] data) {
      this.group = group;
      this.labelOptions = labelOptions;
      this.language = language;
      this.data = data;
    }

    public LightingDataLabel build(LightingCommandTypeContainer commandTypeContainer) {
      LightingDataLabel lightingDataLabel =
          new LightingDataLabel(commandTypeContainer, group, labelOptions, language, data);
      return lightingDataLabel;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LightingDataLabel)) {
      return false;
    }
    LightingDataLabel that = (LightingDataLabel) o;
    return (getGroup() == that.getGroup())
        && (getLabelOptions() == that.getLabelOptions())
        && (getLanguage() == that.getLanguage())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getGroup(), getLabelOptions(), getLanguage(), getData());
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
