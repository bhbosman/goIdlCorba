// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::ComponentIR::FactoryDef", generatedBy by: "WriteInterface"
type CorbaComponentIRFactoryDef interface {
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRFactoryDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRFactoryDefId_Const = "IDL:omg.org/CORBA/ComponentIR/FactoryDef:1.0"

func (self CorbaComponentIRFactoryDef_Helper) Id() string {
	return CorbaComponentIRFactoryDefId_Const
}

func (self CorbaComponentIRFactoryDef_Helper) Read(stream __goidl__.IReadAny) (CorbaComponentIRFactoryDef, error) {
	return nil, nil
}

func (self CorbaComponentIRFactoryDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaComponentIRFactoryDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRFactoryDefHelper = CorbaComponentIRFactoryDef_Helper{}

func init() {
}