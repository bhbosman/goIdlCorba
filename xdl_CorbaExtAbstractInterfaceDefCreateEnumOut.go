// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtAbstractInterfaceDef_create_enum_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtAbstractInterfaceDefCreateEnumOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtAbstractInterfaceDefCreateEnumOutId_Const = "IDL:CORBA/ExtAbstractInterfaceDef_create_enum_Out:1.0"

func (self *CorbaExtAbstractInterfaceDefCreateEnumOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtAbstractInterfaceDefCreateEnumOut) GoString() string {
	return self.String()
}

func (self *CorbaExtAbstractInterfaceDefCreateEnumOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefCreateEnumOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefCreateEnumOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtAbstractInterfaceDefCreateEnumOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtAbstractInterfaceDefCreateEnumOutHelper = CorbaExtAbstractInterfaceDefCreateEnumOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtAbstractInterfaceDefCreateEnumOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtAbstractInterfaceDefCreateEnumOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefCreateEnumOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefCreateEnumOut{}
			},
			__reflect__.TypeOf((*CorbaExtAbstractInterfaceDefCreateEnumOut)(nil))))
}
