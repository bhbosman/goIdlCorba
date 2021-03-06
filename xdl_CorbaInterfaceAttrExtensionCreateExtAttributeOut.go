// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceAttrExtension_create_ext_attribute_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceAttrExtensionCreateExtAttributeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceAttrExtensionCreateExtAttributeOutId_Const = "IDL:CORBA/InterfaceAttrExtension_create_ext_attribute_Out:1.0"

func (self *CorbaInterfaceAttrExtensionCreateExtAttributeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceAttrExtensionCreateExtAttributeOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceAttrExtensionCreateExtAttributeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceAttrExtensionCreateExtAttributeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceAttrExtensionCreateExtAttributeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceAttrExtensionCreateExtAttributeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceAttrExtensionCreateExtAttributeOutHelper = CorbaInterfaceAttrExtensionCreateExtAttributeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceAttrExtensionCreateExtAttributeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceAttrExtensionCreateExtAttributeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceAttrExtensionCreateExtAttributeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceAttrExtensionCreateExtAttributeOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceAttrExtensionCreateExtAttributeOut)(nil))))
}
