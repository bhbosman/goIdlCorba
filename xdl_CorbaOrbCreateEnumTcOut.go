// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_enum_tc_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateEnumTcOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateEnumTcOutId_Const = "IDL:CORBA/ORB_create_enum_tc_Out:1.0"

func (self *CorbaOrbCreateEnumTcOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateEnumTcOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateEnumTcOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateEnumTcOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateEnumTcOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateEnumTcOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateEnumTcOutHelper = CorbaOrbCreateEnumTcOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateEnumTcOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateEnumTcOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateEnumTcOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateEnumTcOut{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateEnumTcOut)(nil))))
}
