// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataOutputStream_write_octet_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataOutputStreamWriteOctetIn struct {
	__goidl__.IdlObject
	Value byte `json:"Value"`
}

//noinspection GoSnakeCaseUsage
const CorbaDataOutputStreamWriteOctetInId_Const = "IDL:CORBA/DataOutputStream_write_octet_In:1.0"

func (self *CorbaDataOutputStreamWriteOctetIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataOutputStreamWriteOctetIn) GoString() string {
	return self.String()
}

func (self *CorbaDataOutputStreamWriteOctetIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlOctetKind)
	self.Value, err = __goidl__.IdlOctetHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteOctetIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteOctetIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlOctetKind)
	err = __goidl__.IdlOctetHelper.Write(stream, self.Value)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataOutputStreamWriteOctetIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataOutputStreamWriteOctetInHelper = CorbaDataOutputStreamWriteOctetIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataOutputStreamWriteOctetInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataOutputStreamWriteOctetIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteOctetIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteOctetIn{}
			},
			__reflect__.TypeOf((*CorbaDataOutputStreamWriteOctetIn)(nil))))
}
