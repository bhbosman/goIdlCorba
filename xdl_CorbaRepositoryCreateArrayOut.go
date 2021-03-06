// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Repository_create_array_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRepositoryCreateArrayOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRepositoryCreateArrayOutId_Const = "IDL:CORBA/Repository_create_array_Out:1.0"

func (self *CorbaRepositoryCreateArrayOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRepositoryCreateArrayOut) GoString() string {
	return self.String()
}

func (self *CorbaRepositoryCreateArrayOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateArrayOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateArrayOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRepositoryCreateArrayOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRepositoryCreateArrayOutHelper = CorbaRepositoryCreateArrayOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRepositoryCreateArrayOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaRepositoryCreateArrayOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaRepositoryCreateArrayOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRepositoryCreateArrayOut{}
			},
			__reflect__.TypeOf((*CorbaRepositoryCreateArrayOut)(nil))))
}
