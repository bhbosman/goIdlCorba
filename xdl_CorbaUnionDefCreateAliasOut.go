// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::UnionDef_create_alias_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaUnionDefCreateAliasOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaUnionDefCreateAliasOutId_Const = "IDL:CORBA/UnionDef_create_alias_Out:1.0"

func (self *CorbaUnionDefCreateAliasOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaUnionDefCreateAliasOut) GoString() string {
	return self.String()
}

func (self *CorbaUnionDefCreateAliasOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateAliasOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateAliasOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaUnionDefCreateAliasOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaUnionDefCreateAliasOutHelper = CorbaUnionDefCreateAliasOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaUnionDefCreateAliasOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaUnionDefCreateAliasOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaUnionDefCreateAliasOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaUnionDefCreateAliasOut{}
			},
			__reflect__.TypeOf((*CorbaUnionDefCreateAliasOut)(nil))))
}
