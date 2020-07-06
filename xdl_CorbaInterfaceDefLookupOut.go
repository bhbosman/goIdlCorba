// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_lookup_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefLookupOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefLookupOutId_Const = "IDL:CORBA/InterfaceDef_lookup_Out:1.0"

func (self *CorbaInterfaceDefLookupOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefLookupOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefLookupOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefLookupOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefLookupOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefLookupOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefLookupOutHelper = CorbaInterfaceDefLookupOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefLookupOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefLookupOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefLookupOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefLookupOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefLookupOut)(nil))))
}