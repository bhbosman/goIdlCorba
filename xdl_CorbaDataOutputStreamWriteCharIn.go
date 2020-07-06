// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataOutputStream_write_char_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataOutputStreamWriteCharIn struct {
	__goidl__.IdlObject
	Value byte `json:"Value"`
}

//noinspection GoSnakeCaseUsage
const CorbaDataOutputStreamWriteCharInId_Const = "IDL:CORBA/DataOutputStream_write_char_In:1.0"

func (self *CorbaDataOutputStreamWriteCharIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataOutputStreamWriteCharIn) GoString() string {
	return self.String()
}

func (self *CorbaDataOutputStreamWriteCharIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(CharType)
	self.Value, err = __goidl__.IdlOctetHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteCharIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteCharIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(CharType)
	err = __goidl__.IdlOctetHelper.Write(stream, self.Value)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataOutputStreamWriteCharIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataOutputStreamWriteCharInHelper = CorbaDataOutputStreamWriteCharIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataOutputStreamWriteCharInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataOutputStreamWriteCharIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteCharIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteCharIn{}
			},
			__reflect__.TypeOf((*CorbaDataOutputStreamWriteCharIn)(nil))))
}
