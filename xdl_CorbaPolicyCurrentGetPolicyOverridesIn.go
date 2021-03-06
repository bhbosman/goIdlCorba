// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::PolicyCurrent_get_policy_overrides_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPolicyCurrentGetPolicyOverridesIn struct {
	__goidl__.IdlObject
	Ts CorbaPolicyTypeSeq `json:"Ts"`
}

//noinspection GoSnakeCaseUsage
const CorbaPolicyCurrentGetPolicyOverridesInId_Const = "IDL:CORBA/PolicyCurrent_get_policy_overrides_In:1.0"

func (self *CorbaPolicyCurrentGetPolicyOverridesIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPolicyCurrentGetPolicyOverridesIn) GoString() string {
	return self.String()
}

func (self *CorbaPolicyCurrentGetPolicyOverridesIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.Ts.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPolicyCurrentGetPolicyOverridesIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPolicyCurrentGetPolicyOverridesIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.Ts.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPolicyCurrentGetPolicyOverridesIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPolicyCurrentGetPolicyOverridesInHelper = CorbaPolicyCurrentGetPolicyOverridesIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPolicyCurrentGetPolicyOverridesInId_Const,
			__goidl__.StructType,
			"CORBA_Policy.idl",
			"xdl_CorbaPolicyCurrentGetPolicyOverridesIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaPolicyCurrentGetPolicyOverridesIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPolicyCurrentGetPolicyOverridesIn{}
			},
			__reflect__.TypeOf((*CorbaPolicyCurrentGetPolicyOverridesIn)(nil))))
}
