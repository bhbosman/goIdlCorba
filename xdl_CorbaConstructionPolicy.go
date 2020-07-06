// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::ConstructionPolicy", generatedBy by: "WriteInterface"
type CorbaConstructionPolicy interface {
	// Original name: "copy"
	Copy(params CorbaConstructionPolicyCopyIn) (CorbaConstructionPolicyCopyOut, error)
	// Original name: "destroy"
	Destroy(params CorbaConstructionPolicyDestroyIn) (CorbaConstructionPolicyDestroyOut, error)
	// Original name: "make_domain_manager"
	MakeDomainManager(params CorbaConstructionPolicyMakeDomainManagerIn) (CorbaConstructionPolicyMakeDomainManagerOut, error)
}

const CorbaConstructionPolicyCopyOperation = "copy"
const CorbaConstructionPolicyDestroyOperation = "destroy"
const CorbaConstructionPolicyMakeDomainManagerOperation = "make_domain_manager"
//noinspection GoSnakeCaseUsage
type CorbaConstructionPolicy_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaConstructionPolicyId_Const = "IDL:omg.org/CORBA/ConstructionPolicy:1.0"

func (self CorbaConstructionPolicy_Helper) Id() string {
	return CorbaConstructionPolicyId_Const
}

func (self CorbaConstructionPolicy_Helper) Read(stream __goidl__.IReadAny) (CorbaConstructionPolicy, error) {
	return nil, nil
}

func (self CorbaConstructionPolicy_Helper) Write(stream __goidl__.IWriteAny, v CorbaConstructionPolicy) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaConstructionPolicyHelper = CorbaConstructionPolicy_Helper{}

func init() {
}
