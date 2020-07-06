// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_is_a_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefIsAOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefIsAOutId_Const = "IDL:CORBA/InterfaceDef_is_a_Out:1.0"

func (self *CorbaInterfaceDefIsAOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefIsAOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefIsAOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefIsAOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefIsAOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefIsAOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefIsAOutHelper = CorbaInterfaceDefIsAOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefIsAOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefIsAOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefIsAOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefIsAOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefIsAOut)(nil))))
}
