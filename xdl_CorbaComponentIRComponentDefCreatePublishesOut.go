// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDef_create_publishes_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDefCreatePublishesOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDefCreatePublishesOutId_Const = "IDL:CORBA/ComponentIR/ComponentDef_create_publishes_Out:1.0"

func (self *CorbaComponentIRComponentDefCreatePublishesOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDefCreatePublishesOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDefCreatePublishesOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefCreatePublishesOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefCreatePublishesOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDefCreatePublishesOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDefCreatePublishesOutHelper = CorbaComponentIRComponentDefCreatePublishesOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDefCreatePublishesOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDefCreatePublishesOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefCreatePublishesOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefCreatePublishesOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDefCreatePublishesOut)(nil))))
}
