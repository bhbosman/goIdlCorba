// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_operation_list_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateOperationListOut struct {
	__goidl__.IdlObject
	NewList CorbaNVList `json:"NewList"`
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateOperationListOutId_Const = "IDL:CORBA/ORB_create_operation_list_Out:1.0"

func (self *CorbaOrbCreateOperationListOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateOperationListOut) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateOperationListOut) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaOrbCreateOperationListOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateOperationListOut) Write(stream __goidl__.IWriteAny) error {
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
type CorbaOrbCreateOperationListOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateOperationListOutHelper = CorbaOrbCreateOperationListOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateOperationListOutId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateOperationListOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateOperationListOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateOperationListOut{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateOperationListOut)(nil))))
}
