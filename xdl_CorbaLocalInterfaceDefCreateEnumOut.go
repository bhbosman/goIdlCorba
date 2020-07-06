// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::LocalInterfaceDef_create_enum_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaLocalInterfaceDefCreateEnumOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaLocalInterfaceDefCreateEnumOutId_Const = "IDL:CORBA/LocalInterfaceDef_create_enum_Out:1.0"

func (self *CorbaLocalInterfaceDefCreateEnumOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaLocalInterfaceDefCreateEnumOut) GoString() string {
	return self.String()
}

func (self *CorbaLocalInterfaceDefCreateEnumOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefCreateEnumOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefCreateEnumOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaLocalInterfaceDefCreateEnumOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaLocalInterfaceDefCreateEnumOutHelper = CorbaLocalInterfaceDefCreateEnumOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaLocalInterfaceDefCreateEnumOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaLocalInterfaceDefCreateEnumOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefCreateEnumOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefCreateEnumOut{}
			},
			__reflect__.TypeOf((*CorbaLocalInterfaceDefCreateEnumOut)(nil))))
}
