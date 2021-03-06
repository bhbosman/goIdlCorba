// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_lookup_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefLookupIn struct {
	__goidl__.IdlObject
	SearchName string `json:"SearchName"`
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefLookupInId_Const = "IDL:CORBA/InterfaceDef_lookup_In:1.0"

func (self *CorbaInterfaceDefLookupIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefLookupIn) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefLookupIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.SearchName, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefLookupIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefLookupIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.SearchName)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefLookupIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefLookupInHelper = CorbaInterfaceDefLookupIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefLookupInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefLookupIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefLookupIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefLookupIn{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefLookupIn)(nil))))
}
