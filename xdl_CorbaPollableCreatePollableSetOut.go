// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Pollable_create_pollable_set_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPollableCreatePollableSetOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaPollableCreatePollableSetOutId_Const = "IDL:CORBA/Pollable_create_pollable_set_Out:1.0"

func (self *CorbaPollableCreatePollableSetOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPollableCreatePollableSetOut) GoString() string {
	return self.String()
}

func (self *CorbaPollableCreatePollableSetOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableCreatePollableSetOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableCreatePollableSetOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPollableCreatePollableSetOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPollableCreatePollableSetOutHelper = CorbaPollableCreatePollableSetOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPollableCreatePollableSetOutId_Const,
			__goidl__.StructType,
			"CORBA_Pollable.idl",
			"xdl_CorbaPollableCreatePollableSetOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaPollableCreatePollableSetOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPollableCreatePollableSetOut{}
			},
			__reflect__.TypeOf((*CorbaPollableCreatePollableSetOut)(nil))))
}
