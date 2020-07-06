// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ValueDef_create_constant_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaValueDefCreateConstantIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	Version string `json:"Version"`
	Type CorbaIDLType `json:"Type"`
	Value __goidl__.IdlAny `json:"Value"`
}

//noinspection GoSnakeCaseUsage
const CorbaValueDefCreateConstantInId_Const = "IDL:CORBA/ValueDef_create_constant_In:1.0"

func (self *CorbaValueDefCreateConstantIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaValueDefCreateConstantIn) GoString() string {
	return self.String()
}

func (self *CorbaValueDefCreateConstantIn) ReadValue(stream __goidl__.IReadAny) error {
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
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.Type, err = CorbaIDLTypeHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(AnyType)
	self.Value, err = __goidl__.IdlAnyHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefCreateConstantIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaValueDefCreateConstantIn) Write(stream __goidl__.IWriteAny) error {
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
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaIDLTypeHelper.Write(stream, self.Type)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(AnyType)
	err = __goidl__.IdlAnyHelper.Write(stream, self.Value)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaValueDefCreateConstantIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaValueDefCreateConstantInHelper = CorbaValueDefCreateConstantIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaValueDefCreateConstantInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaValueDefCreateConstantIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaValueDefCreateConstantIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaValueDefCreateConstantIn{}
			},
			__reflect__.TypeOf((*CorbaValueDefCreateConstantIn)(nil))))
}