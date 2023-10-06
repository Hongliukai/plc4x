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
package org.apache.plc4x.java.bacnetip.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum MaxSegmentsAccepted {
  UNSPECIFIED((byte) 0x0, (short) 255),
  NUM_SEGMENTS_02((byte) 0x1, (short) 2),
  NUM_SEGMENTS_04((byte) 0x2, (short) 4),
  NUM_SEGMENTS_08((byte) 0x3, (short) 8),
  NUM_SEGMENTS_16((byte) 0x4, (short) 16),
  NUM_SEGMENTS_32((byte) 0x5, (short) 32),
  NUM_SEGMENTS_64((byte) 0x6, (short) 64),
  MORE_THAN_64_SEGMENTS((byte) 0x7, (short) 255);
  private static final Map<Byte, MaxSegmentsAccepted> map;

  static {
    map = new HashMap<>();
    for (MaxSegmentsAccepted value : MaxSegmentsAccepted.values()) {
      map.put((byte) value.getValue(), value);
    }
  }

  private final byte value;
  private final short maxSegments;

  MaxSegmentsAccepted(byte value, short maxSegments) {
    this.value = value;
    this.maxSegments = maxSegments;
  }

  public byte getValue() {
    return value;
  }

  public short getMaxSegments() {
    return maxSegments;
  }

  public static MaxSegmentsAccepted firstEnumForFieldMaxSegments(short fieldValue) {
    for (MaxSegmentsAccepted _val : MaxSegmentsAccepted.values()) {
      if (_val.getMaxSegments() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<MaxSegmentsAccepted> enumsForFieldMaxSegments(short fieldValue) {
    List<MaxSegmentsAccepted> _values = new ArrayList<>();
    for (MaxSegmentsAccepted _val : MaxSegmentsAccepted.values()) {
      if (_val.getMaxSegments() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static MaxSegmentsAccepted enumForValue(byte value) {
    return map.get(value);
  }

  public static Boolean isDefined(byte value) {
    return map.containsKey(value);
  }
}
