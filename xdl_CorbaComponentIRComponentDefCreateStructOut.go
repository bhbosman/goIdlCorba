// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDef_create_struct_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDefCreateStructOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDefCreateStructOutId_Const = "IDL:CORBA/ComponentIR/ComponentDef_create_struct_Out:1.0"

func (self *CorbaComponentIRComponentDefCreateStructOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDefCreateStructOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDefCreateStructOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefCreateStructOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefCreateStructOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDefCreateStructOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDefCreateStructOutHelper = CorbaComponentIRComponentDefCreateStructOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDefCreateStructOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDefCreateStructOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefCreateStructOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefCreateStructOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDefCreateStructOut)(nil))))
}