// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_create_value_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefCreateValueOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefCreateValueOutId_Const = "IDL:CORBA/InterfaceDef_create_value_Out:1.0"

func (self *CorbaInterfaceDefCreateValueOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefCreateValueOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefCreateValueOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateValueOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateValueOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefCreateValueOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefCreateValueOutHelper = CorbaInterfaceDefCreateValueOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefCreateValueOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefCreateValueOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateValueOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateValueOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefCreateValueOut)(nil))))
}
