// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_content_type_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeContentTypeIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeContentTypeInId_Const = "IDL:CORBA/TypeCode_content_type_In:1.0"

func (self *CorbaTypeCodeContentTypeIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeContentTypeIn) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeContentTypeIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeContentTypeIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeContentTypeIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeContentTypeIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeContentTypeInHelper = CorbaTypeCodeContentTypeIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeContentTypeInId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeContentTypeIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeContentTypeIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeContentTypeIn{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeContentTypeIn)(nil))))
}
