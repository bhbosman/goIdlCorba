// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ModuleDef_lookup_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaModuleDefLookupOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaModuleDefLookupOutId_Const = "IDL:CORBA/ModuleDef_lookup_Out:1.0"

func (self *CorbaModuleDefLookupOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaModuleDefLookupOut) GoString() string {
	return self.String()
}

func (self *CorbaModuleDefLookupOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefLookupOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefLookupOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaModuleDefLookupOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaModuleDefLookupOutHelper = CorbaModuleDefLookupOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaModuleDefLookupOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaModuleDefLookupOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaModuleDefLookupOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaModuleDefLookupOut{}
			},
			__reflect__.TypeOf((*CorbaModuleDefLookupOut)(nil))))
}
