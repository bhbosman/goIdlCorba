// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtAbstractInterfaceDef_describe_interface_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtAbstractInterfaceDefDescribeInterfaceIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtAbstractInterfaceDefDescribeInterfaceInId_Const = "IDL:CORBA/ExtAbstractInterfaceDef_describe_interface_In:1.0"

func (self *CorbaExtAbstractInterfaceDefDescribeInterfaceIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtAbstractInterfaceDefDescribeInterfaceIn) GoString() string {
	return self.String()
}

func (self *CorbaExtAbstractInterfaceDefDescribeInterfaceIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefDescribeInterfaceIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefDescribeInterfaceIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtAbstractInterfaceDefDescribeInterfaceIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtAbstractInterfaceDefDescribeInterfaceInHelper = CorbaExtAbstractInterfaceDefDescribeInterfaceIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtAbstractInterfaceDefDescribeInterfaceInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtAbstractInterfaceDefDescribeInterfaceIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefDescribeInterfaceIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefDescribeInterfaceIn{}
			},
			__reflect__.TypeOf((*CorbaExtAbstractInterfaceDefDescribeInterfaceIn)(nil))))
}
