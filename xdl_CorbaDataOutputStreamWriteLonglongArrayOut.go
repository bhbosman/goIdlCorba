// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataOutputStream_write_longlong_array_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataOutputStreamWriteLonglongArrayOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaDataOutputStreamWriteLonglongArrayOutId_Const = "IDL:CORBA/DataOutputStream_write_longlong_array_Out:1.0"

func (self *CorbaDataOutputStreamWriteLonglongArrayOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataOutputStreamWriteLonglongArrayOut) GoString() string {
	return self.String()
}

func (self *CorbaDataOutputStreamWriteLonglongArrayOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteLonglongArrayOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteLonglongArrayOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataOutputStreamWriteLonglongArrayOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataOutputStreamWriteLonglongArrayOutHelper = CorbaDataOutputStreamWriteLonglongArrayOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataOutputStreamWriteLonglongArrayOutId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataOutputStreamWriteLonglongArrayOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteLonglongArrayOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteLonglongArrayOut{}
			},
			__reflect__.TypeOf((*CorbaDataOutputStreamWriteLonglongArrayOut)(nil))))
}
