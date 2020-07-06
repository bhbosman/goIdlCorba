// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExceptionDef_create_constant_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExceptionDefCreateConstantOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExceptionDefCreateConstantOutId_Const = "IDL:CORBA/ExceptionDef_create_constant_Out:1.0"

func (self *CorbaExceptionDefCreateConstantOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExceptionDefCreateConstantOut) GoString() string {
	return self.String()
}

func (self *CorbaExceptionDefCreateConstantOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefCreateConstantOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefCreateConstantOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExceptionDefCreateConstantOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExceptionDefCreateConstantOutHelper = CorbaExceptionDefCreateConstantOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExceptionDefCreateConstantOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExceptionDefCreateConstantOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExceptionDefCreateConstantOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExceptionDefCreateConstantOut{}
			},
			__reflect__.TypeOf((*CorbaExceptionDefCreateConstantOut)(nil))))
}