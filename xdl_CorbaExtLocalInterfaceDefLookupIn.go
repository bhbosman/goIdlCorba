// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ExtLocalInterfaceDef_lookup_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaExtLocalInterfaceDefLookupIn struct {
	__goidl__.IdlObject
	SearchName string `json:"SearchName"`
}

//noinspection GoSnakeCaseUsage
const CorbaExtLocalInterfaceDefLookupInId_Const = "IDL:CORBA/ExtLocalInterfaceDef_lookup_In:1.0"

func (self *CorbaExtLocalInterfaceDefLookupIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaExtLocalInterfaceDefLookupIn) GoString() string {
	return self.String()
}

func (self *CorbaExtLocalInterfaceDefLookupIn) ReadValue(stream __goidl__.IReadAny) error {
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

func (self *CorbaExtLocalInterfaceDefLookupIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaExtLocalInterfaceDefLookupIn) Write(stream __goidl__.IWriteAny) error {
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
type CorbaExtLocalInterfaceDefLookupIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaExtLocalInterfaceDefLookupInHelper = CorbaExtLocalInterfaceDefLookupIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaExtLocalInterfaceDefLookupInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaExtLocalInterfaceDefLookupIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefLookupIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaExtLocalInterfaceDefLookupIn{}
			},
			__reflect__.TypeOf((*CorbaExtLocalInterfaceDefLookupIn)(nil))))
}