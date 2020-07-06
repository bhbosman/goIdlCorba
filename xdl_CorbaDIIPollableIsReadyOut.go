// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DIIPollable_is_ready_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDIIPollableIsReadyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDIIPollableIsReadyOutId_Const = "IDL:CORBA/DIIPollable_is_ready_Out:1.0"

func (self *CorbaDIIPollableIsReadyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDIIPollableIsReadyOut) GoString() string {
	return self.String()
}

func (self *CorbaDIIPollableIsReadyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDIIPollableIsReadyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDIIPollableIsReadyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDIIPollableIsReadyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDIIPollableIsReadyOutHelper = CorbaDIIPollableIsReadyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDIIPollableIsReadyOutId_Const,
			__goidl__.StructType,
			"CORBA_Pollable.idl",
			"xdl_CorbaDIIPollableIsReadyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDIIPollableIsReadyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDIIPollableIsReadyOut{}
			},
			__reflect__.TypeOf((*CorbaDIIPollableIsReadyOut)(nil))))
}