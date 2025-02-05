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

// MonitoredItemCreateRequest is the corresponding interface of MonitoredItemCreateRequest
type MonitoredItemCreateRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetItemToMonitor returns ItemToMonitor (property field)
	GetItemToMonitor() ExtensionObjectDefinition
	// GetMonitoringMode returns MonitoringMode (property field)
	GetMonitoringMode() MonitoringMode
	// GetRequestedParameters returns RequestedParameters (property field)
	GetRequestedParameters() ExtensionObjectDefinition
}

// MonitoredItemCreateRequestExactly can be used when we want exactly this type and not a type which fulfills MonitoredItemCreateRequest.
// This is useful for switch cases.
type MonitoredItemCreateRequestExactly interface {
	MonitoredItemCreateRequest
	isMonitoredItemCreateRequest() bool
}

// _MonitoredItemCreateRequest is the data-structure of this message
type _MonitoredItemCreateRequest struct {
	*_ExtensionObjectDefinition
	ItemToMonitor       ExtensionObjectDefinition
	MonitoringMode      MonitoringMode
	RequestedParameters ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemCreateRequest) GetIdentifier() string {
	return "745"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemCreateRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_MonitoredItemCreateRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemCreateRequest) GetItemToMonitor() ExtensionObjectDefinition {
	return m.ItemToMonitor
}

func (m *_MonitoredItemCreateRequest) GetMonitoringMode() MonitoringMode {
	return m.MonitoringMode
}

func (m *_MonitoredItemCreateRequest) GetRequestedParameters() ExtensionObjectDefinition {
	return m.RequestedParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredItemCreateRequest factory function for _MonitoredItemCreateRequest
func NewMonitoredItemCreateRequest(itemToMonitor ExtensionObjectDefinition, monitoringMode MonitoringMode, requestedParameters ExtensionObjectDefinition) *_MonitoredItemCreateRequest {
	_result := &_MonitoredItemCreateRequest{
		ItemToMonitor:              itemToMonitor,
		MonitoringMode:             monitoringMode,
		RequestedParameters:        requestedParameters,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredItemCreateRequest(structType any) MonitoredItemCreateRequest {
	if casted, ok := structType.(MonitoredItemCreateRequest); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemCreateRequest); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemCreateRequest) GetTypeName() string {
	return "MonitoredItemCreateRequest"
}

func (m *_MonitoredItemCreateRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (itemToMonitor)
	lengthInBits += m.ItemToMonitor.GetLengthInBits(ctx)

	// Simple field (monitoringMode)
	lengthInBits += 32

	// Simple field (requestedParameters)
	lengthInBits += m.RequestedParameters.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemCreateRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredItemCreateRequestParse(ctx context.Context, theBytes []byte, identifier string) (MonitoredItemCreateRequest, error) {
	return MonitoredItemCreateRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func MonitoredItemCreateRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (MonitoredItemCreateRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoredItemCreateRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemCreateRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (itemToMonitor)
	if pullErr := readBuffer.PullContext("itemToMonitor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for itemToMonitor")
	}
	_itemToMonitor, _itemToMonitorErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("628"))
	if _itemToMonitorErr != nil {
		return nil, errors.Wrap(_itemToMonitorErr, "Error parsing 'itemToMonitor' field of MonitoredItemCreateRequest")
	}
	itemToMonitor := _itemToMonitor.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("itemToMonitor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for itemToMonitor")
	}

	// Simple Field (monitoringMode)
	if pullErr := readBuffer.PullContext("monitoringMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoringMode")
	}
	_monitoringMode, _monitoringModeErr := MonitoringModeParseWithBuffer(ctx, readBuffer)
	if _monitoringModeErr != nil {
		return nil, errors.Wrap(_monitoringModeErr, "Error parsing 'monitoringMode' field of MonitoredItemCreateRequest")
	}
	monitoringMode := _monitoringMode
	if closeErr := readBuffer.CloseContext("monitoringMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoringMode")
	}

	// Simple Field (requestedParameters)
	if pullErr := readBuffer.PullContext("requestedParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestedParameters")
	}
	_requestedParameters, _requestedParametersErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("742"))
	if _requestedParametersErr != nil {
		return nil, errors.Wrap(_requestedParametersErr, "Error parsing 'requestedParameters' field of MonitoredItemCreateRequest")
	}
	requestedParameters := _requestedParameters.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestedParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestedParameters")
	}

	if closeErr := readBuffer.CloseContext("MonitoredItemCreateRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemCreateRequest")
	}

	// Create a partially initialized instance
	_child := &_MonitoredItemCreateRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ItemToMonitor:              itemToMonitor,
		MonitoringMode:             monitoringMode,
		RequestedParameters:        requestedParameters,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredItemCreateRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemCreateRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemCreateRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemCreateRequest")
		}

		// Simple Field (itemToMonitor)
		if pushErr := writeBuffer.PushContext("itemToMonitor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for itemToMonitor")
		}
		_itemToMonitorErr := writeBuffer.WriteSerializable(ctx, m.GetItemToMonitor())
		if popErr := writeBuffer.PopContext("itemToMonitor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for itemToMonitor")
		}
		if _itemToMonitorErr != nil {
			return errors.Wrap(_itemToMonitorErr, "Error serializing 'itemToMonitor' field")
		}

		// Simple Field (monitoringMode)
		if pushErr := writeBuffer.PushContext("monitoringMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for monitoringMode")
		}
		_monitoringModeErr := writeBuffer.WriteSerializable(ctx, m.GetMonitoringMode())
		if popErr := writeBuffer.PopContext("monitoringMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for monitoringMode")
		}
		if _monitoringModeErr != nil {
			return errors.Wrap(_monitoringModeErr, "Error serializing 'monitoringMode' field")
		}

		// Simple Field (requestedParameters)
		if pushErr := writeBuffer.PushContext("requestedParameters"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestedParameters")
		}
		_requestedParametersErr := writeBuffer.WriteSerializable(ctx, m.GetRequestedParameters())
		if popErr := writeBuffer.PopContext("requestedParameters"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestedParameters")
		}
		if _requestedParametersErr != nil {
			return errors.Wrap(_requestedParametersErr, "Error serializing 'requestedParameters' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemCreateRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemCreateRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemCreateRequest) isMonitoredItemCreateRequest() bool {
	return true
}

func (m *_MonitoredItemCreateRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
