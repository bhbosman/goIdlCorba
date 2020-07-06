// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_sequence_tc_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateSequenceTcIn struct {
	__goidl__.IdlObject
	Bound uint32 `json:"Bound"`
	ElementType CorbaTypeCode `json:"ElementType"`
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateSequenceTcInId_Const = "IDL:CORBA/ORB_create_sequence_tc_In:1.0"

func (self *CorbaOrbCreateSequenceTcIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateSequenceTcIn) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateSequenceTcIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.Bound, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.ElementType, err = CorbaTypeCodeHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateSequenceTcIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateSequenceTcIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Bound)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaTypeCodeHelper.Write(stream, self.ElementType)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateSequenceTcIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateSequenceTcInHelper = CorbaOrbCreateSequenceTcIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateSequenceTcInId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateSequenceTcIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateSequenceTcIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateSequenceTcIn{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateSequenceTcIn)(nil))))
}