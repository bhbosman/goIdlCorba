// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_component_tc_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateComponentTcOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateComponentTcOutId_Const = "IDL:CORBA/ORB_create_component_tc_Out:1.0"

func (self *CorbaOrbCreateComponentTcOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateComponentTcOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateComponentTcOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateComponentTcOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateComponentTcOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateComponentTcOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateComponentTcOutHelper = CorbaOrbCreateComponentTcOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateComponentTcOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateComponentTcOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateComponentTcOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateComponentTcOut{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateComponentTcOut)(nil))))
}