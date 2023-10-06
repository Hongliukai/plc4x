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

// DeleteReferencesItem is the corresponding interface of DeleteReferencesItem
type DeleteReferencesItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSourceNodeId returns SourceNodeId (property field)
	GetSourceNodeId() NodeId
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetTargetNodeId returns TargetNodeId (property field)
	GetTargetNodeId() ExpandedNodeId
	// GetDeleteBidirectional returns DeleteBidirectional (property field)
	GetDeleteBidirectional() bool
}

// DeleteReferencesItemExactly can be used when we want exactly this type and not a type which fulfills DeleteReferencesItem.
// This is useful for switch cases.
type DeleteReferencesItemExactly interface {
	DeleteReferencesItem
	isDeleteReferencesItem() bool
}

// _DeleteReferencesItem is the data-structure of this message
type _DeleteReferencesItem struct {
	*_ExtensionObjectDefinition
	SourceNodeId        NodeId
	ReferenceTypeId     NodeId
	IsForward           bool
	TargetNodeId        ExpandedNodeId
	DeleteBidirectional bool
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteReferencesItem) GetIdentifier() string {
	return "387"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteReferencesItem) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DeleteReferencesItem) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteReferencesItem) GetSourceNodeId() NodeId {
	return m.SourceNodeId
}

func (m *_DeleteReferencesItem) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_DeleteReferencesItem) GetIsForward() bool {
	return m.IsForward
}

func (m *_DeleteReferencesItem) GetTargetNodeId() ExpandedNodeId {
	return m.TargetNodeId
}

func (m *_DeleteReferencesItem) GetDeleteBidirectional() bool {
	return m.DeleteBidirectional
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteReferencesItem factory function for _DeleteReferencesItem
func NewDeleteReferencesItem(sourceNodeId NodeId, referenceTypeId NodeId, isForward bool, targetNodeId ExpandedNodeId, deleteBidirectional bool) *_DeleteReferencesItem {
	_result := &_DeleteReferencesItem{
		SourceNodeId:               sourceNodeId,
		ReferenceTypeId:            referenceTypeId,
		IsForward:                  isForward,
		TargetNodeId:               targetNodeId,
		DeleteBidirectional:        deleteBidirectional,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteReferencesItem(structType any) DeleteReferencesItem {
	if casted, ok := structType.(DeleteReferencesItem); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteReferencesItem); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteReferencesItem) GetTypeName() string {
	return "DeleteReferencesItem"
}

func (m *_DeleteReferencesItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (sourceNodeId)
	lengthInBits += m.SourceNodeId.GetLengthInBits(ctx)

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (targetNodeId)
	lengthInBits += m.TargetNodeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (deleteBidirectional)
	lengthInBits += 1

	return lengthInBits
}

func (m *_DeleteReferencesItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeleteReferencesItemParse(ctx context.Context, theBytes []byte, identifier string) (DeleteReferencesItem, error) {
	return DeleteReferencesItemParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DeleteReferencesItemParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DeleteReferencesItem, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeleteReferencesItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteReferencesItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sourceNodeId)
	if pullErr := readBuffer.PullContext("sourceNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sourceNodeId")
	}
	_sourceNodeId, _sourceNodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _sourceNodeIdErr != nil {
		return nil, errors.Wrap(_sourceNodeIdErr, "Error parsing 'sourceNodeId' field of DeleteReferencesItem")
	}
	sourceNodeId := _sourceNodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("sourceNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sourceNodeId")
	}

	// Simple Field (referenceTypeId)
	if pullErr := readBuffer.PullContext("referenceTypeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceTypeId")
	}
	_referenceTypeId, _referenceTypeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _referenceTypeIdErr != nil {
		return nil, errors.Wrap(_referenceTypeIdErr, "Error parsing 'referenceTypeId' field of DeleteReferencesItem")
	}
	referenceTypeId := _referenceTypeId.(NodeId)
	if closeErr := readBuffer.CloseContext("referenceTypeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceTypeId")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DeleteReferencesItem")
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
		return nil, errors.Wrap(_isForwardErr, "Error parsing 'isForward' field of DeleteReferencesItem")
	}
	isForward := _isForward

	// Simple Field (targetNodeId)
	if pullErr := readBuffer.PullContext("targetNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetNodeId")
	}
	_targetNodeId, _targetNodeIdErr := ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	if _targetNodeIdErr != nil {
		return nil, errors.Wrap(_targetNodeIdErr, "Error parsing 'targetNodeId' field of DeleteReferencesItem")
	}
	targetNodeId := _targetNodeId.(ExpandedNodeId)
	if closeErr := readBuffer.CloseContext("targetNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetNodeId")
	}

	var reservedField1 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of DeleteReferencesItem")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Simple Field (deleteBidirectional)
	_deleteBidirectional, _deleteBidirectionalErr := readBuffer.ReadBit("deleteBidirectional")
	if _deleteBidirectionalErr != nil {
		return nil, errors.Wrap(_deleteBidirectionalErr, "Error parsing 'deleteBidirectional' field of DeleteReferencesItem")
	}
	deleteBidirectional := _deleteBidirectional

	if closeErr := readBuffer.CloseContext("DeleteReferencesItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteReferencesItem")
	}

	// Create a partially initialized instance
	_child := &_DeleteReferencesItem{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		SourceNodeId:               sourceNodeId,
		ReferenceTypeId:            referenceTypeId,
		IsForward:                  isForward,
		TargetNodeId:               targetNodeId,
		DeleteBidirectional:        deleteBidirectional,
		reservedField0:             reservedField0,
		reservedField1:             reservedField1,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DeleteReferencesItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteReferencesItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteReferencesItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteReferencesItem")
		}

		// Simple Field (sourceNodeId)
		if pushErr := writeBuffer.PushContext("sourceNodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sourceNodeId")
		}
		_sourceNodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetSourceNodeId())
		if popErr := writeBuffer.PopContext("sourceNodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sourceNodeId")
		}
		if _sourceNodeIdErr != nil {
			return errors.Wrap(_sourceNodeIdErr, "Error serializing 'sourceNodeId' field")
		}

		// Simple Field (referenceTypeId)
		if pushErr := writeBuffer.PushContext("referenceTypeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceTypeId")
		}
		_referenceTypeIdErr := writeBuffer.WriteSerializable(ctx, m.GetReferenceTypeId())
		if popErr := writeBuffer.PopContext("referenceTypeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceTypeId")
		}
		if _referenceTypeIdErr != nil {
			return errors.Wrap(_referenceTypeIdErr, "Error serializing 'referenceTypeId' field")
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

		// Simple Field (targetNodeId)
		if pushErr := writeBuffer.PushContext("targetNodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetNodeId")
		}
		_targetNodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetTargetNodeId())
		if popErr := writeBuffer.PopContext("targetNodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetNodeId")
		}
		if _targetNodeIdErr != nil {
			return errors.Wrap(_targetNodeIdErr, "Error serializing 'targetNodeId' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField1 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (deleteBidirectional)
		deleteBidirectional := bool(m.GetDeleteBidirectional())
		_deleteBidirectionalErr := writeBuffer.WriteBit("deleteBidirectional", (deleteBidirectional))
		if _deleteBidirectionalErr != nil {
			return errors.Wrap(_deleteBidirectionalErr, "Error serializing 'deleteBidirectional' field")
		}

		if popErr := writeBuffer.PopContext("DeleteReferencesItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteReferencesItem")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteReferencesItem) isDeleteReferencesItem() bool {
	return true
}

func (m *_DeleteReferencesItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
