// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDef_describe_ext_interface_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDefDescribeExtInterfaceIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDefDescribeExtInterfaceInId_Const = "IDL:CORBA/ComponentIR/ComponentDef_describe_ext_interface_In:1.0"

func (self *CorbaComponentIRComponentDefDescribeExtInterfaceIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDefDescribeExtInterfaceIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDefDescribeExtInterfaceIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefDescribeExtInterfaceIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefDescribeExtInterfaceIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDefDescribeExtInterfaceIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDefDescribeExtInterfaceInHelper = CorbaComponentIRComponentDefDescribeExtInterfaceIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDefDescribeExtInterfaceInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDefDescribeExtInterfaceIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefDescribeExtInterfaceIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefDescribeExtInterfaceIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDefDescribeExtInterfaceIn)(nil))))
}
