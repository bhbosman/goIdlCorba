// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_interface_tc_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateInterfaceTcIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateInterfaceTcInId_Const = "IDL:CORBA/ORB_create_interface_tc_In:1.0"

func (self *CorbaOrbCreateInterfaceTcIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateInterfaceTcIn) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateInterfaceTcIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Id, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Name, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateInterfaceTcIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateInterfaceTcIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Id)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Name)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateInterfaceTcIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateInterfaceTcInHelper = CorbaOrbCreateInterfaceTcIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateInterfaceTcInId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateInterfaceTcIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateInterfaceTcIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateInterfaceTcIn{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateInterfaceTcIn)(nil))))
}
