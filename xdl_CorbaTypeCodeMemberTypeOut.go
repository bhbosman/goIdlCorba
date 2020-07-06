// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_member_type_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeMemberTypeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeMemberTypeOutId_Const = "IDL:CORBA/TypeCode_member_type_Out:1.0"

func (self *CorbaTypeCodeMemberTypeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeMemberTypeOut) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeMemberTypeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeMemberTypeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeMemberTypeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeMemberTypeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeMemberTypeOutHelper = CorbaTypeCodeMemberTypeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeMemberTypeOutId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeMemberTypeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeMemberTypeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeMemberTypeOut{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeMemberTypeOut)(nil))))
}