// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ModuleDef_create_module_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaModuleDefCreateModuleOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaModuleDefCreateModuleOutId_Const = "IDL:CORBA/ModuleDef_create_module_Out:1.0"

func (self *CorbaModuleDefCreateModuleOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaModuleDefCreateModuleOut) GoString() string {
	return self.String()
}

func (self *CorbaModuleDefCreateModuleOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateModuleOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateModuleOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaModuleDefCreateModuleOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaModuleDefCreateModuleOutHelper = CorbaModuleDefCreateModuleOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaModuleDefCreateModuleOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaModuleDefCreateModuleOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaModuleDefCreateModuleOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaModuleDefCreateModuleOut{}
			},
			__reflect__.TypeOf((*CorbaModuleDefCreateModuleOut)(nil))))
}
