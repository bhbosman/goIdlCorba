// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Repository_create_module_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRRepositoryCreateModuleOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRRepositoryCreateModuleOutId_Const = "IDL:CORBA/ComponentIR/Repository_create_module_Out:1.0"

func (self *CorbaComponentIRRepositoryCreateModuleOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRRepositoryCreateModuleOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRRepositoryCreateModuleOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateModuleOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateModuleOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRRepositoryCreateModuleOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRRepositoryCreateModuleOutHelper = CorbaComponentIRRepositoryCreateModuleOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRRepositoryCreateModuleOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRRepositoryCreateModuleOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateModuleOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateModuleOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRRepositoryCreateModuleOut)(nil))))
}
