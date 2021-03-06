// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Request_get_response_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRequestGetResponseIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaRequestGetResponseInId_Const = "IDL:CORBA/Request_get_response_In:1.0"

func (self *CorbaRequestGetResponseIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRequestGetResponseIn) GoString() string {
	return self.String()
}

func (self *CorbaRequestGetResponseIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRequestGetResponseIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRequestGetResponseIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRequestGetResponseIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRequestGetResponseInHelper = CorbaRequestGetResponseIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRequestGetResponseInId_Const,
			__goidl__.StructType,
			"CORBA_Request.idl",
			"xdl_CorbaRequestGetResponseIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaRequestGetResponseIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRequestGetResponseIn{}
			},
			__reflect__.TypeOf((*CorbaRequestGetResponseIn)(nil))))
}
