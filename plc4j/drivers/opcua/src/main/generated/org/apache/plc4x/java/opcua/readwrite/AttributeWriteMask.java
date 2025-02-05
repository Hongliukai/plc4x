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
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum AttributeWriteMask {
  attributeWriteMaskNone((long) 0L),
  attributeWriteMaskAccessLevel((long) 1L),
  attributeWriteMaskArrayDimensions((long) 2L),
  attributeWriteMaskBrowseName((long) 4L),
  attributeWriteMaskContainsNoLoops((long) 8L),
  attributeWriteMaskDataType((long) 16L),
  attributeWriteMaskDescription((long) 32L),
  attributeWriteMaskDisplayName((long) 64L),
  attributeWriteMaskEventNotifier((long) 128L),
  attributeWriteMaskExecutable((long) 256L),
  attributeWriteMaskHistorizing((long) 512L),
  attributeWriteMaskInverseName((long) 1024L),
  attributeWriteMaskIsAbstract((long) 2048L),
  attributeWriteMaskMinimumSamplingInterval((long) 4096L),
  attributeWriteMaskNodeClass((long) 8192L),
  attributeWriteMaskNodeId((long) 16384L),
  attributeWriteMaskSymmetric((long) 32768L),
  attributeWriteMaskUserAccessLevel((long) 65536L),
  attributeWriteMaskUserExecutable((long) 131072L),
  attributeWriteMaskUserWriteMask((long) 262144L),
  attributeWriteMaskValueRank((long) 524288L),
  attributeWriteMaskWriteMask((long) 1048576L),
  attributeWriteMaskValueForVariableType((long) 2097152L),
  attributeWriteMaskDataTypeDefinition((long) 4194304L),
  attributeWriteMaskRolePermissions((long) 8388608L),
  attributeWriteMaskAccessRestrictions((long) 16777216L),
  attributeWriteMaskAccessLevelEx((long) 33554432L);
  private static final Map<Long, AttributeWriteMask> map;

  static {
    map = new HashMap<>();
    for (AttributeWriteMask value : AttributeWriteMask.values()) {
      map.put((long) value.getValue(), value);
    }
  }

  private final long value;

  AttributeWriteMask(long value) {
    this.value = value;
  }

  public long getValue() {
    return value;
  }

  public static AttributeWriteMask enumForValue(long value) {
    return map.get(value);
  }

  public static Boolean isDefined(long value) {
    return map.containsKey(value);
  }
}
