// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDef_create_alias_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDefCreateAliasOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDefCreateAliasOutId_Const = "IDL:CORBA/ComponentIR/ComponentDef_create_alias_Out:1.0"

func (self *CorbaComponentIRComponentDefCreateAliasOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDefCreateAliasOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDefCreateAliasOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefCreateAliasOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefCreateAliasOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDefCreateAliasOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDefCreateAliasOutHelper = CorbaComponentIRComponentDefCreateAliasOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDefCreateAliasOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDefCreateAliasOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefCreateAliasOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefCreateAliasOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDefCreateAliasOut)(nil))))
}
