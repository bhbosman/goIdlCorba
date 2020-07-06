// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_member_type_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeMemberTypeIn struct {
	__goidl__.IdlObject
	Index uint32 `json:"Index"`
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeMemberTypeInId_Const = "IDL:CORBA/TypeCode_member_type_In:1.0"

func (self *CorbaTypeCodeMemberTypeIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeMemberTypeIn) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeMemberTypeIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.Index, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeMemberTypeIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeMemberTypeIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Index)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeMemberTypeIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeMemberTypeInHelper = CorbaTypeCodeMemberTypeIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeMemberTypeInId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeMemberTypeIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeMemberTypeIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeMemberTypeIn{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeMemberTypeIn)(nil))))
}
