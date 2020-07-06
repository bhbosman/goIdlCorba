// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::PolicyManager_get_policy_overrides_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaPolicyManagerGetPolicyOverridesOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaPolicyManagerGetPolicyOverridesOutId_Const = "IDL:CORBA/PolicyManager_get_policy_overrides_Out:1.0"

func (self *CorbaPolicyManagerGetPolicyOverridesOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPolicyManagerGetPolicyOverridesOut) GoString() string {
	return self.String()
}

func (self *CorbaPolicyManagerGetPolicyOverridesOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPolicyManagerGetPolicyOverridesOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPolicyManagerGetPolicyOverridesOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPolicyManagerGetPolicyOverridesOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPolicyManagerGetPolicyOverridesOutHelper = CorbaPolicyManagerGetPolicyOverridesOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPolicyManagerGetPolicyOverridesOutId_Const,
			__goidl__.StructType,
			"CORBA_Policy.idl",
			"xdl_CorbaPolicyManagerGetPolicyOverridesOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaPolicyManagerGetPolicyOverridesOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPolicyManagerGetPolicyOverridesOut{}
			},
			__reflect__.TypeOf((*CorbaPolicyManagerGetPolicyOverridesOut)(nil))))
}
