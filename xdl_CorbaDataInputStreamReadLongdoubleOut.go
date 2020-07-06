// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataInputStream_read_longdouble_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataInputStreamReadLongdoubleOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataInputStreamReadLongdoubleOutId_Const = "IDL:CORBA/DataInputStream_read_longdouble_Out:1.0"

func (self *CorbaDataInputStreamReadLongdoubleOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataInputStreamReadLongdoubleOut) GoString() string {
	return self.String()
}

func (self *CorbaDataInputStreamReadLongdoubleOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadLongdoubleOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataInputStreamReadLongdoubleOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataInputStreamReadLongdoubleOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataInputStreamReadLongdoubleOutHelper = CorbaDataInputStreamReadLongdoubleOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataInputStreamReadLongdoubleOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataInputStreamReadLongdoubleOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadLongdoubleOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataInputStreamReadLongdoubleOut{}
			},
			__reflect__.TypeOf((*CorbaDataInputStreamReadLongdoubleOut)(nil))))
}
