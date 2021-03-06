// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Enum declaration: "CORBA::TCKind", generatedBy by: "WriteEnumDcl"
type CorbaTCKind uint32

//noinspection GoUnusedConst
const (
	CorbaTCKindTkNull              CorbaTCKind = 0
	CorbaTCKindTkVoid              CorbaTCKind = 1
	CorbaTCKindTkShort             CorbaTCKind = 2
	CorbaTCKindTkLong              CorbaTCKind = 3
	CorbaTCKindTkUshort            CorbaTCKind = 4
	CorbaTCKindTkUlong             CorbaTCKind = 5
	CorbaTCKindTkFloat             CorbaTCKind = 6
	CorbaTCKindTkDouble            CorbaTCKind = 7
	CorbaTCKindTkBoolean           CorbaTCKind = 8
	CorbaTCKindTkChar              CorbaTCKind = 9
	CorbaTCKindTkOctet             CorbaTCKind = 10
	CorbaTCKindTkAny               CorbaTCKind = 11
	CorbaTCKindTkTypeCode          CorbaTCKind = 12
	CorbaTCKindTkPrincipal         CorbaTCKind = 13
	CorbaTCKindTkObjref            CorbaTCKind = 14
	CorbaTCKindTkStruct            CorbaTCKind = 15
	CorbaTCKindTkUnion             CorbaTCKind = 16
	CorbaTCKindTkEnum              CorbaTCKind = 17
	CorbaTCKindTkString            CorbaTCKind = 18
	CorbaTCKindTkSequence          CorbaTCKind = 19
	CorbaTCKindTkArray             CorbaTCKind = 20
	CorbaTCKindTkAlias             CorbaTCKind = 21
	CorbaTCKindTkExcept            CorbaTCKind = 22
	CorbaTCKindTkLonglong          CorbaTCKind = 23
	CorbaTCKindTkUlonglong         CorbaTCKind = 24
	CorbaTCKindTkLongdouble        CorbaTCKind = 25
	CorbaTCKindTkWchar             CorbaTCKind = 26
	CorbaTCKindTkWstring           CorbaTCKind = 27
	CorbaTCKindTkFixed             CorbaTCKind = 28
	CorbaTCKindTkValue             CorbaTCKind = 29
	CorbaTCKindTkValueBox          CorbaTCKind = 30
	CorbaTCKindTkNative            CorbaTCKind = 31
	CorbaTCKindTkAbstractInterface CorbaTCKind = 32
	CorbaTCKindTkLocalInterface    CorbaTCKind = 33
	CorbaTCKindTkComponent         CorbaTCKind = 34
	CorbaTCKindTkHome              CorbaTCKind = 35
	CorbaTCKindTkEvent             CorbaTCKind = 36
)

//noinspection GoSnakeCaseUsage
type CorbaTCKind_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaTCKindId_Const = "IDL:omg.org/CORBA/TCKind:1.0"

func (self CorbaTCKind_Helper) Id() string {
	return CorbaTCKindId_Const
}

func (self CorbaTCKind_Helper) Read(stream __goidl__.IReadAny) (uint32, error) {
	result, err := __goidl__.IdlUInt32Helper.Read(stream)
	return result, err
}

func (self CorbaTCKind_Helper) Write(stream __goidl__.IWriteAny, v uint32) error {
	return __goidl__.IdlUInt32Helper.Write(stream, v)
}


//noinspection GoUnusedGlobalVariable
var CorbaTCKindHelper = CorbaTCKind_Helper{}

func init() {
}
