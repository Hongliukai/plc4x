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

public class SecurityDataSystemArmedDisarmed extends SecurityData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final SecurityArmCode armCodeType;

  public SecurityDataSystemArmedDisarmed(
      SecurityCommandTypeContainer commandTypeContainer,
      byte argument,
      SecurityArmCode armCodeType) {
    super(commandTypeContainer, argument);
    this.armCodeType = armCodeType;
  }

  public SecurityArmCode getArmCodeType() {
    return armCodeType;
  }

  @Override
  protected void serializeSecurityDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SecurityDataSystemArmedDisarmed");

    // Simple Field (armCodeType)
    writeSimpleField("armCodeType", armCodeType, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("SecurityDataSystemArmedDisarmed");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SecurityDataSystemArmedDisarmed _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (armCodeType)
    lengthInBits += armCodeType.getLengthInBits();

    return lengthInBits;
  }

  public static SecurityDataBuilder staticParseSecurityDataBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("SecurityDataSystemArmedDisarmed");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    SecurityArmCode armCodeType =
        readSimpleField(
            "armCodeType",
            new DataReaderComplexDefault<>(
                () -> SecurityArmCode.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("SecurityDataSystemArmedDisarmed");
    // Create the instance
    return new SecurityDataSystemArmedDisarmedBuilderImpl(armCodeType);
  }

  public static class SecurityDataSystemArmedDisarmedBuilderImpl
      implements SecurityData.SecurityDataBuilder {
    private final SecurityArmCode armCodeType;

    public SecurityDataSystemArmedDisarmedBuilderImpl(SecurityArmCode armCodeType) {
      this.armCodeType = armCodeType;
    }

    public SecurityDataSystemArmedDisarmed build(
        SecurityCommandTypeContainer commandTypeContainer, byte argument) {
      SecurityDataSystemArmedDisarmed securityDataSystemArmedDisarmed =
          new SecurityDataSystemArmedDisarmed(commandTypeContainer, argument, armCodeType);
      return securityDataSystemArmedDisarmed;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SecurityDataSystemArmedDisarmed)) {
      return false;
    }
    SecurityDataSystemArmedDisarmed that = (SecurityDataSystemArmedDisarmed) o;
    return (getArmCodeType() == that.getArmCodeType()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getArmCodeType());
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
