// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Container_create_exception_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContainerCreateExceptionOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaContainerCreateExceptionOutId_Const = "IDL:CORBA/Container_create_exception_Out:1.0"

func (self *CorbaContainerCreateExceptionOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContainerCreateExceptionOut) GoString() string {
	return self.String()
}

func (self *CorbaContainerCreateExceptionOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerCreateExceptionOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerCreateExceptionOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContainerCreateExceptionOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContainerCreateExceptionOutHelper = CorbaContainerCreateExceptionOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContainerCreateExceptionOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaContainerCreateExceptionOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaContainerCreateExceptionOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContainerCreateExceptionOut{}
			},
			__reflect__.TypeOf((*CorbaContainerCreateExceptionOut)(nil))))
}