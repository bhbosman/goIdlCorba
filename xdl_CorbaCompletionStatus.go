// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Enum declaration: "CORBA::completion_status", generatedBy by: "WriteEnumDcl"
type CorbaCompletionStatus uint32

//noinspection GoUnusedConst
const (
	CorbaCompletionStatusCompletedYes   CorbaCompletionStatus = 0
	CorbaCompletionStatusCompletedNo    CorbaCompletionStatus = 1
	CorbaCompletionStatusCompletedMaybe CorbaCompletionStatus = 2
)

//noinspection GoSnakeCaseUsage
type CorbaCompletionStatus_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaCompletionStatusId_Const = "IDL:omg.org/CORBA/completion_status:1.0"

func (self CorbaCompletionStatus_Helper) Id() string {
	return CorbaCompletionStatusId_Const
}

func (self CorbaCompletionStatus_Helper) Read(stream __goidl__.IReadAny) (uint32, error) {
	result, err := __goidl__.IdlUInt32Helper.Read(stream)
	return result, err
}

func (self CorbaCompletionStatus_Helper) Write(stream __goidl__.IWriteAny, v uint32) error {
	return __goidl__.IdlUInt32Helper.Write(stream, v)
}


//noinspection GoUnusedGlobalVariable
var CorbaCompletionStatusHelper = CorbaCompletionStatus_Helper{}

func init() {
}
