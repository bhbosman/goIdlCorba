// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtLocalInterfaceDef_destroy_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtLocalInterfaceDefDestroyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtLocalInterfaceDefDestroyOutId_Const = "IDL:CORBA/ExtLocalInterfaceDef_destroy_Out:1.0"

func (self *CorbaExtLocalInterfaceDefDestroyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtLocalInterfaceDefDestroyOut) GoString() string {
	return self.String()
}

func (self *CorbaExtLocalInterfaceDefDestroyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtLocalInterfaceDefDestroyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtLocalInterfaceDefDestroyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtLocalInterfaceDefDestroyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtLocalInterfaceDefDestroyOutHelper = CorbaExtLocalInterfaceDefDestroyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtLocalInterfaceDefDestroyOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtLocalInterfaceDefDestroyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefDestroyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefDestroyOut{}
			},
			__reflect__.TypeOf((*CorbaExtLocalInterfaceDefDestroyOut)(nil))))
}