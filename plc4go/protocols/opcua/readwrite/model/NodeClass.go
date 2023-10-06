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

// NodeClass is an enum
type NodeClass uint32

type INodeClass interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	NodeClass_nodeClassUnspecified   NodeClass = 0
	NodeClass_nodeClassObject        NodeClass = 1
	NodeClass_nodeClassVariable      NodeClass = 2
	NodeClass_nodeClassMethod        NodeClass = 4
	NodeClass_nodeClassObjectType    NodeClass = 8
	NodeClass_nodeClassVariableType  NodeClass = 16
	NodeClass_nodeClassReferenceType NodeClass = 32
	NodeClass_nodeClassDataType      NodeClass = 64
	NodeClass_nodeClassView          NodeClass = 128
)

var NodeClassValues []NodeClass

func init() {
	_ = errors.New
	NodeClassValues = []NodeClass{
		NodeClass_nodeClassUnspecified,
		NodeClass_nodeClassObject,
		NodeClass_nodeClassVariable,
		NodeClass_nodeClassMethod,
		NodeClass_nodeClassObjectType,
		NodeClass_nodeClassVariableType,
		NodeClass_nodeClassReferenceType,
		NodeClass_nodeClassDataType,
		NodeClass_nodeClassView,
	}
}

func NodeClassByValue(value uint32) (enum NodeClass, ok bool) {
	switch value {
	case 0:
		return NodeClass_nodeClassUnspecified, true
	case 1:
		return NodeClass_nodeClassObject, true
	case 128:
		return NodeClass_nodeClassView, true
	case 16:
		return NodeClass_nodeClassVariableType, true
	case 2:
		return NodeClass_nodeClassVariable, true
	case 32:
		return NodeClass_nodeClassReferenceType, true
	case 4:
		return NodeClass_nodeClassMethod, true
	case 64:
		return NodeClass_nodeClassDataType, true
	case 8:
		return NodeClass_nodeClassObjectType, true
	}
	return 0, false
}

func NodeClassByName(value string) (enum NodeClass, ok bool) {
	switch value {
	case "nodeClassUnspecified":
		return NodeClass_nodeClassUnspecified, true
	case "nodeClassObject":
		return NodeClass_nodeClassObject, true
	case "nodeClassView":
		return NodeClass_nodeClassView, true
	case "nodeClassVariableType":
		return NodeClass_nodeClassVariableType, true
	case "nodeClassVariable":
		return NodeClass_nodeClassVariable, true
	case "nodeClassReferenceType":
		return NodeClass_nodeClassReferenceType, true
	case "nodeClassMethod":
		return NodeClass_nodeClassMethod, true
	case "nodeClassDataType":
		return NodeClass_nodeClassDataType, true
	case "nodeClassObjectType":
		return NodeClass_nodeClassObjectType, true
	}
	return 0, false
}

func NodeClassKnows(value uint32) bool {
	for _, typeValue := range NodeClassValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastNodeClass(structType any) NodeClass {
	castFunc := func(typ any) NodeClass {
		if sNodeClass, ok := typ.(NodeClass); ok {
			return sNodeClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m NodeClass) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m NodeClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeClassParse(ctx context.Context, theBytes []byte) (NodeClass, error) {
	return NodeClassParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeClassParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeClass, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("NodeClass", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading NodeClass")
	}
	if enum, ok := NodeClassByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for NodeClass")
		return NodeClass(val), nil
	} else {
		return enum, nil
	}
}

func (e NodeClass) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e NodeClass) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("NodeClass", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e NodeClass) PLC4XEnumName() string {
	switch e {
	case NodeClass_nodeClassUnspecified:
		return "nodeClassUnspecified"
	case NodeClass_nodeClassObject:
		return "nodeClassObject"
	case NodeClass_nodeClassView:
		return "nodeClassView"
	case NodeClass_nodeClassVariableType:
		return "nodeClassVariableType"
	case NodeClass_nodeClassVariable:
		return "nodeClassVariable"
	case NodeClass_nodeClassReferenceType:
		return "nodeClassReferenceType"
	case NodeClass_nodeClassMethod:
		return "nodeClassMethod"
	case NodeClass_nodeClassDataType:
		return "nodeClassDataType"
	case NodeClass_nodeClassObjectType:
		return "nodeClassObjectType"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e NodeClass) String() string {
	return e.PLC4XEnumName()
}
