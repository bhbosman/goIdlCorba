// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::LocalInterfaceDef_create_enum_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaLocalInterfaceDefCreateEnumIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	Version string `json:"Version"`
	Members CorbaEnumMemberSeq `json:"Members"`
}

//noinspection GoSnakeCaseUsage
const CorbaLocalInterfaceDefCreateEnumInId_Const = "IDL:CORBA/LocalInterfaceDef_create_enum_In:1.0"

func (self *CorbaLocalInterfaceDefCreateEnumIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaLocalInterfaceDefCreateEnumIn) GoString() string {
	return self.String()
}

func (self *CorbaLocalInterfaceDefCreateEnumIn) ReadValue(stream __goidl__.IReadAny) error {
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
	err = self.Members.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefCreateEnumIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefCreateEnumIn) Write(stream __goidl__.IWriteAny) error {
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
	err = self.Members.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaLocalInterfaceDefCreateEnumIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaLocalInterfaceDefCreateEnumInHelper = CorbaLocalInterfaceDefCreateEnumIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaLocalInterfaceDefCreateEnumInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaLocalInterfaceDefCreateEnumIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefCreateEnumIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefCreateEnumIn{}
			},
			__reflect__.TypeOf((*CorbaLocalInterfaceDefCreateEnumIn)(nil))))
}
