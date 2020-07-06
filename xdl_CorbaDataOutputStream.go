// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::DataOutputStream", generatedBy by: "WriteInterface"
type CorbaDataOutputStream interface {
	// Original name: "write_Abstract"
	WriteAbstract(params CorbaDataOutputStreamWriteAbstractIn) (CorbaDataOutputStreamWriteAbstractOut, error)
	// Original name: "write_Object"
	WriteObject(params CorbaDataOutputStreamWriteObjectIn) (CorbaDataOutputStreamWriteObjectOut, error)
	// Original name: "write_TypeCode"
	WriteTypeCode(params CorbaDataOutputStreamWriteTypeCodeIn) (CorbaDataOutputStreamWriteTypeCodeOut, error)
	// Original name: "write_Value"
	WriteValue(params CorbaDataOutputStreamWriteValueIn) (CorbaDataOutputStreamWriteValueOut, error)
	// Original name: "write_any"
	WriteAny(params CorbaDataOutputStreamWriteAnyIn) (CorbaDataOutputStreamWriteAnyOut, error)
	// Original name: "write_any_array"
	WriteAnyArray(params CorbaDataOutputStreamWriteAnyArrayIn) (CorbaDataOutputStreamWriteAnyArrayOut, error)
	// Original name: "write_boolean"
	WriteBoolean(params CorbaDataOutputStreamWriteBooleanIn) (CorbaDataOutputStreamWriteBooleanOut, error)
	// Original name: "write_boolean_array"
	WriteBooleanArray(params CorbaDataOutputStreamWriteBooleanArrayIn) (CorbaDataOutputStreamWriteBooleanArrayOut, error)
	// Original name: "write_char"
	WriteChar(params CorbaDataOutputStreamWriteCharIn) (CorbaDataOutputStreamWriteCharOut, error)
	// Original name: "write_char_array"
	WriteCharArray(params CorbaDataOutputStreamWriteCharArrayIn) (CorbaDataOutputStreamWriteCharArrayOut, error)
	// Original name: "write_double"
	WriteDouble(params CorbaDataOutputStreamWriteDoubleIn) (CorbaDataOutputStreamWriteDoubleOut, error)
	// Original name: "write_double_array"
	WriteDoubleArray(params CorbaDataOutputStreamWriteDoubleArrayIn) (CorbaDataOutputStreamWriteDoubleArrayOut, error)
	//Exceptions for : WriteFixed
	//	CorbaBadFixedValue
	// Original name: "write_fixed"
	WriteFixed(params CorbaDataOutputStreamWriteFixedIn) (CorbaDataOutputStreamWriteFixedOut, error)
	//Exceptions for : WriteFixedArray
	//	CorbaBadFixedValue
	// Original name: "write_fixed_array"
	WriteFixedArray(params CorbaDataOutputStreamWriteFixedArrayIn) (CorbaDataOutputStreamWriteFixedArrayOut, error)
	// Original name: "write_float"
	WriteFloat(params CorbaDataOutputStreamWriteFloatIn) (CorbaDataOutputStreamWriteFloatOut, error)
	// Original name: "write_float_array"
	WriteFloatArray(params CorbaDataOutputStreamWriteFloatArrayIn) (CorbaDataOutputStreamWriteFloatArrayOut, error)
	// Original name: "write_long"
	WriteLong(params CorbaDataOutputStreamWriteLongIn) (CorbaDataOutputStreamWriteLongOut, error)
	// Original name: "write_long_array"
	WriteLongArray(params CorbaDataOutputStreamWriteLongArrayIn) (CorbaDataOutputStreamWriteLongArrayOut, error)
	// Original name: "write_long_double_array"
	WriteLongDoubleArray(params CorbaDataOutputStreamWriteLongDoubleArrayIn) (CorbaDataOutputStreamWriteLongDoubleArrayOut, error)
	// Original name: "write_longdouble"
	WriteLongdouble(params CorbaDataOutputStreamWriteLongdoubleIn) (CorbaDataOutputStreamWriteLongdoubleOut, error)
	// Original name: "write_longlong"
	WriteLonglong(params CorbaDataOutputStreamWriteLonglongIn) (CorbaDataOutputStreamWriteLonglongOut, error)
	// Original name: "write_longlong_array"
	WriteLonglongArray(params CorbaDataOutputStreamWriteLonglongArrayIn) (CorbaDataOutputStreamWriteLonglongArrayOut, error)
	// Original name: "write_octet"
	WriteOctet(params CorbaDataOutputStreamWriteOctetIn) (CorbaDataOutputStreamWriteOctetOut, error)
	// Original name: "write_octet_array"
	WriteOctetArray(params CorbaDataOutputStreamWriteOctetArrayIn) (CorbaDataOutputStreamWriteOctetArrayOut, error)
	// Original name: "write_short"
	WriteShort(params CorbaDataOutputStreamWriteShortIn) (CorbaDataOutputStreamWriteShortOut, error)
	// Original name: "write_short_array"
	WriteShortArray(params CorbaDataOutputStreamWriteShortArrayIn) (CorbaDataOutputStreamWriteShortArrayOut, error)
	// Original name: "write_string"
	WriteString(params CorbaDataOutputStreamWriteStringIn) (CorbaDataOutputStreamWriteStringOut, error)
	// Original name: "write_ulong"
	WriteUlong(params CorbaDataOutputStreamWriteUlongIn) (CorbaDataOutputStreamWriteUlongOut, error)
	// Original name: "write_ulong_array"
	WriteUlongArray(params CorbaDataOutputStreamWriteUlongArrayIn) (CorbaDataOutputStreamWriteUlongArrayOut, error)
	// Original name: "write_ulonglong"
	WriteUlonglong(params CorbaDataOutputStreamWriteUlonglongIn) (CorbaDataOutputStreamWriteUlonglongOut, error)
	// Original name: "write_ulonglong_array"
	WriteUlonglongArray(params CorbaDataOutputStreamWriteUlonglongArrayIn) (CorbaDataOutputStreamWriteUlonglongArrayOut, error)
	// Original name: "write_ushort"
	WriteUshort(params CorbaDataOutputStreamWriteUshortIn) (CorbaDataOutputStreamWriteUshortOut, error)
	// Original name: "write_ushort_array"
	WriteUshortArray(params CorbaDataOutputStreamWriteUshortArrayIn) (CorbaDataOutputStreamWriteUshortArrayOut, error)
	// Original name: "write_wchar"
	WriteWchar(params CorbaDataOutputStreamWriteWcharIn) (CorbaDataOutputStreamWriteWcharOut, error)
	// Original name: "write_wchar_array"
	WriteWcharArray(params CorbaDataOutputStreamWriteWcharArrayIn) (CorbaDataOutputStreamWriteWcharArrayOut, error)
	// Original name: "write_wstring"
	WriteWstring(params CorbaDataOutputStreamWriteWstringIn) (CorbaDataOutputStreamWriteWstringOut, error)
}

