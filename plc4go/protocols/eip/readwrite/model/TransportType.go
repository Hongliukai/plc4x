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

package model

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// TransportType is the corresponding interface of TransportType
type TransportType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDirection returns Direction (property field)
	GetDirection() bool
	// GetTrigger returns Trigger (property field)
	GetTrigger() uint8
	// GetClassTransport returns ClassTransport (property field)
	GetClassTransport() uint8
}

// TransportTypeExactly can be used when we want exactly this type and not a type which fulfills TransportType.
// This is useful for switch cases.
type TransportTypeExactly interface {
	TransportType
	isTransportType() bool
}

// _TransportType is the data-structure of this message
type _TransportType struct {
	Direction      bool
	Trigger        uint8
	ClassTransport uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransportType) GetDirection() bool {
	return m.Direction
}

func (m *_TransportType) GetTrigger() uint8 {
	return m.Trigger
}

func (m *_TransportType) GetClassTransport() uint8 {
	return m.ClassTransport
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTransportType factory function for _TransportType
func NewTransportType(direction bool, trigger uint8, classTransport uint8) *_TransportType {
	return &_TransportType{Direction: direction, Trigger: trigger, ClassTransport: classTransport}
}

// Deprecated: use the interface for direct cast
func CastTransportType(structType any) TransportType {
	if casted, ok := structType.(TransportType); ok {
		return casted
	}
	if casted, ok := structType.(*TransportType); ok {
		return *casted
	}
	return nil
}

func (m *_TransportType) GetTypeName() string {
	return "TransportType"
}

func (m *_TransportType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (direction)
	lengthInBits += 1

	// Simple field (trigger)
	lengthInBits += 3

	// Simple field (classTransport)
	lengthInBits += 4

	return lengthInBits
}

func (m *_TransportType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TransportTypeParse(ctx context.Context, theBytes []byte) (TransportType, error) {
	return TransportTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TransportTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TransportType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TransportType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransportType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (direction)
	_direction, _directionErr := readBuffer.ReadBit("direction")
	if _directionErr != nil {
		return nil, errors.Wrap(_directionErr, "Error parsing 'direction' field of TransportType")
	}
	direction := _direction

	// Simple Field (trigger)
	_trigger, _triggerErr := readBuffer.ReadUint8("trigger", 3)
	if _triggerErr != nil {
		return nil, errors.Wrap(_triggerErr, "Error parsing 'trigger' field of TransportType")
	}
	trigger := _trigger

	// Simple Field (classTransport)
	_classTransport, _classTransportErr := readBuffer.ReadUint8("classTransport", 4)
	if _classTransportErr != nil {
		return nil, errors.Wrap(_classTransportErr, "Error parsing 'classTransport' field of TransportType")
	}
	classTransport := _classTransport

	if closeErr := readBuffer.CloseContext("TransportType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransportType")
	}

	// Create the instance
	return &_TransportType{
		Direction:      direction,
		Trigger:        trigger,
		ClassTransport: classTransport,
	}, nil
}

func (m *_TransportType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransportType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TransportType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TransportType")
	}

	// Simple Field (direction)
	direction := bool(m.GetDirection())
	_directionErr := writeBuffer.WriteBit("direction", (direction))
	if _directionErr != nil {
		return errors.Wrap(_directionErr, "Error serializing 'direction' field")
	}

	// Simple Field (trigger)
	trigger := uint8(m.GetTrigger())
	_triggerErr := writeBuffer.WriteUint8("trigger", 3, (trigger))
	if _triggerErr != nil {
		return errors.Wrap(_triggerErr, "Error serializing 'trigger' field")
	}

	// Simple Field (classTransport)
	classTransport := uint8(m.GetClassTransport())
	_classTransportErr := writeBuffer.WriteUint8("classTransport", 4, (classTransport))
	if _classTransportErr != nil {
		return errors.Wrap(_classTransportErr, "Error serializing 'classTransport' field")
	}

	if popErr := writeBuffer.PopContext("TransportType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TransportType")
	}
	return nil
}

func (m *_TransportType) isTransportType() bool {
	return true
}

func (m *_TransportType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
