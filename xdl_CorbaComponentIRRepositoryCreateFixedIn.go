// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Repository_create_fixed_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRRepositoryCreateFixedIn struct {
	__goidl__.IdlObject
	Digits uint16 `json:"Digits"`
	Scale int16 `json:"Scale"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRRepositoryCreateFixedInId_Const = "IDL:CORBA/ComponentIR/Repository_create_fixed_In:1.0"

func (self *CorbaComponentIRRepositoryCreateFixedIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRRepositoryCreateFixedIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRRepositoryCreateFixedIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedShortType)
	self.Digits, err = __goidl__.IdlUInt16Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(ShortType)
	self.Scale, err = __goidl__.IdlInt16Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateFixedIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateFixedIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedShortType)
	err = __goidl__.IdlUInt16Helper.Write(stream, self.Digits)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(ShortType)
	err = __goidl__.IdlInt16Helper.Write(stream, self.Scale)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRRepositoryCreateFixedIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRRepositoryCreateFixedInHelper = CorbaComponentIRRepositoryCreateFixedIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRRepositoryCreateFixedInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRRepositoryCreateFixedIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateFixedIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateFixedIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIRRepositoryCreateFixedIn)(nil))))
}
