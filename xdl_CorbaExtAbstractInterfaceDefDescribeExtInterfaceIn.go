// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtAbstractInterfaceDef_describe_ext_interface_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtAbstractInterfaceDefDescribeExtInterfaceInId_Const = "IDL:CORBA/ExtAbstractInterfaceDef_describe_ext_interface_In:1.0"

func (self *CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn) GoString() string {
	return self.String()
}

func (self *CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtAbstractInterfaceDefDescribeExtInterfaceInHelper = CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtAbstractInterfaceDefDescribeExtInterfaceInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn{}
			},
			__reflect__.TypeOf((*CorbaExtAbstractInterfaceDefDescribeExtInterfaceIn)(nil))))
}
