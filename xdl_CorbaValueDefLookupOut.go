// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ValueDef_lookup_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaValueDefLookupOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaValueDefLookupOutId_Const = "IDL:CORBA/ValueDef_lookup_Out:1.0"

func (self *CorbaValueDefLookupOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaValueDefLookupOut) GoString() string {
	return self.String()
}

func (self *CorbaValueDefLookupOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefLookupOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefLookupOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaValueDefLookupOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaValueDefLookupOutHelper = CorbaValueDefLookupOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaValueDefLookupOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaValueDefLookupOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaValueDefLookupOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaValueDefLookupOut{}
			},
			__reflect__.TypeOf((*CorbaValueDefLookupOut)(nil))))
}
