// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ModuleDef_contents_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRModuleDefContentsOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRModuleDefContentsOutId_Const = "IDL:CORBA/ComponentIR/ModuleDef_contents_Out:1.0"

func (self *CorbaComponentIRModuleDefContentsOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRModuleDefContentsOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRModuleDefContentsOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefContentsOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRModuleDefContentsOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRModuleDefContentsOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRModuleDefContentsOutHelper = CorbaComponentIRModuleDefContentsOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRModuleDefContentsOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRModuleDefContentsOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefContentsOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRModuleDefContentsOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRModuleDefContentsOut)(nil))))
}