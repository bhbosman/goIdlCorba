// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Container_create_home_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRContainerCreateHomeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRContainerCreateHomeOutId_Const = "IDL:CORBA/ComponentIR/Container_create_home_Out:1.0"

func (self *CorbaComponentIRContainerCreateHomeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRContainerCreateHomeOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRContainerCreateHomeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRContainerCreateHomeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRContainerCreateHomeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRContainerCreateHomeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRContainerCreateHomeOutHelper = CorbaComponentIRContainerCreateHomeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRContainerCreateHomeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRContainerCreateHomeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRContainerCreateHomeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRContainerCreateHomeOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRContainerCreateHomeOut)(nil))))
}
