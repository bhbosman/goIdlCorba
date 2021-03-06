// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtValueDef_create_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtValueDefCreateInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtValueDefCreateInterfaceOutId_Const = "IDL:CORBA/ExtValueDef_create_interface_Out:1.0"

func (self *CorbaExtValueDefCreateInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtValueDefCreateInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaExtValueDefCreateInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefCreateInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefCreateInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtValueDefCreateInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtValueDefCreateInterfaceOutHelper = CorbaExtValueDefCreateInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtValueDefCreateInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtValueDefCreateInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtValueDefCreateInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtValueDefCreateInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaExtValueDefCreateInterfaceOut)(nil))))
}
