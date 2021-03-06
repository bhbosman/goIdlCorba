// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_discriminator_type_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeDiscriminatorTypeIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeDiscriminatorTypeInId_Const = "IDL:CORBA/TypeCode_discriminator_type_In:1.0"

func (self *CorbaTypeCodeDiscriminatorTypeIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeDiscriminatorTypeIn) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeDiscriminatorTypeIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeDiscriminatorTypeIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeDiscriminatorTypeIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeDiscriminatorTypeIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeDiscriminatorTypeInHelper = CorbaTypeCodeDiscriminatorTypeIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeDiscriminatorTypeInId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeDiscriminatorTypeIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeDiscriminatorTypeIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeDiscriminatorTypeIn{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeDiscriminatorTypeIn)(nil))))
}
