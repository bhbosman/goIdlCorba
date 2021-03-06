// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataOutputStream_write_Object_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataOutputStreamWriteObjectIn struct {
	__goidl__.IdlObject
	Value __goidl__.IIdlObject `json:"Value"`
}

//noinspection GoSnakeCaseUsage
const CorbaDataOutputStreamWriteObjectInId_Const = "IDL:CORBA/DataOutputStream_write_Object_In:1.0"

func (self *CorbaDataOutputStreamWriteObjectIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataOutputStreamWriteObjectIn) GoString() string {
	return self.String()
}

func (self *CorbaDataOutputStreamWriteObjectIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlObjectKind)
	self.Value, err = __goidl__.IdlObjectHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteObjectIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteObjectIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlObjectKind)
	err = __goidl__.IdlObjectHelper.Write(stream, self.Value)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataOutputStreamWriteObjectIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataOutputStreamWriteObjectInHelper = CorbaDataOutputStreamWriteObjectIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataOutputStreamWriteObjectInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataOutputStreamWriteObjectIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteObjectIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteObjectIn{}
			},
			__reflect__.TypeOf((*CorbaDataOutputStreamWriteObjectIn)(nil))))
}
