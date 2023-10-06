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

// Code generated by "plc4xgenerator -type=SecureChannelTransactionManager"; DO NOT EDIT.

package opcua

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *SecureChannelTransactionManager) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *SecureChannelTransactionManager) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("SecureChannelTransactionManager"); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt32("transactionIdentifierGenerator", 32, d.transactionIdentifierGenerator.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt32("requestIdentifierGenerator", 32, d.requestIdentifierGenerator.Load()); err != nil {
		return err
	}

	if err := writeBuffer.WriteInt32("activeTransactionId", 32, d.activeTransactionId.Load()); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("queue", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _name, elem := range d.queue {
		name := fmt.Sprintf("%v", _name)
		_value := fmt.Sprintf("%v", elem)

		if err := writeBuffer.WriteString(name, uint32(len(_value)*8), "UTF-8", _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("queue", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("SecureChannelTransactionManager"); err != nil {
		return err
	}
	return nil
}

func (d *SecureChannelTransactionManager) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
