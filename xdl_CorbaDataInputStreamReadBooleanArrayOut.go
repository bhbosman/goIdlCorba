// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_boolean_array_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadBooleanArrayOut struct {
	__goidl__.IdlObject
	Seq CorbaBooleanSeq `json:"Seq"`
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadBooleanArrayOutId_Const = "IDL:CORBA/DataInputStream_read_boolean_array_Out:1.0"

func (self *CorbaDataInputStreamReadBooleanArrayOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadBooleanArrayOut) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadBooleanArrayOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.Seq.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadBooleanArrayOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadBooleanArrayOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.Seq.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadBooleanArrayOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadBooleanArrayOutHelper = CorbaDataInputStreamReadBooleanArrayOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadBooleanArrayOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadBooleanArrayOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadBooleanArrayOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadBooleanArrayOut{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadBooleanArrayOut)(nil))))
}
