// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Container_create_component_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRContainerCreateComponentOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRContainerCreateComponentOutId_Const = "IDL:CORBA/ComponentIR/Container_create_component_Out:1.0"

func (self *CorbaComponentIRContainerCreateComponentOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRContainerCreateComponentOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRContainerCreateComponentOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRContainerCreateComponentOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRContainerCreateComponentOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRContainerCreateComponentOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRContainerCreateComponentOutHelper = CorbaComponentIRContainerCreateComponentOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRContainerCreateComponentOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRContainerCreateComponentOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRContainerCreateComponentOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRContainerCreateComponentOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRContainerCreateComponentOut)(nil))))
}
