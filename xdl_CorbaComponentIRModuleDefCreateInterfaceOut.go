// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ModuleDef_create_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRModuleDefCreateInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRModuleDefCreateInterfaceOutId_Const = "IDL:CORBA/ComponentIR/ModuleDef_create_interface_Out:1.0"

func (self *CorbaComponentIRModuleDefCreateInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRModuleDefCreateInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRModuleDefCreateInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefCreateInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefCreateInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRModuleDefCreateInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRModuleDefCreateInterfaceOutHelper = CorbaComponentIRModuleDefCreateInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRModuleDefCreateInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRModuleDefCreateInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefCreateInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefCreateInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRModuleDefCreateInterfaceOut)(nil))))
}