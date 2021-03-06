// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DIIPollable_create_pollable_set_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDIIPollableCreatePollableSetIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDIIPollableCreatePollableSetInId_Const = "IDL:CORBA/DIIPollable_create_pollable_set_In:1.0"

func (self *CorbaDIIPollableCreatePollableSetIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDIIPollableCreatePollableSetIn) GoString() string {
	return self.String()
}

func (self *CorbaDIIPollableCreatePollableSetIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDIIPollableCreatePollableSetIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDIIPollableCreatePollableSetIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDIIPollableCreatePollableSetIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDIIPollableCreatePollableSetInHelper = CorbaDIIPollableCreatePollableSetIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDIIPollableCreatePollableSetInId_Const,
			__goidl__.StructType,
			"CORBA_Pollable.idl",
			"xdl_CorbaDIIPollableCreatePollableSetIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDIIPollableCreatePollableSetIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDIIPollableCreatePollableSetIn{}
			},
			__reflect__.TypeOf((*CorbaDIIPollableCreatePollableSetIn)(nil))))
}
