// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Container_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContainerDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaContainerDestroyInId_Const = "IDL:CORBA/Container_destroy_In:1.0"

func (self *CorbaContainerDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContainerDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaContainerDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContainerDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContainerDestroyInHelper = CorbaContainerDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContainerDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaContainerDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaContainerDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContainerDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaContainerDestroyIn)(nil))))
}