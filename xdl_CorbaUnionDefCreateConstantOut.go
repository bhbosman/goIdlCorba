// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::UnionDef_create_constant_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaUnionDefCreateConstantOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaUnionDefCreateConstantOutId_Const = "IDL:CORBA/UnionDef_create_constant_Out:1.0"

func (self *CorbaUnionDefCreateConstantOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaUnionDefCreateConstantOut) GoString() string {
	return self.String()
}

func (self *CorbaUnionDefCreateConstantOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateConstantOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateConstantOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaUnionDefCreateConstantOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaUnionDefCreateConstantOutHelper = CorbaUnionDefCreateConstantOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaUnionDefCreateConstantOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaUnionDefCreateConstantOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaUnionDefCreateConstantOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaUnionDefCreateConstantOut{}
			},
			__reflect__.TypeOf((*CorbaUnionDefCreateConstantOut)(nil))))
}
