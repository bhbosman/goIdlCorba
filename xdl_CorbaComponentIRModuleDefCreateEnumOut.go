// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ModuleDef_create_enum_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRModuleDefCreateEnumOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRModuleDefCreateEnumOutId_Const = "IDL:CORBA/ComponentIR/ModuleDef_create_enum_Out:1.0"

func (self *CorbaComponentIRModuleDefCreateEnumOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRModuleDefCreateEnumOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRModuleDefCreateEnumOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefCreateEnumOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefCreateEnumOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRModuleDefCreateEnumOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRModuleDefCreateEnumOutHelper = CorbaComponentIRModuleDefCreateEnumOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRModuleDefCreateEnumOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRModuleDefCreateEnumOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefCreateEnumOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefCreateEnumOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRModuleDefCreateEnumOut)(nil))))
}
