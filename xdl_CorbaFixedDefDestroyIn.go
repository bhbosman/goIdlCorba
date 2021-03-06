// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::FixedDef_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaFixedDefDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaFixedDefDestroyInId_Const = "IDL:CORBA/FixedDef_destroy_In:1.0"

func (self *CorbaFixedDefDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaFixedDefDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaFixedDefDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaFixedDefDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaFixedDefDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaFixedDefDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaFixedDefDestroyInHelper = CorbaFixedDefDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaFixedDefDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaFixedDefDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaFixedDefDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaFixedDefDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaFixedDefDestroyIn)(nil))))
}