const CorbaDataOutputStreamWriteAbstractOperation = "write_Abstract"
const CorbaDataOutputStreamWriteObjectOperation = "write_Object"
const CorbaDataOutputStreamWriteTypeCodeOperation = "write_TypeCode"
const CorbaDataOutputStreamWriteValueOperation = "write_Value"
const CorbaDataOutputStreamWriteAnyOperation = "write_any"
const CorbaDataOutputStreamWriteAnyArrayOperation = "write_any_array"
const CorbaDataOutputStreamWriteBooleanOperation = "write_boolean"
const CorbaDataOutputStreamWriteBooleanArrayOperation = "write_boolean_array"
const CorbaDataOutputStreamWriteCharOperation = "write_char"
const CorbaDataOutputStreamWriteCharArrayOperation = "write_char_array"
const CorbaDataOutputStreamWriteDoubleOperation = "write_double"
const CorbaDataOutputStreamWriteDoubleArrayOperation = "write_double_array"
const CorbaDataOutputStreamWriteFixedOperation = "write_fixed"
const CorbaDataOutputStreamWriteFixedArrayOperation = "write_fixed_array"
const CorbaDataOutputStreamWriteFloatOperation = "write_float"
const CorbaDataOutputStreamWriteFloatArrayOperation = "write_float_array"
const CorbaDataOutputStreamWriteLongOperation = "write_long"
const CorbaDataOutputStreamWriteLongArrayOperation = "write_long_array"
const CorbaDataOutputStreamWriteLongDoubleArrayOperation = "write_long_double_array"
const CorbaDataOutputStreamWriteLongdoubleOperation = "write_longdouble"
const CorbaDataOutputStreamWriteLonglongOperation = "write_longlong"
const CorbaDataOutputStreamWriteLonglongArrayOperation = "write_longlong_array"
const CorbaDataOutputStreamWriteOctetOperation = "write_octet"
const CorbaDataOutputStreamWriteOctetArrayOperation = "write_octet_array"
const CorbaDataOutputStreamWriteShortOperation = "write_short"
const CorbaDataOutputStreamWriteShortArrayOperation = "write_short_array"
const CorbaDataOutputStreamWriteStringOperation = "write_string"
const CorbaDataOutputStreamWriteUlongOperation = "write_ulong"
const CorbaDataOutputStreamWriteUlongArrayOperation = "write_ulong_array"
const CorbaDataOutputStreamWriteUlonglongOperation = "write_ulonglong"
const CorbaDataOutputStreamWriteUlonglongArrayOperation = "write_ulonglong_array"
const CorbaDataOutputStreamWriteUshortOperation = "write_ushort"
const CorbaDataOutputStreamWriteUshortArrayOperation = "write_ushort_array"
const CorbaDataOutputStreamWriteWcharOperation = "write_wchar"
const CorbaDataOutputStreamWriteWcharArrayOperation = "write_wchar_array"
const CorbaDataOutputStreamWriteWstringOperation = "write_wstring"
//noinspection GoSnakeCaseUsage
type CorbaDataOutputStream_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaDataOutputStreamId_Const = "IDL:omg.org/CORBA/DataOutputStream:1.0"

func (self CorbaDataOutputStream_Helper) Id() string {
	return CorbaDataOutputStreamId_Const
}

func (self CorbaDataOutputStream_Helper) Read(stream __goidl__.IReadAny) (CorbaDataOutputStream, error) {
	return nil, nil
}

func (self CorbaDataOutputStream_Helper) Write(stream __goidl__.IWriteAny, v CorbaDataOutputStream) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaDataOutputStreamHelper = CorbaDataOutputStream_Helper{}

func init() {
}
