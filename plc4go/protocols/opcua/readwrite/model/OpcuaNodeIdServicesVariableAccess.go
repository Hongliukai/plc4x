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

// OpcuaNodeIdServicesVariableAccess is an enum
type OpcuaNodeIdServicesVariableAccess int32

type IOpcuaNodeIdServicesVariableAccess interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableAccess_AccessLevelType_OptionSetValues       OpcuaNodeIdServicesVariableAccess = 15032
	OpcuaNodeIdServicesVariableAccess_AccessRestrictionType_OptionSetValues OpcuaNodeIdServicesVariableAccess = 15035
	OpcuaNodeIdServicesVariableAccess_AccessLevelExType_OptionSetValues     OpcuaNodeIdServicesVariableAccess = 15407
)

var OpcuaNodeIdServicesVariableAccessValues []OpcuaNodeIdServicesVariableAccess

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableAccessValues = []OpcuaNodeIdServicesVariableAccess{
		OpcuaNodeIdServicesVariableAccess_AccessLevelType_OptionSetValues,
		OpcuaNodeIdServicesVariableAccess_AccessRestrictionType_OptionSetValues,
		OpcuaNodeIdServicesVariableAccess_AccessLevelExType_OptionSetValues,
	}
}

func OpcuaNodeIdServicesVariableAccessByValue(value int32) (enum OpcuaNodeIdServicesVariableAccess, ok bool) {
	switch value {
	case 15032:
		return OpcuaNodeIdServicesVariableAccess_AccessLevelType_OptionSetValues, true
	case 15035:
		return OpcuaNodeIdServicesVariableAccess_AccessRestrictionType_OptionSetValues, true
	case 15407:
		return OpcuaNodeIdServicesVariableAccess_AccessLevelExType_OptionSetValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAccessByName(value string) (enum OpcuaNodeIdServicesVariableAccess, ok bool) {
	switch value {
	case "AccessLevelType_OptionSetValues":
		return OpcuaNodeIdServicesVariableAccess_AccessLevelType_OptionSetValues, true
	case "AccessRestrictionType_OptionSetValues":
		return OpcuaNodeIdServicesVariableAccess_AccessRestrictionType_OptionSetValues, true
	case "AccessLevelExType_OptionSetValues":
		return OpcuaNodeIdServicesVariableAccess_AccessLevelExType_OptionSetValues, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAccessKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableAccessValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableAccess(structType any) OpcuaNodeIdServicesVariableAccess {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableAccess {
		if sOpcuaNodeIdServicesVariableAccess, ok := typ.(OpcuaNodeIdServicesVariableAccess); ok {
			return sOpcuaNodeIdServicesVariableAccess
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableAccess) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableAccess) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableAccessParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableAccess, error) {
	return OpcuaNodeIdServicesVariableAccessParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableAccessParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableAccess, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableAccess", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableAccess")
	}
	if enum, ok := OpcuaNodeIdServicesVariableAccessByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableAccess")
		return OpcuaNodeIdServicesVariableAccess(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableAccess) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableAccess) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableAccess", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableAccess) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableAccess_AccessLevelType_OptionSetValues:
		return "AccessLevelType_OptionSetValues"
	case OpcuaNodeIdServicesVariableAccess_AccessRestrictionType_OptionSetValues:
		return "AccessRestrictionType_OptionSetValues"
	case OpcuaNodeIdServicesVariableAccess_AccessLevelExType_OptionSetValues:
		return "AccessLevelExType_OptionSetValues"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableAccess) String() string {
	return e.PLC4XEnumName()
}
