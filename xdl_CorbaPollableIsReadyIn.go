// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Pollable_is_ready_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPollableIsReadyIn struct {
	__goidl__.IdlObject
	Timeout uint32 `json:"Timeout"`
}

//noinspection GoSnakeCaseUsage
const CorbaPollableIsReadyInId_Const = "IDL:CORBA/Pollable_is_ready_In:1.0"

func (self *CorbaPollableIsReadyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPollableIsReadyIn) GoString() string {
	return self.String()
}

func (self *CorbaPollableIsReadyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.Timeout, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableIsReadyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableIsReadyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Timeout)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPollableIsReadyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPollableIsReadyInHelper = CorbaPollableIsReadyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPollableIsReadyInId_Const,
			__goidl__.StructType,
			"CORBA_Pollable.idl",
			"xdl_CorbaPollableIsReadyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaPollableIsReadyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPollableIsReadyIn{}
			},
			__reflect__.TypeOf((*CorbaPollableIsReadyIn)(nil))))
}
