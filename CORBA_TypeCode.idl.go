// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// TypeDef declaration: "CORBA::ValueModifier", generatedBy by: "WriteTypeDefOfPrimitiveDcl"
//Typedef Primitive declaration: "CORBA::ValueModifier" from: "int16"
type CorbaValueModifier int16

//noinspection GoSnakeCaseUsage
type CorbaValueModifier_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaValueModifierId_Const = "IDL:omg.org/CORBA/ValueModifier:1.0"

func (self CorbaValueModifier_Helper) Id() string {
	return CorbaValueModifierId_Const
}

func (self CorbaValueModifier_Helper) Read(stream __goidl__.IReadAny) (int16, error) {
	result, err := __goidl__.IdlInt16Helper.Read(stream)
	return result, err
}

func (self CorbaValueModifier_Helper) Write(stream __goidl__.IWriteAny, v int16) error {
	return __goidl__.IdlInt16Helper.Write(stream, v)
}

// Constant declaration: "CORBA::VM_NONE", generatedBy by: "WriteConstantValue"
//noinspection GoUnusedConst
const CorbaVmNone int16 = 0

// Constant declaration: "CORBA::VM_CUSTOM", generatedBy by: "WriteConstantValue"
//noinspection GoUnusedConst
const CorbaVmCustom int16 = 1

// Constant declaration: "CORBA::VM_ABSTRACT", generatedBy by: "WriteConstantValue"
//noinspection GoUnusedConst
const CorbaVmAbstract int16 = 2

// Constant declaration: "CORBA::VM_TRUNCATABLE", generatedBy by: "WriteConstantValue"
//noinspection GoUnusedConst
const CorbaVmTruncatable int16 = 3


//noinspection GoUnusedGlobalVariable
var CorbaValueModifierHelper = CorbaValueModifier_Helper{}

func init() {
}
