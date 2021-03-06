// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_concrete_base_type_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeConcreteBaseTypeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeConcreteBaseTypeOutId_Const = "IDL:CORBA/TypeCode_concrete_base_type_Out:1.0"

func (self *CorbaTypeCodeConcreteBaseTypeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeConcreteBaseTypeOut) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeConcreteBaseTypeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeConcreteBaseTypeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeConcreteBaseTypeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeConcreteBaseTypeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeConcreteBaseTypeOutHelper = CorbaTypeCodeConcreteBaseTypeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeConcreteBaseTypeOutId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeConcreteBaseTypeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeConcreteBaseTypeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeConcreteBaseTypeOut{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeConcreteBaseTypeOut)(nil))))
}
