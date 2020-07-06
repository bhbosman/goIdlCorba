// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_lookup_value_factory_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbLookupValueFactoryOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaOrbLookupValueFactoryOutId_Const = "IDL:CORBA/ORB_lookup_value_factory_Out:1.0"

func (self *CorbaOrbLookupValueFactoryOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbLookupValueFactoryOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbLookupValueFactoryOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbLookupValueFactoryOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbLookupValueFactoryOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbLookupValueFactoryOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbLookupValueFactoryOutHelper = CorbaOrbLookupValueFactoryOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbLookupValueFactoryOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbLookupValueFactoryOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbLookupValueFactoryOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbLookupValueFactoryOut{}
			},
			__reflect__.TypeOf((*CorbaOrbLookupValueFactoryOut)(nil))))
}