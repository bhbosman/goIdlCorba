// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::EventDef_create_struct_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIREventDefCreateStructIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	Version string `json:"Version"`
	Members CorbaStructMemberSeq `json:"Members"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIREventDefCreateStructInId_Const = "IDL:CORBA/ComponentIR/EventDef_create_struct_In:1.0"

func (self *CorbaComponentIREventDefCreateStructIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIREventDefCreateStructIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIREventDefCreateStructIn) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaComponentIREventDefCreateStructIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIREventDefCreateStructIn) Write(stream __goidl__.IWriteAny) error {
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
type CorbaComponentIREventDefCreateStructIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIREventDefCreateStructInHelper = CorbaComponentIREventDefCreateStructIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIREventDefCreateStructInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIREventDefCreateStructIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateStructIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIREventDefCreateStructIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIREventDefCreateStructIn)(nil))))
}