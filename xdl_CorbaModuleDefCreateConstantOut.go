// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ModuleDef_create_constant_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaModuleDefCreateConstantOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaModuleDefCreateConstantOutId_Const = "IDL:CORBA/ModuleDef_create_constant_Out:1.0"

func (self *CorbaModuleDefCreateConstantOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaModuleDefCreateConstantOut) GoString() string {
	return self.String()
}

func (self *CorbaModuleDefCreateConstantOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateConstantOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateConstantOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaModuleDefCreateConstantOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaModuleDefCreateConstantOutHelper = CorbaModuleDefCreateConstantOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaModuleDefCreateConstantOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaModuleDefCreateConstantOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaModuleDefCreateConstantOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaModuleDefCreateConstantOut{}
			},
			__reflect__.TypeOf((*CorbaModuleDefCreateConstantOut)(nil))))
}
