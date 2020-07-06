// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_type_modifier_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeTypeModifierIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeTypeModifierInId_Const = "IDL:CORBA/TypeCode_type_modifier_In:1.0"

func (self *CorbaTypeCodeTypeModifierIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeTypeModifierIn) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeTypeModifierIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeTypeModifierIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeTypeModifierIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeTypeModifierIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeTypeModifierInHelper = CorbaTypeCodeTypeModifierIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeTypeModifierInId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeTypeModifierIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeTypeModifierIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeTypeModifierIn{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeTypeModifierIn)(nil))))
}