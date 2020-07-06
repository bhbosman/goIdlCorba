// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_octet_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadOctetIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadOctetInId_Const = "IDL:CORBA/DataInputStream_read_octet_In:1.0"

func (self *CorbaDataInputStreamReadOctetIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadOctetIn) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadOctetIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadOctetIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadOctetIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadOctetIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadOctetInHelper = CorbaDataInputStreamReadOctetIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadOctetInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadOctetIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadOctetIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadOctetIn{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadOctetIn)(nil))))
}
