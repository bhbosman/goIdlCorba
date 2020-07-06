// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Container_create_abstract_interface_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContainerCreateAbstractInterfaceOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaContainerCreateAbstractInterfaceOutId_Const = "IDL:CORBA/Container_create_abstract_interface_Out:1.0"

func (self *CorbaContainerCreateAbstractInterfaceOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContainerCreateAbstractInterfaceOut) GoString() string {
	return self.String()
}

func (self *CorbaContainerCreateAbstractInterfaceOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerCreateAbstractInterfaceOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerCreateAbstractInterfaceOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContainerCreateAbstractInterfaceOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContainerCreateAbstractInterfaceOutHelper = CorbaContainerCreateAbstractInterfaceOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContainerCreateAbstractInterfaceOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaContainerCreateAbstractInterfaceOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaContainerCreateAbstractInterfaceOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContainerCreateAbstractInterfaceOut{}
			},
			__reflect__.TypeOf((*CorbaContainerCreateAbstractInterfaceOut)(nil))))
}
