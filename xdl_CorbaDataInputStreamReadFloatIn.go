// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_float_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadFloatIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadFloatInId_Const = "IDL:CORBA/DataInputStream_read_float_In:1.0"

func (self *CorbaDataInputStreamReadFloatIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadFloatIn) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadFloatIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadFloatIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadFloatIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadFloatIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadFloatInHelper = CorbaDataInputStreamReadFloatIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadFloatInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadFloatIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadFloatIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadFloatIn{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadFloatIn)(nil))))
}