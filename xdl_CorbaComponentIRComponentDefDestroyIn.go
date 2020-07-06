// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDef_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDefDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDefDestroyInId_Const = "IDL:CORBA/ComponentIR/ComponentDef_destroy_In:1.0"

func (self *CorbaComponentIRComponentDefDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDefDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDefDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDefDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDefDestroyInHelper = CorbaComponentIRComponentDefDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDefDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDefDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDefDestroyIn)(nil))))
}
