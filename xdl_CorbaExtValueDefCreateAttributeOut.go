// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtValueDef_create_attribute_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtValueDefCreateAttributeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtValueDefCreateAttributeOutId_Const = "IDL:CORBA/ExtValueDef_create_attribute_Out:1.0"

func (self *CorbaExtValueDefCreateAttributeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtValueDefCreateAttributeOut) GoString() string {
	return self.String()
}

func (self *CorbaExtValueDefCreateAttributeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefCreateAttributeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefCreateAttributeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtValueDefCreateAttributeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtValueDefCreateAttributeOutHelper = CorbaExtValueDefCreateAttributeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtValueDefCreateAttributeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtValueDefCreateAttributeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtValueDefCreateAttributeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtValueDefCreateAttributeOut{}
			},
			__reflect__.TypeOf((*CorbaExtValueDefCreateAttributeOut)(nil))))
}
