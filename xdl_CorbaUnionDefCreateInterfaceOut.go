// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::UnionDef_create_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaUnionDefCreateInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaUnionDefCreateInterfaceOutId_Const = "IDL:CORBA/UnionDef_create_interface_Out:1.0"

func (self *CorbaUnionDefCreateInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaUnionDefCreateInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaUnionDefCreateInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaUnionDefCreateInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaUnionDefCreateInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaUnionDefCreateInterfaceOutHelper = CorbaUnionDefCreateInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaUnionDefCreateInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaUnionDefCreateInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaUnionDefCreateInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaUnionDefCreateInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaUnionDefCreateInterfaceOut)(nil))))
}
