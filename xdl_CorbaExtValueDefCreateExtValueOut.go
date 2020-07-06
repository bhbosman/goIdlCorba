// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtValueDef_create_ext_value_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtValueDefCreateExtValueOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtValueDefCreateExtValueOutId_Const = "IDL:CORBA/ExtValueDef_create_ext_value_Out:1.0"

func (self *CorbaExtValueDefCreateExtValueOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtValueDefCreateExtValueOut) GoString() string {
	return self.String()
}

func (self *CorbaExtValueDefCreateExtValueOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefCreateExtValueOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefCreateExtValueOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtValueDefCreateExtValueOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtValueDefCreateExtValueOutHelper = CorbaExtValueDefCreateExtValueOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtValueDefCreateExtValueOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtValueDefCreateExtValueOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtValueDefCreateExtValueOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtValueDefCreateExtValueOut{}
			},
			__reflect__.TypeOf((*CorbaExtValueDefCreateExtValueOut)(nil))))
}