// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExceptionDef_lookup_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExceptionDefLookupIn struct {
	__goidl__.IdlObject
	SearchName string `json:"SearchName"`
}

//noinspection GoSnakeCaseUsage
const CorbaExceptionDefLookupInId_Const = "IDL:CORBA/ExceptionDef_lookup_In:1.0"

func (self *CorbaExceptionDefLookupIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExceptionDefLookupIn) GoString() string {
	return self.String()
}

func (self *CorbaExceptionDefLookupIn) ReadValue(stream __goidl__.IReadAny) error {
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
	return nil
}

func (self *CorbaExceptionDefLookupIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefLookupIn) Write(stream __goidl__.IWriteAny) error {
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
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExceptionDefLookupIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExceptionDefLookupInHelper = CorbaExceptionDefLookupIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExceptionDefLookupInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExceptionDefLookupIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExceptionDefLookupIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExceptionDefLookupIn{}
			},
			__reflect__.TypeOf((*CorbaExceptionDefLookupIn)(nil))))
}
