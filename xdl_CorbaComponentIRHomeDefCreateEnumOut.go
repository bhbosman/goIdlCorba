// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::HomeDef_create_enum_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRHomeDefCreateEnumOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRHomeDefCreateEnumOutId_Const = "IDL:CORBA/ComponentIR/HomeDef_create_enum_Out:1.0"

func (self *CorbaComponentIRHomeDefCreateEnumOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRHomeDefCreateEnumOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRHomeDefCreateEnumOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefCreateEnumOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefCreateEnumOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRHomeDefCreateEnumOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRHomeDefCreateEnumOutHelper = CorbaComponentIRHomeDefCreateEnumOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRHomeDefCreateEnumOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRHomeDefCreateEnumOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefCreateEnumOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefCreateEnumOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRHomeDefCreateEnumOut)(nil))))
}