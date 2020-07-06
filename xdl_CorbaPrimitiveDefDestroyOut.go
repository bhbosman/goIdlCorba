// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::PrimitiveDef_destroy_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPrimitiveDefDestroyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaPrimitiveDefDestroyOutId_Const = "IDL:CORBA/PrimitiveDef_destroy_Out:1.0"

func (self *CorbaPrimitiveDefDestroyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPrimitiveDefDestroyOut) GoString() string {
	return self.String()
}

func (self *CorbaPrimitiveDefDestroyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPrimitiveDefDestroyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPrimitiveDefDestroyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPrimitiveDefDestroyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPrimitiveDefDestroyOutHelper = CorbaPrimitiveDefDestroyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPrimitiveDefDestroyOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaPrimitiveDefDestroyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaPrimitiveDefDestroyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPrimitiveDefDestroyOut{}
			},
			__reflect__.TypeOf((*CorbaPrimitiveDefDestroyOut)(nil))))
}