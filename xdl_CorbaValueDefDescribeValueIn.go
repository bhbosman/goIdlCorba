// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ValueDef_describe_value_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaValueDefDescribeValueIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaValueDefDescribeValueInId_Const = "IDL:CORBA/ValueDef_describe_value_In:1.0"

func (self *CorbaValueDefDescribeValueIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaValueDefDescribeValueIn) GoString() string {
	return self.String()
}

func (self *CorbaValueDefDescribeValueIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefDescribeValueIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefDescribeValueIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaValueDefDescribeValueIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaValueDefDescribeValueInHelper = CorbaValueDefDescribeValueIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaValueDefDescribeValueInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaValueDefDescribeValueIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaValueDefDescribeValueIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaValueDefDescribeValueIn{}
			},
			__reflect__.TypeOf((*CorbaValueDefDescribeValueIn)(nil))))
}
