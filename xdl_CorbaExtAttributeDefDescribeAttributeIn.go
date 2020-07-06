// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtAttributeDef_describe_attribute_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtAttributeDefDescribeAttributeIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtAttributeDefDescribeAttributeInId_Const = "IDL:CORBA/ExtAttributeDef_describe_attribute_In:1.0"

func (self *CorbaExtAttributeDefDescribeAttributeIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtAttributeDefDescribeAttributeIn) GoString() string {
	return self.String()
}

func (self *CorbaExtAttributeDefDescribeAttributeIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAttributeDefDescribeAttributeIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAttributeDefDescribeAttributeIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtAttributeDefDescribeAttributeIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtAttributeDefDescribeAttributeInHelper = CorbaExtAttributeDefDescribeAttributeIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtAttributeDefDescribeAttributeInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtAttributeDefDescribeAttributeIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtAttributeDefDescribeAttributeIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtAttributeDefDescribeAttributeIn{}
			},
			__reflect__.TypeOf((*CorbaExtAttributeDefDescribeAttributeIn)(nil))))
}
