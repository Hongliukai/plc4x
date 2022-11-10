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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// S7VarRequestParameterItem is the corresponding interface of S7VarRequestParameterItem
type S7VarRequestParameterItem interface {
	utils.LengthAware
	utils.Serializable
	// GetItemType returns ItemType (discriminator field)
	GetItemType() uint8
}

// S7VarRequestParameterItemExactly can be used when we want exactly this type and not a type which fulfills S7VarRequestParameterItem.
// This is useful for switch cases.
type S7VarRequestParameterItemExactly interface {
	S7VarRequestParameterItem
	isS7VarRequestParameterItem() bool
}

// _S7VarRequestParameterItem is the data-structure of this message
type _S7VarRequestParameterItem struct {
	_S7VarRequestParameterItemChildRequirements
}

type _S7VarRequestParameterItemChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetItemType() uint8
}

type S7VarRequestParameterItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child S7VarRequestParameterItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7VarRequestParameterItemChild interface {
	utils.Serializable
	InitializeParent(parent S7VarRequestParameterItem)
	GetParent() *S7VarRequestParameterItem

	GetTypeName() string
	S7VarRequestParameterItem
}

// NewS7VarRequestParameterItem factory function for _S7VarRequestParameterItem
func NewS7VarRequestParameterItem() *_S7VarRequestParameterItem {
	return &_S7VarRequestParameterItem{}
}

// Deprecated: use the interface for direct cast
func CastS7VarRequestParameterItem(structType interface{}) S7VarRequestParameterItem {
	if casted, ok := structType.(S7VarRequestParameterItem); ok {
		return casted
	}
	if casted, ok := structType.(*S7VarRequestParameterItem); ok {
		return *casted
	}
	return nil
}

func (m *_S7VarRequestParameterItem) GetTypeName() string {
	return "S7VarRequestParameterItem"
}

func (m *_S7VarRequestParameterItem) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7VarRequestParameterItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7VarRequestParameterItemParse(theBytes []byte) (S7VarRequestParameterItem, error) {
	return S7VarRequestParameterItemParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func S7VarRequestParameterItemParseWithBuffer(readBuffer utils.ReadBuffer) (S7VarRequestParameterItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarRequestParameterItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7VarRequestParameterItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType, _itemTypeErr := readBuffer.ReadUint8("itemType", 8)
	if _itemTypeErr != nil {
		return nil, errors.Wrap(_itemTypeErr, "Error parsing 'itemType' field of S7VarRequestParameterItem")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7VarRequestParameterItemChildSerializeRequirement interface {
		S7VarRequestParameterItem
		InitializeParent(S7VarRequestParameterItem)
		GetParent() S7VarRequestParameterItem
	}
	var _childTemp interface{}
	var _child S7VarRequestParameterItemChildSerializeRequirement
	var typeSwitchError error
	switch {
	case itemType == 0x12: // S7VarRequestParameterItemAddress
		_childTemp, typeSwitchError = S7VarRequestParameterItemAddressParseWithBuffer(readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [itemType=%v]", itemType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of S7VarRequestParameterItem")
	}
	_child = _childTemp.(S7VarRequestParameterItemChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("S7VarRequestParameterItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7VarRequestParameterItem")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_S7VarRequestParameterItem) SerializeParent(writeBuffer utils.WriteBuffer, child S7VarRequestParameterItem, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7VarRequestParameterItem"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7VarRequestParameterItem")
	}

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType := uint8(child.GetItemType())
	_itemTypeErr := writeBuffer.WriteUint8("itemType", 8, (itemType))

	if _itemTypeErr != nil {
		return errors.Wrap(_itemTypeErr, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7VarRequestParameterItem"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7VarRequestParameterItem")
	}
	return nil
}

func (m *_S7VarRequestParameterItem) isS7VarRequestParameterItem() bool {
	return true
}

func (m *_S7VarRequestParameterItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
