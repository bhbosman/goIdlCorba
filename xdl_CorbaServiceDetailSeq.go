// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ServiceDetailSeq", generatedBy by: "WriteStructSequenceDcl"
type CorbaServiceDetailSeq struct {
	__goidl__.IdlObject
	Array []*CorbaServiceDetail `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaServiceDetailSeqId_Const = "IDL:CORBA/ServiceDetailSeq:1.0"

func (self *CorbaServiceDetailSeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaServiceDetailSeq) GoString() string {
	return self.String()
}

func (self *CorbaServiceDetailSeq) ReadValue(stream __goidl__.IReadAny) error {
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
		self.Array = make([]*CorbaServiceDetail, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i] = &CorbaServiceDetail{}
			err = self.Array[i].ReadValue(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *CorbaServiceDetailSeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaServiceDetailSeq) Write(stream __goidl__.IWriteAny) error {
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
type CorbaServiceDetailSeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaServiceDetailSeqHelper = CorbaServiceDetailSeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaServiceDetailSeqId_Const,
			__goidl__.SequenceType,
			"CORBA_ORB.idl",
			"xdl_CorbaServiceDetailSeq.go",
			func() __goidl__.IIdlObject {
				return &CorbaServiceDetailSeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaServiceDetailSeq{}
			},
			__reflect__.TypeOf((*CorbaServiceDetailSeq)(nil))))
}
