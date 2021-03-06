// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Repository_contents_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRepositoryContentsOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRepositoryContentsOutId_Const = "IDL:CORBA/Repository_contents_Out:1.0"

func (self *CorbaRepositoryContentsOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRepositoryContentsOut) GoString() string {
	return self.String()
}

func (self *CorbaRepositoryContentsOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryContentsOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryContentsOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRepositoryContentsOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRepositoryContentsOutHelper = CorbaRepositoryContentsOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRepositoryContentsOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaRepositoryContentsOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaRepositoryContentsOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRepositoryContentsOut{}
			},
			__reflect__.TypeOf((*CorbaRepositoryContentsOut)(nil))))
}
