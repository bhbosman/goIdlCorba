// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CORBA::ComponentIR::Repository_get_canonical_typecode_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CorbaComponentIRRepositoryGetCanonicalTypecodeOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CorbaComponentIRRepositoryGetCanonicalTypecodeOutId_Const = "IDL:CORBA/ComponentIR/Repository_get_canonical_typecode_Out:1.0"

func (self *CorbaComponentIRRepositoryGetCanonicalTypecodeOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CorbaComponentIRRepositoryGetCanonicalTypecodeOut) GoString() string {
	return self.String()
}

func (self *CorbaComponentIRRepositoryGetCanonicalTypecodeOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryGetCanonicalTypecodeOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CorbaComponentIRRepositoryGetCanonicalTypecodeOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CorbaComponentIRRepositoryGetCanonicalTypecodeOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CorbaComponentIRRepositoryGetCanonicalTypecodeOutHelper = CorbaComponentIRRepositoryGetCanonicalTypecodeOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CorbaComponentIRRepositoryGetCanonicalTypecodeOutId_Const,
			__goidl__.StructType,
			"CORBA_InterfaceRepository.idl",
			"xdl_CorbaComponentIRRepositoryGetCanonicalTypecodeOut.go",
			func() __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryGetCanonicalTypecodeOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CorbaComponentIRRepositoryGetCanonicalTypecodeOut{}
			},
			__reflect__.TypeOf((*CorbaComponentIRRepositoryGetCanonicalTypecodeOut)(nil))))
}
