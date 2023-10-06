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
package org.apache.plc4x.java.opcua.readwrite;

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

public class SamplingIntervalDiagnosticsDataType extends ExtensionObjectDefinition
    implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "858";
  }

  // Properties.
  protected final double samplingInterval;
  protected final long monitoredItemCount;
  protected final long maxMonitoredItemCount;
  protected final long disabledMonitoredItemCount;

  public SamplingIntervalDiagnosticsDataType(
      double samplingInterval,
      long monitoredItemCount,
      long maxMonitoredItemCount,
      long disabledMonitoredItemCount) {
    super();
    this.samplingInterval = samplingInterval;
    this.monitoredItemCount = monitoredItemCount;
    this.maxMonitoredItemCount = maxMonitoredItemCount;
    this.disabledMonitoredItemCount = disabledMonitoredItemCount;
  }

  public double getSamplingInterval() {
    return samplingInterval;
  }

  public long getMonitoredItemCount() {
    return monitoredItemCount;
  }

  public long getMaxMonitoredItemCount() {
    return maxMonitoredItemCount;
  }

  public long getDisabledMonitoredItemCount() {
    return disabledMonitoredItemCount;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SamplingIntervalDiagnosticsDataType");

    // Simple Field (samplingInterval)
    writeSimpleField("samplingInterval", samplingInterval, writeDouble(writeBuffer, 64));

    // Simple Field (monitoredItemCount)
    writeSimpleField("monitoredItemCount", monitoredItemCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxMonitoredItemCount)
    writeSimpleField(
        "maxMonitoredItemCount", maxMonitoredItemCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (disabledMonitoredItemCount)
    writeSimpleField(
        "disabledMonitoredItemCount",
        disabledMonitoredItemCount,
        writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("SamplingIntervalDiagnosticsDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SamplingIntervalDiagnosticsDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (samplingInterval)
    lengthInBits += 64;

    // Simple field (monitoredItemCount)
    lengthInBits += 32;

    // Simple field (maxMonitoredItemCount)
    lengthInBits += 32;

    // Simple field (disabledMonitoredItemCount)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("SamplingIntervalDiagnosticsDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    double samplingInterval = readSimpleField("samplingInterval", readDouble(readBuffer, 64));

    long monitoredItemCount =
        readSimpleField("monitoredItemCount", readUnsignedLong(readBuffer, 32));

    long maxMonitoredItemCount =
        readSimpleField("maxMonitoredItemCount", readUnsignedLong(readBuffer, 32));

    long disabledMonitoredItemCount =
        readSimpleField("disabledMonitoredItemCount", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("SamplingIntervalDiagnosticsDataType");
    // Create the instance
    return new SamplingIntervalDiagnosticsDataTypeBuilderImpl(
        samplingInterval, monitoredItemCount, maxMonitoredItemCount, disabledMonitoredItemCount);
  }

  public static class SamplingIntervalDiagnosticsDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final double samplingInterval;
    private final long monitoredItemCount;
    private final long maxMonitoredItemCount;
    private final long disabledMonitoredItemCount;

    public SamplingIntervalDiagnosticsDataTypeBuilderImpl(
        double samplingInterval,
        long monitoredItemCount,
        long maxMonitoredItemCount,
        long disabledMonitoredItemCount) {
      this.samplingInterval = samplingInterval;
      this.monitoredItemCount = monitoredItemCount;
      this.maxMonitoredItemCount = maxMonitoredItemCount;
      this.disabledMonitoredItemCount = disabledMonitoredItemCount;
    }

    public SamplingIntervalDiagnosticsDataType build() {
      SamplingIntervalDiagnosticsDataType samplingIntervalDiagnosticsDataType =
          new SamplingIntervalDiagnosticsDataType(
              samplingInterval,
              monitoredItemCount,
              maxMonitoredItemCount,
              disabledMonitoredItemCount);
      return samplingIntervalDiagnosticsDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SamplingIntervalDiagnosticsDataType)) {
      return false;
    }
    SamplingIntervalDiagnosticsDataType that = (SamplingIntervalDiagnosticsDataType) o;
    return (getSamplingInterval() == that.getSamplingInterval())
        && (getMonitoredItemCount() == that.getMonitoredItemCount())
        && (getMaxMonitoredItemCount() == that.getMaxMonitoredItemCount())
        && (getDisabledMonitoredItemCount() == that.getDisabledMonitoredItemCount())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSamplingInterval(),
        getMonitoredItemCount(),
        getMaxMonitoredItemCount(),
        getDisabledMonitoredItemCount());
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
