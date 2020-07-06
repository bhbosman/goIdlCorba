// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::StructDef_create_alias_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaStructDefCreateAliasOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaStructDefCreateAliasOutId_Const = "IDL:CORBA/StructDef_create_alias_Out:1.0"

func (self *CorbaStructDefCreateAliasOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaStructDefCreateAliasOut) GoString() string {
	return self.String()
}

func (self *CorbaStructDefCreateAliasOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaStructDefCreateAliasOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaStructDefCreateAliasOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaStructDefCreateAliasOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaStructDefCreateAliasOutHelper = CorbaStructDefCreateAliasOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaStructDefCreateAliasOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaStructDefCreateAliasOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaStructDefCreateAliasOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaStructDefCreateAliasOut{}
			},
			__reflect__.TypeOf((*CorbaStructDefCreateAliasOut)(nil))))
}
