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

public class AdditionalParametersType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "16315";
  }

  // Properties.
  protected final int noOfParameters;
  protected final List<ExtensionObjectDefinition> parameters;

  public AdditionalParametersType(int noOfParameters, List<ExtensionObjectDefinition> parameters) {
    super();
    this.noOfParameters = noOfParameters;
    this.parameters = parameters;
  }

  public int getNoOfParameters() {
    return noOfParameters;
  }

  public List<ExtensionObjectDefinition> getParameters() {
    return parameters;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AdditionalParametersType");

    // Simple Field (noOfParameters)
    writeSimpleField("noOfParameters", noOfParameters, writeSignedInt(writeBuffer, 32));

    // Array Field (parameters)
    writeComplexTypeArrayField("parameters", parameters, writeBuffer);

    writeBuffer.popContext("AdditionalParametersType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdditionalParametersType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (noOfParameters)
    lengthInBits += 32;

    // Array field
    if (parameters != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : parameters) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= parameters.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("AdditionalParametersType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfParameters = readSimpleField("noOfParameters", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> parameters =
        readCountArrayField(
            "parameters",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("14535")),
                readBuffer),
            noOfParameters);

    readBuffer.closeContext("AdditionalParametersType");
    // Create the instance
    return new AdditionalParametersTypeBuilderImpl(noOfParameters, parameters);
  }

  public static class AdditionalParametersTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final int noOfParameters;
    private final List<ExtensionObjectDefinition> parameters;

    public AdditionalParametersTypeBuilderImpl(
        int noOfParameters, List<ExtensionObjectDefinition> parameters) {
      this.noOfParameters = noOfParameters;
      this.parameters = parameters;
    }

    public AdditionalParametersType build() {
      AdditionalParametersType additionalParametersType =
          new AdditionalParametersType(noOfParameters, parameters);
      return additionalParametersType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdditionalParametersType)) {
      return false;
    }
    AdditionalParametersType that = (AdditionalParametersType) o;
    return (getNoOfParameters() == that.getNoOfParameters())
        && (getParameters() == that.getParameters())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNoOfParameters(), getParameters());
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
