// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Enum declaration: "CORBA::SetOverrideType", generatedBy by: "WriteEnumDcl"
type CorbaSetOverrideType uint32

//noinspection GoUnusedConst
const (
	CorbaSetOverrideTypeSetOverride CorbaSetOverrideType = 0
	CorbaSetOverrideTypeAddOverride CorbaSetOverrideType = 1
)

//noinspection GoSnakeCaseUsage
type CorbaSetOverrideType_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaSetOverrideTypeId_Const = "IDL:omg.org/CORBA/SetOverrideType:1.0"

func (self CorbaSetOverrideType_Helper) Id() string {
	return CorbaSetOverrideTypeId_Const
}

func (self CorbaSetOverrideType_Helper) Read(stream __goidl__.IReadAny) (uint32, error) {
	result, err := __goidl__.IdlUInt32Helper.Read(stream)
	return result, err
}

func (self CorbaSetOverrideType_Helper) Write(stream __goidl__.IWriteAny, v uint32) error {
	return __goidl__.IdlUInt32Helper.Write(stream, v)
}


//noinspection GoUnusedGlobalVariable
var CorbaSetOverrideTypeHelper = CorbaSetOverrideType_Helper{}

func init() {
}
