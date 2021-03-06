// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ShortSeq", generatedBy by: "WriteStructSequenceDcl"
type CorbaShortSeq struct {
	__goidl__.IdlObject
	Array []int16 `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaShortSeqId_Const = "IDL:CORBA/ShortSeq:1.0"

func (self *CorbaShortSeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaShortSeq) GoString() string {
	return self.String()
}

func (self *CorbaShortSeq) ReadValue(stream __goidl__.IReadAny) error {
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
		self.Array = make([]int16, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i], err = __goidl__.IdlInt16Helper.Read(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *CorbaShortSeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaShortSeq) Write(stream __goidl__.IWriteAny) error {
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
			err = __goidl__.IdlInt16Helper.Write(stream, item)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaShortSeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaShortSeqHelper = CorbaShortSeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaShortSeqId_Const,
			__goidl__.SequenceType,
			"CORBA_Stream.idl",
			"xdl_CorbaShortSeq.go",
			func() __goidl__.IIdlObject {
				return &CorbaShortSeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaShortSeq{}
			},
			__reflect__.TypeOf((*CorbaShortSeq)(nil))))
}
