// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::ValueMemberDef", generatedBy by: "WriteInterface"
type CorbaValueMemberDef interface {
	// Property Type
	// Get Property Type
	GetType() (CorbaTypeCode, error)
	// Property TypeDef
	// Get Property TypeDef
	GetTypeDef() (CorbaIDLType, error)
	// Set Property TypeDef
	SetTypeDef(value CorbaIDLType) error
	// Property Access
	// Get Property Access
	GetAccess() (int16, error)
	// Set Property Access
	SetAccess(value int16) error
}

//noinspection GoSnakeCaseUsage
type CorbaValueMemberDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaValueMemberDefId_Const = "IDL:omg.org/CORBA/ValueMemberDef:1.0"

func (self CorbaValueMemberDef_Helper) Id() string {
	return CorbaValueMemberDefId_Const
}

func (self CorbaValueMemberDef_Helper) Read(stream __goidl__.IReadAny) (CorbaValueMemberDef, error) {
	return nil, nil
}

func (self CorbaValueMemberDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaValueMemberDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaValueMemberDefHelper = CorbaValueMemberDef_Helper{}

func init() {
}
