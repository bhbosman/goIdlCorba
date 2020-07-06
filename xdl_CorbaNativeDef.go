// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::NativeDef", generatedBy by: "WriteInterface"
type CorbaNativeDef interface {
	// Original name: "destroy"
	Destroy(params CorbaNativeDefDestroyIn) (CorbaNativeDefDestroyOut, error)
}

const CorbaNativeDefDestroyOperation = "destroy"
//noinspection GoSnakeCaseUsage
type CorbaNativeDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaNativeDefId_Const = "IDL:omg.org/CORBA/NativeDef:1.0"

func (self CorbaNativeDef_Helper) Id() string {
	return CorbaNativeDefId_Const
}

func (self CorbaNativeDef_Helper) Read(stream __goidl__.IReadAny) (CorbaNativeDef, error) {
	return nil, nil
}

func (self CorbaNativeDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaNativeDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaNativeDefHelper = CorbaNativeDef_Helper{}

func init() {
}
