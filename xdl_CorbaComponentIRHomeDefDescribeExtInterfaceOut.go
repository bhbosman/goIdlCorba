// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::HomeDef_describe_ext_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRHomeDefDescribeExtInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRHomeDefDescribeExtInterfaceOutId_Const = "IDL:CORBA/ComponentIR/HomeDef_describe_ext_interface_Out:1.0"

func (self *CorbaComponentIRHomeDefDescribeExtInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRHomeDefDescribeExtInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRHomeDefDescribeExtInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefDescribeExtInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefDescribeExtInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRHomeDefDescribeExtInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRHomeDefDescribeExtInterfaceOutHelper = CorbaComponentIRHomeDefDescribeExtInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRHomeDefDescribeExtInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRHomeDefDescribeExtInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefDescribeExtInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefDescribeExtInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRHomeDefDescribeExtInterfaceOut)(nil))))
}
