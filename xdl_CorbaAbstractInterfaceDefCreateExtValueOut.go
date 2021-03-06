// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::AbstractInterfaceDef_create_ext_value_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaAbstractInterfaceDefCreateExtValueOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaAbstractInterfaceDefCreateExtValueOutId_Const = "IDL:CORBA/AbstractInterfaceDef_create_ext_value_Out:1.0"

func (self *CorbaAbstractInterfaceDefCreateExtValueOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaAbstractInterfaceDefCreateExtValueOut) GoString() string {
	return self.String()
}

func (self *CorbaAbstractInterfaceDefCreateExtValueOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaAbstractInterfaceDefCreateExtValueOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaAbstractInterfaceDefCreateExtValueOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaAbstractInterfaceDefCreateExtValueOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaAbstractInterfaceDefCreateExtValueOutHelper = CorbaAbstractInterfaceDefCreateExtValueOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaAbstractInterfaceDefCreateExtValueOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaAbstractInterfaceDefCreateExtValueOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaAbstractInterfaceDefCreateExtValueOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaAbstractInterfaceDefCreateExtValueOut{}
			},
			__reflect__.TypeOf((*CorbaAbstractInterfaceDefCreateExtValueOut)(nil))))
}
