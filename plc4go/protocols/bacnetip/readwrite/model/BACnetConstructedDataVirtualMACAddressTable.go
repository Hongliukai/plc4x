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

// BACnetConstructedDataVirtualMACAddressTable is the corresponding interface of BACnetConstructedDataVirtualMACAddressTable
type BACnetConstructedDataVirtualMACAddressTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetVirtualMacAddressTable returns VirtualMacAddressTable (property field)
	GetVirtualMacAddressTable() []BACnetVMACEntry
}

// BACnetConstructedDataVirtualMACAddressTableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataVirtualMACAddressTable.
// This is useful for switch cases.
type BACnetConstructedDataVirtualMACAddressTableExactly interface {
	BACnetConstructedDataVirtualMACAddressTable
	isBACnetConstructedDataVirtualMACAddressTable() bool
}

// _BACnetConstructedDataVirtualMACAddressTable is the data-structure of this message
type _BACnetConstructedDataVirtualMACAddressTable struct {
	*_BACnetConstructedData
	VirtualMacAddressTable []BACnetVMACEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VIRTUAL_MAC_ADDRESS_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVirtualMACAddressTable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetVirtualMacAddressTable() []BACnetVMACEntry {
	return m.VirtualMacAddressTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVirtualMACAddressTable factory function for _BACnetConstructedDataVirtualMACAddressTable
func NewBACnetConstructedDataVirtualMACAddressTable(virtualMacAddressTable []BACnetVMACEntry, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVirtualMACAddressTable {
	_result := &_BACnetConstructedDataVirtualMACAddressTable{
		VirtualMacAddressTable: virtualMacAddressTable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVirtualMACAddressTable(structType any) BACnetConstructedDataVirtualMACAddressTable {
	if casted, ok := structType.(BACnetConstructedDataVirtualMACAddressTable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVirtualMACAddressTable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetTypeName() string {
	return "BACnetConstructedDataVirtualMACAddressTable"
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.VirtualMacAddressTable) > 0 {
		for _, element := range m.VirtualMacAddressTable {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataVirtualMACAddressTableParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVirtualMACAddressTable, error) {
	return BACnetConstructedDataVirtualMACAddressTableParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataVirtualMACAddressTableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVirtualMACAddressTable, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVirtualMACAddressTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVirtualMACAddressTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (virtualMacAddressTable)
	if pullErr := readBuffer.PullContext("virtualMacAddressTable", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for virtualMacAddressTable")
	}
	// Terminated array
	var virtualMacAddressTable []BACnetVMACEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetVMACEntryParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'virtualMacAddressTable' field of BACnetConstructedDataVirtualMACAddressTable")
			}
			virtualMacAddressTable = append(virtualMacAddressTable, _item.(BACnetVMACEntry))
		}
	}
	if closeErr := readBuffer.CloseContext("virtualMacAddressTable", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for virtualMacAddressTable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVirtualMACAddressTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVirtualMACAddressTable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataVirtualMACAddressTable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		VirtualMacAddressTable: virtualMacAddressTable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVirtualMACAddressTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVirtualMACAddressTable")
		}

		// Array Field (virtualMacAddressTable)
		if pushErr := writeBuffer.PushContext("virtualMacAddressTable", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for virtualMacAddressTable")
		}
		for _curItem, _element := range m.GetVirtualMacAddressTable() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetVirtualMacAddressTable()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'virtualMacAddressTable' field")
			}
		}
		if popErr := writeBuffer.PopContext("virtualMacAddressTable", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for virtualMacAddressTable")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVirtualMACAddressTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVirtualMACAddressTable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) isBACnetConstructedDataVirtualMACAddressTable() bool {
	return true
}

func (m *_BACnetConstructedDataVirtualMACAddressTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
