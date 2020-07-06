// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::HomeDef_create_alias_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRHomeDefCreateAliasOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRHomeDefCreateAliasOutId_Const = "IDL:CORBA/ComponentIR/HomeDef_create_alias_Out:1.0"

func (self *CorbaComponentIRHomeDefCreateAliasOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRHomeDefCreateAliasOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRHomeDefCreateAliasOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefCreateAliasOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefCreateAliasOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRHomeDefCreateAliasOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRHomeDefCreateAliasOutHelper = CorbaComponentIRHomeDefCreateAliasOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRHomeDefCreateAliasOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRHomeDefCreateAliasOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefCreateAliasOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefCreateAliasOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRHomeDefCreateAliasOut)(nil))))
}
