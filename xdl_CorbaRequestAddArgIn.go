// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Request_add_arg_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaRequestAddArgIn struct {
	__goidl__.IdlObject
	Name string `json:"Name"`
	ArgType CorbaTypeCode `json:"ArgType"`
	Value CorbaOpaqueValue `json:"Value"`
	Len int32 `json:"Len"`
	ArgFlags uint32 `json:"ArgFlags"`
}

//noinspection GoSnakeCaseUsage
const CorbaRequestAddArgInId_Const = "IDL:CORBA/Request_add_arg_In:1.0"

func (self *CorbaRequestAddArgIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRequestAddArgIn) GoString() string {
	return self.String()
}

func (self *CorbaRequestAddArgIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Name, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.ArgType, err = CorbaTypeCodeHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlNative)
	// WriteStructHelper::WriteStructMemberExtractValue(LongType)
	self.Len, err = __goidl__.IdlInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.ArgFlags, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRequestAddArgIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRequestAddArgIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Name)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaTypeCodeHelper.Write(stream, self.ArgType)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlNative)
	// WriteStructHelper::WriteStructMemberInsert(LongType)
	err = __goidl__.IdlInt32Helper.Write(stream, self.Len)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.ArgFlags)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRequestAddArgIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRequestAddArgInHelper = CorbaRequestAddArgIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRequestAddArgInId_Const,
			__goidl__.StructType,
			"CORBA_Request.idl",
			"xdl_CorbaRequestAddArgIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaRequestAddArgIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRequestAddArgIn{}
			},
			__reflect__.TypeOf((*CorbaRequestAddArgIn)(nil))))
}
