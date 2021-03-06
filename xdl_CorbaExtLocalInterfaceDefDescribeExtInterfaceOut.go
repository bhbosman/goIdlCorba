// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtLocalInterfaceDef_describe_ext_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtLocalInterfaceDefDescribeExtInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtLocalInterfaceDefDescribeExtInterfaceOutId_Const = "IDL:CORBA/ExtLocalInterfaceDef_describe_ext_interface_Out:1.0"

func (self *CorbaExtLocalInterfaceDefDescribeExtInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtLocalInterfaceDefDescribeExtInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaExtLocalInterfaceDefDescribeExtInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtLocalInterfaceDefDescribeExtInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtLocalInterfaceDefDescribeExtInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtLocalInterfaceDefDescribeExtInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtLocalInterfaceDefDescribeExtInterfaceOutHelper = CorbaExtLocalInterfaceDefDescribeExtInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtLocalInterfaceDefDescribeExtInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtLocalInterfaceDefDescribeExtInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefDescribeExtInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefDescribeExtInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaExtLocalInterfaceDefDescribeExtInterfaceOut)(nil))))
}
