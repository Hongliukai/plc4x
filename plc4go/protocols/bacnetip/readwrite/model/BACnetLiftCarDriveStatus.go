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

// BACnetLiftCarDriveStatus is an enum
type BACnetLiftCarDriveStatus uint16

type IBACnetLiftCarDriveStatus interface {
	utils.Serializable
}

const (
	BACnetLiftCarDriveStatus_UNKNOWN                  BACnetLiftCarDriveStatus = 0
	BACnetLiftCarDriveStatus_STATIONARY               BACnetLiftCarDriveStatus = 1
	BACnetLiftCarDriveStatus_BRAKING                  BACnetLiftCarDriveStatus = 2
	BACnetLiftCarDriveStatus_ACCELERATE               BACnetLiftCarDriveStatus = 3
	BACnetLiftCarDriveStatus_DECELERATE               BACnetLiftCarDriveStatus = 4
	BACnetLiftCarDriveStatus_RATED_SPEED              BACnetLiftCarDriveStatus = 5
	BACnetLiftCarDriveStatus_SINGLE_FLOOR_JUMP        BACnetLiftCarDriveStatus = 6
	BACnetLiftCarDriveStatus_TWO_FLOOR_JUMP           BACnetLiftCarDriveStatus = 7
	BACnetLiftCarDriveStatus_THREE_FLOOR_JUMP         BACnetLiftCarDriveStatus = 8
	BACnetLiftCarDriveStatus_MULTI_FLOOR_JUMP         BACnetLiftCarDriveStatus = 9
	BACnetLiftCarDriveStatus_VENDOR_PROPRIETARY_VALUE BACnetLiftCarDriveStatus = 0xFFFF
)

var BACnetLiftCarDriveStatusValues []BACnetLiftCarDriveStatus

func init() {
	_ = errors.New
	BACnetLiftCarDriveStatusValues = []BACnetLiftCarDriveStatus{
		BACnetLiftCarDriveStatus_UNKNOWN,
		BACnetLiftCarDriveStatus_STATIONARY,
		BACnetLiftCarDriveStatus_BRAKING,
		BACnetLiftCarDriveStatus_ACCELERATE,
		BACnetLiftCarDriveStatus_DECELERATE,
		BACnetLiftCarDriveStatus_RATED_SPEED,
		BACnetLiftCarDriveStatus_SINGLE_FLOOR_JUMP,
		BACnetLiftCarDriveStatus_TWO_FLOOR_JUMP,
		BACnetLiftCarDriveStatus_THREE_FLOOR_JUMP,
		BACnetLiftCarDriveStatus_MULTI_FLOOR_JUMP,
		BACnetLiftCarDriveStatus_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLiftCarDriveStatusByValue(value uint16) (enum BACnetLiftCarDriveStatus, ok bool) {
	switch value {
	case 0:
		return BACnetLiftCarDriveStatus_UNKNOWN, true
	case 0xFFFF:
		return BACnetLiftCarDriveStatus_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetLiftCarDriveStatus_STATIONARY, true
	case 2:
		return BACnetLiftCarDriveStatus_BRAKING, true
	case 3:
		return BACnetLiftCarDriveStatus_ACCELERATE, true
	case 4:
		return BACnetLiftCarDriveStatus_DECELERATE, true
	case 5:
		return BACnetLiftCarDriveStatus_RATED_SPEED, true
	case 6:
		return BACnetLiftCarDriveStatus_SINGLE_FLOOR_JUMP, true
	case 7:
		return BACnetLiftCarDriveStatus_TWO_FLOOR_JUMP, true
	case 8:
		return BACnetLiftCarDriveStatus_THREE_FLOOR_JUMP, true
	case 9:
		return BACnetLiftCarDriveStatus_MULTI_FLOOR_JUMP, true
	}
	return 0, false
}

func BACnetLiftCarDriveStatusByName(value string) (enum BACnetLiftCarDriveStatus, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetLiftCarDriveStatus_UNKNOWN, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLiftCarDriveStatus_VENDOR_PROPRIETARY_VALUE, true
	case "STATIONARY":
		return BACnetLiftCarDriveStatus_STATIONARY, true
	case "BRAKING":
		return BACnetLiftCarDriveStatus_BRAKING, true
	case "ACCELERATE":
		return BACnetLiftCarDriveStatus_ACCELERATE, true
	case "DECELERATE":
		return BACnetLiftCarDriveStatus_DECELERATE, true
	case "RATED_SPEED":
		return BACnetLiftCarDriveStatus_RATED_SPEED, true
	case "SINGLE_FLOOR_JUMP":
		return BACnetLiftCarDriveStatus_SINGLE_FLOOR_JUMP, true
	case "TWO_FLOOR_JUMP":
		return BACnetLiftCarDriveStatus_TWO_FLOOR_JUMP, true
	case "THREE_FLOOR_JUMP":
		return BACnetLiftCarDriveStatus_THREE_FLOOR_JUMP, true
	case "MULTI_FLOOR_JUMP":
		return BACnetLiftCarDriveStatus_MULTI_FLOOR_JUMP, true
	}
	return 0, false
}

func BACnetLiftCarDriveStatusKnows(value uint16) bool {
	for _, typeValue := range BACnetLiftCarDriveStatusValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLiftCarDriveStatus(structType interface{}) BACnetLiftCarDriveStatus {
	castFunc := func(typ interface{}) BACnetLiftCarDriveStatus {
		if sBACnetLiftCarDriveStatus, ok := typ.(BACnetLiftCarDriveStatus); ok {
			return sBACnetLiftCarDriveStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLiftCarDriveStatus) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetLiftCarDriveStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLiftCarDriveStatusParse(theBytes []byte) (BACnetLiftCarDriveStatus, error) {
	return BACnetLiftCarDriveStatusParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetLiftCarDriveStatusParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetLiftCarDriveStatus, error) {
	val, err := readBuffer.ReadUint16("BACnetLiftCarDriveStatus", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLiftCarDriveStatus")
	}
	if enum, ok := BACnetLiftCarDriveStatusByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetLiftCarDriveStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLiftCarDriveStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLiftCarDriveStatus) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetLiftCarDriveStatus", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLiftCarDriveStatus) PLC4XEnumName() string {
	switch e {
	case BACnetLiftCarDriveStatus_UNKNOWN:
		return "UNKNOWN"
	case BACnetLiftCarDriveStatus_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLiftCarDriveStatus_STATIONARY:
		return "STATIONARY"
	case BACnetLiftCarDriveStatus_BRAKING:
		return "BRAKING"
	case BACnetLiftCarDriveStatus_ACCELERATE:
		return "ACCELERATE"
	case BACnetLiftCarDriveStatus_DECELERATE:
		return "DECELERATE"
	case BACnetLiftCarDriveStatus_RATED_SPEED:
		return "RATED_SPEED"
	case BACnetLiftCarDriveStatus_SINGLE_FLOOR_JUMP:
		return "SINGLE_FLOOR_JUMP"
	case BACnetLiftCarDriveStatus_TWO_FLOOR_JUMP:
		return "TWO_FLOOR_JUMP"
	case BACnetLiftCarDriveStatus_THREE_FLOOR_JUMP:
		return "THREE_FLOOR_JUMP"
	case BACnetLiftCarDriveStatus_MULTI_FLOOR_JUMP:
		return "MULTI_FLOOR_JUMP"
	}
	return ""
}

func (e BACnetLiftCarDriveStatus) String() string {
	return e.PLC4XEnumName()
}
