// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::TypeCode_get_compact_typecode_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaTypeCodeGetCompactTypecodeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeGetCompactTypecodeOutId_Const = "IDL:CORBA/TypeCode_get_compact_typecode_Out:1.0"

func (self *CorbaTypeCodeGetCompactTypecodeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeGetCompactTypecodeOut) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeGetCompactTypecodeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeGetCompactTypecodeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeGetCompactTypecodeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeGetCompactTypecodeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeGetCompactTypecodeOutHelper = CorbaTypeCodeGetCompactTypecodeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeGetCompactTypecodeOutId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeGetCompactTypecodeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeGetCompactTypecodeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeGetCompactTypecodeOut{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeGetCompactTypecodeOut)(nil))))
}