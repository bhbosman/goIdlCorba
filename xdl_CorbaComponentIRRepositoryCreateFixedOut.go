// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Repository_create_fixed_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRRepositoryCreateFixedOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRRepositoryCreateFixedOutId_Const = "IDL:CORBA/ComponentIR/Repository_create_fixed_Out:1.0"

func (self *CorbaComponentIRRepositoryCreateFixedOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRRepositoryCreateFixedOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRRepositoryCreateFixedOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateFixedOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateFixedOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRRepositoryCreateFixedOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRRepositoryCreateFixedOutHelper = CorbaComponentIRRepositoryCreateFixedOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRRepositoryCreateFixedOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRRepositoryCreateFixedOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateFixedOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateFixedOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRRepositoryCreateFixedOut)(nil))))
}