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
// Code generated by "stringer -type TagType"; DO NOT EDIT.

package s7

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[S7Tag-0]
	_ = x[S7StringTag-1]
}

const _TagType_name = "S7TagS7StringTag"

var _TagType_index = [...]uint8{0, 5, 16}

func (i TagType) String() string {
	if i >= TagType(len(_TagType_index)-1) {
		return "TagType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TagType_name[_TagType_index[i]:_TagType_index[i+1]]
}
