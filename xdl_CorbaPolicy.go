// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::Policy", generatedBy by: "WriteInterface"
type CorbaPolicy interface {
	// Original name: "copy"
	Copy(params CorbaPolicyCopyIn) (CorbaPolicyCopyOut, error)
	// Original name: "destroy"
	Destroy(params CorbaPolicyDestroyIn) (CorbaPolicyDestroyOut, error)
	// Property PolicyType
	// Get Property PolicyType
	GetPolicyType() (uint32, error)
}

const CorbaPolicyCopyOperation = "copy"
const CorbaPolicyDestroyOperation = "destroy"
//noinspection GoSnakeCaseUsage
type CorbaPolicy_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaPolicyId_Const = "IDL:omg.org/CORBA/Policy:1.0"

func (self CorbaPolicy_Helper) Id() string {
	return CorbaPolicyId_Const
}

func (self CorbaPolicy_Helper) Read(stream __goidl__.IReadAny) (CorbaPolicy, error) {
	return nil, nil
}

func (self CorbaPolicy_Helper) Write(stream __goidl__.IWriteAny, v CorbaPolicy) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaPolicyHelper = CorbaPolicy_Helper{}

func init() {
}
