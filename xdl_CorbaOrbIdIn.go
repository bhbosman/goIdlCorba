// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_id_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbIdIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbIdInId_Const = "IDL:CORBA/ORB_id_In:1.0"

func (self *CorbaOrbIdIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbIdIn) GoString() string {
	return self.String()
}

func (self *CorbaOrbIdIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbIdIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbIdIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbIdIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbIdInHelper = CorbaOrbIdIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbIdInId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbIdIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbIdIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbIdIn{}
			},
			__reflect__.TypeOf((*CorbaOrbIdIn)(nil))))
}
