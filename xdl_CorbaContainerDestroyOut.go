// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Container_destroy_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContainerDestroyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaContainerDestroyOutId_Const = "IDL:CORBA/Container_destroy_Out:1.0"

func (self *CorbaContainerDestroyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContainerDestroyOut) GoString() string {
	return self.String()
}

func (self *CorbaContainerDestroyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerDestroyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerDestroyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContainerDestroyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContainerDestroyOutHelper = CorbaContainerDestroyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContainerDestroyOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaContainerDestroyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaContainerDestroyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContainerDestroyOut{}
			},
			__reflect__.TypeOf((*CorbaContainerDestroyOut)(nil))))
}