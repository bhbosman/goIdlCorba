// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::OpDescriptionSeq", generatedBy by: "WriteStructSequenceDcl"
type CorbaOpDescriptionSeq struct {
	__goidl__.IdlObject
	Array []*CorbaOperationDescription `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaOpDescriptionSeqId_Const = "IDL:CORBA/OpDescriptionSeq:1.0"

func (self *CorbaOpDescriptionSeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOpDescriptionSeq) GoString() string {
	return self.String()
}

func (self *CorbaOpDescriptionSeq) ReadValue(stream __goidl__.IReadAny) error {
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
		self.Array = make([]*CorbaOperationDescription, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i] = &CorbaOperationDescription{}
			err = self.Array[i].ReadValue(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *CorbaOpDescriptionSeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOpDescriptionSeq) Write(stream __goidl__.IWriteAny) error {
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
type CorbaOpDescriptionSeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOpDescriptionSeqHelper = CorbaOpDescriptionSeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOpDescriptionSeqId_Const,
			__goidl__.SequenceType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaOpDescriptionSeq.go",
			func() __goidl__.IIdlObject {
				return &CorbaOpDescriptionSeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOpDescriptionSeq{}
			},
			__reflect__.TypeOf((*CorbaOpDescriptionSeq)(nil))))
}
