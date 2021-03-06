// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::InterfaceDef_create_value_box_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaInterfaceDefCreateValueBoxIn struct {
	__goidl__.IdlObject
	Id string `json:"Id"`
	Name string `json:"Name"`
	Version string `json:"Version"`
	OriginalTypeDef CorbaIDLType `json:"OriginalTypeDef"`
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefCreateValueBoxInId_Const = "IDL:CORBA/InterfaceDef_create_value_box_In:1.0"

func (self *CorbaInterfaceDefCreateValueBoxIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaInterfaceDefCreateValueBoxIn) GoString() string {
	return self.String()
}

func (self *CorbaInterfaceDefCreateValueBoxIn) ReadValue(stream __goidl__.IReadAny) error {
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
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Version, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlInterface)
	self.OriginalTypeDef, err = CorbaIDLTypeHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateValueBoxIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaInterfaceDefCreateValueBoxIn) Write(stream __goidl__.IWriteAny) error {
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
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Version)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlInterface)
	err = CorbaIDLTypeHelper.Write(stream, self.OriginalTypeDef)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaInterfaceDefCreateValueBoxIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefCreateValueBoxInHelper = CorbaInterfaceDefCreateValueBoxIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaInterfaceDefCreateValueBoxInId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaInterfaceDefCreateValueBoxIn.go",
			func() __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateValueBoxIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaInterfaceDefCreateValueBoxIn{}
			},
			__reflect__.TypeOf((*CorbaInterfaceDefCreateValueBoxIn)(nil))))
}
