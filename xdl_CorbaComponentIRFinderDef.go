// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::ComponentIR::FinderDef", generatedBy by: "WriteInterface"
type CorbaComponentIRFinderDef interface {
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRFinderDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRFinderDefId_Const = "IDL:omg.org/CORBA/ComponentIR/FinderDef:1.0"

func (self CorbaComponentIRFinderDef_Helper) Id() string {
	return CorbaComponentIRFinderDefId_Const
}

func (self CorbaComponentIRFinderDef_Helper) Read(stream __goidl__.IReadAny) (CorbaComponentIRFinderDef, error) {
	return nil, nil
}

func (self CorbaComponentIRFinderDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaComponentIRFinderDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRFinderDefHelper = CorbaComponentIRFinderDef_Helper{}

func init() {
}