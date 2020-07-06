// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __fmt__ "fmt"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Exception declaration: "CORBA::TypeCode::BadKind", generatedBy by: "WriteStructDcl"
// Exception Decl: true
type CorbaTypeCodeBadKind struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaTypeCodeBadKindId_Const = "IDL:omg.org/CORBA/TypeCode/BadKind:1.0"

func (self *CorbaTypeCodeBadKind) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaTypeCodeBadKind) Error() string{
	return 	__fmt__.Sprintf("Error of type CorbaTypeCodeBadKind(%v)", self.String())
}
func (self *CorbaTypeCodeBadKind) GoString() string {
	return self.String()
}

func (self *CorbaTypeCodeBadKind) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeBadKind) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaTypeCodeBadKind) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaTypeCodeBadKind_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaTypeCodeBadKindHelper = CorbaTypeCodeBadKind_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaTypeCodeBadKindId_Const,
			__goidl__.StructType,
			"CORBA_TypeCode.idl",
			"xdl_CorbaTypeCodeBadKind.go",
			func() __goidl__.IIdlObject {
				return &CorbaTypeCodeBadKind{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaTypeCodeBadKind{}
			},
			__reflect__.TypeOf((*CorbaTypeCodeBadKind)(nil))))
}