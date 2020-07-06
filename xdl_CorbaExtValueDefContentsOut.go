// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtValueDef_contents_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtValueDefContentsOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtValueDefContentsOutId_Const = "IDL:CORBA/ExtValueDef_contents_Out:1.0"

func (self *CorbaExtValueDefContentsOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtValueDefContentsOut) GoString() string {
	return self.String()
}

func (self *CorbaExtValueDefContentsOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefContentsOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtValueDefContentsOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtValueDefContentsOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtValueDefContentsOutHelper = CorbaExtValueDefContentsOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtValueDefContentsOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtValueDefContentsOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtValueDefContentsOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtValueDefContentsOut{}
			},
			__reflect__.TypeOf((*CorbaExtValueDefContentsOut)(nil))))
}
