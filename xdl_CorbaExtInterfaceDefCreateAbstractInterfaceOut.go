// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtInterfaceDef_create_abstract_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtInterfaceDefCreateAbstractInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtInterfaceDefCreateAbstractInterfaceOutId_Const = "IDL:CORBA/ExtInterfaceDef_create_abstract_interface_Out:1.0"

func (self *CorbaExtInterfaceDefCreateAbstractInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtInterfaceDefCreateAbstractInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaExtInterfaceDefCreateAbstractInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtInterfaceDefCreateAbstractInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtInterfaceDefCreateAbstractInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtInterfaceDefCreateAbstractInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtInterfaceDefCreateAbstractInterfaceOutHelper = CorbaExtInterfaceDefCreateAbstractInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtInterfaceDefCreateAbstractInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtInterfaceDefCreateAbstractInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtInterfaceDefCreateAbstractInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtInterfaceDefCreateAbstractInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaExtInterfaceDefCreateAbstractInterfaceOut)(nil))))
}
