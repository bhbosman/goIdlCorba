// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::EventDef_describe_contents_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIREventDefDescribeContentsOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIREventDefDescribeContentsOutId_Const = "IDL:CORBA/ComponentIR/EventDef_describe_contents_Out:1.0"

func (self *CorbaComponentIREventDefDescribeContentsOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIREventDefDescribeContentsOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIREventDefDescribeContentsOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefDescribeContentsOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefDescribeContentsOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIREventDefDescribeContentsOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIREventDefDescribeContentsOutHelper = CorbaComponentIREventDefDescribeContentsOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIREventDefDescribeContentsOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIREventDefDescribeContentsOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIREventDefDescribeContentsOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIREventDefDescribeContentsOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIREventDefDescribeContentsOut)(nil))))
}