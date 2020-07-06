// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_fixed_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadFixedOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadFixedOutId_Const = "IDL:CORBA/DataInputStream_read_fixed_Out:1.0"

func (self *CorbaDataInputStreamReadFixedOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadFixedOut) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadFixedOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadFixedOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadFixedOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadFixedOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadFixedOutHelper = CorbaDataInputStreamReadFixedOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadFixedOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadFixedOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadFixedOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadFixedOut{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadFixedOut)(nil))))
}