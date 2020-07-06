// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::PrimitiveDef_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPrimitiveDefDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaPrimitiveDefDestroyInId_Const = "IDL:CORBA/PrimitiveDef_destroy_In:1.0"

func (self *CorbaPrimitiveDefDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPrimitiveDefDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaPrimitiveDefDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPrimitiveDefDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPrimitiveDefDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPrimitiveDefDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPrimitiveDefDestroyInHelper = CorbaPrimitiveDefDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPrimitiveDefDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaPrimitiveDefDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaPrimitiveDefDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPrimitiveDefDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaPrimitiveDefDestroyIn)(nil))))
}
