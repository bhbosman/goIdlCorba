// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::AbstractInterfaceDef_create_attribute_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaAbstractInterfaceDefCreateAttributeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaAbstractInterfaceDefCreateAttributeOutId_Const = "IDL:CORBA/AbstractInterfaceDef_create_attribute_Out:1.0"

func (self *CorbaAbstractInterfaceDefCreateAttributeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaAbstractInterfaceDefCreateAttributeOut) GoString() string {
	return self.String()
}

func (self *CorbaAbstractInterfaceDefCreateAttributeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaAbstractInterfaceDefCreateAttributeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaAbstractInterfaceDefCreateAttributeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaAbstractInterfaceDefCreateAttributeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaAbstractInterfaceDefCreateAttributeOutHelper = CorbaAbstractInterfaceDefCreateAttributeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaAbstractInterfaceDefCreateAttributeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaAbstractInterfaceDefCreateAttributeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaAbstractInterfaceDefCreateAttributeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaAbstractInterfaceDefCreateAttributeOut{}
			},
			__reflect__.TypeOf((*CorbaAbstractInterfaceDefCreateAttributeOut)(nil))))
}
