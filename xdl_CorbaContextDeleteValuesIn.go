// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Context_delete_values_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContextDeleteValuesIn struct {
	__goidl__.IdlObject
	PropName string `json:"PropName"`
}

//noinspection GoSnakeCaseUsage
const CorbaContextDeleteValuesInId_Const = "IDL:CORBA/Context_delete_values_In:1.0"

func (self *CorbaContextDeleteValuesIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContextDeleteValuesIn) GoString() string {
	return self.String()
}

func (self *CorbaContextDeleteValuesIn) ReadValue(stream __goidl__.IReadAny) error {
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
	return nil
}

func (self *CorbaContextDeleteValuesIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContextDeleteValuesIn) Write(stream __goidl__.IWriteAny) error {
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
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContextDeleteValuesIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContextDeleteValuesInHelper = CorbaContextDeleteValuesIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContextDeleteValuesInId_Const,
			__goidl__.StructType,
			"CORBA_Context.idl",
			"xdl_CorbaContextDeleteValuesIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaContextDeleteValuesIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContextDeleteValuesIn{}
			},
			__reflect__.TypeOf((*CorbaContextDeleteValuesIn)(nil))))
}
