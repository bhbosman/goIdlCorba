// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::StructDef_describe_contents_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaStructDefDescribeContentsOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaStructDefDescribeContentsOutId_Const = "IDL:CORBA/StructDef_describe_contents_Out:1.0"

func (self *CorbaStructDefDescribeContentsOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaStructDefDescribeContentsOut) GoString() string {
	return self.String()
}

func (self *CorbaStructDefDescribeContentsOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaStructDefDescribeContentsOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaStructDefDescribeContentsOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaStructDefDescribeContentsOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaStructDefDescribeContentsOutHelper = CorbaStructDefDescribeContentsOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaStructDefDescribeContentsOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaStructDefDescribeContentsOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaStructDefDescribeContentsOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaStructDefDescribeContentsOut{}
			},
			__reflect__.TypeOf((*CorbaStructDefDescribeContentsOut)(nil))))
}
