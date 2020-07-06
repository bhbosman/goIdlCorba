// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_create_native_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefCreateNativeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefCreateNativeOutId_Const = "IDL:CORBA/InterfaceDef_create_native_Out:1.0"

func (self *CorbaInterfaceDefCreateNativeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefCreateNativeOut) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefCreateNativeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateNativeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateNativeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefCreateNativeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefCreateNativeOutHelper = CorbaInterfaceDefCreateNativeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefCreateNativeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefCreateNativeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateNativeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateNativeOut{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefCreateNativeOut)(nil))))
}
