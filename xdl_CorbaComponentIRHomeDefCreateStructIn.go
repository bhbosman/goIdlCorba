// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::HomeDef_create_struct_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRHomeDefCreateStructIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	Version string `json:"Version"`
	Members CorbaStructMemberSeq `json:"Members"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRHomeDefCreateStructInId_Const = "IDL:CORBA/ComponentIR/HomeDef_create_struct_In:1.0"

func (self *CorbaComponentIRHomeDefCreateStructIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRHomeDefCreateStructIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRHomeDefCreateStructIn) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaComponentIRHomeDefCreateStructIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRHomeDefCreateStructIn) Write(stream __goidl__.IWriteAny) error {
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
type CorbaComponentIRHomeDefCreateStructIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRHomeDefCreateStructInHelper = CorbaComponentIRHomeDefCreateStructIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRHomeDefCreateStructInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRHomeDefCreateStructIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefCreateStructIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRHomeDefCreateStructIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIRHomeDefCreateStructIn)(nil))))
}
