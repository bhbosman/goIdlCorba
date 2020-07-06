// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_long_double_array_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadLongDoubleArrayOut struct {
	__goidl__.IdlObject
	Seq CorbaDoubleSeq `json:"Seq"`
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadLongDoubleArrayOutId_Const = "IDL:CORBA/DataInputStream_read_long_double_array_Out:1.0"

func (self *CorbaDataInputStreamReadLongDoubleArrayOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadLongDoubleArrayOut) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadLongDoubleArrayOut) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaDataInputStreamReadLongDoubleArrayOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadLongDoubleArrayOut) Write(stream __goidl__.IWriteAny) error {
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
type CorbaDataInputStreamReadLongDoubleArrayOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadLongDoubleArrayOutHelper = CorbaDataInputStreamReadLongDoubleArrayOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadLongDoubleArrayOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadLongDoubleArrayOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadLongDoubleArrayOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadLongDoubleArrayOut{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadLongDoubleArrayOut)(nil))))
}
