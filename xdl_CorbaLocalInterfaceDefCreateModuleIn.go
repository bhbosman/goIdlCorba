// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::LocalInterfaceDef_create_module_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaLocalInterfaceDefCreateModuleIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	Version string `json:"Version"`
}

//noinspection GoSnakeCaseUsage
const CorbaLocalInterfaceDefCreateModuleInId_Const = "IDL:CORBA/LocalInterfaceDef_create_module_In:1.0"

func (self *CorbaLocalInterfaceDefCreateModuleIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaLocalInterfaceDefCreateModuleIn) GoString() string {
	return self.String()
}

func (self *CorbaLocalInterfaceDefCreateModuleIn) ReadValue(stream __goidl__.IReadAny) error {
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
	return nil
}

func (self *CorbaLocalInterfaceDefCreateModuleIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaLocalInterfaceDefCreateModuleIn) Write(stream __goidl__.IWriteAny) error {
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
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaLocalInterfaceDefCreateModuleIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaLocalInterfaceDefCreateModuleInHelper = CorbaLocalInterfaceDefCreateModuleIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaLocalInterfaceDefCreateModuleInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaLocalInterfaceDefCreateModuleIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefCreateModuleIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaLocalInterfaceDefCreateModuleIn{}
			},
			__reflect__.TypeOf((*CorbaLocalInterfaceDefCreateModuleIn)(nil))))
}
