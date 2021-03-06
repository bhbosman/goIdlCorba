// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::UnionDef_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaUnionDefDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaUnionDefDestroyInId_Const = "IDL:CORBA/UnionDef_destroy_In:1.0"

func (self *CorbaUnionDefDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaUnionDefDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaUnionDefDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaUnionDefDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaUnionDefDestroyInHelper = CorbaUnionDefDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaUnionDefDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaUnionDefDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaUnionDefDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaUnionDefDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaUnionDefDestroyIn)(nil))))
}
