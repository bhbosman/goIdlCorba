// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ConstructionPolicy_copy_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaConstructionPolicyCopyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaConstructionPolicyCopyOutId_Const = "IDL:CORBA/ConstructionPolicy_copy_Out:1.0"

func (self *CorbaConstructionPolicyCopyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaConstructionPolicyCopyOut) GoString() string {
	return self.String()
}

func (self *CorbaConstructionPolicyCopyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaConstructionPolicyCopyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaConstructionPolicyCopyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaConstructionPolicyCopyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaConstructionPolicyCopyOutHelper = CorbaConstructionPolicyCopyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaConstructionPolicyCopyOutId_Const,
			__goidl__.StructType,
			"CORBA_DomainManager.idl",
			"xdl_CorbaConstructionPolicyCopyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaConstructionPolicyCopyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaConstructionPolicyCopyOut{}
			},
			__reflect__.TypeOf((*CorbaConstructionPolicyCopyOut)(nil))))
}
