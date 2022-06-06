/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
// package org.apache.plc4x.rust.s7.readwrite;

// Code generated by code-generation. DO NOT EDIT.
use std::io::{Error, ErrorKind, Read, Write};
use plc4rust::{Message, NoOption};
use plc4rust::read_buffer::ReadBuffer;
use plc4rust::write_buffer::WriteBuffer;

use crate::SyntaxIdType::SyntaxIdType;
use crate::QueryType::QueryType;
use crate::AlarmType::AlarmType;

#[derive(PartialEq, Debug, Clone)]
pub struct S7PayloadUserDataItemCpuFunctionAlarmQueryOptions {
    pub cpuFunctionType: u8, 
    pub cpuSubfunction: u8
}
#[derive(PartialEq, Debug, Clone)]
pub struct S7PayloadUserDataItemCpuFunctionAlarmQuery {
    // Intentionally do nothing
    // Intentionally do nothing
    // Intentionally do nothing
    // Intentionally do nothing
    pub syntaxId: SyntaxIdType,
        // -> DefaultReservedField{referenceValue=0x00} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
    pub queryType: QueryType,
        // -> DefaultReservedField{referenceValue=0x34} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
    pub alarmType: AlarmType
}

impl S7PayloadUserDataItemCpuFunctionAlarmQuery {
}

impl Message for S7PayloadUserDataItemCpuFunctionAlarmQuery {
    type M = S7PayloadUserDataItemCpuFunctionAlarmQuery;
    type P = S7PayloadUserDataItemCpuFunctionAlarmQueryOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        writer.write_u8(0x00)?;
        writer.write_u8(0x01)?;
        writer.write_u8(0x12)?;
        writer.write_u8(0x08)?;
        self.syntaxId.serialize(writer)?;
        ---> DefaultReservedField{referenceValue=0x00} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
        self.queryType.serialize(writer)?;
        ---> DefaultReservedField{referenceValue=0x34} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
        self.alarmType.serialize(writer)?;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let cpuFunctionType = parameter.cpuFunctionType;
        let cpuSubfunction = parameter.cpuSubfunction;
        let functionId = reader.read_u8()?;
        // assert value of constant
        assert_eq!(0x00, functionId);
        let numberMessageObj = reader.read_u8()?;
        // assert value of constant
        assert_eq!(0x01, numberMessageObj);
        let variableSpec = reader.read_u8()?;
        // assert value of constant
        assert_eq!(0x12, variableSpec);
        let length = reader.read_u8()?;
        // assert value of constant
        assert_eq!(0x08, length);
        let syntaxId = SyntaxIdType::parse(reader, None)?;
        // Reserved field
        let _ = reader.read_u8()?;
        let queryType = QueryType::parse(reader, None)?;
        // Reserved field
        let _ = reader.read_u8()?;
        let alarmType = AlarmType::parse(reader, None)?;
        Ok(Self::M {
            syntaxId,
            queryType,
            alarmType
        })
    }
}


