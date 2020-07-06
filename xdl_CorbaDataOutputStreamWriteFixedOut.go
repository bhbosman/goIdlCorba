// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataOutputStream_write_fixed_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataOutputStreamWriteFixedOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataOutputStreamWriteFixedOutId_Const = "IDL:CORBA/DataOutputStream_write_fixed_Out:1.0"

func (self *CorbaDataOutputStreamWriteFixedOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataOutputStreamWriteFixedOut) GoString() string {
	return self.String()
}

func (self *CorbaDataOutputStreamWriteFixedOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteFixedOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteFixedOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataOutputStreamWriteFixedOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataOutputStreamWriteFixedOutHelper = CorbaDataOutputStreamWriteFixedOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataOutputStreamWriteFixedOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataOutputStreamWriteFixedOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteFixedOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteFixedOut{}
			},
			__reflect__.TypeOf((*CorbaDataOutputStreamWriteFixedOut)(nil))))
}