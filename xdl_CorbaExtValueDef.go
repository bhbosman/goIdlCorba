// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::ExtValueDef", generatedBy by: "WriteInterface"
type CorbaExtValueDef interface {
	// Original name: "contents"
	Contents(params CorbaExtValueDefContentsIn) (CorbaExtValueDefContentsOut, error)
	// Original name: "create_abstract_interface"
	CreateAbstractInterface(params CorbaExtValueDefCreateAbstractInterfaceIn) (CorbaExtValueDefCreateAbstractInterfaceOut, error)
	// Original name: "create_alias"
	CreateAlias(params CorbaExtValueDefCreateAliasIn) (CorbaExtValueDefCreateAliasOut, error)
	// Original name: "create_attribute"
	CreateAttribute(params CorbaExtValueDefCreateAttributeIn) (CorbaExtValueDefCreateAttributeOut, error)
	// Original name: "create_constant"
	CreateConstant(params CorbaExtValueDefCreateConstantIn) (CorbaExtValueDefCreateConstantOut, error)
	// Original name: "create_enum"
	CreateEnum(params CorbaExtValueDefCreateEnumIn) (CorbaExtValueDefCreateEnumOut, error)
	// Original name: "create_exception"
	CreateException(params CorbaExtValueDefCreateExceptionIn) (CorbaExtValueDefCreateExceptionOut, error)
	// Original name: "create_ext_attribute"
	CreateExtAttribute(params CorbaExtValueDefCreateExtAttributeIn) (CorbaExtValueDefCreateExtAttributeOut, error)
	// Original name: "create_ext_value"
	CreateExtValue(params CorbaExtValueDefCreateExtValueIn) (CorbaExtValueDefCreateExtValueOut, error)
	// Original name: "create_interface"
	CreateInterface(params CorbaExtValueDefCreateInterfaceIn) (CorbaExtValueDefCreateInterfaceOut, error)
	// Original name: "create_local_interface"
	CreateLocalInterface(params CorbaExtValueDefCreateLocalInterfaceIn) (CorbaExtValueDefCreateLocalInterfaceOut, error)
	// Original name: "create_module"
	CreateModule(params CorbaExtValueDefCreateModuleIn) (CorbaExtValueDefCreateModuleOut, error)
	// Original name: "create_native"
	CreateNative(params CorbaExtValueDefCreateNativeIn) (CorbaExtValueDefCreateNativeOut, error)
	// Original name: "create_operation"
	CreateOperation(params CorbaExtValueDefCreateOperationIn) (CorbaExtValueDefCreateOperationOut, error)
	// Original name: "create_struct"
	CreateStruct(params CorbaExtValueDefCreateStructIn) (CorbaExtValueDefCreateStructOut, error)
	// Original name: "create_union"
	CreateUnion(params CorbaExtValueDefCreateUnionIn) (CorbaExtValueDefCreateUnionOut, error)
	// Original name: "create_value"
	CreateValue(params CorbaExtValueDefCreateValueIn) (CorbaExtValueDefCreateValueOut, error)
	// Original name: "create_value_box"
	CreateValueBox(params CorbaExtValueDefCreateValueBoxIn) (CorbaExtValueDefCreateValueBoxOut, error)
	// Original name: "create_value_member"
	CreateValueMember(params CorbaExtValueDefCreateValueMemberIn) (CorbaExtValueDefCreateValueMemberOut, error)
	// Original name: "describe_contents"
	DescribeContents(params CorbaExtValueDefDescribeContentsIn) (CorbaExtValueDefDescribeContentsOut, error)
	// Original name: "describe_ext_value"
	DescribeExtValue(params CorbaExtValueDefDescribeExtValueIn) (CorbaExtValueDefDescribeExtValueOut, error)
	// Original name: "describe_value"
	DescribeValue(params CorbaExtValueDefDescribeValueIn) (CorbaExtValueDefDescribeValueOut, error)
	// Original name: "destroy"
	Destroy(params CorbaExtValueDefDestroyIn) (CorbaExtValueDefDestroyOut, error)
	// Original name: "is_a"
	IsA(params CorbaExtValueDefIsAIn) (CorbaExtValueDefIsAOut, error)
	// Original name: "lookup"
	Lookup(params CorbaExtValueDefLookupIn) (CorbaExtValueDefLookupOut, error)
	// Original name: "lookup_name"
	LookupName(params CorbaExtValueDefLookupNameIn) (CorbaExtValueDefLookupNameOut, error)
	// Property ExtInitializers
	// Get Property ExtInitializers
	GetExtInitializers() (CorbaExtInitializerSeq, error)
	// Set Property ExtInitializers
	SetExtInitializers(value CorbaExtInitializerSeq) error
}

const CorbaExtValueDefContentsOperation = "contents"
const CorbaExtValueDefCreateAbstractInterfaceOperation = "create_abstract_interface"
const CorbaExtValueDefCreateAliasOperation = "create_alias"
const CorbaExtValueDefCreateAttributeOperation = "create_attribute"
const CorbaExtValueDefCreateConstantOperation = "create_constant"
const CorbaExtValueDefCreateEnumOperation = "create_enum"
const CorbaExtValueDefCreateExceptionOperation = "create_exception"
const CorbaExtValueDefCreateExtAttributeOperation = "create_ext_attribute"
const CorbaExtValueDefCreateExtValueOperation = "create_ext_value"
const CorbaExtValueDefCreateInterfaceOperation = "create_interface"
const CorbaExtValueDefCreateLocalInterfaceOperation = "create_local_interface"
const CorbaExtValueDefCreateModuleOperation = "create_module"
const CorbaExtValueDefCreateNativeOperation = "create_native"
const CorbaExtValueDefCreateOperationOperation = "create_operation"
const CorbaExtValueDefCreateStructOperation = "create_struct"
const CorbaExtValueDefCreateUnionOperation = "create_union"
const CorbaExtValueDefCreateValueOperation = "create_value"
const CorbaExtValueDefCreateValueBoxOperation = "create_value_box"
const CorbaExtValueDefCreateValueMemberOperation = "create_value_member"
const CorbaExtValueDefDescribeContentsOperation = "describe_contents"
const CorbaExtValueDefDescribeExtValueOperation = "describe_ext_value"
const CorbaExtValueDefDescribeValueOperation = "describe_value"
const CorbaExtValueDefDestroyOperation = "destroy"
const CorbaExtValueDefIsAOperation = "is_a"
const CorbaExtValueDefLookupOperation = "lookup"
const CorbaExtValueDefLookupNameOperation = "lookup_name"
//noinspection GoSnakeCaseUsage
type CorbaExtValueDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaExtValueDefId_Const = "IDL:omg.org/CORBA/ExtValueDef:1.0"

func (self CorbaExtValueDef_Helper) Id() string {
	return CorbaExtValueDefId_Const
}

func (self CorbaExtValueDef_Helper) Read(stream __goidl__.IReadAny) (CorbaExtValueDef, error) {
	return nil, nil
}

func (self CorbaExtValueDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaExtValueDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaExtValueDefHelper = CorbaExtValueDef_Helper{}

func init() {
}
