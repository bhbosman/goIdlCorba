// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::StructDef_create_union_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaStructDefCreateUnionOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaStructDefCreateUnionOutId_Const = "IDL:CORBA/StructDef_create_union_Out:1.0"

func (self *CorbaStructDefCreateUnionOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaStructDefCreateUnionOut) GoString() string {
	return self.String()
}

func (self *CorbaStructDefCreateUnionOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaStructDefCreateUnionOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaStructDefCreateUnionOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaStructDefCreateUnionOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaStructDefCreateUnionOutHelper = CorbaStructDefCreateUnionOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaStructDefCreateUnionOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaStructDefCreateUnionOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaStructDefCreateUnionOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaStructDefCreateUnionOut{}
			},
			__reflect__.TypeOf((*CorbaStructDefCreateUnionOut)(nil))))
}
