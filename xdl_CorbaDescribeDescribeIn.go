// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::describe_describe_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDescribeDescribeIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDescribeDescribeInId_Const = "IDL:CORBA/describe_describe_In:1.0"

func (self *CorbaDescribeDescribeIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDescribeDescribeIn) GoString() string {
	return self.String()
}

func (self *CorbaDescribeDescribeIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDescribeDescribeIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDescribeDescribeIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDescribeDescribeIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDescribeDescribeInHelper = CorbaDescribeDescribeIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDescribeDescribeInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaDescribeDescribeIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDescribeDescribeIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDescribeDescribeIn{}
			},
			__reflect__.TypeOf((*CorbaDescribeDescribeIn)(nil))))
}
