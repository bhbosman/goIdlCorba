// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Request_delete_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRequestDeleteIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRequestDeleteInId_Const = "IDL:CORBA/Request_delete_In:1.0"

func (self *CorbaRequestDeleteIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRequestDeleteIn) GoString() string {
	return self.String()
}

func (self *CorbaRequestDeleteIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRequestDeleteIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRequestDeleteIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRequestDeleteIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRequestDeleteInHelper = CorbaRequestDeleteIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRequestDeleteInId_Const,
			__goidl__.StructType,
			"CORBA_Request.idl",
			"xdl_CorbaRequestDeleteIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaRequestDeleteIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRequestDeleteIn{}
			},
			__reflect__.TypeOf((*CorbaRequestDeleteIn)(nil))))
}
