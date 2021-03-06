// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Repository_destroy_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRepositoryDestroyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRepositoryDestroyOutId_Const = "IDL:CORBA/Repository_destroy_Out:1.0"

func (self *CorbaRepositoryDestroyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRepositoryDestroyOut) GoString() string {
	return self.String()
}

func (self *CorbaRepositoryDestroyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryDestroyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryDestroyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRepositoryDestroyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRepositoryDestroyOutHelper = CorbaRepositoryDestroyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRepositoryDestroyOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaRepositoryDestroyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaRepositoryDestroyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRepositoryDestroyOut{}
			},
			__reflect__.TypeOf((*CorbaRepositoryDestroyOut)(nil))))
}
