// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_long_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadLongIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadLongInId_Const = "IDL:CORBA/DataInputStream_read_long_In:1.0"

func (self *CorbaDataInputStreamReadLongIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadLongIn) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadLongIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadLongIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadLongIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadLongIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadLongInHelper = CorbaDataInputStreamReadLongIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadLongInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadLongIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadLongIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadLongIn{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadLongIn)(nil))))
}
