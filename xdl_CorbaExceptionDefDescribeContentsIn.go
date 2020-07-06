// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExceptionDef_describe_contents_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExceptionDefDescribeContentsIn struct {
	__goidl__.IdlObject
	LimitType uint32 `json:"LimitType"`
	ExcludeInherited bool `json:"ExcludeInherited"`
	MaxReturnedObjs int32 `json:"MaxReturnedObjs"`
}

//noinspection GoSnakeCaseUsage
const CorbaExceptionDefDescribeContentsInId_Const = "IDL:CORBA/ExceptionDef_describe_contents_In:1.0"

func (self *CorbaExceptionDefDescribeContentsIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExceptionDefDescribeContentsIn) GoString() string {
	return self.String()
}

func (self *CorbaExceptionDefDescribeContentsIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlEnum)
	self.LimitType, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(BooleanType)
	self.ExcludeInherited, err = __goidl__.IdlBooleanHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(LongType)
	self.MaxReturnedObjs, err = __goidl__.IdlInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefDescribeContentsIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExceptionDefDescribeContentsIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlEnum)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.LimitType)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(BooleanType)
	err = __goidl__.IdlBooleanHelper.Write(stream, self.ExcludeInherited)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(LongType)
	err = __goidl__.IdlInt32Helper.Write(stream, self.MaxReturnedObjs)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaExceptionDefDescribeContentsIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExceptionDefDescribeContentsInHelper = CorbaExceptionDefDescribeContentsIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExceptionDefDescribeContentsInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExceptionDefDescribeContentsIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExceptionDefDescribeContentsIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExceptionDefDescribeContentsIn{}
			},
			__reflect__.TypeOf((*CorbaExceptionDefDescribeContentsIn)(nil))))
}
