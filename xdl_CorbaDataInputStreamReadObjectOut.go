// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_Object_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadObjectOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadObjectOutId_Const = "IDL:CORBA/DataInputStream_read_Object_Out:1.0"

func (self *CorbaDataInputStreamReadObjectOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadObjectOut) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadObjectOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadObjectOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadObjectOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadObjectOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadObjectOutHelper = CorbaDataInputStreamReadObjectOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadObjectOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadObjectOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadObjectOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadObjectOut{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadObjectOut)(nil))))
}
