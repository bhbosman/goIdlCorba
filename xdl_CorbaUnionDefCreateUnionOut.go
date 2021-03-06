// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::UnionDef_create_union_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaUnionDefCreateUnionOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaUnionDefCreateUnionOutId_Const = "IDL:CORBA/UnionDef_create_union_Out:1.0"

func (self *CorbaUnionDefCreateUnionOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaUnionDefCreateUnionOut) GoString() string {
	return self.String()
}

func (self *CorbaUnionDefCreateUnionOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateUnionOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateUnionOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaUnionDefCreateUnionOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaUnionDefCreateUnionOutHelper = CorbaUnionDefCreateUnionOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaUnionDefCreateUnionOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaUnionDefCreateUnionOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaUnionDefCreateUnionOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaUnionDefCreateUnionOut{}
			},
			__reflect__.TypeOf((*CorbaUnionDefCreateUnionOut)(nil))))
}
