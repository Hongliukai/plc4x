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

// COTPParameterCallingTsap is the corresponding interface of COTPParameterCallingTsap
type COTPParameterCallingTsap interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	COTPParameter
	// GetTsapId returns TsapId (property field)
	GetTsapId() uint16
}

// COTPParameterCallingTsapExactly can be used when we want exactly this type and not a type which fulfills COTPParameterCallingTsap.
// This is useful for switch cases.
type COTPParameterCallingTsapExactly interface {
	COTPParameterCallingTsap
	isCOTPParameterCallingTsap() bool
}

// _COTPParameterCallingTsap is the data-structure of this message
type _COTPParameterCallingTsap struct {
	*_COTPParameter
	TsapId uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPParameterCallingTsap) GetParameterType() uint8 {
	return 0xC1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPParameterCallingTsap) InitializeParent(parent COTPParameter) {}

func (m *_COTPParameterCallingTsap) GetParent() COTPParameter {
	return m._COTPParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPParameterCallingTsap) GetTsapId() uint16 {
	return m.TsapId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPParameterCallingTsap factory function for _COTPParameterCallingTsap
func NewCOTPParameterCallingTsap(tsapId uint16, rest uint8) *_COTPParameterCallingTsap {
	_result := &_COTPParameterCallingTsap{
		TsapId:         tsapId,
		_COTPParameter: NewCOTPParameter(rest),
	}
	_result._COTPParameter._COTPParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPParameterCallingTsap(structType any) COTPParameterCallingTsap {
	if casted, ok := structType.(COTPParameterCallingTsap); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameterCallingTsap); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameterCallingTsap) GetTypeName() string {
	return "COTPParameterCallingTsap"
}

func (m *_COTPParameterCallingTsap) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (tsapId)
	lengthInBits += 16

	return lengthInBits
}

func (m *_COTPParameterCallingTsap) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPParameterCallingTsapParse(ctx context.Context, theBytes []byte, rest uint8) (COTPParameterCallingTsap, error) {
	return COTPParameterCallingTsapParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), rest)
}

func COTPParameterCallingTsapParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, rest uint8) (COTPParameterCallingTsap, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("COTPParameterCallingTsap"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameterCallingTsap")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (tsapId)
	_tsapId, _tsapIdErr := readBuffer.ReadUint16("tsapId", 16)
	if _tsapIdErr != nil {
		return nil, errors.Wrap(_tsapIdErr, "Error parsing 'tsapId' field of COTPParameterCallingTsap")
	}
	tsapId := _tsapId

	if closeErr := readBuffer.CloseContext("COTPParameterCallingTsap"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameterCallingTsap")
	}

	// Create a partially initialized instance
	_child := &_COTPParameterCallingTsap{
		_COTPParameter: &_COTPParameter{
			Rest: rest,
		},
		TsapId: tsapId,
	}
	_child._COTPParameter._COTPParameterChildRequirements = _child
	return _child, nil
}

func (m *_COTPParameterCallingTsap) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPParameterCallingTsap) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPParameterCallingTsap"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPParameterCallingTsap")
		}

		// Simple Field (tsapId)
		tsapId := uint16(m.GetTsapId())
		_tsapIdErr := writeBuffer.WriteUint16("tsapId", 16, (tsapId))
		if _tsapIdErr != nil {
			return errors.Wrap(_tsapIdErr, "Error serializing 'tsapId' field")
		}

		if popErr := writeBuffer.PopContext("COTPParameterCallingTsap"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPParameterCallingTsap")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPParameterCallingTsap) isCOTPParameterCallingTsap() bool {
	return true
}

func (m *_COTPParameterCallingTsap) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
