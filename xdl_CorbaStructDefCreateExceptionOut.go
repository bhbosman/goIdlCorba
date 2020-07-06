// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::StructDef_create_exception_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaStructDefCreateExceptionOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaStructDefCreateExceptionOutId_Const = "IDL:CORBA/StructDef_create_exception_Out:1.0"

func (self *CorbaStructDefCreateExceptionOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaStructDefCreateExceptionOut) GoString() string {
	return self.String()
}

func (self *CorbaStructDefCreateExceptionOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaStructDefCreateExceptionOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaStructDefCreateExceptionOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaStructDefCreateExceptionOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaStructDefCreateExceptionOutHelper = CorbaStructDefCreateExceptionOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaStructDefCreateExceptionOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaStructDefCreateExceptionOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaStructDefCreateExceptionOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaStructDefCreateExceptionOut{}
			},
			__reflect__.TypeOf((*CorbaStructDefCreateExceptionOut)(nil))))
}