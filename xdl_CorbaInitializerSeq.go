// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InitializerSeq", generatedBy by: "WriteStructSequenceDcl"
type CorbaInitializerSeq struct {
	__goidl__.IdlObject
	Array []*CorbaInitializer `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaInitializerSeqId_Const = "IDL:CORBA/InitializerSeq:1.0"

func (self *CorbaInitializerSeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInitializerSeq) GoString() string {
	return self.String()
}

func (self *CorbaInitializerSeq) ReadValue(stream __goidl__.IReadAny) error {
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
		self.Array = make([]*CorbaInitializer, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i] = &CorbaInitializer{}
			err = self.Array[i].ReadValue(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *CorbaInitializerSeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInitializerSeq) Write(stream __goidl__.IWriteAny) error {
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
type CorbaInitializerSeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInitializerSeqHelper = CorbaInitializerSeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInitializerSeqId_Const,
			__goidl__.SequenceType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInitializerSeq.go",
			func() __goidl__.IIdlObject {
				return &CorbaInitializerSeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInitializerSeq{}
			},
			__reflect__.TypeOf((*CorbaInitializerSeq)(nil))))
}
