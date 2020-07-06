// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_TypeCode_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadTypeCodeIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadTypeCodeInId_Const = "IDL:CORBA/DataInputStream_read_TypeCode_In:1.0"

func (self *CorbaDataInputStreamReadTypeCodeIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadTypeCodeIn) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadTypeCodeIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadTypeCodeIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadTypeCodeIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadTypeCodeIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadTypeCodeInHelper = CorbaDataInputStreamReadTypeCodeIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadTypeCodeInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadTypeCodeIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadTypeCodeIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadTypeCodeIn{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadTypeCodeIn)(nil))))
}