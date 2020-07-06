// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtAbstractInterfaceDef_describe_contents_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtAbstractInterfaceDefDescribeContentsIn struct {
	__goidl__.IdlObject
	LimitType uint32 `json:"LimitType"`
	ExcludeInherited bool `json:"ExcludeInherited"`
	MaxReturnedObjs int32 `json:"MaxReturnedObjs"`
}

//noinspection GoSnakeCaseUsage
const CorbaExtAbstractInterfaceDefDescribeContentsInId_Const = "IDL:CORBA/ExtAbstractInterfaceDef_describe_contents_In:1.0"

func (self *CorbaExtAbstractInterfaceDefDescribeContentsIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtAbstractInterfaceDefDescribeContentsIn) GoString() string {
	return self.String()
}

func (self *CorbaExtAbstractInterfaceDefDescribeContentsIn) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaExtAbstractInterfaceDefDescribeContentsIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtAbstractInterfaceDefDescribeContentsIn) Write(stream __goidl__.IWriteAny) error {
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
type CorbaExtAbstractInterfaceDefDescribeContentsIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtAbstractInterfaceDefDescribeContentsInHelper = CorbaExtAbstractInterfaceDefDescribeContentsIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtAbstractInterfaceDefDescribeContentsInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtAbstractInterfaceDefDescribeContentsIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefDescribeContentsIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtAbstractInterfaceDefDescribeContentsIn{}
			},
			__reflect__.TypeOf((*CorbaExtAbstractInterfaceDefDescribeContentsIn)(nil))))
}
