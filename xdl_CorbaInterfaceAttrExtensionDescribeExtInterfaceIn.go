// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceAttrExtension_describe_ext_interface_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceAttrExtensionDescribeExtInterfaceIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceAttrExtensionDescribeExtInterfaceInId_Const = "IDL:CORBA/InterfaceAttrExtension_describe_ext_interface_In:1.0"

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceIn) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceAttrExtensionDescribeExtInterfaceIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceAttrExtensionDescribeExtInterfaceIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceAttrExtensionDescribeExtInterfaceInHelper = CorbaInterfaceAttrExtensionDescribeExtInterfaceIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceAttrExtensionDescribeExtInterfaceInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceAttrExtensionDescribeExtInterfaceIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceAttrExtensionDescribeExtInterfaceIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceAttrExtensionDescribeExtInterfaceIn{}
			},
			__reflect__.TypeOf((*CorbaInterfaceAttrExtensionDescribeExtInterfaceIn)(nil))))
}
