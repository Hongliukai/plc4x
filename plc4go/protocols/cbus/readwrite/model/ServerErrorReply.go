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

// Constant values.
const ServerErrorReply_ERRORMARKER byte = 0x21

// ServerErrorReply is the corresponding interface of ServerErrorReply
type ServerErrorReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ReplyOrConfirmation
}

// ServerErrorReplyExactly can be used when we want exactly this type and not a type which fulfills ServerErrorReply.
// This is useful for switch cases.
type ServerErrorReplyExactly interface {
	ServerErrorReply
	isServerErrorReply() bool
}

// _ServerErrorReply is the data-structure of this message
type _ServerErrorReply struct {
	*_ReplyOrConfirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServerErrorReply) InitializeParent(parent ReplyOrConfirmation, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ServerErrorReply) GetParent() ReplyOrConfirmation {
	return m._ReplyOrConfirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ServerErrorReply) GetErrorMarker() byte {
	return ServerErrorReply_ERRORMARKER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewServerErrorReply factory function for _ServerErrorReply
func NewServerErrorReply(peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ServerErrorReply {
	_result := &_ServerErrorReply{
		_ReplyOrConfirmation: NewReplyOrConfirmation(peekedByte, cBusOptions, requestContext),
	}
	_result._ReplyOrConfirmation._ReplyOrConfirmationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastServerErrorReply(structType any) ServerErrorReply {
	if casted, ok := structType.(ServerErrorReply); ok {
		return casted
	}
	if casted, ok := structType.(*ServerErrorReply); ok {
		return *casted
	}
	return nil
}

func (m *_ServerErrorReply) GetTypeName() string {
	return "ServerErrorReply"
}

func (m *_ServerErrorReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (errorMarker)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ServerErrorReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ServerErrorReplyParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (ServerErrorReply, error) {
	return ServerErrorReplyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func ServerErrorReplyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (ServerErrorReply, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ServerErrorReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServerErrorReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (errorMarker)
	errorMarker, _errorMarkerErr := readBuffer.ReadByte("errorMarker")
	if _errorMarkerErr != nil {
		return nil, errors.Wrap(_errorMarkerErr, "Error parsing 'errorMarker' field of ServerErrorReply")
	}
	if errorMarker != ServerErrorReply_ERRORMARKER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ServerErrorReply_ERRORMARKER) + " but got " + fmt.Sprintf("%d", errorMarker))
	}

	if closeErr := readBuffer.CloseContext("ServerErrorReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServerErrorReply")
	}

	// Create a partially initialized instance
	_child := &_ServerErrorReply{
		_ReplyOrConfirmation: &_ReplyOrConfirmation{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
	}
	_child._ReplyOrConfirmation._ReplyOrConfirmationChildRequirements = _child
	return _child, nil
}

func (m *_ServerErrorReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServerErrorReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServerErrorReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServerErrorReply")
		}

		// Const Field (errorMarker)
		_errorMarkerErr := writeBuffer.WriteByte("errorMarker", 0x21)
		if _errorMarkerErr != nil {
			return errors.Wrap(_errorMarkerErr, "Error serializing 'errorMarker' field")
		}

		if popErr := writeBuffer.PopContext("ServerErrorReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServerErrorReply")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServerErrorReply) isServerErrorReply() bool {
	return true
}

func (m *_ServerErrorReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
