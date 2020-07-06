// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_create_attribute_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefCreateAttributeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefCreateAttributeOutId_Const = "IDL:CORBA/InterfaceDef_create_attribute_Out:1.0"

func (self *CorbaInterfaceDefCreateAttributeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefCreateAttributeOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefCreateAttributeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateAttributeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateAttributeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefCreateAttributeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefCreateAttributeOutHelper = CorbaInterfaceDefCreateAttributeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefCreateAttributeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefCreateAttributeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateAttributeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateAttributeOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefCreateAttributeOut)(nil))))
}
