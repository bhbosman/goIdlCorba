// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_fixed_scale_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeFixedScaleIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeFixedScaleInId_Const = "IDL:CORBA/TypeCode_fixed_scale_In:1.0"

func (self *CorbaTypeCodeFixedScaleIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeFixedScaleIn) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeFixedScaleIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeFixedScaleIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeFixedScaleIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeFixedScaleIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeFixedScaleInHelper = CorbaTypeCodeFixedScaleIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeFixedScaleInId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeFixedScaleIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeFixedScaleIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeFixedScaleIn{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeFixedScaleIn)(nil))))
}