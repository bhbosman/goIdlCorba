// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceAttrExtension_describe_ext_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceAttrExtensionDescribeExtInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceAttrExtensionDescribeExtInterfaceOutId_Const = "IDL:CORBA/InterfaceAttrExtension_describe_ext_interface_Out:1.0"

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceAttrExtensionDescribeExtInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceAttrExtensionDescribeExtInterfaceOutHelper = CorbaInterfaceAttrExtensionDescribeExtInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceAttrExtensionDescribeExtInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceAttrExtensionDescribeExtInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceAttrExtensionDescribeExtInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceAttrExtensionDescribeExtInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceAttrExtensionDescribeExtInterfaceOut)(nil))))
}
