// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::EventDef_create_ext_value_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIREventDefCreateExtValueOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIREventDefCreateExtValueOutId_Const = "IDL:CORBA/ComponentIR/EventDef_create_ext_value_Out:1.0"

func (self *CorbaComponentIREventDefCreateExtValueOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIREventDefCreateExtValueOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIREventDefCreateExtValueOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateExtValueOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateExtValueOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIREventDefCreateExtValueOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIREventDefCreateExtValueOutHelper = CorbaComponentIREventDefCreateExtValueOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIREventDefCreateExtValueOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIREventDefCreateExtValueOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateExtValueOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateExtValueOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIREventDefCreateExtValueOut)(nil))))
}
