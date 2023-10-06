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

// ReferenceListEntryDataType is the corresponding interface of ReferenceListEntryDataType
type ReferenceListEntryDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetTargetNode returns TargetNode (property field)
	GetTargetNode() ExpandedNodeId
}

// ReferenceListEntryDataTypeExactly can be used when we want exactly this type and not a type which fulfills ReferenceListEntryDataType.
// This is useful for switch cases.
type ReferenceListEntryDataTypeExactly interface {
	ReferenceListEntryDataType
	isReferenceListEntryDataType() bool
}

// _ReferenceListEntryDataType is the data-structure of this message
type _ReferenceListEntryDataType struct {
	*_ExtensionObjectDefinition
	ReferenceType NodeId
	IsForward     bool
	TargetNode    ExpandedNodeId
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReferenceListEntryDataType) GetIdentifier() string {
	return "32662"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReferenceListEntryDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ReferenceListEntryDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReferenceListEntryDataType) GetReferenceType() NodeId {
	return m.ReferenceType
}

func (m *_ReferenceListEntryDataType) GetIsForward() bool {
	return m.IsForward
}

func (m *_ReferenceListEntryDataType) GetTargetNode() ExpandedNodeId {
	return m.TargetNode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReferenceListEntryDataType factory function for _ReferenceListEntryDataType
func NewReferenceListEntryDataType(referenceType NodeId, isForward bool, targetNode ExpandedNodeId) *_ReferenceListEntryDataType {
	_result := &_ReferenceListEntryDataType{
		ReferenceType:              referenceType,
		IsForward:                  isForward,
		TargetNode:                 targetNode,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReferenceListEntryDataType(structType any) ReferenceListEntryDataType {
	if casted, ok := structType.(ReferenceListEntryDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ReferenceListEntryDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ReferenceListEntryDataType) GetTypeName() string {
	return "ReferenceListEntryDataType"
}

func (m *_ReferenceListEntryDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (referenceType)
	lengthInBits += m.ReferenceType.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (targetNode)
	lengthInBits += m.TargetNode.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReferenceListEntryDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReferenceListEntryDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (ReferenceListEntryDataType, error) {
	return ReferenceListEntryDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ReferenceListEntryDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ReferenceListEntryDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ReferenceListEntryDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReferenceListEntryDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceType)
	if pullErr := readBuffer.PullContext("referenceType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceType")
	}
	_referenceType, _referenceTypeErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _referenceTypeErr != nil {
		return nil, errors.Wrap(_referenceTypeErr, "Error parsing 'referenceType' field of ReferenceListEntryDataType")
	}
	referenceType := _referenceType.(NodeId)
	if closeErr := readBuffer.CloseContext("referenceType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceType")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ReferenceListEntryDataType")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (isForward)
	_isForward, _isForwardErr := readBuffer.ReadBit("isForward")
	if _isForwardErr != nil {
		return nil, errors.Wrap(_isForwardErr, "Error parsing 'isForward' field of ReferenceListEntryDataType")
	}
	isForward := _isForward

	// Simple Field (targetNode)
	if pullErr := readBuffer.PullContext("targetNode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetNode")
	}
	_targetNode, _targetNodeErr := ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	if _targetNodeErr != nil {
		return nil, errors.Wrap(_targetNodeErr, "Error parsing 'targetNode' field of ReferenceListEntryDataType")
	}
	targetNode := _targetNode.(ExpandedNodeId)
	if closeErr := readBuffer.CloseContext("targetNode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetNode")
	}

	if closeErr := readBuffer.CloseContext("ReferenceListEntryDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReferenceListEntryDataType")
	}

	// Create a partially initialized instance
	_child := &_ReferenceListEntryDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ReferenceType:              referenceType,
		IsForward:                  isForward,
		TargetNode:                 targetNode,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ReferenceListEntryDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReferenceListEntryDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReferenceListEntryDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReferenceListEntryDataType")
		}

		// Simple Field (referenceType)
		if pushErr := writeBuffer.PushContext("referenceType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceType")
		}
		_referenceTypeErr := writeBuffer.WriteSerializable(ctx, m.GetReferenceType())
		if popErr := writeBuffer.PopContext("referenceType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceType")
		}
		if _referenceTypeErr != nil {
			return errors.Wrap(_referenceTypeErr, "Error serializing 'referenceType' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (isForward)
		isForward := bool(m.GetIsForward())
		_isForwardErr := writeBuffer.WriteBit("isForward", (isForward))
		if _isForwardErr != nil {
			return errors.Wrap(_isForwardErr, "Error serializing 'isForward' field")
		}

		// Simple Field (targetNode)
		if pushErr := writeBuffer.PushContext("targetNode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetNode")
		}
		_targetNodeErr := writeBuffer.WriteSerializable(ctx, m.GetTargetNode())
		if popErr := writeBuffer.PopContext("targetNode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetNode")
		}
		if _targetNodeErr != nil {
			return errors.Wrap(_targetNodeErr, "Error serializing 'targetNode' field")
		}

		if popErr := writeBuffer.PopContext("ReferenceListEntryDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReferenceListEntryDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReferenceListEntryDataType) isReferenceListEntryDataType() bool {
	return true
}

func (m *_ReferenceListEntryDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
