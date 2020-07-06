// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::Context_delete_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaContextDeleteIn struct {
	__goidl__.IdlObject
	DelFlags uint32 `json:"DelFlags"`
}

//noinspection GoSnakeCaseUsage
const CorbaContextDeleteInId_Const = "IDL:CORBA/Context_delete_In:1.0"

func (self *CorbaContextDeleteIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaContextDeleteIn) GoString() string {
	return self.String()
}

func (self *CorbaContextDeleteIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.DelFlags, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContextDeleteIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaContextDeleteIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.DelFlags)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaContextDeleteIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaContextDeleteInHelper = CorbaContextDeleteIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaContextDeleteInId_Const,
			__goidl__.StructType,
			"CORBA_Context.idl",
			"xdl_CorbaContextDeleteIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaContextDeleteIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaContextDeleteIn{}
			},
			__reflect__.TypeOf((*CorbaContextDeleteIn)(nil))))
}