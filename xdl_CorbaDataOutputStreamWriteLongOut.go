// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataOutputStream_write_long_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataOutputStreamWriteLongOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataOutputStreamWriteLongOutId_Const = "IDL:CORBA/DataOutputStream_write_long_Out:1.0"

func (self *CorbaDataOutputStreamWriteLongOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataOutputStreamWriteLongOut) GoString() string {
	return self.String()
}

func (self *CorbaDataOutputStreamWriteLongOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteLongOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteLongOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataOutputStreamWriteLongOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataOutputStreamWriteLongOutHelper = CorbaDataOutputStreamWriteLongOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataOutputStreamWriteLongOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataOutputStreamWriteLongOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteLongOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteLongOut{}
			},
			__reflect__.TypeOf((*CorbaDataOutputStreamWriteLongOut)(nil))))
}
