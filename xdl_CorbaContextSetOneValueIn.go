// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Context_set_one_value_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContextSetOneValueIn struct {
	__goidl__.IdlObject
	PropName string `json:"PropName"`
	Value string `json:"Value"`
}

//noinspection GoSnakeCaseUsage
const CorbaContextSetOneValueInId_Const = "IDL:CORBA/Context_set_one_value_In:1.0"

func (self *CorbaContextSetOneValueIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContextSetOneValueIn) GoString() string {
	return self.String()
}

func (self *CorbaContextSetOneValueIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.PropName, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Value, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContextSetOneValueIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContextSetOneValueIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.PropName)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Value)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContextSetOneValueIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContextSetOneValueInHelper = CorbaContextSetOneValueIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContextSetOneValueInId_Const,
			__goidl__.StructType,
			"CORBA_Context.idl",
			"xdl_CorbaContextSetOneValueIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaContextSetOneValueIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContextSetOneValueIn{}
			},
			__reflect__.TypeOf((*CorbaContextSetOneValueIn)(nil))))
}
