// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtAbstractInterfaceDef_create_abstract_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOutId_Const = "IDL:CORBA/ExtAbstractInterfaceDef_create_abstract_interface_Out:1.0"

func (self *CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOutHelper = CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaExtAbstractInterfaceDefCreateAbstractInterfaceOut)(nil))))
}
