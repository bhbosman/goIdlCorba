// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_create_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefCreateInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefCreateInterfaceOutId_Const = "IDL:CORBA/InterfaceDef_create_interface_Out:1.0"

func (self *CorbaInterfaceDefCreateInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefCreateInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefCreateInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefCreateInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefCreateInterfaceOutHelper = CorbaInterfaceDefCreateInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefCreateInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefCreateInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefCreateInterfaceOut)(nil))))
}