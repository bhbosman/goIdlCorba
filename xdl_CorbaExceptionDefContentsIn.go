// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExceptionDef_contents_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExceptionDefContentsIn struct {
	__goidl__.IdlObject
	LimitType uint32 `json:"LimitType"`
	ExcludeInherited bool `json:"ExcludeInherited"`
}

//noinspection GoSnakeCaseUsage
const CorbaExceptionDefContentsInId_Const = "IDL:CORBA/ExceptionDef_contents_In:1.0"

func (self *CorbaExceptionDefContentsIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExceptionDefContentsIn) GoString() string {
	return self.String()
}

func (self *CorbaExceptionDefContentsIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
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

func (self *CorbaExceptionDefContentsIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefContentsIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
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
type CorbaExceptionDefContentsIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExceptionDefContentsInHelper = CorbaExceptionDefContentsIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExceptionDefContentsInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExceptionDefContentsIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExceptionDefContentsIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExceptionDefContentsIn{}
			},
			__reflect__.TypeOf((*CorbaExceptionDefContentsIn)(nil))))
}