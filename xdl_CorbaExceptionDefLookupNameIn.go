// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExceptionDef_lookup_name_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExceptionDefLookupNameIn struct {
	__goidl__.IdlObject
	SearchName string `json:"SearchName"`
	LevelsToSearch int32 `json:"LevelsToSearch"`
	LimitType uint32 `json:"LimitType"`
	ExcludeInherited bool `json:"ExcludeInherited"`
}

//noinspection GoSnakeCaseUsage
const CorbaExceptionDefLookupNameInId_Const = "IDL:CORBA/ExceptionDef_lookup_name_In:1.0"

func (self *CorbaExceptionDefLookupNameIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExceptionDefLookupNameIn) GoString() string {
	return self.String()
}

func (self *CorbaExceptionDefLookupNameIn) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaExceptionDefLookupNameIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefLookupNameIn) Write(stream __goidl__.IWriteAny) error {
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
type CorbaExceptionDefLookupNameIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExceptionDefLookupNameInHelper = CorbaExceptionDefLookupNameIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExceptionDefLookupNameInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExceptionDefLookupNameIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExceptionDefLookupNameIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExceptionDefLookupNameIn{}
			},
			__reflect__.TypeOf((*CorbaExceptionDefLookupNameIn)(nil))))
}
