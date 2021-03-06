// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_member_count_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeMemberCountOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeMemberCountOutId_Const = "IDL:CORBA/TypeCode_member_count_Out:1.0"

func (self *CorbaTypeCodeMemberCountOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeMemberCountOut) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeMemberCountOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeMemberCountOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeMemberCountOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeMemberCountOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeMemberCountOutHelper = CorbaTypeCodeMemberCountOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeMemberCountOutId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeMemberCountOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeMemberCountOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeMemberCountOut{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeMemberCountOut)(nil))))
}
