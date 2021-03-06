// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::ComponentIR::EventPortDef", generatedBy by: "WriteInterface"
type CorbaComponentIREventPortDef interface {
	// Original name: "is_a"
	IsA(params CorbaComponentIREventPortDefIsAIn) (CorbaComponentIREventPortDefIsAOut, error)
	// Property Event
	// Get Property Event
	GetEvent() (CorbaComponentIREventDef, error)
	// Set Property Event
	SetEvent(value CorbaComponentIREventDef) error
}

const CorbaComponentIREventPortDefIsAOperation = "is_a"
//noinspection GoSnakeCaseUsage
type CorbaComponentIREventPortDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIREventPortDefId_Const = "IDL:omg.org/CORBA/ComponentIR/EventPortDef:1.0"

func (self CorbaComponentIREventPortDef_Helper) Id() string {
	return CorbaComponentIREventPortDefId_Const
}

func (self CorbaComponentIREventPortDef_Helper) Read(stream __goidl__.IReadAny) (CorbaComponentIREventPortDef, error) {
	return nil, nil
}

func (self CorbaComponentIREventPortDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaComponentIREventPortDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIREventPortDefHelper = CorbaComponentIREventPortDef_Helper{}

func init() {
}
