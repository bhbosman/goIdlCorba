// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Request_prepare_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRequestPrepareIn struct {
	__goidl__.IdlObject
	P __goidl__.IIdlObject `json:"P"`
}

//noinspection GoSnakeCaseUsage
const CorbaRequestPrepareInId_Const = "IDL:CORBA/Request_prepare_In:1.0"

func (self *CorbaRequestPrepareIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRequestPrepareIn) GoString() string {
	return self.String()
}

func (self *CorbaRequestPrepareIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlObjectKind)
	self.P, err = __goidl__.IdlObjectHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRequestPrepareIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRequestPrepareIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlObjectKind)
	err = __goidl__.IdlObjectHelper.Write(stream, self.P)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRequestPrepareIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRequestPrepareInHelper = CorbaRequestPrepareIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRequestPrepareInId_Const,
			__goidl__.StructType,
			"CORBA_Request.idl",
			"xdl_CorbaRequestPrepareIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaRequestPrepareIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRequestPrepareIn{}
			},
			__reflect__.TypeOf((*CorbaRequestPrepareIn)(nil))))
}
