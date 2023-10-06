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

// MediaTransportControlDataSetSelection is the corresponding interface of MediaTransportControlDataSetSelection
type MediaTransportControlDataSetSelection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetSelectionHi returns SelectionHi (property field)
	GetSelectionHi() byte
	// GetSelectionLo returns SelectionLo (property field)
	GetSelectionLo() byte
}

// MediaTransportControlDataSetSelectionExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataSetSelection.
// This is useful for switch cases.
type MediaTransportControlDataSetSelectionExactly interface {
	MediaTransportControlDataSetSelection
	isMediaTransportControlDataSetSelection() bool
}

// _MediaTransportControlDataSetSelection is the data-structure of this message
type _MediaTransportControlDataSetSelection struct {
	*_MediaTransportControlData
	SelectionHi byte
	SelectionLo byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataSetSelection) InitializeParent(parent MediaTransportControlData, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataSetSelection) GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataSetSelection) GetSelectionHi() byte {
	return m.SelectionHi
}

func (m *_MediaTransportControlDataSetSelection) GetSelectionLo() byte {
	return m.SelectionLo
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataSetSelection factory function for _MediaTransportControlDataSetSelection
func NewMediaTransportControlDataSetSelection(selectionHi byte, selectionLo byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataSetSelection {
	_result := &_MediaTransportControlDataSetSelection{
		SelectionHi:                selectionHi,
		SelectionLo:                selectionLo,
		_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataSetSelection(structType any) MediaTransportControlDataSetSelection {
	if casted, ok := structType.(MediaTransportControlDataSetSelection); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataSetSelection); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataSetSelection) GetTypeName() string {
	return "MediaTransportControlDataSetSelection"
}

func (m *_MediaTransportControlDataSetSelection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (selectionHi)
	lengthInBits += 8

	// Simple field (selectionLo)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlDataSetSelection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataSetSelectionParse(ctx context.Context, theBytes []byte) (MediaTransportControlDataSetSelection, error) {
	return MediaTransportControlDataSetSelectionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataSetSelectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataSetSelection, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MediaTransportControlDataSetSelection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataSetSelection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (selectionHi)
	_selectionHi, _selectionHiErr := readBuffer.ReadByte("selectionHi")
	if _selectionHiErr != nil {
		return nil, errors.Wrap(_selectionHiErr, "Error parsing 'selectionHi' field of MediaTransportControlDataSetSelection")
	}
	selectionHi := _selectionHi

	// Simple Field (selectionLo)
	_selectionLo, _selectionLoErr := readBuffer.ReadByte("selectionLo")
	if _selectionLoErr != nil {
		return nil, errors.Wrap(_selectionLoErr, "Error parsing 'selectionLo' field of MediaTransportControlDataSetSelection")
	}
	selectionLo := _selectionLo

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataSetSelection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataSetSelection")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataSetSelection{
		_MediaTransportControlData: &_MediaTransportControlData{},
		SelectionHi:                selectionHi,
		SelectionLo:                selectionLo,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataSetSelection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataSetSelection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataSetSelection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataSetSelection")
		}

		// Simple Field (selectionHi)
		selectionHi := byte(m.GetSelectionHi())
		_selectionHiErr := writeBuffer.WriteByte("selectionHi", (selectionHi))
		if _selectionHiErr != nil {
			return errors.Wrap(_selectionHiErr, "Error serializing 'selectionHi' field")
		}

		// Simple Field (selectionLo)
		selectionLo := byte(m.GetSelectionLo())
		_selectionLoErr := writeBuffer.WriteByte("selectionLo", (selectionLo))
		if _selectionLoErr != nil {
			return errors.Wrap(_selectionLoErr, "Error serializing 'selectionLo' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataSetSelection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataSetSelection")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataSetSelection) isMediaTransportControlDataSetSelection() bool {
	return true
}

func (m *_MediaTransportControlDataSetSelection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
