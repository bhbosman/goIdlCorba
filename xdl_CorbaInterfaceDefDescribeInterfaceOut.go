// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_describe_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefDescribeInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefDescribeInterfaceOutId_Const = "IDL:CORBA/InterfaceDef_describe_interface_Out:1.0"

func (self *CorbaInterfaceDefDescribeInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefDescribeInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefDescribeInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefDescribeInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefDescribeInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefDescribeInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefDescribeInterfaceOutHelper = CorbaInterfaceDefDescribeInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefDescribeInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefDescribeInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefDescribeInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefDescribeInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefDescribeInterfaceOut)(nil))))
}
