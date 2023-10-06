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

// UadpDataSetMessageContentMask is an enum
type UadpDataSetMessageContentMask uint32

type IUadpDataSetMessageContentMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskNone           UadpDataSetMessageContentMask = 0
	UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskTimestamp      UadpDataSetMessageContentMask = 1
	UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskPicoSeconds    UadpDataSetMessageContentMask = 2
	UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskStatus         UadpDataSetMessageContentMask = 4
	UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMajorVersion   UadpDataSetMessageContentMask = 8
	UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMinorVersion   UadpDataSetMessageContentMask = 16
	UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskSequenceNumber UadpDataSetMessageContentMask = 32
)

var UadpDataSetMessageContentMaskValues []UadpDataSetMessageContentMask

func init() {
	_ = errors.New
	UadpDataSetMessageContentMaskValues = []UadpDataSetMessageContentMask{
		UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskNone,
		UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskTimestamp,
		UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskPicoSeconds,
		UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskStatus,
		UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMajorVersion,
		UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMinorVersion,
		UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskSequenceNumber,
	}
}

func UadpDataSetMessageContentMaskByValue(value uint32) (enum UadpDataSetMessageContentMask, ok bool) {
	switch value {
	case 0:
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskNone, true
	case 1:
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskTimestamp, true
	case 16:
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMinorVersion, true
	case 2:
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskPicoSeconds, true
	case 32:
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskSequenceNumber, true
	case 4:
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskStatus, true
	case 8:
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMajorVersion, true
	}
	return 0, false
}

func UadpDataSetMessageContentMaskByName(value string) (enum UadpDataSetMessageContentMask, ok bool) {
	switch value {
	case "uadpDataSetMessageContentMaskNone":
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskNone, true
	case "uadpDataSetMessageContentMaskTimestamp":
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskTimestamp, true
	case "uadpDataSetMessageContentMaskMinorVersion":
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMinorVersion, true
	case "uadpDataSetMessageContentMaskPicoSeconds":
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskPicoSeconds, true
	case "uadpDataSetMessageContentMaskSequenceNumber":
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskSequenceNumber, true
	case "uadpDataSetMessageContentMaskStatus":
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskStatus, true
	case "uadpDataSetMessageContentMaskMajorVersion":
		return UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMajorVersion, true
	}
	return 0, false
}

func UadpDataSetMessageContentMaskKnows(value uint32) bool {
	for _, typeValue := range UadpDataSetMessageContentMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastUadpDataSetMessageContentMask(structType any) UadpDataSetMessageContentMask {
	castFunc := func(typ any) UadpDataSetMessageContentMask {
		if sUadpDataSetMessageContentMask, ok := typ.(UadpDataSetMessageContentMask); ok {
			return sUadpDataSetMessageContentMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m UadpDataSetMessageContentMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m UadpDataSetMessageContentMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UadpDataSetMessageContentMaskParse(ctx context.Context, theBytes []byte) (UadpDataSetMessageContentMask, error) {
	return UadpDataSetMessageContentMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func UadpDataSetMessageContentMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (UadpDataSetMessageContentMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("UadpDataSetMessageContentMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading UadpDataSetMessageContentMask")
	}
	if enum, ok := UadpDataSetMessageContentMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for UadpDataSetMessageContentMask")
		return UadpDataSetMessageContentMask(val), nil
	} else {
		return enum, nil
	}
}

func (e UadpDataSetMessageContentMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e UadpDataSetMessageContentMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("UadpDataSetMessageContentMask", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e UadpDataSetMessageContentMask) PLC4XEnumName() string {
	switch e {
	case UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskNone:
		return "uadpDataSetMessageContentMaskNone"
	case UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskTimestamp:
		return "uadpDataSetMessageContentMaskTimestamp"
	case UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMinorVersion:
		return "uadpDataSetMessageContentMaskMinorVersion"
	case UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskPicoSeconds:
		return "uadpDataSetMessageContentMaskPicoSeconds"
	case UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskSequenceNumber:
		return "uadpDataSetMessageContentMaskSequenceNumber"
	case UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskStatus:
		return "uadpDataSetMessageContentMaskStatus"
	case UadpDataSetMessageContentMask_uadpDataSetMessageContentMaskMajorVersion:
		return "uadpDataSetMessageContentMaskMajorVersion"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e UadpDataSetMessageContentMask) String() string {
	return e.PLC4XEnumName()
}
