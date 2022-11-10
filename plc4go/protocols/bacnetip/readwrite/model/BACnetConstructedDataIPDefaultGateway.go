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

// BACnetConstructedDataIPDefaultGateway is the corresponding interface of BACnetConstructedDataIPDefaultGateway
type BACnetConstructedDataIPDefaultGateway interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpDefaultGateway returns IpDefaultGateway (property field)
	GetIpDefaultGateway() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataIPDefaultGatewayExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPDefaultGateway.
// This is useful for switch cases.
type BACnetConstructedDataIPDefaultGatewayExactly interface {
	BACnetConstructedDataIPDefaultGateway
	isBACnetConstructedDataIPDefaultGateway() bool
}

// _BACnetConstructedDataIPDefaultGateway is the data-structure of this message
type _BACnetConstructedDataIPDefaultGateway struct {
	*_BACnetConstructedData
	IpDefaultGateway BACnetApplicationTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPDefaultGateway) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPDefaultGateway) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DEFAULT_GATEWAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPDefaultGateway) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPDefaultGateway) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPDefaultGateway) GetIpDefaultGateway() BACnetApplicationTagOctetString {
	return m.IpDefaultGateway
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPDefaultGateway) GetActualValue() BACnetApplicationTagOctetString {
	return CastBACnetApplicationTagOctetString(m.GetIpDefaultGateway())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPDefaultGateway factory function for _BACnetConstructedDataIPDefaultGateway
func NewBACnetConstructedDataIPDefaultGateway(ipDefaultGateway BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPDefaultGateway {
	_result := &_BACnetConstructedDataIPDefaultGateway{
		IpDefaultGateway:       ipDefaultGateway,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPDefaultGateway(structType interface{}) BACnetConstructedDataIPDefaultGateway {
	if casted, ok := structType.(BACnetConstructedDataIPDefaultGateway); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDefaultGateway); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPDefaultGateway) GetTypeName() string {
	return "BACnetConstructedDataIPDefaultGateway"
}

func (m *_BACnetConstructedDataIPDefaultGateway) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIPDefaultGateway) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipDefaultGateway)
	lengthInBits += m.IpDefaultGateway.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPDefaultGateway) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPDefaultGatewayParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPDefaultGateway, error) {
	return BACnetConstructedDataIPDefaultGatewayParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIPDefaultGatewayParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPDefaultGateway, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDefaultGateway"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDefaultGateway")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipDefaultGateway)
	if pullErr := readBuffer.PullContext("ipDefaultGateway"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipDefaultGateway")
	}
	_ipDefaultGateway, _ipDefaultGatewayErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _ipDefaultGatewayErr != nil {
		return nil, errors.Wrap(_ipDefaultGatewayErr, "Error parsing 'ipDefaultGateway' field of BACnetConstructedDataIPDefaultGateway")
	}
	ipDefaultGateway := _ipDefaultGateway.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("ipDefaultGateway"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipDefaultGateway")
	}

	// Virtual field
	_actualValue := ipDefaultGateway
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDefaultGateway"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDefaultGateway")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPDefaultGateway{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IpDefaultGateway: ipDefaultGateway,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPDefaultGateway) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPDefaultGateway) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDefaultGateway"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDefaultGateway")
		}

		// Simple Field (ipDefaultGateway)
		if pushErr := writeBuffer.PushContext("ipDefaultGateway"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipDefaultGateway")
		}
		_ipDefaultGatewayErr := writeBuffer.WriteSerializable(m.GetIpDefaultGateway())
		if popErr := writeBuffer.PopContext("ipDefaultGateway"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipDefaultGateway")
		}
		if _ipDefaultGatewayErr != nil {
			return errors.Wrap(_ipDefaultGatewayErr, "Error serializing 'ipDefaultGateway' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDefaultGateway"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDefaultGateway")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPDefaultGateway) isBACnetConstructedDataIPDefaultGateway() bool {
	return true
}

func (m *_BACnetConstructedDataIPDefaultGateway) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
