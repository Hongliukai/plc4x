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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CEMIAdditionalInformationBusmonitorInfo_LEN uint8 = uint8(1)

// CEMIAdditionalInformationBusmonitorInfo is the corresponding interface of CEMIAdditionalInformationBusmonitorInfo
type CEMIAdditionalInformationBusmonitorInfo interface {
	utils.LengthAware
	utils.Serializable
	CEMIAdditionalInformation
	// GetFrameErrorFlag returns FrameErrorFlag (property field)
	GetFrameErrorFlag() bool
	// GetBitErrorFlag returns BitErrorFlag (property field)
	GetBitErrorFlag() bool
	// GetParityErrorFlag returns ParityErrorFlag (property field)
	GetParityErrorFlag() bool
	// GetUnknownFlag returns UnknownFlag (property field)
	GetUnknownFlag() bool
	// GetLostFlag returns LostFlag (property field)
	GetLostFlag() bool
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
}

// CEMIAdditionalInformationBusmonitorInfoExactly can be used when we want exactly this type and not a type which fulfills CEMIAdditionalInformationBusmonitorInfo.
// This is useful for switch cases.
type CEMIAdditionalInformationBusmonitorInfoExactly interface {
	CEMIAdditionalInformationBusmonitorInfo
	isCEMIAdditionalInformationBusmonitorInfo() bool
}

