// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtAbstractInterfaceDef_create_operation_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtAbstractInterfaceDefCreateOperationOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaExtAbstractInterfaceDefCreateOperationOutId_Const = "IDL:CORBA/ExtAbstractInterfaceDef_create_operation_Out:1.0"

func (self *CorbaExtAbstractInterfaceDefCreateOperationOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtAbstractInterfaceDefCreateOperationOut) GoString() string {
	return self.String()
}

func (self *CorbaExtAbstractInterfaceDefCreateOperationOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefCreateOperationOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefCreateOperationOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtAbstractInterfaceDefCreateOperationOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtAbstractInterfaceDefCreateOperationOutHelper = CorbaExtAbstractInterfaceDefCreateOperationOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtAbstractInterfaceDefCreateOperationOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtAbstractInterfaceDefCreateOperationOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefCreateOperationOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefCreateOperationOut{}
			},
			__reflect__.TypeOf((*CorbaExtAbstractInterfaceDefCreateOperationOut)(nil))))
}
