// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::EnumDef_destroy_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaEnumDefDestroyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaEnumDefDestroyOutId_Const = "IDL:CORBA/EnumDef_destroy_Out:1.0"

func (self *CorbaEnumDefDestroyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaEnumDefDestroyOut) GoString() string {
	return self.String()
}

func (self *CorbaEnumDefDestroyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaEnumDefDestroyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaEnumDefDestroyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaEnumDefDestroyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaEnumDefDestroyOutHelper = CorbaEnumDefDestroyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaEnumDefDestroyOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaEnumDefDestroyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaEnumDefDestroyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaEnumDefDestroyOut{}
			},
			__reflect__.TypeOf((*CorbaEnumDefDestroyOut)(nil))))
}
