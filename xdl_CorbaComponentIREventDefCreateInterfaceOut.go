// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::EventDef_create_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIREventDefCreateInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIREventDefCreateInterfaceOutId_Const = "IDL:CORBA/ComponentIR/EventDef_create_interface_Out:1.0"

func (self *CorbaComponentIREventDefCreateInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIREventDefCreateInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIREventDefCreateInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIREventDefCreateInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIREventDefCreateInterfaceOutHelper = CorbaComponentIREventDefCreateInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIREventDefCreateInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIREventDefCreateInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIREventDefCreateInterfaceOut)(nil))))
}
