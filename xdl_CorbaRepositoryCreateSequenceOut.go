// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Repository_create_sequence_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRepositoryCreateSequenceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRepositoryCreateSequenceOutId_Const = "IDL:CORBA/Repository_create_sequence_Out:1.0"

func (self *CorbaRepositoryCreateSequenceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRepositoryCreateSequenceOut) GoString() string {
	return self.String()
}

func (self *CorbaRepositoryCreateSequenceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateSequenceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateSequenceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRepositoryCreateSequenceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRepositoryCreateSequenceOutHelper = CorbaRepositoryCreateSequenceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRepositoryCreateSequenceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaRepositoryCreateSequenceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaRepositoryCreateSequenceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRepositoryCreateSequenceOut{}
			},
			__reflect__.TypeOf((*CorbaRepositoryCreateSequenceOut)(nil))))
}
