// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtInterfaceDef_create_module_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtInterfaceDefCreateModuleOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtInterfaceDefCreateModuleOutId_Const = "IDL:CORBA/ExtInterfaceDef_create_module_Out:1.0"

func (self *CorbaExtInterfaceDefCreateModuleOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtInterfaceDefCreateModuleOut) GoString() string {
	return self.String()
}

func (self *CorbaExtInterfaceDefCreateModuleOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtInterfaceDefCreateModuleOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtInterfaceDefCreateModuleOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtInterfaceDefCreateModuleOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtInterfaceDefCreateModuleOutHelper = CorbaExtInterfaceDefCreateModuleOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtInterfaceDefCreateModuleOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtInterfaceDefCreateModuleOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtInterfaceDefCreateModuleOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtInterfaceDefCreateModuleOut{}
			},
			__reflect__.TypeOf((*CorbaExtInterfaceDefCreateModuleOut)(nil))))
}
