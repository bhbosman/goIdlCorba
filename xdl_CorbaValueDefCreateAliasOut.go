// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ValueDef_create_alias_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaValueDefCreateAliasOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaValueDefCreateAliasOutId_Const = "IDL:CORBA/ValueDef_create_alias_Out:1.0"

func (self *CorbaValueDefCreateAliasOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaValueDefCreateAliasOut) GoString() string {
	return self.String()
}

func (self *CorbaValueDefCreateAliasOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefCreateAliasOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefCreateAliasOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaValueDefCreateAliasOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaValueDefCreateAliasOutHelper = CorbaValueDefCreateAliasOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaValueDefCreateAliasOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaValueDefCreateAliasOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaValueDefCreateAliasOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaValueDefCreateAliasOut{}
			},
			__reflect__.TypeOf((*CorbaValueDefCreateAliasOut)(nil))))
}
