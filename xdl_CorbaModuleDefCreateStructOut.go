// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ModuleDef_create_struct_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaModuleDefCreateStructOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaModuleDefCreateStructOutId_Const = "IDL:CORBA/ModuleDef_create_struct_Out:1.0"

func (self *CorbaModuleDefCreateStructOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaModuleDefCreateStructOut) GoString() string {
	return self.String()
}

func (self *CorbaModuleDefCreateStructOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateStructOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaModuleDefCreateStructOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaModuleDefCreateStructOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaModuleDefCreateStructOutHelper = CorbaModuleDefCreateStructOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaModuleDefCreateStructOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaModuleDefCreateStructOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaModuleDefCreateStructOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaModuleDefCreateStructOut{}
			},
			__reflect__.TypeOf((*CorbaModuleDefCreateStructOut)(nil))))
}
