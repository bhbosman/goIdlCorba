// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ModuleDef_create_exception_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRModuleDefCreateExceptionOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRModuleDefCreateExceptionOutId_Const = "IDL:CORBA/ComponentIR/ModuleDef_create_exception_Out:1.0"

func (self *CorbaComponentIRModuleDefCreateExceptionOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRModuleDefCreateExceptionOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRModuleDefCreateExceptionOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefCreateExceptionOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefCreateExceptionOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRModuleDefCreateExceptionOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRModuleDefCreateExceptionOutHelper = CorbaComponentIRModuleDefCreateExceptionOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRModuleDefCreateExceptionOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRModuleDefCreateExceptionOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefCreateExceptionOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefCreateExceptionOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRModuleDefCreateExceptionOut)(nil))))
}
