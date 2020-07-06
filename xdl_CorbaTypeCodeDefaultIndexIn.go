// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_default_index_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeDefaultIndexIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeDefaultIndexInId_Const = "IDL:CORBA/TypeCode_default_index_In:1.0"

func (self *CorbaTypeCodeDefaultIndexIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeDefaultIndexIn) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeDefaultIndexIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeDefaultIndexIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeDefaultIndexIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeDefaultIndexIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeDefaultIndexInHelper = CorbaTypeCodeDefaultIndexIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeDefaultIndexInId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeDefaultIndexIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeDefaultIndexIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeDefaultIndexIn{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeDefaultIndexIn)(nil))))
}
