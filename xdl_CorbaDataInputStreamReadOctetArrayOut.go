// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_octet_array_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadOctetArrayOut struct {
	__goidl__.IdlObject
	Seq CorbaOctetSeq `json:"Seq"`
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadOctetArrayOutId_Const = "IDL:CORBA/DataInputStream_read_octet_array_Out:1.0"

func (self *CorbaDataInputStreamReadOctetArrayOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadOctetArrayOut) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadOctetArrayOut) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaDataInputStreamReadOctetArrayOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadOctetArrayOut) Write(stream __goidl__.IWriteAny) error {
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
type CorbaDataInputStreamReadOctetArrayOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadOctetArrayOutHelper = CorbaDataInputStreamReadOctetArrayOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadOctetArrayOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadOctetArrayOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadOctetArrayOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadOctetArrayOut{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadOctetArrayOut)(nil))))
}
