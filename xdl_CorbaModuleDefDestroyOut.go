// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ModuleDef_destroy_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaModuleDefDestroyOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaModuleDefDestroyOutId_Const = "IDL:CORBA/ModuleDef_destroy_Out:1.0"

func (self *CorbaModuleDefDestroyOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaModuleDefDestroyOut) GoString() string {
	return self.String()
}

func (self *CorbaModuleDefDestroyOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefDestroyOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefDestroyOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaModuleDefDestroyOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaModuleDefDestroyOutHelper = CorbaModuleDefDestroyOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaModuleDefDestroyOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaModuleDefDestroyOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaModuleDefDestroyOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaModuleDefDestroyOut{}
			},
			__reflect__.TypeOf((*CorbaModuleDefDestroyOut)(nil))))
}
