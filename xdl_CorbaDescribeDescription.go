// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::describe::Description", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaDescribeDescription struct {
	__goidl__.IdlObject
	Kind uint32 `json:"Kind"`
	Value __goidl__.IdlAny `json:"Value"`
}

//noinspection GoSnakeCaseUsage
const CorbaDescribeDescriptionId_Const = "IDL:omg.org/CORBA/describe/Description:1.0"

func (self *CorbaDescribeDescription) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaDescribeDescription) GoString() string {
	return self.String()
}

func (self *CorbaDescribeDescription) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlEnum)
	self.Kind, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(AnyType)
	self.Value, err = __goidl__.IdlAnyHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDescribeDescription) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaDescribeDescription) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlEnum)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Kind)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(AnyType)
	err = __goidl__.IdlAnyHelper.Write(stream, self.Value)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaDescribeDescription_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaDescribeDescriptionHelper = CorbaDescribeDescription_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaDescribeDescriptionId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaDescribeDescription.go",
			func() __goidl__.IIdlObject {
				return &CorbaDescribeDescription{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaDescribeDescription{}
			},
			__reflect__.TypeOf((*CorbaDescribeDescription)(nil))))
}
