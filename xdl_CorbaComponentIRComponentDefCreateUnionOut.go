// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDef_create_union_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDefCreateUnionOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDefCreateUnionOutId_Const = "IDL:CORBA/ComponentIR/ComponentDef_create_union_Out:1.0"

func (self *CorbaComponentIRComponentDefCreateUnionOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDefCreateUnionOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDefCreateUnionOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefCreateUnionOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefCreateUnionOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDefCreateUnionOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDefCreateUnionOutHelper = CorbaComponentIRComponentDefCreateUnionOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDefCreateUnionOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDefCreateUnionOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefCreateUnionOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefCreateUnionOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDefCreateUnionOut)(nil))))
}
