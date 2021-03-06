// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_equal_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeEqualOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeEqualOutId_Const = "IDL:CORBA/TypeCode_equal_Out:1.0"

func (self *CorbaTypeCodeEqualOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeEqualOut) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeEqualOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeEqualOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeEqualOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeEqualOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeEqualOutHelper = CorbaTypeCodeEqualOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeEqualOutId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeEqualOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeEqualOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeEqualOut{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeEqualOut)(nil))))
}
