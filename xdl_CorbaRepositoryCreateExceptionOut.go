// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Repository_create_exception_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRepositoryCreateExceptionOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRepositoryCreateExceptionOutId_Const = "IDL:CORBA/Repository_create_exception_Out:1.0"

func (self *CorbaRepositoryCreateExceptionOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRepositoryCreateExceptionOut) GoString() string {
	return self.String()
}

func (self *CorbaRepositoryCreateExceptionOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateExceptionOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateExceptionOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRepositoryCreateExceptionOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRepositoryCreateExceptionOutHelper = CorbaRepositoryCreateExceptionOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRepositoryCreateExceptionOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaRepositoryCreateExceptionOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaRepositoryCreateExceptionOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRepositoryCreateExceptionOut{}
			},
			__reflect__.TypeOf((*CorbaRepositoryCreateExceptionOut)(nil))))
}
