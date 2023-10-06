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

// BACnetAuthenticationPolicyListEntry is the corresponding interface of BACnetAuthenticationPolicyListEntry
type BACnetAuthenticationPolicyListEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCredentialDataInput returns CredentialDataInput (property field)
	GetCredentialDataInput() BACnetDeviceObjectReferenceEnclosed
	// GetIndex returns Index (property field)
	GetIndex() BACnetContextTagUnsignedInteger
}

// BACnetAuthenticationPolicyListEntryExactly can be used when we want exactly this type and not a type which fulfills BACnetAuthenticationPolicyListEntry.
// This is useful for switch cases.
type BACnetAuthenticationPolicyListEntryExactly interface {
	BACnetAuthenticationPolicyListEntry
	isBACnetAuthenticationPolicyListEntry() bool
}

// _BACnetAuthenticationPolicyListEntry is the data-structure of this message
type _BACnetAuthenticationPolicyListEntry struct {
	CredentialDataInput BACnetDeviceObjectReferenceEnclosed
	Index               BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationPolicyListEntry) GetCredentialDataInput() BACnetDeviceObjectReferenceEnclosed {
	return m.CredentialDataInput
}

func (m *_BACnetAuthenticationPolicyListEntry) GetIndex() BACnetContextTagUnsignedInteger {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationPolicyListEntry factory function for _BACnetAuthenticationPolicyListEntry
func NewBACnetAuthenticationPolicyListEntry(credentialDataInput BACnetDeviceObjectReferenceEnclosed, index BACnetContextTagUnsignedInteger) *_BACnetAuthenticationPolicyListEntry {
	return &_BACnetAuthenticationPolicyListEntry{CredentialDataInput: credentialDataInput, Index: index}
}

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationPolicyListEntry(structType any) BACnetAuthenticationPolicyListEntry {
	if casted, ok := structType.(BACnetAuthenticationPolicyListEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationPolicyListEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationPolicyListEntry) GetTypeName() string {
	return "BACnetAuthenticationPolicyListEntry"
}

func (m *_BACnetAuthenticationPolicyListEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (credentialDataInput)
	lengthInBits += m.CredentialDataInput.GetLengthInBits(ctx)

	// Simple field (index)
	lengthInBits += m.Index.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAuthenticationPolicyListEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationPolicyListEntryParse(ctx context.Context, theBytes []byte) (BACnetAuthenticationPolicyListEntry, error) {
	return BACnetAuthenticationPolicyListEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthenticationPolicyListEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationPolicyListEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetAuthenticationPolicyListEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationPolicyListEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (credentialDataInput)
	if pullErr := readBuffer.PullContext("credentialDataInput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for credentialDataInput")
	}
	_credentialDataInput, _credentialDataInputErr := BACnetDeviceObjectReferenceEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _credentialDataInputErr != nil {
		return nil, errors.Wrap(_credentialDataInputErr, "Error parsing 'credentialDataInput' field of BACnetAuthenticationPolicyListEntry")
	}
	credentialDataInput := _credentialDataInput.(BACnetDeviceObjectReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("credentialDataInput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for credentialDataInput")
	}

	// Simple Field (index)
	if pullErr := readBuffer.PullContext("index"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for index")
	}
	_index, _indexErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _indexErr != nil {
		return nil, errors.Wrap(_indexErr, "Error parsing 'index' field of BACnetAuthenticationPolicyListEntry")
	}
	index := _index.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("index"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for index")
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationPolicyListEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationPolicyListEntry")
	}

	// Create the instance
	return &_BACnetAuthenticationPolicyListEntry{
		CredentialDataInput: credentialDataInput,
		Index:               index,
	}, nil
}

func (m *_BACnetAuthenticationPolicyListEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationPolicyListEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationPolicyListEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationPolicyListEntry")
	}

	// Simple Field (credentialDataInput)
	if pushErr := writeBuffer.PushContext("credentialDataInput"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for credentialDataInput")
	}
	_credentialDataInputErr := writeBuffer.WriteSerializable(ctx, m.GetCredentialDataInput())
	if popErr := writeBuffer.PopContext("credentialDataInput"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for credentialDataInput")
	}
	if _credentialDataInputErr != nil {
		return errors.Wrap(_credentialDataInputErr, "Error serializing 'credentialDataInput' field")
	}

	// Simple Field (index)
	if pushErr := writeBuffer.PushContext("index"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for index")
	}
	_indexErr := writeBuffer.WriteSerializable(ctx, m.GetIndex())
	if popErr := writeBuffer.PopContext("index"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for index")
	}
	if _indexErr != nil {
		return errors.Wrap(_indexErr, "Error serializing 'index' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationPolicyListEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationPolicyListEntry")
	}
	return nil
}

func (m *_BACnetAuthenticationPolicyListEntry) isBACnetAuthenticationPolicyListEntry() bool {
	return true
}

func (m *_BACnetAuthenticationPolicyListEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
