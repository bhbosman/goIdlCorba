// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::DataOutputStream_write_longdouble_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDataOutputStreamWriteLongdoubleIn struct {
	__goidl__.IdlObject
	Value float64 `json:"Value"`
}

//noinspection GoSnakeCaseUsage
const CorbaDataOutputStreamWriteLongdoubleInId_Const = "IDL:CORBA/DataOutputStream_write_longdouble_In:1.0"

func (self *CorbaDataOutputStreamWriteLongdoubleIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDataOutputStreamWriteLongdoubleIn) GoString() string {
	return self.String()
}

func (self *CorbaDataOutputStreamWriteLongdoubleIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(LongDoubleType)
	self.Value, err = __goidl__.IdlFloat64Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteLongdoubleIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDataOutputStreamWriteLongdoubleIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(LongDoubleType)
	err = __goidl__.IdlFloat64Helper.Write(stream, self.Value)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDataOutputStreamWriteLongdoubleIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDataOutputStreamWriteLongdoubleInHelper = CorbaDataOutputStreamWriteLongdoubleIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDataOutputStreamWriteLongdoubleInId_Const,
			__goidl__.StructType,
			"CORBA_Stream.idl",
			"xdl_CorbaDataOutputStreamWriteLongdoubleIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteLongdoubleIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDataOutputStreamWriteLongdoubleIn{}
			},
			__reflect__.TypeOf((*CorbaDataOutputStreamWriteLongdoubleIn)(nil))))
}