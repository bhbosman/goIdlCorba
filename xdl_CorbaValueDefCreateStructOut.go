// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ValueDef_create_struct_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaValueDefCreateStructOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaValueDefCreateStructOutId_Const = "IDL:CORBA/ValueDef_create_struct_Out:1.0"

func (self *CorbaValueDefCreateStructOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaValueDefCreateStructOut) GoString() string {
	return self.String()
}

func (self *CorbaValueDefCreateStructOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefCreateStructOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefCreateStructOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaValueDefCreateStructOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaValueDefCreateStructOutHelper = CorbaValueDefCreateStructOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaValueDefCreateStructOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaValueDefCreateStructOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaValueDefCreateStructOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaValueDefCreateStructOut{}
			},
			__reflect__.TypeOf((*CorbaValueDefCreateStructOut)(nil))))
}
