// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_float_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadFloatOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadFloatOutId_Const = "IDL:CORBA/DataInputStream_read_float_Out:1.0"

func (self *CorbaDataInputStreamReadFloatOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadFloatOut) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadFloatOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadFloatOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadFloatOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadFloatOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadFloatOutHelper = CorbaDataInputStreamReadFloatOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadFloatOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadFloatOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadFloatOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadFloatOut{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadFloatOut)(nil))))
}
