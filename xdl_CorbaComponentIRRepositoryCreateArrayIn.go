// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Repository_create_array_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRRepositoryCreateArrayIn struct {
	__goidl__.IdlObject
	Length uint32 `json:"Length"`
	ElementType CorbaIDLType `json:"ElementType"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRRepositoryCreateArrayInId_Const = "IDL:CORBA/ComponentIR/Repository_create_array_In:1.0"

func (self *CorbaComponentIRRepositoryCreateArrayIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRRepositoryCreateArrayIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRRepositoryCreateArrayIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.Length, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.ElementType, err = CorbaIDLTypeHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateArrayIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryCreateArrayIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Length)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaIDLTypeHelper.Write(stream, self.ElementType)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRRepositoryCreateArrayIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRRepositoryCreateArrayInHelper = CorbaComponentIRRepositoryCreateArrayIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRRepositoryCreateArrayInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRRepositoryCreateArrayIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateArrayIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryCreateArrayIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIRRepositoryCreateArrayIn)(nil))))
}
