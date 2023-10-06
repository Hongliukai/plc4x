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

public class AirConditioningDataZoneTemperature extends AirConditioningData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte zoneGroup;
  protected final HVACZoneList zoneList;
  protected final HVACTemperature temperature;
  protected final HVACSensorStatus sensorStatus;

  public AirConditioningDataZoneTemperature(
      AirConditioningCommandTypeContainer commandTypeContainer,
      byte zoneGroup,
      HVACZoneList zoneList,
      HVACTemperature temperature,
      HVACSensorStatus sensorStatus) {
    super(commandTypeContainer);
    this.zoneGroup = zoneGroup;
    this.zoneList = zoneList;
    this.temperature = temperature;
    this.sensorStatus = sensorStatus;
  }

  public byte getZoneGroup() {
    return zoneGroup;
  }

  public HVACZoneList getZoneList() {
    return zoneList;
  }

  public HVACTemperature getTemperature() {
    return temperature;
  }

  public HVACSensorStatus getSensorStatus() {
    return sensorStatus;
  }

  @Override
  protected void serializeAirConditioningDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AirConditioningDataZoneTemperature");

    // Simple Field (zoneGroup)
    writeSimpleField("zoneGroup", zoneGroup, writeByte(writeBuffer, 8));

    // Simple Field (zoneList)
    writeSimpleField("zoneList", zoneList, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (temperature)
    writeSimpleField("temperature", temperature, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (sensorStatus)
    writeSimpleEnumField(
        "sensorStatus",
        "HVACSensorStatus",
        sensorStatus,
        new DataWriterEnumDefault<>(
            HVACSensorStatus::getValue,
            HVACSensorStatus::name,
            writeUnsignedShort(writeBuffer, 8)));

    writeBuffer.popContext("AirConditioningDataZoneTemperature");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AirConditioningDataZoneTemperature _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (zoneGroup)
    lengthInBits += 8;

    // Simple field (zoneList)
    lengthInBits += zoneList.getLengthInBits();

    // Simple field (temperature)
    lengthInBits += temperature.getLengthInBits();

    // Simple field (sensorStatus)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static AirConditioningDataBuilder staticParseAirConditioningDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("AirConditioningDataZoneTemperature");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte zoneGroup = readSimpleField("zoneGroup", readByte(readBuffer, 8));

    HVACZoneList zoneList =
        readSimpleField(
            "zoneList",
            new DataReaderComplexDefault<>(() -> HVACZoneList.staticParse(readBuffer), readBuffer));

    HVACTemperature temperature =
        readSimpleField(
            "temperature",
            new DataReaderComplexDefault<>(
                () -> HVACTemperature.staticParse(readBuffer), readBuffer));

    HVACSensorStatus sensorStatus =
        readEnumField(
            "sensorStatus",
            "HVACSensorStatus",
            new DataReaderEnumDefault<>(
                HVACSensorStatus::enumForValue, readUnsignedShort(readBuffer, 8)));

    readBuffer.closeContext("AirConditioningDataZoneTemperature");
    // Create the instance
    return new AirConditioningDataZoneTemperatureBuilderImpl(
        zoneGroup, zoneList, temperature, sensorStatus);
  }

  public static class AirConditioningDataZoneTemperatureBuilderImpl
      implements AirConditioningData.AirConditioningDataBuilder {
    private final byte zoneGroup;
    private final HVACZoneList zoneList;
    private final HVACTemperature temperature;
    private final HVACSensorStatus sensorStatus;

    public AirConditioningDataZoneTemperatureBuilderImpl(
        byte zoneGroup,
        HVACZoneList zoneList,
        HVACTemperature temperature,
        HVACSensorStatus sensorStatus) {
      this.zoneGroup = zoneGroup;
      this.zoneList = zoneList;
      this.temperature = temperature;
      this.sensorStatus = sensorStatus;
    }

    public AirConditioningDataZoneTemperature build(
        AirConditioningCommandTypeContainer commandTypeContainer) {
      AirConditioningDataZoneTemperature airConditioningDataZoneTemperature =
          new AirConditioningDataZoneTemperature(
              commandTypeContainer, zoneGroup, zoneList, temperature, sensorStatus);
      return airConditioningDataZoneTemperature;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AirConditioningDataZoneTemperature)) {
      return false;
    }
    AirConditioningDataZoneTemperature that = (AirConditioningDataZoneTemperature) o;
    return (getZoneGroup() == that.getZoneGroup())
        && (getZoneList() == that.getZoneList())
        && (getTemperature() == that.getTemperature())
        && (getSensorStatus() == that.getSensorStatus())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getZoneGroup(), getZoneList(), getTemperature(), getSensorStatus());
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
