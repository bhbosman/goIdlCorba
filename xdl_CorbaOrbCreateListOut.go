// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_list_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateListOut struct {
	__goidl__.IdlObject
	NewList CorbaNVList `json:"NewList"`
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateListOutId_Const = "IDL:CORBA/ORB_create_list_Out:1.0"

func (self *CorbaOrbCreateListOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateListOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateListOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.NewList, err = CorbaNVListHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateListOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateListOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaNVListHelper.Write(stream, self.NewList)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateListOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateListOutHelper = CorbaOrbCreateListOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateListOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateListOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateListOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateListOut{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateListOut)(nil))))
}
