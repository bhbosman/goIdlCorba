// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDef_is_a_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDefIsAOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDefIsAOutId_Const = "IDL:CORBA/ComponentIR/ComponentDef_is_a_Out:1.0"

func (self *CorbaComponentIRComponentDefIsAOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDefIsAOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDefIsAOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefIsAOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefIsAOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDefIsAOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDefIsAOutHelper = CorbaComponentIRComponentDefIsAOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDefIsAOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDefIsAOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefIsAOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefIsAOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDefIsAOut)(nil))))
}
