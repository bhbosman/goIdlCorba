// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::BooleanSeq", generatedBy by: "WriteStructSequenceDcl"
type CorbaBooleanSeq struct {
	__goidl__.IdlObject
	Array []bool `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaBooleanSeqId_Const = "IDL:CORBA/BooleanSeq:1.0"

func (self *CorbaBooleanSeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaBooleanSeq) GoString() string {
	return self.String()
}

func (self *CorbaBooleanSeq) ReadValue(stream __goidl__.IReadAny) error {
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
		self.Array = make([]bool, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i], err = __goidl__.IdlBooleanHelper.Read(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *CorbaBooleanSeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaBooleanSeq) Write(stream __goidl__.IWriteAny) error {
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
			err = __goidl__.IdlBooleanHelper.Write(stream, item)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaBooleanSeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaBooleanSeqHelper = CorbaBooleanSeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaBooleanSeqId_Const,
			__goidl__.SequenceType,
			"CORBA_Stream.idl",
			"xdl_CorbaBooleanSeq.go",
			func() __goidl__.IIdlObject {
				return &CorbaBooleanSeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaBooleanSeq{}
			},
			__reflect__.TypeOf((*CorbaBooleanSeq)(nil))))
}
