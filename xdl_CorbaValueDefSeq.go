// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ValueDefSeq", generatedBy by: "WriteStructSequenceDcl"
type CorbaValueDefSeq struct {
	__goidl__.IdlObject
	Array []CorbaValueDef `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaValueDefSeqId_Const = "IDL:CORBA/ValueDefSeq:1.0"

func (self *CorbaValueDefSeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaValueDefSeq) GoString() string {
	return self.String()
}

func (self *CorbaValueDefSeq) ReadValue(stream __goidl__.IReadAny) error {
	return nil
}

func (self *CorbaValueDefSeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefSeq) Write(stream __goidl__.IWriteAny) error {
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaValueDefSeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaValueDefSeqHelper = CorbaValueDefSeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaValueDefSeqId_Const,
			__goidl__.SequenceType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaValueDefSeq.go",
			func() __goidl__.IIdlObject {
				return &CorbaValueDefSeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaValueDefSeq{}
			},
			__reflect__.TypeOf((*CorbaValueDefSeq)(nil))))
}