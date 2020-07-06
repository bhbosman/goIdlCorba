// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtLocalInterfaceDefSeq", generatedBy by: "WriteStructSequenceDcl"
type CorbaExtLocalInterfaceDefSeq struct {
	__goidl__.IdlObject
	Array []CorbaExtLocalInterfaceDef `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaExtLocalInterfaceDefSeqId_Const = "IDL:CORBA/ExtLocalInterfaceDefSeq:1.0"

func (self *CorbaExtLocalInterfaceDefSeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtLocalInterfaceDefSeq) GoString() string {
	return self.String()
}

func (self *CorbaExtLocalInterfaceDefSeq) ReadValue(stream __goidl__.IReadAny) error {
	return nil
}

func (self *CorbaExtLocalInterfaceDefSeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtLocalInterfaceDefSeq) Write(stream __goidl__.IWriteAny) error {
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtLocalInterfaceDefSeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtLocalInterfaceDefSeqHelper = CorbaExtLocalInterfaceDefSeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtLocalInterfaceDefSeqId_Const,
			__goidl__.SequenceType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtLocalInterfaceDefSeq.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefSeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefSeq{}
			},
			__reflect__.TypeOf((*CorbaExtLocalInterfaceDefSeq)(nil))))
}
