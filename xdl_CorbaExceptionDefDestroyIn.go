// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExceptionDef_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExceptionDefDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExceptionDefDestroyInId_Const = "IDL:CORBA/ExceptionDef_destroy_In:1.0"

func (self *CorbaExceptionDefDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExceptionDefDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaExceptionDefDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExceptionDefDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExceptionDefDestroyInHelper = CorbaExceptionDefDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExceptionDefDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExceptionDefDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExceptionDefDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExceptionDefDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaExceptionDefDestroyIn)(nil))))
}