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
pub struct AlarmMessageAckResponseTypeOptions {
}
#[derive(PartialEq, Debug, Clone)]
pub struct AlarmMessageAckResponseType {
    pub functionId: u8,
    pub numberOfObjects: u8,
    pub messageObjects: Vec<u8>
}

impl AlarmMessageAckResponseType {
}

impl Message for AlarmMessageAckResponseType {
    type M = AlarmMessageAckResponseType;
    type P = AlarmMessageAckResponseTypeOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        writer.write_u8(self.functionId)?;
        writer.write_u8(self.numberOfObjects)?;
        // not handled yet;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let functionId = reader.read_u8()?;
        let numberOfObjects = reader.read_u8()?;
        let messageObjects = vec![];
        let messageObjects_read = 0 as usize;
        // for _ in 0..(DefaultVariableLiteral{name='numberOfObjects', typeReference='AbstractSimpleTypeReference{baseType=UINT, sizeInBits=8}', args=null, index=null, child=null}) {
            // do something
        // }
        Ok(Self::M {
            functionId,
            numberOfObjects,
            messageObjects
        })
    }
}


