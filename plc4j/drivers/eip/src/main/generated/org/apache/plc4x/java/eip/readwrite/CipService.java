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

public abstract class CipService implements Message {

  // Abstract accessors for discriminator values.
  public abstract Short getService();

  public CipService() {
    super();
  }

  protected abstract void serializeCipServiceChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CipService");

    // Discriminator Field (service) (Used as input to a switch field)
    writeDiscriminatorField("service", getService(), writeUnsignedShort(writeBuffer, 8));

    // Switch field (Serialize the sub-type)
    serializeCipServiceChild(writeBuffer);

    writeBuffer.popContext("CipService");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CipService _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (service)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static CipService staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 1)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 1, but got " + args.length);
    }
    Integer serviceLen;
    if (args[0] instanceof Integer) {
      serviceLen = (Integer) args[0];
    } else if (args[0] instanceof String) {
      serviceLen = Integer.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Integer or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    return staticParse(readBuffer, serviceLen);
  }

  public static CipService staticParse(ReadBuffer readBuffer, Integer serviceLen)
      throws ParseException {
    readBuffer.pullContext("CipService");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short service = readDiscriminatorField("service", readUnsignedShort(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    CipServiceBuilder builder = null;
    if (EvaluationHelper.equals(service, (short) 0x4C)) {
      builder = CipReadRequest.staticParseCipServiceBuilder(readBuffer, serviceLen);
    } else if (EvaluationHelper.equals(service, (short) 0xCC)) {
      builder = CipReadResponse.staticParseCipServiceBuilder(readBuffer, serviceLen);
    } else if (EvaluationHelper.equals(service, (short) 0x4D)) {
      builder = CipWriteRequest.staticParseCipServiceBuilder(readBuffer, serviceLen);
    } else if (EvaluationHelper.equals(service, (short) 0xCD)) {
      builder = CipWriteResponse.staticParseCipServiceBuilder(readBuffer, serviceLen);
    } else if (EvaluationHelper.equals(service, (short) 0x0A)) {
      builder = MultipleServiceRequest.staticParseCipServiceBuilder(readBuffer, serviceLen);
    } else if (EvaluationHelper.equals(service, (short) 0x8A)) {
      builder = MultipleServiceResponse.staticParseCipServiceBuilder(readBuffer, serviceLen);
    } else if (EvaluationHelper.equals(service, (short) 0x52)) {
      builder = CipUnconnectedRequest.staticParseCipServiceBuilder(readBuffer, serviceLen);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type" + " parameters [" + "service=" + service + "]");
    }

    readBuffer.closeContext("CipService");
    // Create the instance
    CipService _cipService = builder.build();
    return _cipService;
  }

  public interface CipServiceBuilder {
    CipService build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipService)) {
      return false;
    }
    CipService that = (CipService) o;
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
