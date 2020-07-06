// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtInterfaceDef_create_ext_value_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtInterfaceDefCreateExtValueIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	Version string `json:"Version"`
	IsCustom bool `json:"IsCustom"`
	IsAbstract bool `json:"IsAbstract"`
	BaseValue CorbaValueDef `json:"BaseValue"`
	IsTruncatable bool `json:"IsTruncatable"`
	AbstractBaseValues CorbaValueDefSeq `json:"AbstractBaseValues"`
	SupportedInterfaces CorbaInterfaceDefSeq `json:"SupportedInterfaces"`
	Initializers CorbaExtInitializerSeq `json:"Initializers"`
}

//noinspection GoSnakeCaseUsage
const CorbaExtInterfaceDefCreateExtValueInId_Const = "IDL:CORBA/ExtInterfaceDef_create_ext_value_In:1.0"

func (self *CorbaExtInterfaceDefCreateExtValueIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtInterfaceDefCreateExtValueIn) GoString() string {
	return self.String()
}

func (self *CorbaExtInterfaceDefCreateExtValueIn) ReadValue(stream __goidl__.IReadAny) error {
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
	// WriteStructHelper::WriteStructMemberExtractValue(BooleanType)
	self.IsCustom, err = __goidl__.IdlBooleanHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(BooleanType)
	self.IsAbstract, err = __goidl__.IdlBooleanHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.BaseValue, err = CorbaValueDefHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(BooleanType)
	self.IsTruncatable, err = __goidl__.IdlBooleanHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.AbstractBaseValues.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.SupportedInterfaces.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.Initializers.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtInterfaceDefCreateExtValueIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtInterfaceDefCreateExtValueIn) Write(stream __goidl__.IWriteAny) error {
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
	// WriteStructHelper::WriteStructMemberInsert(BooleanType)
	err = __goidl__.IdlBooleanHelper.Write(stream, self.IsCustom)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(BooleanType)
	err = __goidl__.IdlBooleanHelper.Write(stream, self.IsAbstract)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaValueDefHelper.Write(stream, self.BaseValue)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(BooleanType)
	err = __goidl__.IdlBooleanHelper.Write(stream, self.IsTruncatable)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.AbstractBaseValues.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.SupportedInterfaces.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.Initializers.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExtInterfaceDefCreateExtValueIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtInterfaceDefCreateExtValueInHelper = CorbaExtInterfaceDefCreateExtValueIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtInterfaceDefCreateExtValueInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtInterfaceDefCreateExtValueIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtInterfaceDefCreateExtValueIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtInterfaceDefCreateExtValueIn{}
			},
			__reflect__.TypeOf((*CorbaExtInterfaceDefCreateExtValueIn)(nil))))
}