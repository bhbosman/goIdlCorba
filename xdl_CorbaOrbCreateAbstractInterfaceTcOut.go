// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_abstract_interface_tc_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateAbstractInterfaceTcOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateAbstractInterfaceTcOutId_Const = "IDL:CORBA/ORB_create_abstract_interface_tc_Out:1.0"

func (self *CorbaOrbCreateAbstractInterfaceTcOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateAbstractInterfaceTcOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateAbstractInterfaceTcOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateAbstractInterfaceTcOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateAbstractInterfaceTcOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateAbstractInterfaceTcOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateAbstractInterfaceTcOutHelper = CorbaOrbCreateAbstractInterfaceTcOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateAbstractInterfaceTcOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateAbstractInterfaceTcOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateAbstractInterfaceTcOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateAbstractInterfaceTcOut{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateAbstractInterfaceTcOut)(nil))))
}
