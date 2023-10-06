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

// SysexCommandAnalogMappingQueryRequest is the corresponding interface of SysexCommandAnalogMappingQueryRequest
type SysexCommandAnalogMappingQueryRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandAnalogMappingQueryRequestExactly can be used when we want exactly this type and not a type which fulfills SysexCommandAnalogMappingQueryRequest.
// This is useful for switch cases.
type SysexCommandAnalogMappingQueryRequestExactly interface {
	SysexCommandAnalogMappingQueryRequest
	isSysexCommandAnalogMappingQueryRequest() bool
}

// _SysexCommandAnalogMappingQueryRequest is the data-structure of this message
type _SysexCommandAnalogMappingQueryRequest struct {
	*_SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandAnalogMappingQueryRequest) GetCommandType() uint8 {
	return 0x69
}

func (m *_SysexCommandAnalogMappingQueryRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandAnalogMappingQueryRequest) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandAnalogMappingQueryRequest) GetParent() SysexCommand {
	return m._SysexCommand
}

// NewSysexCommandAnalogMappingQueryRequest factory function for _SysexCommandAnalogMappingQueryRequest
func NewSysexCommandAnalogMappingQueryRequest() *_SysexCommandAnalogMappingQueryRequest {
	_result := &_SysexCommandAnalogMappingQueryRequest{
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandAnalogMappingQueryRequest(structType any) SysexCommandAnalogMappingQueryRequest {
	if casted, ok := structType.(SysexCommandAnalogMappingQueryRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandAnalogMappingQueryRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandAnalogMappingQueryRequest) GetTypeName() string {
	return "SysexCommandAnalogMappingQueryRequest"
}

func (m *_SysexCommandAnalogMappingQueryRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandAnalogMappingQueryRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SysexCommandAnalogMappingQueryRequestParse(ctx context.Context, theBytes []byte, response bool) (SysexCommandAnalogMappingQueryRequest, error) {
	return SysexCommandAnalogMappingQueryRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandAnalogMappingQueryRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (SysexCommandAnalogMappingQueryRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingQueryRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandAnalogMappingQueryRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingQueryRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandAnalogMappingQueryRequest")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandAnalogMappingQueryRequest{
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandAnalogMappingQueryRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandAnalogMappingQueryRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingQueryRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandAnalogMappingQueryRequest")
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingQueryRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandAnalogMappingQueryRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandAnalogMappingQueryRequest) isSysexCommandAnalogMappingQueryRequest() bool {
	return true
}

func (m *_SysexCommandAnalogMappingQueryRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
