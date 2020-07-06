// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Container_create_ext_value_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContainerCreateExtValueOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaContainerCreateExtValueOutId_Const = "IDL:CORBA/Container_create_ext_value_Out:1.0"

func (self *CorbaContainerCreateExtValueOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContainerCreateExtValueOut) GoString() string {
	return self.String()
}

func (self *CorbaContainerCreateExtValueOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerCreateExtValueOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContainerCreateExtValueOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContainerCreateExtValueOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContainerCreateExtValueOutHelper = CorbaContainerCreateExtValueOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContainerCreateExtValueOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaContainerCreateExtValueOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaContainerCreateExtValueOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContainerCreateExtValueOut{}
			},
			__reflect__.TypeOf((*CorbaContainerCreateExtValueOut)(nil))))
}
