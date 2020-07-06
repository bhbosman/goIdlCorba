// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_local_interface_tc_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateLocalInterfaceTcOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateLocalInterfaceTcOutId_Const = "IDL:CORBA/ORB_create_local_interface_tc_Out:1.0"

func (self *CorbaOrbCreateLocalInterfaceTcOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateLocalInterfaceTcOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateLocalInterfaceTcOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateLocalInterfaceTcOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateLocalInterfaceTcOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateLocalInterfaceTcOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateLocalInterfaceTcOutHelper = CorbaOrbCreateLocalInterfaceTcOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateLocalInterfaceTcOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateLocalInterfaceTcOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateLocalInterfaceTcOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateLocalInterfaceTcOut{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateLocalInterfaceTcOut)(nil))))
}
