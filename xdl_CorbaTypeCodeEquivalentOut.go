// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_equivalent_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeEquivalentOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeEquivalentOutId_Const = "IDL:CORBA/TypeCode_equivalent_Out:1.0"

func (self *CorbaTypeCodeEquivalentOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeEquivalentOut) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeEquivalentOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeEquivalentOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeEquivalentOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeEquivalentOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeEquivalentOutHelper = CorbaTypeCodeEquivalentOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeEquivalentOutId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeEquivalentOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeEquivalentOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeEquivalentOut{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeEquivalentOut)(nil))))
}
