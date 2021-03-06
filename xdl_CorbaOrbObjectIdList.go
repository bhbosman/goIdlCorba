// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB::ObjectIdList", generatedBy by: "WriteStructSequenceDcl"
type CorbaOrbObjectIdList struct {
	__goidl__.IdlObject
	Array []string `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CorbaOrbObjectIdListId_Const = "IDL:CORBA/ORB/ObjectIdList:1.0"

func (self *CorbaOrbObjectIdList) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbObjectIdList) GoString() string {
	return self.String()
}

func (self *CorbaOrbObjectIdList) ReadValue(stream __goidl__.IReadAny) error {
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
		self.Array = make([]string, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i], err = __goidl__.IdlStringHelper.Read(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *CorbaOrbObjectIdList) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbObjectIdList) Write(stream __goidl__.IWriteAny) error {
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
			err = __goidl__.IdlStringHelper.Write(stream, item)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbObjectIdList_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbObjectIdListHelper = CorbaOrbObjectIdList_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbObjectIdListId_Const,
			__goidl__.SequenceType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbObjectIdList.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbObjectIdList{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbObjectIdList{}
			},
			__reflect__.TypeOf((*CorbaOrbObjectIdList)(nil))))
}
