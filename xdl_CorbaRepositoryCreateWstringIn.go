// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Repository_create_wstring_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRepositoryCreateWstringIn struct {
	__goidl__.IdlObject
	Bound uint32 `json:"Bound"`
}

//noinspection GoSnakeCaseUsage
const CorbaRepositoryCreateWstringInId_Const = "IDL:CORBA/Repository_create_wstring_In:1.0"

func (self *CorbaRepositoryCreateWstringIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRepositoryCreateWstringIn) GoString() string {
	return self.String()
}

func (self *CorbaRepositoryCreateWstringIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.Bound, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateWstringIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRepositoryCreateWstringIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Bound)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRepositoryCreateWstringIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRepositoryCreateWstringInHelper = CorbaRepositoryCreateWstringIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRepositoryCreateWstringInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaRepositoryCreateWstringIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaRepositoryCreateWstringIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRepositoryCreateWstringIn{}
			},
			__reflect__.TypeOf((*CorbaRepositoryCreateWstringIn)(nil))))
}
