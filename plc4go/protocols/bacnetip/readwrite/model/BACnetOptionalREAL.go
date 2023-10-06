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

// BACnetOptionalREAL is the corresponding interface of BACnetOptionalREAL
type BACnetOptionalREAL interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetOptionalREALExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalREAL.
// This is useful for switch cases.
type BACnetOptionalREALExactly interface {
	BACnetOptionalREAL
	isBACnetOptionalREAL() bool
}

// _BACnetOptionalREAL is the data-structure of this message
type _BACnetOptionalREAL struct {
	_BACnetOptionalREALChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetOptionalREALChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetOptionalREALParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetOptionalREAL, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetOptionalREALChild interface {
	utils.Serializable
	InitializeParent(parent BACnetOptionalREAL, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetOptionalREAL

	GetTypeName() string
	BACnetOptionalREAL
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalREAL) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetOptionalREAL) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalREAL factory function for _BACnetOptionalREAL
func NewBACnetOptionalREAL(peekedTagHeader BACnetTagHeader) *_BACnetOptionalREAL {
	return &_BACnetOptionalREAL{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalREAL(structType any) BACnetOptionalREAL {
	if casted, ok := structType.(BACnetOptionalREAL); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalREAL); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalREAL) GetTypeName() string {
	return "BACnetOptionalREAL"
}

func (m *_BACnetOptionalREAL) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetOptionalREAL) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetOptionalREALParse(ctx context.Context, theBytes []byte) (BACnetOptionalREAL, error) {
	return BACnetOptionalREALParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetOptionalREALParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetOptionalREAL, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetOptionalREAL"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalREAL")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetOptionalREALChildSerializeRequirement interface {
		BACnetOptionalREAL
		InitializeParent(BACnetOptionalREAL, BACnetTagHeader)
		GetParent() BACnetOptionalREAL
	}
	var _childTemp any
	var _child BACnetOptionalREALChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetOptionalREALNull
		_childTemp, typeSwitchError = BACnetOptionalREALNullParseWithBuffer(ctx, readBuffer)
	case 0 == 0: // BACnetOptionalREALValue
		_childTemp, typeSwitchError = BACnetOptionalREALValueParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetOptionalREAL")
	}
	_child = _childTemp.(BACnetOptionalREALChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetOptionalREAL"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalREAL")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetOptionalREAL) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetOptionalREAL, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetOptionalREAL"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOptionalREAL")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetOptionalREAL"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOptionalREAL")
	}
	return nil
}

func (m *_BACnetOptionalREAL) isBACnetOptionalREAL() bool {
	return true
}

func (m *_BACnetOptionalREAL) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
