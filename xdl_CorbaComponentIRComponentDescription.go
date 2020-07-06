// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDescription", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDescription struct {
	__goidl__.IdlObject
	Name string `json:"Name"`
	Id string `json:"Id"`
	DefinedIn string `json:"DefinedIn"`
	Version string `json:"Version"`
	BaseComponent string `json:"BaseComponent"`
	SupportedInterfaces CorbaRepositoryIdSeq `json:"SupportedInterfaces"`
	ProvidedInterfaces CorbaComponentIRProvidesDescriptionSeq `json:"ProvidedInterfaces"`
	UsedInterfaces CorbaComponentIRUsesDescriptionSeq `json:"UsedInterfaces"`
	EmitsEvents CorbaComponentIREventPortDescriptionSeq `json:"EmitsEvents"`
	PublishesEvents CorbaComponentIREventPortDescriptionSeq `json:"PublishesEvents"`
	ConsumesEvents CorbaComponentIREventPortDescriptionSeq `json:"ConsumesEvents"`
	Attributes CorbaExtAttrDescriptionSeq `json:"Attributes"`
	Type CorbaTypeCode `json:"Type"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDescriptionId_Const = "IDL:omg.org/CORBA/ComponentIR/ComponentDescription:1.0"

func (self *CorbaComponentIRComponentDescription) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDescription) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDescription) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Name, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Id, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.DefinedIn, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Version, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.BaseComponent, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.SupportedInterfaces.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.ProvidedInterfaces.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.UsedInterfaces.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.EmitsEvents.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.PublishesEvents.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.ConsumesEvents.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.Attributes.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.Type, err = CorbaTypeCodeHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDescription) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDescription) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Name)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Id)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.DefinedIn)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Version)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.BaseComponent)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.SupportedInterfaces.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.ProvidedInterfaces.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.UsedInterfaces.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.EmitsEvents.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.PublishesEvents.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.ConsumesEvents.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.Attributes.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaTypeCodeHelper.Write(stream, self.Type)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDescription_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDescriptionHelper = CorbaComponentIRComponentDescription_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDescriptionId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDescription.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDescription{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDescription{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDescription)(nil))))
}
