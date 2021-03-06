// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_object_to_string_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbObjectToStringOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbObjectToStringOutId_Const = "IDL:CORBA/ORB_object_to_string_Out:1.0"

func (self *CorbaOrbObjectToStringOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbObjectToStringOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbObjectToStringOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbObjectToStringOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbObjectToStringOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbObjectToStringOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbObjectToStringOutHelper = CorbaOrbObjectToStringOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbObjectToStringOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbObjectToStringOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbObjectToStringOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbObjectToStringOut{}
			},
			__reflect__.TypeOf((*CorbaOrbObjectToStringOut)(nil))))
}
