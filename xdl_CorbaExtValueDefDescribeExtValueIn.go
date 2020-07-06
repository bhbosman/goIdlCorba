// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtValueDef_describe_ext_value_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtValueDefDescribeExtValueIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtValueDefDescribeExtValueInId_Const = "IDL:CORBA/ExtValueDef_describe_ext_value_In:1.0"

func (self *CorbaExtValueDefDescribeExtValueIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtValueDefDescribeExtValueIn) GoString() string {
	return self.String()
}

func (self *CorbaExtValueDefDescribeExtValueIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefDescribeExtValueIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefDescribeExtValueIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtValueDefDescribeExtValueIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtValueDefDescribeExtValueInHelper = CorbaExtValueDefDescribeExtValueIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtValueDefDescribeExtValueInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtValueDefDescribeExtValueIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtValueDefDescribeExtValueIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtValueDefDescribeExtValueIn{}
			},
			__reflect__.TypeOf((*CorbaExtValueDefDescribeExtValueIn)(nil))))
}
