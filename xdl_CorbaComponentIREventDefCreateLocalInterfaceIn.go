// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::EventDef_create_local_interface_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIREventDefCreateLocalInterfaceIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	Version string `json:"Version"`
	BaseInterfaces CorbaInterfaceDefSeq `json:"BaseInterfaces"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIREventDefCreateLocalInterfaceInId_Const = "IDL:CORBA/ComponentIR/EventDef_create_local_interface_In:1.0"

func (self *CorbaComponentIREventDefCreateLocalInterfaceIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIREventDefCreateLocalInterfaceIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIREventDefCreateLocalInterfaceIn) ReadValue(stream __goidl__.IReadAny) error {
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
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Version, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.BaseInterfaces.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateLocalInterfaceIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateLocalInterfaceIn) Write(stream __goidl__.IWriteAny) error {
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
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Version)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.BaseInterfaces.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIREventDefCreateLocalInterfaceIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIREventDefCreateLocalInterfaceInHelper = CorbaComponentIREventDefCreateLocalInterfaceIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIREventDefCreateLocalInterfaceInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIREventDefCreateLocalInterfaceIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateLocalInterfaceIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateLocalInterfaceIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIREventDefCreateLocalInterfaceIn)(nil))))
}
