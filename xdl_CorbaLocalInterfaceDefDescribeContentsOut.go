// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::LocalInterfaceDef_describe_contents_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaLocalInterfaceDefDescribeContentsOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaLocalInterfaceDefDescribeContentsOutId_Const = "IDL:CORBA/LocalInterfaceDef_describe_contents_Out:1.0"

func (self *CorbaLocalInterfaceDefDescribeContentsOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaLocalInterfaceDefDescribeContentsOut) GoString() string {
	return self.String()
}

func (self *CorbaLocalInterfaceDefDescribeContentsOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefDescribeContentsOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefDescribeContentsOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaLocalInterfaceDefDescribeContentsOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaLocalInterfaceDefDescribeContentsOutHelper = CorbaLocalInterfaceDefDescribeContentsOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaLocalInterfaceDefDescribeContentsOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaLocalInterfaceDefDescribeContentsOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefDescribeContentsOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefDescribeContentsOut{}
			},
			__reflect__.TypeOf((*CorbaLocalInterfaceDefDescribeContentsOut)(nil))))
}