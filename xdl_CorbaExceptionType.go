// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Enum declaration: "CORBA::exception_type", generatedBy by: "WriteEnumDcl"
type CorbaExceptionType uint32

//noinspection GoUnusedConst
const (
	CorbaExceptionTypeNoException     CorbaExceptionType = 0
	CorbaExceptionTypeUserException   CorbaExceptionType = 1
	CorbaExceptionTypeSystemException CorbaExceptionType = 2
)

//noinspection GoSnakeCaseUsage
type CorbaExceptionType_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaExceptionTypeId_Const = "IDL:omg.org/CORBA/exception_type:1.0"

func (self CorbaExceptionType_Helper) Id() string {
	return CorbaExceptionTypeId_Const
}

func (self CorbaExceptionType_Helper) Read(stream __goidl__.IReadAny) (uint32, error) {
	result, err := __goidl__.IdlUInt32Helper.Read(stream)
	return result, err
}

func (self CorbaExceptionType_Helper) Write(stream __goidl__.IWriteAny, v uint32) error {
	return __goidl__.IdlUInt32Helper.Write(stream, v)
}


//noinspection GoUnusedGlobalVariable
var CorbaExceptionTypeHelper = CorbaExceptionType_Helper{}

func init() {
}