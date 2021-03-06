// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::IRObject_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaIRObjectDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaIRObjectDestroyInId_Const = "IDL:CORBA/IRObject_destroy_In:1.0"

func (self *CorbaIRObjectDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaIRObjectDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaIRObjectDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaIRObjectDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaIRObjectDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaIRObjectDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaIRObjectDestroyInHelper = CorbaIRObjectDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaIRObjectDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaIRObjectDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaIRObjectDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaIRObjectDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaIRObjectDestroyIn)(nil))))
}
