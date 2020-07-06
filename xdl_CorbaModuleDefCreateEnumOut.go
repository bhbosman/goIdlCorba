// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ModuleDef_create_enum_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaModuleDefCreateEnumOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaModuleDefCreateEnumOutId_Const = "IDL:CORBA/ModuleDef_create_enum_Out:1.0"

func (self *CorbaModuleDefCreateEnumOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaModuleDefCreateEnumOut) GoString() string {
	return self.String()
}

func (self *CorbaModuleDefCreateEnumOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateEnumOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateEnumOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaModuleDefCreateEnumOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaModuleDefCreateEnumOutHelper = CorbaModuleDefCreateEnumOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaModuleDefCreateEnumOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaModuleDefCreateEnumOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaModuleDefCreateEnumOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaModuleDefCreateEnumOut{}
			},
			__reflect__.TypeOf((*CorbaModuleDefCreateEnumOut)(nil))))
}