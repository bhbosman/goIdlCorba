// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::AbstractInterfaceDef_describe_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaAbstractInterfaceDefDescribeInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaAbstractInterfaceDefDescribeInterfaceOutId_Const = "IDL:CORBA/AbstractInterfaceDef_describe_interface_Out:1.0"

func (self *CorbaAbstractInterfaceDefDescribeInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaAbstractInterfaceDefDescribeInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaAbstractInterfaceDefDescribeInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaAbstractInterfaceDefDescribeInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaAbstractInterfaceDefDescribeInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaAbstractInterfaceDefDescribeInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaAbstractInterfaceDefDescribeInterfaceOutHelper = CorbaAbstractInterfaceDefDescribeInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaAbstractInterfaceDefDescribeInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaAbstractInterfaceDefDescribeInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaAbstractInterfaceDefDescribeInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaAbstractInterfaceDefDescribeInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaAbstractInterfaceDefDescribeInterfaceOut)(nil))))
}
