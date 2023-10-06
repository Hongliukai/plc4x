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

public class BACnetConfirmedServiceRequestConfirmedEventNotification
    extends BACnetConfirmedServiceRequest implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.CONFIRMED_EVENT_NOTIFICATION;
  }

  // Properties.
  protected final BACnetContextTagUnsignedInteger processIdentifier;
  protected final BACnetContextTagObjectIdentifier initiatingDeviceIdentifier;
  protected final BACnetContextTagObjectIdentifier eventObjectIdentifier;
  protected final BACnetTimeStampEnclosed timestamp;
  protected final BACnetContextTagUnsignedInteger notificationClass;
  protected final BACnetContextTagUnsignedInteger priority;
  protected final BACnetEventTypeTagged eventType;
  protected final BACnetContextTagCharacterString messageText;
  protected final BACnetNotifyTypeTagged notifyType;
  protected final BACnetContextTagBoolean ackRequired;
  protected final BACnetEventStateTagged fromState;
  protected final BACnetEventStateTagged toState;
  protected final BACnetNotificationParameters eventValues;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestConfirmedEventNotification(
      BACnetContextTagUnsignedInteger processIdentifier,
      BACnetContextTagObjectIdentifier initiatingDeviceIdentifier,
      BACnetContextTagObjectIdentifier eventObjectIdentifier,
      BACnetTimeStampEnclosed timestamp,
      BACnetContextTagUnsignedInteger notificationClass,
      BACnetContextTagUnsignedInteger priority,
      BACnetEventTypeTagged eventType,
      BACnetContextTagCharacterString messageText,
      BACnetNotifyTypeTagged notifyType,
      BACnetContextTagBoolean ackRequired,
      BACnetEventStateTagged fromState,
      BACnetEventStateTagged toState,
      BACnetNotificationParameters eventValues,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.processIdentifier = processIdentifier;
    this.initiatingDeviceIdentifier = initiatingDeviceIdentifier;
    this.eventObjectIdentifier = eventObjectIdentifier;
    this.timestamp = timestamp;
    this.notificationClass = notificationClass;
    this.priority = priority;
    this.eventType = eventType;
    this.messageText = messageText;
    this.notifyType = notifyType;
    this.ackRequired = ackRequired;
    this.fromState = fromState;
    this.toState = toState;
    this.eventValues = eventValues;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetContextTagUnsignedInteger getProcessIdentifier() {
    return processIdentifier;
  }

  public BACnetContextTagObjectIdentifier getInitiatingDeviceIdentifier() {
    return initiatingDeviceIdentifier;
  }

  public BACnetContextTagObjectIdentifier getEventObjectIdentifier() {
    return eventObjectIdentifier;
  }

  public BACnetTimeStampEnclosed getTimestamp() {
    return timestamp;
  }

  public BACnetContextTagUnsignedInteger getNotificationClass() {
    return notificationClass;
  }

  public BACnetContextTagUnsignedInteger getPriority() {
    return priority;
  }

  public BACnetEventTypeTagged getEventType() {
    return eventType;
  }

  public BACnetContextTagCharacterString getMessageText() {
    return messageText;
  }

  public BACnetNotifyTypeTagged getNotifyType() {
    return notifyType;
  }

  public BACnetContextTagBoolean getAckRequired() {
    return ackRequired;
  }

  public BACnetEventStateTagged getFromState() {
    return fromState;
  }

  public BACnetEventStateTagged getToState() {
    return toState;
  }

  public BACnetNotificationParameters getEventValues() {
    return eventValues;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestConfirmedEventNotification");

    // Simple Field (processIdentifier)
    writeSimpleField(
        "processIdentifier", processIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (initiatingDeviceIdentifier)
    writeSimpleField(
        "initiatingDeviceIdentifier",
        initiatingDeviceIdentifier,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (eventObjectIdentifier)
    writeSimpleField(
        "eventObjectIdentifier",
        eventObjectIdentifier,
        new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timestamp)
    writeSimpleField("timestamp", timestamp, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (notificationClass)
    writeSimpleField(
        "notificationClass", notificationClass, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (priority)
    writeSimpleField("priority", priority, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (eventType)
    writeSimpleField("eventType", eventType, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (messageText) (Can be skipped, if the value is null)
    writeOptionalField("messageText", messageText, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (notifyType)
    writeSimpleField("notifyType", notifyType, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (ackRequired) (Can be skipped, if the value is null)
    writeOptionalField("ackRequired", ackRequired, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (fromState) (Can be skipped, if the value is null)
    writeOptionalField("fromState", fromState, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (toState)
    writeSimpleField("toState", toState, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (eventValues) (Can be skipped, if the value is null)
    writeOptionalField("eventValues", eventValues, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestConfirmedEventNotification");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestConfirmedEventNotification _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (processIdentifier)
    lengthInBits += processIdentifier.getLengthInBits();

    // Simple field (initiatingDeviceIdentifier)
    lengthInBits += initiatingDeviceIdentifier.getLengthInBits();

    // Simple field (eventObjectIdentifier)
    lengthInBits += eventObjectIdentifier.getLengthInBits();

    // Simple field (timestamp)
    lengthInBits += timestamp.getLengthInBits();

    // Simple field (notificationClass)
    lengthInBits += notificationClass.getLengthInBits();

    // Simple field (priority)
    lengthInBits += priority.getLengthInBits();

    // Simple field (eventType)
    lengthInBits += eventType.getLengthInBits();

    // Optional Field (messageText)
    if (messageText != null) {
      lengthInBits += messageText.getLengthInBits();
    }

    // Simple field (notifyType)
    lengthInBits += notifyType.getLengthInBits();

    // Optional Field (ackRequired)
    if (ackRequired != null) {
      lengthInBits += ackRequired.getLengthInBits();
    }

    // Optional Field (fromState)
    if (fromState != null) {
      lengthInBits += fromState.getLengthInBits();
    }

    // Simple field (toState)
    lengthInBits += toState.getLengthInBits();

    // Optional Field (eventValues)
    if (eventValues != null) {
      lengthInBits += eventValues.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestConfirmedEventNotification");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagUnsignedInteger processIdentifier =
        readSimpleField(
            "processIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagObjectIdentifier initiatingDeviceIdentifier =
        readSimpleField(
            "initiatingDeviceIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetContextTagObjectIdentifier eventObjectIdentifier =
        readSimpleField(
            "eventObjectIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetTimeStampEnclosed timestamp =
        readSimpleField(
            "timestamp",
            new DataReaderComplexDefault<>(
                () -> BACnetTimeStampEnclosed.staticParse(readBuffer, (short) (3)), readBuffer));

    BACnetContextTagUnsignedInteger notificationClass =
        readSimpleField(
            "notificationClass",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (4),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagUnsignedInteger priority =
        readSimpleField(
            "priority",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (5),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetEventTypeTagged eventType =
        readSimpleField(
            "eventType",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetEventTypeTagged.staticParse(
                        readBuffer, (short) (6), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagCharacterString messageText =
        readOptionalField(
            "messageText",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagCharacterString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (7),
                            (BACnetDataType) (BACnetDataType.CHARACTER_STRING)),
                readBuffer));

    BACnetNotifyTypeTagged notifyType =
        readSimpleField(
            "notifyType",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetNotifyTypeTagged.staticParse(
                        readBuffer, (short) (8), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagBoolean ackRequired =
        readOptionalField(
            "ackRequired",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagBoolean)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (9), (BACnetDataType) (BACnetDataType.BOOLEAN)),
                readBuffer));

    BACnetEventStateTagged fromState =
        readOptionalField(
            "fromState",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetEventStateTagged.staticParse(
                        readBuffer, (short) (10), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetEventStateTagged toState =
        readSimpleField(
            "toState",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetEventStateTagged.staticParse(
                        readBuffer, (short) (11), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetNotificationParameters eventValues =
        readOptionalField(
            "eventValues",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetNotificationParameters.staticParse(
                        readBuffer,
                        (short) (12),
                        (BACnetObjectType) (eventObjectIdentifier.getObjectType())),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestConfirmedEventNotification");
    // Create the instance
    return new BACnetConfirmedServiceRequestConfirmedEventNotificationBuilderImpl(
        processIdentifier,
        initiatingDeviceIdentifier,
        eventObjectIdentifier,
        timestamp,
        notificationClass,
        priority,
        eventType,
        messageText,
        notifyType,
        ackRequired,
        fromState,
        toState,
        eventValues,
        serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestConfirmedEventNotificationBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetContextTagUnsignedInteger processIdentifier;
    private final BACnetContextTagObjectIdentifier initiatingDeviceIdentifier;
    private final BACnetContextTagObjectIdentifier eventObjectIdentifier;
    private final BACnetTimeStampEnclosed timestamp;
    private final BACnetContextTagUnsignedInteger notificationClass;
    private final BACnetContextTagUnsignedInteger priority;
    private final BACnetEventTypeTagged eventType;
    private final BACnetContextTagCharacterString messageText;
    private final BACnetNotifyTypeTagged notifyType;
    private final BACnetContextTagBoolean ackRequired;
    private final BACnetEventStateTagged fromState;
    private final BACnetEventStateTagged toState;
    private final BACnetNotificationParameters eventValues;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestConfirmedEventNotificationBuilderImpl(
        BACnetContextTagUnsignedInteger processIdentifier,
        BACnetContextTagObjectIdentifier initiatingDeviceIdentifier,
        BACnetContextTagObjectIdentifier eventObjectIdentifier,
        BACnetTimeStampEnclosed timestamp,
        BACnetContextTagUnsignedInteger notificationClass,
        BACnetContextTagUnsignedInteger priority,
        BACnetEventTypeTagged eventType,
        BACnetContextTagCharacterString messageText,
        BACnetNotifyTypeTagged notifyType,
        BACnetContextTagBoolean ackRequired,
        BACnetEventStateTagged fromState,
        BACnetEventStateTagged toState,
        BACnetNotificationParameters eventValues,
        Long serviceRequestLength) {
      this.processIdentifier = processIdentifier;
      this.initiatingDeviceIdentifier = initiatingDeviceIdentifier;
      this.eventObjectIdentifier = eventObjectIdentifier;
      this.timestamp = timestamp;
      this.notificationClass = notificationClass;
      this.priority = priority;
      this.eventType = eventType;
      this.messageText = messageText;
      this.notifyType = notifyType;
      this.ackRequired = ackRequired;
      this.fromState = fromState;
      this.toState = toState;
      this.eventValues = eventValues;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestConfirmedEventNotification build(
        Long serviceRequestLength) {
      BACnetConfirmedServiceRequestConfirmedEventNotification
          bACnetConfirmedServiceRequestConfirmedEventNotification =
              new BACnetConfirmedServiceRequestConfirmedEventNotification(
                  processIdentifier,
                  initiatingDeviceIdentifier,
                  eventObjectIdentifier,
                  timestamp,
                  notificationClass,
                  priority,
                  eventType,
                  messageText,
                  notifyType,
                  ackRequired,
                  fromState,
                  toState,
                  eventValues,
                  serviceRequestLength);
      return bACnetConfirmedServiceRequestConfirmedEventNotification;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestConfirmedEventNotification)) {
      return false;
    }
    BACnetConfirmedServiceRequestConfirmedEventNotification that =
        (BACnetConfirmedServiceRequestConfirmedEventNotification) o;
    return (getProcessIdentifier() == that.getProcessIdentifier())
        && (getInitiatingDeviceIdentifier() == that.getInitiatingDeviceIdentifier())
        && (getEventObjectIdentifier() == that.getEventObjectIdentifier())
        && (getTimestamp() == that.getTimestamp())
        && (getNotificationClass() == that.getNotificationClass())
        && (getPriority() == that.getPriority())
        && (getEventType() == that.getEventType())
        && (getMessageText() == that.getMessageText())
        && (getNotifyType() == that.getNotifyType())
        && (getAckRequired() == that.getAckRequired())
        && (getFromState() == that.getFromState())
        && (getToState() == that.getToState())
        && (getEventValues() == that.getEventValues())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getProcessIdentifier(),
        getInitiatingDeviceIdentifier(),
        getEventObjectIdentifier(),
        getTimestamp(),
        getNotificationClass(),
        getPriority(),
        getEventType(),
        getMessageText(),
        getNotifyType(),
        getAckRequired(),
        getFromState(),
        getToState(),
        getEventValues());
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
