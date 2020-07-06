// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __fmt__ "fmt"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Exception declaration: "CORBA::PolicyError", generatedBy by: "WriteStructDcl"
// Exception Decl: true
type CorbaPolicyError struct {
	__goidl__.IdlObject
	Reason int16 `json:"Reason"`
}

//noinspection GoSnakeCaseUsage
const CorbaPolicyErrorId_Const = "IDL:omg.org/CORBA/PolicyError:1.0"

func (self *CorbaPolicyError) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaPolicyError) Error() string{
	return 	__fmt__.Sprintf("Error of type CorbaPolicyError(%v)", self.String())
}
func (self *CorbaPolicyError) GoString() string {
	return self.String()
}

func (self *CorbaPolicyError) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(ShortType)
	self.Reason, err = __goidl__.IdlInt16Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPolicyError) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaPolicyError) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(ShortType)
	err = __goidl__.IdlInt16Helper.Write(stream, self.Reason)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaPolicyError_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaPolicyErrorHelper = CorbaPolicyError_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaPolicyErrorId_Const,
			__goidl__.StructType,
			"CORBA_Policy.idl",
			"xdl_CorbaPolicyError.go",
			func() __goidl__.IIdlObject {
				return &CorbaPolicyError{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaPolicyError{}
			},
			__reflect__.TypeOf((*CorbaPolicyError)(nil))))
}
