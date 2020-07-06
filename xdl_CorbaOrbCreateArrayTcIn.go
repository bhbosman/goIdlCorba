// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_array_tc_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateArrayTcIn struct {
	__goidl__.IdlObject
	Length uint32 `json:"Length"`
	ElementType CorbaTypeCode `json:"ElementType"`
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateArrayTcInId_Const = "IDL:CORBA/ORB_create_array_tc_In:1.0"

func (self *CorbaOrbCreateArrayTcIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateArrayTcIn) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateArrayTcIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.Length, err = __goidl__.IdlUInt32Helper.Read(stream)
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

func (self *CorbaOrbCreateArrayTcIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateArrayTcIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Length)
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
type CorbaOrbCreateArrayTcIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateArrayTcInHelper = CorbaOrbCreateArrayTcIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateArrayTcInId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateArrayTcIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateArrayTcIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateArrayTcIn{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateArrayTcIn)(nil))))
}