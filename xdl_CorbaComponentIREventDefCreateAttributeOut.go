// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::EventDef_create_attribute_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIREventDefCreateAttributeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIREventDefCreateAttributeOutId_Const = "IDL:CORBA/ComponentIR/EventDef_create_attribute_Out:1.0"

func (self *CorbaComponentIREventDefCreateAttributeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIREventDefCreateAttributeOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIREventDefCreateAttributeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateAttributeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateAttributeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIREventDefCreateAttributeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIREventDefCreateAttributeOutHelper = CorbaComponentIREventDefCreateAttributeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIREventDefCreateAttributeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIREventDefCreateAttributeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateAttributeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateAttributeOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIREventDefCreateAttributeOut)(nil))))
}
