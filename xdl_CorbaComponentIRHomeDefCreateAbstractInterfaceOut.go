// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::HomeDef_create_abstract_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRHomeDefCreateAbstractInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRHomeDefCreateAbstractInterfaceOutId_Const = "IDL:CORBA/ComponentIR/HomeDef_create_abstract_interface_Out:1.0"

func (self *CorbaComponentIRHomeDefCreateAbstractInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRHomeDefCreateAbstractInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRHomeDefCreateAbstractInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefCreateAbstractInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefCreateAbstractInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRHomeDefCreateAbstractInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRHomeDefCreateAbstractInterfaceOutHelper = CorbaComponentIRHomeDefCreateAbstractInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRHomeDefCreateAbstractInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRHomeDefCreateAbstractInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefCreateAbstractInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefCreateAbstractInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRHomeDefCreateAbstractInterfaceOut)(nil))))
}