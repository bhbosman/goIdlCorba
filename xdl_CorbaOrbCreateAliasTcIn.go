// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ORB_create_alias_tc_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaOrbCreateAliasTcIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	OriginalType CorbaTypeCode `json:"OriginalType"`
}

//noinspection GoSnakeCaseUsage
const CorbaOrbCreateAliasTcInId_Const = "IDL:CORBA/ORB_create_alias_tc_In:1.0"

func (self *CorbaOrbCreateAliasTcIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaOrbCreateAliasTcIn) GoString() string {
	return self.String()
}

func (self *CorbaOrbCreateAliasTcIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Id, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Name, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.OriginalType, err = CorbaTypeCodeHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateAliasTcIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaOrbCreateAliasTcIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Id)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Name)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaTypeCodeHelper.Write(stream, self.OriginalType)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaOrbCreateAliasTcIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaOrbCreateAliasTcInHelper = CorbaOrbCreateAliasTcIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaOrbCreateAliasTcInId_Const,
			__goidl__.StructType,
			"CORBA_ORB.idl",
			"xdl_CorbaOrbCreateAliasTcIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaOrbCreateAliasTcIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaOrbCreateAliasTcIn{}
			},
			__reflect__.TypeOf((*CorbaOrbCreateAliasTcIn)(nil))))
}
