// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_exception_tc_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateExceptionTcOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateExceptionTcOutId_Const = "IDL:CORBA/ORB_create_exception_tc_Out:1.0"

func (self *CorbaOrbCreateExceptionTcOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateExceptionTcOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateExceptionTcOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateExceptionTcOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateExceptionTcOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateExceptionTcOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateExceptionTcOutHelper = CorbaOrbCreateExceptionTcOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateExceptionTcOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateExceptionTcOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateExceptionTcOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateExceptionTcOut{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateExceptionTcOut)(nil))))
}
