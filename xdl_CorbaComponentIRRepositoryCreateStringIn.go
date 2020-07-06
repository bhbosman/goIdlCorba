// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Repository_create_string_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRRepositoryCreateStringIn struct {
	__goidl__.IdlObject
	Bound uint32 `json:"Bound"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRRepositoryCreateStringInId_Const = "IDL:CORBA/ComponentIR/Repository_create_string_In:1.0"

func (self *CorbaComponentIRRepositoryCreateStringIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRRepositoryCreateStringIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRRepositoryCreateStringIn) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaComponentIRRepositoryCreateStringIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateStringIn) Write(stream __goidl__.IWriteAny) error {
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
type CorbaComponentIRRepositoryCreateStringIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRRepositoryCreateStringInHelper = CorbaComponentIRRepositoryCreateStringIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRRepositoryCreateStringInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRRepositoryCreateStringIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateStringIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateStringIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIRRepositoryCreateStringIn)(nil))))
}
