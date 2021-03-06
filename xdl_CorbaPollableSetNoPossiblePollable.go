// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __fmt__ "fmt"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Exception declaration: "CORBA::PollableSet::NoPossiblePollable", generatedBy by: "WriteStructDcl"
// Exception Decl: true
type CorbaPollableSetNoPossiblePollable struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaPollableSetNoPossiblePollableId_Const = "IDL:omg.org/CORBA/PollableSet/NoPossiblePollable:1.0"

func (self *CorbaPollableSetNoPossiblePollable) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPollableSetNoPossiblePollable) Error() string{
	return 	__fmt__.Sprintf("Error of type CorbaPollableSetNoPossiblePollable(%v)", self.String())
}
func (self *CorbaPollableSetNoPossiblePollable) GoString() string {
	return self.String()
}

func (self *CorbaPollableSetNoPossiblePollable) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableSetNoPossiblePollable) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPollableSetNoPossiblePollable) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPollableSetNoPossiblePollable_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPollableSetNoPossiblePollableHelper = CorbaPollableSetNoPossiblePollable_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPollableSetNoPossiblePollableId_Const,
			__goidl__.StructType,
			"CORBA_Pollable.idl",
			"xdl_CorbaPollableSetNoPossiblePollable.go",
			func() __goidl__.IIdlObject {
				return &CorbaPollableSetNoPossiblePollable{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPollableSetNoPossiblePollable{}
			},
			__reflect__.TypeOf((*CorbaPollableSetNoPossiblePollable)(nil))))
}
