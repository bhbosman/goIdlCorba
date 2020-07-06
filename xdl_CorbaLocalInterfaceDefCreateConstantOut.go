// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::LocalInterfaceDef_create_constant_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaLocalInterfaceDefCreateConstantOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaLocalInterfaceDefCreateConstantOutId_Const = "IDL:CORBA/LocalInterfaceDef_create_constant_Out:1.0"

func (self *CorbaLocalInterfaceDefCreateConstantOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaLocalInterfaceDefCreateConstantOut) GoString() string {
	return self.String()
}

func (self *CorbaLocalInterfaceDefCreateConstantOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefCreateConstantOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefCreateConstantOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaLocalInterfaceDefCreateConstantOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaLocalInterfaceDefCreateConstantOutHelper = CorbaLocalInterfaceDefCreateConstantOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaLocalInterfaceDefCreateConstantOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaLocalInterfaceDefCreateConstantOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefCreateConstantOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefCreateConstantOut{}
			},
			__reflect__.TypeOf((*CorbaLocalInterfaceDefCreateConstantOut)(nil))))
}
