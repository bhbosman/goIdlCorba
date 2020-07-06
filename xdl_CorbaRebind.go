// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __fmt__ "fmt"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Exception declaration: "CORBA::REBIND", generatedBy by: "WriteStructDcl"
// Exception Decl: true
type CorbaRebind struct {
	__goidl__.IdlObject
	Minor uint32 `json:"Minor"`
	Completed uint32 `json:"Completed"`
}

//noinspection GoSnakeCaseUsage
const CorbaRebindId_Const = "IDL:omg.org/CORBA/REBIND:1.0"

func (self *CorbaRebind) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaRebind) Error() string{
	return 	__fmt__.Sprintf("Error of type CorbaRebind(%v)", self.String())
}
func (self *CorbaRebind) GoString() string {
	return self.String()
}

func (self *CorbaRebind) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedLongType)
	self.Minor, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlEnum)
	self.Completed, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRebind) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaRebind) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedLongType)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Minor)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlEnum)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Completed)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaRebind_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaRebindHelper = CorbaRebind_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaRebindId_Const,
			__goidl__.StructType,
			"CORBA_StandardExceptions.idl",
			"xdl_CorbaRebind.go",
			func() __goidl__.IIdlObject {
				return &CorbaRebind{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaRebind{}
			},
			__reflect__.TypeOf((*CorbaRebind)(nil))))
}