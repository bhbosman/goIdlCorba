// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::EnumDef_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaEnumDefDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaEnumDefDestroyInId_Const = "IDL:CORBA/EnumDef_destroy_In:1.0"

func (self *CorbaEnumDefDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaEnumDefDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaEnumDefDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaEnumDefDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaEnumDefDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaEnumDefDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaEnumDefDestroyInHelper = CorbaEnumDefDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaEnumDefDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaEnumDefDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaEnumDefDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaEnumDefDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaEnumDefDestroyIn)(nil))))
}
