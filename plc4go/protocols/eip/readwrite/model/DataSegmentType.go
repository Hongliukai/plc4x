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

// DataSegmentType is the corresponding interface of DataSegmentType
type DataSegmentType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDataSegmentType returns DataSegmentType (discriminator field)
	GetDataSegmentType() uint8
}

// DataSegmentTypeExactly can be used when we want exactly this type and not a type which fulfills DataSegmentType.
// This is useful for switch cases.
type DataSegmentTypeExactly interface {
	DataSegmentType
	isDataSegmentType() bool
}

// _DataSegmentType is the data-structure of this message
type _DataSegmentType struct {
	_DataSegmentTypeChildRequirements
}

type _DataSegmentTypeChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetDataSegmentType() uint8
}

type DataSegmentTypeParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DataSegmentType, serializeChildFunction func() error) error
	GetTypeName() string
}

type DataSegmentTypeChild interface {
	utils.Serializable
	InitializeParent(parent DataSegmentType)
	GetParent() *DataSegmentType

	GetTypeName() string
	DataSegmentType
}

// NewDataSegmentType factory function for _DataSegmentType
func NewDataSegmentType() *_DataSegmentType {
	return &_DataSegmentType{}
}

// Deprecated: use the interface for direct cast
func CastDataSegmentType(structType any) DataSegmentType {
	if casted, ok := structType.(DataSegmentType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSegmentType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSegmentType) GetTypeName() string {
	return "DataSegmentType"
}

func (m *_DataSegmentType) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (dataSegmentType)
	lengthInBits += 5

	return lengthInBits
}

func (m *_DataSegmentType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataSegmentTypeParse(ctx context.Context, theBytes []byte) (DataSegmentType, error) {
	return DataSegmentTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DataSegmentTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DataSegmentType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DataSegmentType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSegmentType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (dataSegmentType) (Used as input to a switch field)
	dataSegmentType, _dataSegmentTypeErr := readBuffer.ReadUint8("dataSegmentType", 5)
	if _dataSegmentTypeErr != nil {
		return nil, errors.Wrap(_dataSegmentTypeErr, "Error parsing 'dataSegmentType' field of DataSegmentType")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type DataSegmentTypeChildSerializeRequirement interface {
		DataSegmentType
		InitializeParent(DataSegmentType)
		GetParent() DataSegmentType
	}
	var _childTemp any
	var _child DataSegmentTypeChildSerializeRequirement
	var typeSwitchError error
	switch {
	case dataSegmentType == 0x11: // AnsiExtendedSymbolSegment
		_childTemp, typeSwitchError = AnsiExtendedSymbolSegmentParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [dataSegmentType=%v]", dataSegmentType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of DataSegmentType")
	}
	_child = _childTemp.(DataSegmentTypeChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("DataSegmentType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSegmentType")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_DataSegmentType) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child DataSegmentType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("DataSegmentType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DataSegmentType")
	}

	// Discriminator Field (dataSegmentType) (Used as input to a switch field)
	dataSegmentType := uint8(child.GetDataSegmentType())
	_dataSegmentTypeErr := writeBuffer.WriteUint8("dataSegmentType", 5, (dataSegmentType))

	if _dataSegmentTypeErr != nil {
		return errors.Wrap(_dataSegmentTypeErr, "Error serializing 'dataSegmentType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("DataSegmentType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DataSegmentType")
	}
	return nil
}

func (m *_DataSegmentType) isDataSegmentType() bool {
	return true
}

func (m *_DataSegmentType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
