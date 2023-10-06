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

// BACnetConstructedDataChangesPending is the corresponding interface of BACnetConstructedDataChangesPending
type BACnetConstructedDataChangesPending interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetChangesPending returns ChangesPending (property field)
	GetChangesPending() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataChangesPendingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataChangesPending.
// This is useful for switch cases.
type BACnetConstructedDataChangesPendingExactly interface {
	BACnetConstructedDataChangesPending
	isBACnetConstructedDataChangesPending() bool
}

// _BACnetConstructedDataChangesPending is the data-structure of this message
type _BACnetConstructedDataChangesPending struct {
	*_BACnetConstructedData
	ChangesPending BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChangesPending) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataChangesPending) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANGES_PENDING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChangesPending) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataChangesPending) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChangesPending) GetChangesPending() BACnetApplicationTagBoolean {
	return m.ChangesPending
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChangesPending) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetChangesPending())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataChangesPending factory function for _BACnetConstructedDataChangesPending
func NewBACnetConstructedDataChangesPending(changesPending BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChangesPending {
	_result := &_BACnetConstructedDataChangesPending{
		ChangesPending:         changesPending,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChangesPending(structType any) BACnetConstructedDataChangesPending {
	if casted, ok := structType.(BACnetConstructedDataChangesPending); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChangesPending); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChangesPending) GetTypeName() string {
	return "BACnetConstructedDataChangesPending"
}

func (m *_BACnetConstructedDataChangesPending) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (changesPending)
	lengthInBits += m.ChangesPending.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataChangesPending) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataChangesPendingParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataChangesPending, error) {
	return BACnetConstructedDataChangesPendingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataChangesPendingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataChangesPending, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChangesPending"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChangesPending")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (changesPending)
	if pullErr := readBuffer.PullContext("changesPending"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for changesPending")
	}
	_changesPending, _changesPendingErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _changesPendingErr != nil {
		return nil, errors.Wrap(_changesPendingErr, "Error parsing 'changesPending' field of BACnetConstructedDataChangesPending")
	}
	changesPending := _changesPending.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("changesPending"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for changesPending")
	}

	// Virtual field
	_actualValue := changesPending
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChangesPending"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChangesPending")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataChangesPending{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ChangesPending: changesPending,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataChangesPending) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChangesPending) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChangesPending"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChangesPending")
		}

		// Simple Field (changesPending)
		if pushErr := writeBuffer.PushContext("changesPending"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for changesPending")
		}
		_changesPendingErr := writeBuffer.WriteSerializable(ctx, m.GetChangesPending())
		if popErr := writeBuffer.PopContext("changesPending"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for changesPending")
		}
		if _changesPendingErr != nil {
			return errors.Wrap(_changesPendingErr, "Error serializing 'changesPending' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChangesPending"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChangesPending")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChangesPending) isBACnetConstructedDataChangesPending() bool {
	return true
}

func (m *_BACnetConstructedDataChangesPending) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
