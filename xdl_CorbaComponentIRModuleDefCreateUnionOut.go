// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ModuleDef_create_union_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRModuleDefCreateUnionOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRModuleDefCreateUnionOutId_Const = "IDL:CORBA/ComponentIR/ModuleDef_create_union_Out:1.0"

func (self *CorbaComponentIRModuleDefCreateUnionOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRModuleDefCreateUnionOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRModuleDefCreateUnionOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefCreateUnionOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefCreateUnionOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRModuleDefCreateUnionOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRModuleDefCreateUnionOutHelper = CorbaComponentIRModuleDefCreateUnionOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRModuleDefCreateUnionOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRModuleDefCreateUnionOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefCreateUnionOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefCreateUnionOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRModuleDefCreateUnionOut)(nil))))
}
