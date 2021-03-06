// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ValueDef_create_operation_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaValueDefCreateOperationOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaValueDefCreateOperationOutId_Const = "IDL:CORBA/ValueDef_create_operation_Out:1.0"

func (self *CorbaValueDefCreateOperationOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaValueDefCreateOperationOut) GoString() string {
	return self.String()
}

func (self *CorbaValueDefCreateOperationOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefCreateOperationOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefCreateOperationOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaValueDefCreateOperationOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaValueDefCreateOperationOutHelper = CorbaValueDefCreateOperationOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaValueDefCreateOperationOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaValueDefCreateOperationOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaValueDefCreateOperationOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaValueDefCreateOperationOut{}
			},
			__reflect__.TypeOf((*CorbaValueDefCreateOperationOut)(nil))))
}
