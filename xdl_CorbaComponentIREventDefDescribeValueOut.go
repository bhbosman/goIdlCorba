// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::EventDef_describe_value_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIREventDefDescribeValueOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIREventDefDescribeValueOutId_Const = "IDL:CORBA/ComponentIR/EventDef_describe_value_Out:1.0"

func (self *CorbaComponentIREventDefDescribeValueOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIREventDefDescribeValueOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIREventDefDescribeValueOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefDescribeValueOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefDescribeValueOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIREventDefDescribeValueOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIREventDefDescribeValueOutHelper = CorbaComponentIREventDefDescribeValueOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIREventDefDescribeValueOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIREventDefDescribeValueOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIREventDefDescribeValueOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIREventDefDescribeValueOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIREventDefDescribeValueOut)(nil))))
}
