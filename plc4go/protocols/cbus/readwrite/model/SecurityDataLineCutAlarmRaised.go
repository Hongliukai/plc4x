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

// SecurityDataLineCutAlarmRaised is the corresponding interface of SecurityDataLineCutAlarmRaised
type SecurityDataLineCutAlarmRaised interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataLineCutAlarmRaisedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataLineCutAlarmRaised.
// This is useful for switch cases.
type SecurityDataLineCutAlarmRaisedExactly interface {
	SecurityDataLineCutAlarmRaised
	isSecurityDataLineCutAlarmRaised() bool
}

// _SecurityDataLineCutAlarmRaised is the data-structure of this message
type _SecurityDataLineCutAlarmRaised struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataLineCutAlarmRaised) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataLineCutAlarmRaised) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataLineCutAlarmRaised factory function for _SecurityDataLineCutAlarmRaised
func NewSecurityDataLineCutAlarmRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataLineCutAlarmRaised {
	_result := &_SecurityDataLineCutAlarmRaised{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataLineCutAlarmRaised(structType interface{}) SecurityDataLineCutAlarmRaised {
	if casted, ok := structType.(SecurityDataLineCutAlarmRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataLineCutAlarmRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataLineCutAlarmRaised) GetTypeName() string {
	return "SecurityDataLineCutAlarmRaised"
}

func (m *_SecurityDataLineCutAlarmRaised) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataLineCutAlarmRaised) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SecurityDataLineCutAlarmRaised) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataLineCutAlarmRaisedParse(theBytes []byte) (SecurityDataLineCutAlarmRaised, error) {
	return SecurityDataLineCutAlarmRaisedParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataLineCutAlarmRaisedParseWithBuffer(readBuffer utils.ReadBuffer) (SecurityDataLineCutAlarmRaised, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataLineCutAlarmRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataLineCutAlarmRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataLineCutAlarmRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataLineCutAlarmRaised")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataLineCutAlarmRaised{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataLineCutAlarmRaised) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataLineCutAlarmRaised) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataLineCutAlarmRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataLineCutAlarmRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataLineCutAlarmRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataLineCutAlarmRaised")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataLineCutAlarmRaised) isSecurityDataLineCutAlarmRaised() bool {
	return true
}

func (m *_SecurityDataLineCutAlarmRaised) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
