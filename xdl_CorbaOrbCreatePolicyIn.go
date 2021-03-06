// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_policy_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreatePolicyIn struct {
	__goidl__.IdlObject
	Type uint32 `json:"Type"`
	Val __goidl__.IdlAny `json:"Val"`
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreatePolicyInId_Const = "IDL:CORBA/ORB_create_policy_In:1.0"

func (self *CorbaOrbCreatePolicyIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreatePolicyIn) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreatePolicyIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.Type, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(AnyType)
	self.Val, err = __goidl__.IdlAnyHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreatePolicyIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreatePolicyIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Type)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(AnyType)
	err = __goidl__.IdlAnyHelper.Write(stream, self.Val)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreatePolicyIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreatePolicyInHelper = CorbaOrbCreatePolicyIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreatePolicyInId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreatePolicyIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreatePolicyIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreatePolicyIn{}
			},
			__reflect__.TypeOf((*CorbaOrbCreatePolicyIn)(nil))))
}
