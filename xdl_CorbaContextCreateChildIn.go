// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Context_create_child_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContextCreateChildIn struct {
	__goidl__.IdlObject
	CtxName string `json:"CtxName"`
}

//noinspection GoSnakeCaseUsage
const CorbaContextCreateChildInId_Const = "IDL:CORBA/Context_create_child_In:1.0"

func (self *CorbaContextCreateChildIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContextCreateChildIn) GoString() string {
	return self.String()
}

func (self *CorbaContextCreateChildIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.CtxName, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContextCreateChildIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContextCreateChildIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.CtxName)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContextCreateChildIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContextCreateChildInHelper = CorbaContextCreateChildIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContextCreateChildInId_Const,
			__goidl__.StructType,
			"CORBA_Context.idl",
			"xdl_CorbaContextCreateChildIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaContextCreateChildIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContextCreateChildIn{}
			},
			__reflect__.TypeOf((*CorbaContextCreateChildIn)(nil))))
}
