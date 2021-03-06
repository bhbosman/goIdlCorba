// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Repository_get_primitive_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRepositoryGetPrimitiveOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRepositoryGetPrimitiveOutId_Const = "IDL:CORBA/Repository_get_primitive_Out:1.0"

func (self *CorbaRepositoryGetPrimitiveOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRepositoryGetPrimitiveOut) GoString() string {
	return self.String()
}

func (self *CorbaRepositoryGetPrimitiveOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryGetPrimitiveOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryGetPrimitiveOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRepositoryGetPrimitiveOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRepositoryGetPrimitiveOutHelper = CorbaRepositoryGetPrimitiveOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRepositoryGetPrimitiveOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaRepositoryGetPrimitiveOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaRepositoryGetPrimitiveOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRepositoryGetPrimitiveOut{}
			},
			__reflect__.TypeOf((*CorbaRepositoryGetPrimitiveOut)(nil))))
}
