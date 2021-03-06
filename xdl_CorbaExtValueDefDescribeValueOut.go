// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtValueDef_describe_value_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtValueDefDescribeValueOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtValueDefDescribeValueOutId_Const = "IDL:CORBA/ExtValueDef_describe_value_Out:1.0"

func (self *CorbaExtValueDefDescribeValueOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtValueDefDescribeValueOut) GoString() string {
	return self.String()
}

func (self *CorbaExtValueDefDescribeValueOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefDescribeValueOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefDescribeValueOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtValueDefDescribeValueOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtValueDefDescribeValueOutHelper = CorbaExtValueDefDescribeValueOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtValueDefDescribeValueOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtValueDefDescribeValueOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtValueDefDescribeValueOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtValueDefDescribeValueOut{}
			},
			__reflect__.TypeOf((*CorbaExtValueDefDescribeValueOut)(nil))))
}
