// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_content_type_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeContentTypeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeContentTypeOutId_Const = "IDL:CORBA/TypeCode_content_type_Out:1.0"

func (self *CorbaTypeCodeContentTypeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeContentTypeOut) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeContentTypeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeContentTypeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeContentTypeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeContentTypeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeContentTypeOutHelper = CorbaTypeCodeContentTypeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeContentTypeOutId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeContentTypeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeContentTypeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeContentTypeOut{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeContentTypeOut)(nil))))
}
