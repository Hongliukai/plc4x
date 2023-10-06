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

// Orientation is the corresponding interface of Orientation
type Orientation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
}

// OrientationExactly can be used when we want exactly this type and not a type which fulfills Orientation.
// This is useful for switch cases.
type OrientationExactly interface {
	Orientation
	isOrientation() bool
}

// _Orientation is the data-structure of this message
type _Orientation struct {
	*_ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Orientation) GetIdentifier() string {
	return "18813"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Orientation) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_Orientation) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

// NewOrientation factory function for _Orientation
func NewOrientation() *_Orientation {
	_result := &_Orientation{
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOrientation(structType any) Orientation {
	if casted, ok := structType.(Orientation); ok {
		return casted
	}
	if casted, ok := structType.(*Orientation); ok {
		return *casted
	}
	return nil
}

func (m *_Orientation) GetTypeName() string {
	return "Orientation"
}

func (m *_Orientation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_Orientation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OrientationParse(ctx context.Context, theBytes []byte, identifier string) (Orientation, error) {
	return OrientationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func OrientationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (Orientation, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("Orientation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Orientation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Orientation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Orientation")
	}

	// Create a partially initialized instance
	_child := &_Orientation{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_Orientation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Orientation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Orientation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Orientation")
		}

		if popErr := writeBuffer.PopContext("Orientation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Orientation")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Orientation) isOrientation() bool {
	return true
}

func (m *_Orientation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
