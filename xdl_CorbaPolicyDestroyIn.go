// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Policy_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPolicyDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaPolicyDestroyInId_Const = "IDL:CORBA/Policy_destroy_In:1.0"

func (self *CorbaPolicyDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPolicyDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaPolicyDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPolicyDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPolicyDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPolicyDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPolicyDestroyInHelper = CorbaPolicyDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPolicyDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_Policy.idl",
			"xdl_CorbaPolicyDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaPolicyDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPolicyDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaPolicyDestroyIn)(nil))))
}
