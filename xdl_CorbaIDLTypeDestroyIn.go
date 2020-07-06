// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::IDLType_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaIDLTypeDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaIDLTypeDestroyInId_Const = "IDL:CORBA/IDLType_destroy_In:1.0"

func (self *CorbaIDLTypeDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaIDLTypeDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaIDLTypeDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaIDLTypeDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaIDLTypeDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaIDLTypeDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaIDLTypeDestroyInHelper = CorbaIDLTypeDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaIDLTypeDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaIDLTypeDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaIDLTypeDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaIDLTypeDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaIDLTypeDestroyIn)(nil))))
}