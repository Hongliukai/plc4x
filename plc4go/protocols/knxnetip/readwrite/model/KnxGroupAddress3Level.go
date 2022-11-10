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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// KnxGroupAddress3Level is the corresponding interface of KnxGroupAddress3Level
type KnxGroupAddress3Level interface {
	utils.LengthAware
	utils.Serializable
	KnxGroupAddress
	// GetMainGroup returns MainGroup (property field)
	GetMainGroup() uint8
	// GetMiddleGroup returns MiddleGroup (property field)
	GetMiddleGroup() uint8
	// GetSubGroup returns SubGroup (property field)
	GetSubGroup() uint8
}

// KnxGroupAddress3LevelExactly can be used when we want exactly this type and not a type which fulfills KnxGroupAddress3Level.
// This is useful for switch cases.
type KnxGroupAddress3LevelExactly interface {
	KnxGroupAddress3Level
	isKnxGroupAddress3Level() bool
}

// _KnxGroupAddress3Level is the data-structure of this message
type _KnxGroupAddress3Level struct {
	*_KnxGroupAddress
	MainGroup   uint8
	MiddleGroup uint8
	SubGroup    uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxGroupAddress3Level) GetNumLevels() uint8 {
	return uint8(3)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxGroupAddress3Level) InitializeParent(parent KnxGroupAddress) {}

func (m *_KnxGroupAddress3Level) GetParent() KnxGroupAddress {
	return m._KnxGroupAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxGroupAddress3Level) GetMainGroup() uint8 {
	return m.MainGroup
}

func (m *_KnxGroupAddress3Level) GetMiddleGroup() uint8 {
	return m.MiddleGroup
}

func (m *_KnxGroupAddress3Level) GetSubGroup() uint8 {
	return m.SubGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxGroupAddress3Level factory function for _KnxGroupAddress3Level
func NewKnxGroupAddress3Level(mainGroup uint8, middleGroup uint8, subGroup uint8) *_KnxGroupAddress3Level {
	_result := &_KnxGroupAddress3Level{
		MainGroup:        mainGroup,
		MiddleGroup:      middleGroup,
		SubGroup:         subGroup,
		_KnxGroupAddress: NewKnxGroupAddress(),
	}
	_result._KnxGroupAddress._KnxGroupAddressChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxGroupAddress3Level(structType interface{}) KnxGroupAddress3Level {
	if casted, ok := structType.(KnxGroupAddress3Level); ok {
		return casted
	}
	if casted, ok := structType.(*KnxGroupAddress3Level); ok {
		return *casted
	}
	return nil
}

func (m *_KnxGroupAddress3Level) GetTypeName() string {
	return "KnxGroupAddress3Level"
}

func (m *_KnxGroupAddress3Level) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxGroupAddress3Level) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (mainGroup)
	lengthInBits += 5

	// Simple field (middleGroup)
	lengthInBits += 3

	// Simple field (subGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxGroupAddress3Level) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxGroupAddress3LevelParse(theBytes []byte, numLevels uint8) (KnxGroupAddress3Level, error) {
	return KnxGroupAddress3LevelParseWithBuffer(utils.NewReadBufferByteBased(theBytes), numLevels)
}

func KnxGroupAddress3LevelParseWithBuffer(readBuffer utils.ReadBuffer, numLevels uint8) (KnxGroupAddress3Level, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxGroupAddress3Level"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxGroupAddress3Level")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (mainGroup)
	_mainGroup, _mainGroupErr := readBuffer.ReadUint8("mainGroup", 5)
	if _mainGroupErr != nil {
		return nil, errors.Wrap(_mainGroupErr, "Error parsing 'mainGroup' field of KnxGroupAddress3Level")
	}
	mainGroup := _mainGroup

	// Simple Field (middleGroup)
	_middleGroup, _middleGroupErr := readBuffer.ReadUint8("middleGroup", 3)
	if _middleGroupErr != nil {
		return nil, errors.Wrap(_middleGroupErr, "Error parsing 'middleGroup' field of KnxGroupAddress3Level")
	}
	middleGroup := _middleGroup

	// Simple Field (subGroup)
	_subGroup, _subGroupErr := readBuffer.ReadUint8("subGroup", 8)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field of KnxGroupAddress3Level")
	}
	subGroup := _subGroup

	if closeErr := readBuffer.CloseContext("KnxGroupAddress3Level"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxGroupAddress3Level")
	}

	// Create a partially initialized instance
	_child := &_KnxGroupAddress3Level{
		_KnxGroupAddress: &_KnxGroupAddress{},
		MainGroup:        mainGroup,
		MiddleGroup:      middleGroup,
		SubGroup:         subGroup,
	}
	_child._KnxGroupAddress._KnxGroupAddressChildRequirements = _child
	return _child, nil
}

func (m *_KnxGroupAddress3Level) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxGroupAddress3Level) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxGroupAddress3Level"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxGroupAddress3Level")
		}

		// Simple Field (mainGroup)
		mainGroup := uint8(m.GetMainGroup())
		_mainGroupErr := writeBuffer.WriteUint8("mainGroup", 5, (mainGroup))
		if _mainGroupErr != nil {
			return errors.Wrap(_mainGroupErr, "Error serializing 'mainGroup' field")
		}

		// Simple Field (middleGroup)
		middleGroup := uint8(m.GetMiddleGroup())
		_middleGroupErr := writeBuffer.WriteUint8("middleGroup", 3, (middleGroup))
		if _middleGroupErr != nil {
			return errors.Wrap(_middleGroupErr, "Error serializing 'middleGroup' field")
		}

		// Simple Field (subGroup)
		subGroup := uint8(m.GetSubGroup())
		_subGroupErr := writeBuffer.WriteUint8("subGroup", 8, (subGroup))
		if _subGroupErr != nil {
			return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
		}

		if popErr := writeBuffer.PopContext("KnxGroupAddress3Level"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxGroupAddress3Level")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_KnxGroupAddress3Level) isKnxGroupAddress3Level() bool {
	return true
}

func (m *_KnxGroupAddress3Level) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
