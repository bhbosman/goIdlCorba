// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Repository_create_native_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRepositoryCreateNativeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRepositoryCreateNativeOutId_Const = "IDL:CORBA/Repository_create_native_Out:1.0"

func (self *CorbaRepositoryCreateNativeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRepositoryCreateNativeOut) GoString() string {
	return self.String()
}

func (self *CorbaRepositoryCreateNativeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateNativeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateNativeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRepositoryCreateNativeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRepositoryCreateNativeOutHelper = CorbaRepositoryCreateNativeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRepositoryCreateNativeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaRepositoryCreateNativeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaRepositoryCreateNativeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRepositoryCreateNativeOut{}
			},
			__reflect__.TypeOf((*CorbaRepositoryCreateNativeOut)(nil))))
}