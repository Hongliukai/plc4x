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

public enum OpcuaNodeIdServicesVariableRole {
  RoleSetType_RoleName_Placeholder_ApplicationsExclude((int) 15408L),
  RoleSetType_RoleName_Placeholder_EndpointsExclude((int) 15409L),
  RoleType_ApplicationsExclude((int) 15410L),
  RoleType_EndpointsExclude((int) 15411L),
  RoleSetType_RoleName_Placeholder_AddIdentity_InputArguments((int) 15613L),
  RoleSetType_RoleName_Placeholder_RemoveIdentity_InputArguments((int) 15615L),
  RoleType_AddIdentity_InputArguments((int) 15625L),
  RoleType_RemoveIdentity_InputArguments((int) 15627L),
  RoleSetType_AddRole_InputArguments((int) 15998L),
  RoleSetType_AddRole_OutputArguments((int) 15999L),
  RoleSetType_RemoveRole_InputArguments((int) 16001L),
  RoleSetType_RoleName_Placeholder_Identities((int) 16162L),
  RoleSetType_RoleName_Placeholder_Applications((int) 16163L),
  RoleSetType_RoleName_Placeholder_Endpoints((int) 16164L),
  RoleSetType_RoleName_Placeholder_AddApplication_InputArguments((int) 16166L),
  RoleSetType_RoleName_Placeholder_RemoveApplication_InputArguments((int) 16168L),
  RoleSetType_RoleName_Placeholder_AddEndpoint_InputArguments((int) 16170L),
  RoleSetType_RoleName_Placeholder_RemoveEndpoint_InputArguments((int) 16172L),
  RoleType_Identities((int) 16173L),
  RoleType_Applications((int) 16174L),
  RoleType_Endpoints((int) 16175L),
  RoleType_AddApplication_InputArguments((int) 16177L),
  RoleType_RemoveApplication_InputArguments((int) 16179L),
  RoleType_AddEndpoint_InputArguments((int) 16181L),
  RoleType_RemoveEndpoint_InputArguments((int) 16183L),
  RoleMappingRuleChangedAuditEventType_EventId((int) 17642L),
  RoleMappingRuleChangedAuditEventType_EventType((int) 17643L),
  RoleMappingRuleChangedAuditEventType_SourceNode((int) 17644L),
  RoleMappingRuleChangedAuditEventType_SourceName((int) 17645L),
  RoleMappingRuleChangedAuditEventType_Time((int) 17646L),
  RoleMappingRuleChangedAuditEventType_ReceiveTime((int) 17647L),
  RoleMappingRuleChangedAuditEventType_LocalTime((int) 17648L),
  RoleMappingRuleChangedAuditEventType_Message((int) 17649L),
  RoleMappingRuleChangedAuditEventType_Severity((int) 17650L),
  RoleMappingRuleChangedAuditEventType_ActionTimeStamp((int) 17651L),
  RoleMappingRuleChangedAuditEventType_Status((int) 17652L),
  RoleMappingRuleChangedAuditEventType_ServerId((int) 17653L),
  RoleMappingRuleChangedAuditEventType_ClientAuditEntryId((int) 17654L),
  RoleMappingRuleChangedAuditEventType_ClientUserId((int) 17655L),
  RoleMappingRuleChangedAuditEventType_MethodId((int) 17656L),
  RoleMappingRuleChangedAuditEventType_InputArguments((int) 17657L),
  RoleSetType_RoleName_Placeholder_CustomConfiguration((int) 24138L),
  RoleType_CustomConfiguration((int) 24139L),
  RoleMappingRuleChangedAuditEventType_ConditionClassId((int) 31927L),
  RoleMappingRuleChangedAuditEventType_ConditionClassName((int) 31928L),
  RoleMappingRuleChangedAuditEventType_ConditionSubClassId((int) 31929L),
  RoleMappingRuleChangedAuditEventType_ConditionSubClassName((int) 31930L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableRole> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableRole value : OpcuaNodeIdServicesVariableRole.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableRole(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableRole enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
