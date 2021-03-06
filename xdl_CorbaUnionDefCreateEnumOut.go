// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::UnionDef_create_enum_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaUnionDefCreateEnumOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaUnionDefCreateEnumOutId_Const = "IDL:CORBA/UnionDef_create_enum_Out:1.0"

func (self *CorbaUnionDefCreateEnumOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaUnionDefCreateEnumOut) GoString() string {
	return self.String()
}

func (self *CorbaUnionDefCreateEnumOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateEnumOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateEnumOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaUnionDefCreateEnumOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaUnionDefCreateEnumOutHelper = CorbaUnionDefCreateEnumOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaUnionDefCreateEnumOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaUnionDefCreateEnumOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaUnionDefCreateEnumOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaUnionDefCreateEnumOut{}
			},
			__reflect__.TypeOf((*CorbaUnionDefCreateEnumOut)(nil))))
}
