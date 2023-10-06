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

// SysexCommandAnalogMappingResponse is the corresponding interface of SysexCommandAnalogMappingResponse
type SysexCommandAnalogMappingResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandAnalogMappingResponseExactly can be used when we want exactly this type and not a type which fulfills SysexCommandAnalogMappingResponse.
// This is useful for switch cases.
type SysexCommandAnalogMappingResponseExactly interface {
	SysexCommandAnalogMappingResponse
	isSysexCommandAnalogMappingResponse() bool
}

// _SysexCommandAnalogMappingResponse is the data-structure of this message
type _SysexCommandAnalogMappingResponse struct {
	*_SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandAnalogMappingResponse) GetCommandType() uint8 {
	return 0x6A
}

func (m *_SysexCommandAnalogMappingResponse) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandAnalogMappingResponse) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandAnalogMappingResponse) GetParent() SysexCommand {
	return m._SysexCommand
}

// NewSysexCommandAnalogMappingResponse factory function for _SysexCommandAnalogMappingResponse
func NewSysexCommandAnalogMappingResponse() *_SysexCommandAnalogMappingResponse {
	_result := &_SysexCommandAnalogMappingResponse{
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandAnalogMappingResponse(structType any) SysexCommandAnalogMappingResponse {
	if casted, ok := structType.(SysexCommandAnalogMappingResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandAnalogMappingResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandAnalogMappingResponse) GetTypeName() string {
	return "SysexCommandAnalogMappingResponse"
}

func (m *_SysexCommandAnalogMappingResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandAnalogMappingResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SysexCommandAnalogMappingResponseParse(ctx context.Context, theBytes []byte, response bool) (SysexCommandAnalogMappingResponse, error) {
	return SysexCommandAnalogMappingResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandAnalogMappingResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (SysexCommandAnalogMappingResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandAnalogMappingResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandAnalogMappingResponse")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandAnalogMappingResponse{
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandAnalogMappingResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandAnalogMappingResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandAnalogMappingResponse")
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandAnalogMappingResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandAnalogMappingResponse) isSysexCommandAnalogMappingResponse() bool {
	return true
}

func (m *_SysexCommandAnalogMappingResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
