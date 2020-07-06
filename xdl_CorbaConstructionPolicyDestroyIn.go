// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ConstructionPolicy_destroy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaConstructionPolicyDestroyIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaConstructionPolicyDestroyInId_Const = "IDL:CORBA/ConstructionPolicy_destroy_In:1.0"

func (self *CorbaConstructionPolicyDestroyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaConstructionPolicyDestroyIn) GoString() string {
	return self.String()
}

func (self *CorbaConstructionPolicyDestroyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaConstructionPolicyDestroyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaConstructionPolicyDestroyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaConstructionPolicyDestroyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaConstructionPolicyDestroyInHelper = CorbaConstructionPolicyDestroyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaConstructionPolicyDestroyInId_Const,
			__goidl__.StructType,
			"CORBA_DomainManager.idl",
			"xdl_CorbaConstructionPolicyDestroyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaConstructionPolicyDestroyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaConstructionPolicyDestroyIn{}
			},
			__reflect__.TypeOf((*CorbaConstructionPolicyDestroyIn)(nil))))
}
