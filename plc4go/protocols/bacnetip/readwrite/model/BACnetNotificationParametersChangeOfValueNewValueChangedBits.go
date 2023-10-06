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

// BACnetNotificationParametersChangeOfValueNewValueChangedBits is the corresponding interface of BACnetNotificationParametersChangeOfValueNewValueChangedBits
type BACnetNotificationParametersChangeOfValueNewValueChangedBits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfValueNewValue
	// GetChangedBits returns ChangedBits (property field)
	GetChangedBits() BACnetContextTagBitString
}

// BACnetNotificationParametersChangeOfValueNewValueChangedBitsExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfValueNewValueChangedBits.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfValueNewValueChangedBitsExactly interface {
	BACnetNotificationParametersChangeOfValueNewValueChangedBits
	isBACnetNotificationParametersChangeOfValueNewValueChangedBits() bool
}

// _BACnetNotificationParametersChangeOfValueNewValueChangedBits is the data-structure of this message
type _BACnetNotificationParametersChangeOfValueNewValueChangedBits struct {
	*_BACnetNotificationParametersChangeOfValueNewValue
	ChangedBits BACnetContextTagBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) InitializeParent(parent BACnetNotificationParametersChangeOfValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetParent() BACnetNotificationParametersChangeOfValueNewValue {
	return m._BACnetNotificationParametersChangeOfValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetChangedBits() BACnetContextTagBitString {
	return m.ChangedBits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfValueNewValueChangedBits factory function for _BACnetNotificationParametersChangeOfValueNewValueChangedBits
func NewBACnetNotificationParametersChangeOfValueNewValueChangedBits(changedBits BACnetContextTagBitString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfValueNewValueChangedBits {
	_result := &_BACnetNotificationParametersChangeOfValueNewValueChangedBits{
		ChangedBits: changedBits,
		_BACnetNotificationParametersChangeOfValueNewValue: NewBACnetNotificationParametersChangeOfValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfValueNewValue._BACnetNotificationParametersChangeOfValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfValueNewValueChangedBits(structType any) BACnetNotificationParametersChangeOfValueNewValueChangedBits {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValueChangedBits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValueChangedBits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValueChangedBits"
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (changedBits)
	lengthInBits += m.ChangedBits.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfValueNewValueChangedBitsParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8, tagNumber uint8) (BACnetNotificationParametersChangeOfValueNewValueChangedBits, error) {
	return BACnetNotificationParametersChangeOfValueNewValueChangedBitsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber, tagNumber)
}

func BACnetNotificationParametersChangeOfValueNewValueChangedBitsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8, tagNumber uint8) (BACnetNotificationParametersChangeOfValueNewValueChangedBits, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfValueNewValueChangedBits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (changedBits)
	if pullErr := readBuffer.PullContext("changedBits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for changedBits")
	}
	_changedBits, _changedBitsErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BIT_STRING))
	if _changedBitsErr != nil {
		return nil, errors.Wrap(_changedBitsErr, "Error parsing 'changedBits' field of BACnetNotificationParametersChangeOfValueNewValueChangedBits")
	}
	changedBits := _changedBits.(BACnetContextTagBitString)
	if closeErr := readBuffer.CloseContext("changedBits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for changedBits")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfValueNewValueChangedBits")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfValueNewValueChangedBits{
		_BACnetNotificationParametersChangeOfValueNewValue: &_BACnetNotificationParametersChangeOfValueNewValue{
			TagNumber: tagNumber,
		},
		ChangedBits: changedBits,
	}
	_child._BACnetNotificationParametersChangeOfValueNewValue._BACnetNotificationParametersChangeOfValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfValueNewValueChangedBits")
		}

		// Simple Field (changedBits)
		if pushErr := writeBuffer.PushContext("changedBits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for changedBits")
		}
		_changedBitsErr := writeBuffer.WriteSerializable(ctx, m.GetChangedBits())
		if popErr := writeBuffer.PopContext("changedBits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for changedBits")
		}
		if _changedBitsErr != nil {
			return errors.Wrap(_changedBitsErr, "Error serializing 'changedBits' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfValueNewValueChangedBits")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) isBACnetNotificationParametersChangeOfValueNewValueChangedBits() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
