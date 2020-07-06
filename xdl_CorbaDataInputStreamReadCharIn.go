// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_char_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadCharIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadCharInId_Const = "IDL:CORBA/DataInputStream_read_char_In:1.0"

func (self *CorbaDataInputStreamReadCharIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadCharIn) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadCharIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadCharIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadCharIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadCharIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadCharInHelper = CorbaDataInputStreamReadCharIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadCharInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadCharIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadCharIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadCharIn{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadCharIn)(nil))))
}
