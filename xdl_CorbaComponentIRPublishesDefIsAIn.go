// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::PublishesDef_is_a_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRPublishesDefIsAIn struct {
	__goidl__.IdlObject
	EventId string `json:"EventId"`
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRPublishesDefIsAInId_Const = "IDL:CORBA/ComponentIR/PublishesDef_is_a_In:1.0"

func (self *CorbaComponentIRPublishesDefIsAIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRPublishesDefIsAIn) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRPublishesDefIsAIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.EventId, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRPublishesDefIsAIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRPublishesDefIsAIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.EventId)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRPublishesDefIsAIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRPublishesDefIsAInHelper = CorbaComponentIRPublishesDefIsAIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRPublishesDefIsAInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRPublishesDefIsAIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRPublishesDefIsAIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRPublishesDefIsAIn{}
			},
			__reflect__.TypeOf((*CorbaComponentIRPublishesDefIsAIn)(nil))))
}
