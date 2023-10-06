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

// IdentifyReplyCommandMaximumLevels is the corresponding interface of IdentifyReplyCommandMaximumLevels
type IdentifyReplyCommandMaximumLevels interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetMaximumLevels returns MaximumLevels (property field)
	GetMaximumLevels() []byte
}

// IdentifyReplyCommandMaximumLevelsExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandMaximumLevels.
// This is useful for switch cases.
type IdentifyReplyCommandMaximumLevelsExactly interface {
	IdentifyReplyCommandMaximumLevels
	isIdentifyReplyCommandMaximumLevels() bool
}

// _IdentifyReplyCommandMaximumLevels is the data-structure of this message
type _IdentifyReplyCommandMaximumLevels struct {
	*_IdentifyReplyCommand
	MaximumLevels []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandMaximumLevels) GetAttribute() Attribute {
	return Attribute_MaximumLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandMaximumLevels) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandMaximumLevels) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandMaximumLevels) GetMaximumLevels() []byte {
	return m.MaximumLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandMaximumLevels factory function for _IdentifyReplyCommandMaximumLevels
func NewIdentifyReplyCommandMaximumLevels(maximumLevels []byte, numBytes uint8) *_IdentifyReplyCommandMaximumLevels {
	_result := &_IdentifyReplyCommandMaximumLevels{
		MaximumLevels:         maximumLevels,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandMaximumLevels(structType any) IdentifyReplyCommandMaximumLevels {
	if casted, ok := structType.(IdentifyReplyCommandMaximumLevels); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandMaximumLevels); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandMaximumLevels) GetTypeName() string {
	return "IdentifyReplyCommandMaximumLevels"
}

func (m *_IdentifyReplyCommandMaximumLevels) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.MaximumLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.MaximumLevels))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandMaximumLevels) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandMaximumLevelsParse(ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandMaximumLevels, error) {
	return IdentifyReplyCommandMaximumLevelsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandMaximumLevelsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandMaximumLevels, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandMaximumLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandMaximumLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (maximumLevels)
	numberOfBytesmaximumLevels := int(numBytes)
	maximumLevels, _readArrayErr := readBuffer.ReadByteArray("maximumLevels", numberOfBytesmaximumLevels)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'maximumLevels' field of IdentifyReplyCommandMaximumLevels")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandMaximumLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandMaximumLevels")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandMaximumLevels{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		MaximumLevels: maximumLevels,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandMaximumLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandMaximumLevels) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandMaximumLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandMaximumLevels")
		}

		// Array Field (maximumLevels)
		// Byte Array field (maximumLevels)
		if err := writeBuffer.WriteByteArray("maximumLevels", m.GetMaximumLevels()); err != nil {
			return errors.Wrap(err, "Error serializing 'maximumLevels' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandMaximumLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandMaximumLevels")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandMaximumLevels) isIdentifyReplyCommandMaximumLevels() bool {
	return true
}

func (m *_IdentifyReplyCommandMaximumLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
