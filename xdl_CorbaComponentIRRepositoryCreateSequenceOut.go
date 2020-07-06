// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Repository_create_sequence_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRRepositoryCreateSequenceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRRepositoryCreateSequenceOutId_Const = "IDL:CORBA/ComponentIR/Repository_create_sequence_Out:1.0"

func (self *CorbaComponentIRRepositoryCreateSequenceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRRepositoryCreateSequenceOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRRepositoryCreateSequenceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateSequenceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateSequenceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRRepositoryCreateSequenceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRRepositoryCreateSequenceOutHelper = CorbaComponentIRRepositoryCreateSequenceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRRepositoryCreateSequenceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRRepositoryCreateSequenceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateSequenceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateSequenceOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRRepositoryCreateSequenceOut)(nil))))
}
