// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_work_pending_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbWorkPendingIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbWorkPendingInId_Const = "IDL:CORBA/ORB_work_pending_In:1.0"

func (self *CorbaOrbWorkPendingIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbWorkPendingIn) GoString() string {
	return self.String()
}

func (self *CorbaOrbWorkPendingIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbWorkPendingIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbWorkPendingIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbWorkPendingIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbWorkPendingInHelper = CorbaOrbWorkPendingIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbWorkPendingInId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbWorkPendingIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbWorkPendingIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbWorkPendingIn{}
			},
			__reflect__.TypeOf((*CorbaOrbWorkPendingIn)(nil))))
}
