// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtLocalInterfaceDef_create_local_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtLocalInterfaceDefCreateLocalInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtLocalInterfaceDefCreateLocalInterfaceOutId_Const = "IDL:CORBA/ExtLocalInterfaceDef_create_local_interface_Out:1.0"

func (self *CorbaExtLocalInterfaceDefCreateLocalInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtLocalInterfaceDefCreateLocalInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaExtLocalInterfaceDefCreateLocalInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtLocalInterfaceDefCreateLocalInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtLocalInterfaceDefCreateLocalInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtLocalInterfaceDefCreateLocalInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtLocalInterfaceDefCreateLocalInterfaceOutHelper = CorbaExtLocalInterfaceDefCreateLocalInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtLocalInterfaceDefCreateLocalInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtLocalInterfaceDefCreateLocalInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefCreateLocalInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefCreateLocalInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaExtLocalInterfaceDefCreateLocalInterfaceOut)(nil))))
}