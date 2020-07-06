// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::PollableSet_add_pollable_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPollableSetAddPollableIn struct {
	__goidl__.IdlObject
	Potential CorbaPollable `json:"Potential"`
}

//noinspection GoSnakeCaseUsage
const CorbaPollableSetAddPollableInId_Const = "IDL:CORBA/PollableSet_add_pollable_In:1.0"

func (self *CorbaPollableSetAddPollableIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPollableSetAddPollableIn) GoString() string {
	return self.String()
}

func (self *CorbaPollableSetAddPollableIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.Potential, err = CorbaPollableHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableSetAddPollableIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableSetAddPollableIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaPollableHelper.Write(stream, self.Potential)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPollableSetAddPollableIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPollableSetAddPollableInHelper = CorbaPollableSetAddPollableIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPollableSetAddPollableInId_Const,
			__goidl__.StructType,
			"CORBA_Pollable.idl",
			"xdl_CorbaPollableSetAddPollableIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaPollableSetAddPollableIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPollableSetAddPollableIn{}
			},
			__reflect__.TypeOf((*CorbaPollableSetAddPollableIn)(nil))))
}
