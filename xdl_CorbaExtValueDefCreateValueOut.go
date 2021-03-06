// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtValueDef_create_value_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtValueDefCreateValueOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtValueDefCreateValueOutId_Const = "IDL:CORBA/ExtValueDef_create_value_Out:1.0"

func (self *CorbaExtValueDefCreateValueOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtValueDefCreateValueOut) GoString() string {
	return self.String()
}

func (self *CorbaExtValueDefCreateValueOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefCreateValueOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefCreateValueOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtValueDefCreateValueOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtValueDefCreateValueOutHelper = CorbaExtValueDefCreateValueOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtValueDefCreateValueOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtValueDefCreateValueOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtValueDefCreateValueOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtValueDefCreateValueOut{}
			},
			__reflect__.TypeOf((*CorbaExtValueDefCreateValueOut)(nil))))
}
