// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ParDescriptionSeq", generatedBy by: "WriteStructSequenceDcl"
type CorbaParDescriptionSeq struct {
	__goidl__.IdlObject
	Array []*CorbaParameterDescription `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaParDescriptionSeqId_Const = "IDL:CORBA/ParDescriptionSeq:1.0"

func (self *CorbaParDescriptionSeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaParDescriptionSeq) GoString() string {
	return self.String()
}

func (self *CorbaParDescriptionSeq) ReadValue(stream __goidl__.IReadAny) error {
	err := self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	var n uint32
	n, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	if n > 0 {
		self.Array = make([]*CorbaParameterDescription, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i] = &CorbaParameterDescription{}
			err = self.Array[i].ReadValue(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *CorbaParDescriptionSeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaParDescriptionSeq) Write(stream __goidl__.IWriteAny) error {
	err := self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	err = __goidl__.IdlUInt32Helper.Write(stream, uint32(len(self.Array)))
	if err != nil {
	return err
	}
	if len(self.Array) > 0 {
		for _, item := range self.Array {
			err = item.Write(stream)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaParDescriptionSeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaParDescriptionSeqHelper = CorbaParDescriptionSeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaParDescriptionSeqId_Const,
			__goidl__.SequenceType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaParDescriptionSeq.go",
			func() __goidl__.IIdlObject {
				return &CorbaParDescriptionSeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaParDescriptionSeq{}
			},
			__reflect__.TypeOf((*CorbaParDescriptionSeq)(nil))))
}
