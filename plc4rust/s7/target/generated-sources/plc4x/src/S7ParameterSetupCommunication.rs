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


#[derive(PartialEq, Debug, Clone)]
pub struct S7ParameterSetupCommunicationOptions {
    pub messageType: u8
}
#[derive(PartialEq, Debug, Clone)]
pub struct S7ParameterSetupCommunication {
        // -> DefaultReservedField{referenceValue=0x00} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
    pub maxAmqCaller: u16,
    pub maxAmqCallee: u16,
    pub pduLength: u16
}

impl S7ParameterSetupCommunication {
}

impl Message for S7ParameterSetupCommunication {
    type M = S7ParameterSetupCommunication;
    type P = S7ParameterSetupCommunicationOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        ---> DefaultReservedField{referenceValue=0x00} DefaultTypedField{type=AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}} DefaultField{attributes={}}
        writer.write_u16(self.maxAmqCaller)?;
        writer.write_u16(self.maxAmqCallee)?;
        writer.write_u16(self.pduLength)?;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let messageType = parameter.messageType;
        // Reserved field
        let _ = reader.read_u8()?;
        let maxAmqCaller = reader.read_u16()?;
        let maxAmqCallee = reader.read_u16()?;
        let pduLength = reader.read_u16()?;
        Ok(Self::M {
            maxAmqCaller,
            maxAmqCallee,
            pduLength
        })
    }
}


