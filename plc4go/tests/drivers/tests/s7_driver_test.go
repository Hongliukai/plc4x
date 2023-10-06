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

package tests

import (
	"context"
	"testing"

	"github.com/apache/plc4x/plc4go/internal/s7"
	s7IO "github.com/apache/plc4x/plc4go/protocols/s7/readwrite"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/s7/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/testutils"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

func TestS7Driver(t *testing.T) {
	parser := func(readBufferByteBased utils.ReadBufferByteBased) (any, error) {
		return readWriteModel.TPKTPacketParseWithBuffer(context.Background(), readBufferByteBased)
	}
	optionsForTesting := testutils.EnrichOptionsWithOptionsForTesting(t)
	testutils.RunDriverTestsuite(
		t,
		s7.NewDriver(optionsForTesting...),
		"assets/testing/protocols/s7/DriverTestsuite.xml",
		s7IO.S7XmlParserHelper{},
		append(optionsForTesting,
			testutils.WithRootTypeParser(parser),
			testutils.WithSkippedTestCases(
				// TODO: ignored due to carcia changes
				"Single element read request",
				"Single element read request with disabled PUT/GET",
			),
		)...,
	)
}
