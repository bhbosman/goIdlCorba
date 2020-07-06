// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Pollable_is_ready_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPollableIsReadyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaPollableIsReadyOutId_Const = "IDL:CORBA/Pollable_is_ready_Out:1.0"

func (self *CorbaPollableIsReadyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPollableIsReadyOut) GoString() string {
	return self.String()
}

func (self *CorbaPollableIsReadyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableIsReadyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableIsReadyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPollableIsReadyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPollableIsReadyOutHelper = CorbaPollableIsReadyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPollableIsReadyOutId_Const,
			__goidl__.StructType,
			"CORBA_Pollable.idl",
			"xdl_CorbaPollableIsReadyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaPollableIsReadyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPollableIsReadyOut{}
			},
			__reflect__.TypeOf((*CorbaPollableIsReadyOut)(nil))))
}
