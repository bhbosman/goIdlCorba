// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::CustomMarshal_marshal_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaCustomMarshalMarshalIn struct {
	__goidl__.IdlObject
	Os CorbaDataOutputStream `json:"Os"`
}

//noinspection GoSnakeCaseUsage
const CorbaCustomMarshalMarshalInId_Const = "IDL:CORBA/CustomMarshal_marshal_In:1.0"

func (self *CorbaCustomMarshalMarshalIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaCustomMarshalMarshalIn) GoString() string {
	return self.String()
}

func (self *CorbaCustomMarshalMarshalIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.Os, err = CorbaDataOutputStreamHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaCustomMarshalMarshalIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaCustomMarshalMarshalIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaDataOutputStreamHelper.Write(stream, self.Os)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaCustomMarshalMarshalIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaCustomMarshalMarshalInHelper = CorbaCustomMarshalMarshalIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaCustomMarshalMarshalInId_Const,
			__goidl__.StructType,
			"CORBA_CustomMarshal.idl",
			"xdl_CorbaCustomMarshalMarshalIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaCustomMarshalMarshalIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaCustomMarshalMarshalIn{}
			},
			__reflect__.TypeOf((*CorbaCustomMarshalMarshalIn)(nil))))
}
