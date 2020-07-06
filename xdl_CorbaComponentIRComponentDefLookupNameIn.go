// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::ComponentDef_lookup_name_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRComponentDefLookupNameIn struct {
	__goidl__.IdlObject
	SearchName string `json:"SearchName"`
	LevelsToSearch int32 `json:"LevelsToSearch"`
	LimitType uint32 `json:"LimitType"`
	ExcludeInherited bool `json:"ExcludeInherited"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRComponentDefLookupNameInId_Const = "IDL:CORBA/ComponentIR/ComponentDef_lookup_name_In:1.0"

func (self *CorbaComponentIRComponentDefLookupNameIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRComponentDefLookupNameIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRComponentDefLookupNameIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.SearchName, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(LongType)
	self.LevelsToSearch, err = __goidl__.IdlInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlEnum)
	self.LimitType, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(BooleanType)
	self.ExcludeInherited, err = __goidl__.IdlBooleanHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefLookupNameIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRComponentDefLookupNameIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.SearchName)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(LongType)
	err = __goidl__.IdlInt32Helper.Write(stream, self.LevelsToSearch)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlEnum)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.LimitType)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(BooleanType)
	err = __goidl__.IdlBooleanHelper.Write(stream, self.ExcludeInherited)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRComponentDefLookupNameIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRComponentDefLookupNameInHelper = CorbaComponentIRComponentDefLookupNameIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRComponentDefLookupNameInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRComponentDefLookupNameIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefLookupNameIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRComponentDefLookupNameIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIRComponentDefLookupNameIn)(nil))))
}
