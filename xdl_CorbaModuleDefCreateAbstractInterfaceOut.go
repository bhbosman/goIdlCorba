// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ModuleDef_create_abstract_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaModuleDefCreateAbstractInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaModuleDefCreateAbstractInterfaceOutId_Const = "IDL:CORBA/ModuleDef_create_abstract_interface_Out:1.0"

func (self *CorbaModuleDefCreateAbstractInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaModuleDefCreateAbstractInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaModuleDefCreateAbstractInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateAbstractInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateAbstractInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaModuleDefCreateAbstractInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaModuleDefCreateAbstractInterfaceOutHelper = CorbaModuleDefCreateAbstractInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaModuleDefCreateAbstractInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaModuleDefCreateAbstractInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaModuleDefCreateAbstractInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaModuleDefCreateAbstractInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaModuleDefCreateAbstractInterfaceOut)(nil))))
}