// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_policy_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreatePolicyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreatePolicyOutId_Const = "IDL:CORBA/ORB_create_policy_Out:1.0"

func (self *CorbaOrbCreatePolicyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreatePolicyOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreatePolicyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreatePolicyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreatePolicyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreatePolicyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreatePolicyOutHelper = CorbaOrbCreatePolicyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreatePolicyOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreatePolicyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreatePolicyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreatePolicyOut{}
			},
			__reflect__.TypeOf((*CorbaOrbCreatePolicyOut)(nil))))
}
