// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_describe_contents_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefDescribeContentsOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefDescribeContentsOutId_Const = "IDL:CORBA/InterfaceDef_describe_contents_Out:1.0"

func (self *CorbaInterfaceDefDescribeContentsOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefDescribeContentsOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefDescribeContentsOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefDescribeContentsOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefDescribeContentsOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefDescribeContentsOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefDescribeContentsOutHelper = CorbaInterfaceDefDescribeContentsOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefDescribeContentsOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefDescribeContentsOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefDescribeContentsOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefDescribeContentsOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefDescribeContentsOut)(nil))))
}