// _CEMIAdditionalInformationBusmonitorInfo is the data-structure of this message
type _CEMIAdditionalInformationBusmonitorInfo struct {
	*_CEMIAdditionalInformation
	FrameErrorFlag  bool
	BitErrorFlag    bool
	ParityErrorFlag bool
	UnknownFlag     bool
	LostFlag        bool
	SequenceNumber  uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetAdditionalInformationType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CEMIAdditionalInformationBusmonitorInfo) InitializeParent(parent CEMIAdditionalInformation) {
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetParent() CEMIAdditionalInformation {
	return m._CEMIAdditionalInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetFrameErrorFlag() bool {
	return m.FrameErrorFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetBitErrorFlag() bool {
	return m.BitErrorFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetParityErrorFlag() bool {
	return m.ParityErrorFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetUnknownFlag() bool {
	return m.UnknownFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLostFlag() bool {
	return m.LostFlag
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLen() uint8 {
	return CEMIAdditionalInformationBusmonitorInfo_LEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCEMIAdditionalInformationBusmonitorInfo factory function for _CEMIAdditionalInformationBusmonitorInfo
func NewCEMIAdditionalInformationBusmonitorInfo(frameErrorFlag bool, bitErrorFlag bool, parityErrorFlag bool, unknownFlag bool, lostFlag bool, sequenceNumber uint8) *_CEMIAdditionalInformationBusmonitorInfo {
	_result := &_CEMIAdditionalInformationBusmonitorInfo{
		FrameErrorFlag:             frameErrorFlag,
		BitErrorFlag:               bitErrorFlag,
		ParityErrorFlag:            parityErrorFlag,
		UnknownFlag:                unknownFlag,
		LostFlag:                   lostFlag,
		SequenceNumber:             sequenceNumber,
		_CEMIAdditionalInformation: NewCEMIAdditionalInformation(),
	}
	_result._CEMIAdditionalInformation._CEMIAdditionalInformationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCEMIAdditionalInformationBusmonitorInfo(structType interface{}) CEMIAdditionalInformationBusmonitorInfo {
	if casted, ok := structType.(CEMIAdditionalInformationBusmonitorInfo); ok {
		return casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformationBusmonitorInfo); ok {
		return *casted
	}
	return nil
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetTypeName() string {
	return "CEMIAdditionalInformationBusmonitorInfo"
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (len)
	lengthInBits += 8

	// Simple field (frameErrorFlag)
	lengthInBits += 1

	// Simple field (bitErrorFlag)
	lengthInBits += 1

	// Simple field (parityErrorFlag)
	lengthInBits += 1

	// Simple field (unknownFlag)
	lengthInBits += 1

	// Simple field (lostFlag)
	lengthInBits += 1

	// Simple field (sequenceNumber)
	lengthInBits += 3

	return lengthInBits
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CEMIAdditionalInformationBusmonitorInfoParse(theBytes []byte) (CEMIAdditionalInformationBusmonitorInfo, error) {
	return CEMIAdditionalInformationBusmonitorInfoParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func CEMIAdditionalInformationBusmonitorInfoParseWithBuffer(readBuffer utils.ReadBuffer) (CEMIAdditionalInformationBusmonitorInfo, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformationBusmonitorInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMIAdditionalInformationBusmonitorInfo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (len)
	len, _lenErr := readBuffer.ReadUint8("len", 8)
	if _lenErr != nil {
		return nil, errors.Wrap(_lenErr, "Error parsing 'len' field of CEMIAdditionalInformationBusmonitorInfo")
	}
	if len != CEMIAdditionalInformationBusmonitorInfo_LEN {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CEMIAdditionalInformationBusmonitorInfo_LEN) + " but got " + fmt.Sprintf("%d", len))
	}

	// Simple Field (frameErrorFlag)
	_frameErrorFlag, _frameErrorFlagErr := readBuffer.ReadBit("frameErrorFlag")
	if _frameErrorFlagErr != nil {
		return nil, errors.Wrap(_frameErrorFlagErr, "Error parsing 'frameErrorFlag' field of CEMIAdditionalInformationBusmonitorInfo")
	}
	frameErrorFlag := _frameErrorFlag

	// Simple Field (bitErrorFlag)
	_bitErrorFlag, _bitErrorFlagErr := readBuffer.ReadBit("bitErrorFlag")
	if _bitErrorFlagErr != nil {
		return nil, errors.Wrap(_bitErrorFlagErr, "Error parsing 'bitErrorFlag' field of CEMIAdditionalInformationBusmonitorInfo")
	}
	bitErrorFlag := _bitErrorFlag

	// Simple Field (parityErrorFlag)
	_parityErrorFlag, _parityErrorFlagErr := readBuffer.ReadBit("parityErrorFlag")
	if _parityErrorFlagErr != nil {
		return nil, errors.Wrap(_parityErrorFlagErr, "Error parsing 'parityErrorFlag' field of CEMIAdditionalInformationBusmonitorInfo")
	}
	parityErrorFlag := _parityErrorFlag

	// Simple Field (unknownFlag)
	_unknownFlag, _unknownFlagErr := readBuffer.ReadBit("unknownFlag")
	if _unknownFlagErr != nil {
		return nil, errors.Wrap(_unknownFlagErr, "Error parsing 'unknownFlag' field of CEMIAdditionalInformationBusmonitorInfo")
	}
	unknownFlag := _unknownFlag

	// Simple Field (lostFlag)
	_lostFlag, _lostFlagErr := readBuffer.ReadBit("lostFlag")
	if _lostFlagErr != nil {
		return nil, errors.Wrap(_lostFlagErr, "Error parsing 'lostFlag' field of CEMIAdditionalInformationBusmonitorInfo")
	}
	lostFlag := _lostFlag

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadUint8("sequenceNumber", 3)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of CEMIAdditionalInformationBusmonitorInfo")
	}
	sequenceNumber := _sequenceNumber

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformationBusmonitorInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMIAdditionalInformationBusmonitorInfo")
	}

	// Create a partially initialized instance
	_child := &_CEMIAdditionalInformationBusmonitorInfo{
		_CEMIAdditionalInformation: &_CEMIAdditionalInformation{},
		FrameErrorFlag:             frameErrorFlag,
		BitErrorFlag:               bitErrorFlag,
		ParityErrorFlag:            parityErrorFlag,
		UnknownFlag:                unknownFlag,
		LostFlag:                   lostFlag,
		SequenceNumber:             sequenceNumber,
	}
	_child._CEMIAdditionalInformation._CEMIAdditionalInformationChildRequirements = _child
	return _child, nil
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CEMIAdditionalInformationBusmonitorInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CEMIAdditionalInformationBusmonitorInfo")
		}

		// Const Field (len)
		_lenErr := writeBuffer.WriteUint8("len", 8, 1)
		if _lenErr != nil {
			return errors.Wrap(_lenErr, "Error serializing 'len' field")
		}

		// Simple Field (frameErrorFlag)
		frameErrorFlag := bool(m.GetFrameErrorFlag())
		_frameErrorFlagErr := writeBuffer.WriteBit("frameErrorFlag", (frameErrorFlag))
		if _frameErrorFlagErr != nil {
			return errors.Wrap(_frameErrorFlagErr, "Error serializing 'frameErrorFlag' field")
		}

		// Simple Field (bitErrorFlag)
		bitErrorFlag := bool(m.GetBitErrorFlag())
		_bitErrorFlagErr := writeBuffer.WriteBit("bitErrorFlag", (bitErrorFlag))
		if _bitErrorFlagErr != nil {
			return errors.Wrap(_bitErrorFlagErr, "Error serializing 'bitErrorFlag' field")
		}

		// Simple Field (parityErrorFlag)
		parityErrorFlag := bool(m.GetParityErrorFlag())
		_parityErrorFlagErr := writeBuffer.WriteBit("parityErrorFlag", (parityErrorFlag))
		if _parityErrorFlagErr != nil {
			return errors.Wrap(_parityErrorFlagErr, "Error serializing 'parityErrorFlag' field")
		}

		// Simple Field (unknownFlag)
		unknownFlag := bool(m.GetUnknownFlag())
		_unknownFlagErr := writeBuffer.WriteBit("unknownFlag", (unknownFlag))
		if _unknownFlagErr != nil {
			return errors.Wrap(_unknownFlagErr, "Error serializing 'unknownFlag' field")
		}

		// Simple Field (lostFlag)
		lostFlag := bool(m.GetLostFlag())
		_lostFlagErr := writeBuffer.WriteBit("lostFlag", (lostFlag))
		if _lostFlagErr != nil {
			return errors.Wrap(_lostFlagErr, "Error serializing 'lostFlag' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.GetSequenceNumber())
		_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 3, (sequenceNumber))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("CEMIAdditionalInformationBusmonitorInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CEMIAdditionalInformationBusmonitorInfo")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) isCEMIAdditionalInformationBusmonitorInfo() bool {
	return true
}

func (m *_CEMIAdditionalInformationBusmonitorInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
