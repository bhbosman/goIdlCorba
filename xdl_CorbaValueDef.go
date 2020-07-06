// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::ValueDef", generatedBy by: "WriteInterface"
type CorbaValueDef interface {
	// Original name: "contents"
	Contents(params CorbaValueDefContentsIn) (CorbaValueDefContentsOut, error)
	// Original name: "create_abstract_interface"
	CreateAbstractInterface(params CorbaValueDefCreateAbstractInterfaceIn) (CorbaValueDefCreateAbstractInterfaceOut, error)
	// Original name: "create_alias"
	CreateAlias(params CorbaValueDefCreateAliasIn) (CorbaValueDefCreateAliasOut, error)
	// Original name: "create_attribute"
	CreateAttribute(params CorbaValueDefCreateAttributeIn) (CorbaValueDefCreateAttributeOut, error)
	// Original name: "create_constant"
	CreateConstant(params CorbaValueDefCreateConstantIn) (CorbaValueDefCreateConstantOut, error)
	// Original name: "create_enum"
	CreateEnum(params CorbaValueDefCreateEnumIn) (CorbaValueDefCreateEnumOut, error)
	// Original name: "create_exception"
	CreateException(params CorbaValueDefCreateExceptionIn) (CorbaValueDefCreateExceptionOut, error)
	// Original name: "create_ext_value"
	CreateExtValue(params CorbaValueDefCreateExtValueIn) (CorbaValueDefCreateExtValueOut, error)
	// Original name: "create_interface"
	CreateInterface(params CorbaValueDefCreateInterfaceIn) (CorbaValueDefCreateInterfaceOut, error)
	// Original name: "create_local_interface"
	CreateLocalInterface(params CorbaValueDefCreateLocalInterfaceIn) (CorbaValueDefCreateLocalInterfaceOut, error)
	// Original name: "create_module"
	CreateModule(params CorbaValueDefCreateModuleIn) (CorbaValueDefCreateModuleOut, error)
	// Original name: "create_native"
	CreateNative(params CorbaValueDefCreateNativeIn) (CorbaValueDefCreateNativeOut, error)
	// Original name: "create_operation"
	CreateOperation(params CorbaValueDefCreateOperationIn) (CorbaValueDefCreateOperationOut, error)
	// Original name: "create_struct"
	CreateStruct(params CorbaValueDefCreateStructIn) (CorbaValueDefCreateStructOut, error)
	// Original name: "create_union"
	CreateUnion(params CorbaValueDefCreateUnionIn) (CorbaValueDefCreateUnionOut, error)
	// Original name: "create_value"
	CreateValue(params CorbaValueDefCreateValueIn) (CorbaValueDefCreateValueOut, error)
	// Original name: "create_value_box"
	CreateValueBox(params CorbaValueDefCreateValueBoxIn) (CorbaValueDefCreateValueBoxOut, error)
	// Original name: "create_value_member"
	CreateValueMember(params CorbaValueDefCreateValueMemberIn) (CorbaValueDefCreateValueMemberOut, error)
	// Original name: "describe_contents"
	DescribeContents(params CorbaValueDefDescribeContentsIn) (CorbaValueDefDescribeContentsOut, error)
	// Original name: "describe_value"
	DescribeValue(params CorbaValueDefDescribeValueIn) (CorbaValueDefDescribeValueOut, error)
	// Original name: "destroy"
	Destroy(params CorbaValueDefDestroyIn) (CorbaValueDefDestroyOut, error)
	// Original name: "is_a"
	IsA(params CorbaValueDefIsAIn) (CorbaValueDefIsAOut, error)
	// Original name: "lookup"
	Lookup(params CorbaValueDefLookupIn) (CorbaValueDefLookupOut, error)
	// Original name: "lookup_name"
	LookupName(params CorbaValueDefLookupNameIn) (CorbaValueDefLookupNameOut, error)
	// Property SupportedInterfaces
	// Get Property SupportedInterfaces
	GetSupportedInterfaces() (CorbaInterfaceDefSeq, error)
	// Set Property SupportedInterfaces
	SetSupportedInterfaces(value CorbaInterfaceDefSeq) error
	// Property Initializers
	// Get Property Initializers
	GetInitializers() (CorbaInitializerSeq, error)
	// Set Property Initializers
	SetInitializers(value CorbaInitializerSeq) error
	// Property BaseValue
	// Get Property BaseValue
	GetBaseValue() (CorbaValueDef, error)
	// Set Property BaseValue
	SetBaseValue(value CorbaValueDef) error
	// Property AbstractBaseValues
	// Get Property AbstractBaseValues
	GetAbstractBaseValues() (CorbaValueDefSeq, error)
	// Set Property AbstractBaseValues
	SetAbstractBaseValues(value CorbaValueDefSeq) error
	// Property IsAbstract
	// Get Property IsAbstract
	GetIsAbstract() (bool, error)
	// Set Property IsAbstract
	SetIsAbstract(value bool) error
	// Property IsCustom
	// Get Property IsCustom
	GetIsCustom() (bool, error)
	// Set Property IsCustom
	SetIsCustom(value bool) error
	// Property IsTruncatable
	// Get Property IsTruncatable
	GetIsTruncatable() (bool, error)
	// Set Property IsTruncatable
	SetIsTruncatable(value bool) error
}

const CorbaValueDefContentsOperation = "contents"
const CorbaValueDefCreateAbstractInterfaceOperation = "create_abstract_interface"
const CorbaValueDefCreateAliasOperation = "create_alias"
const CorbaValueDefCreateAttributeOperation = "create_attribute"
const CorbaValueDefCreateConstantOperation = "create_constant"
const CorbaValueDefCreateEnumOperation = "create_enum"
const CorbaValueDefCreateExceptionOperation = "create_exception"
const CorbaValueDefCreateExtValueOperation = "create_ext_value"
const CorbaValueDefCreateInterfaceOperation = "create_interface"
const CorbaValueDefCreateLocalInterfaceOperation = "create_local_interface"
const CorbaValueDefCreateModuleOperation = "create_module"
const CorbaValueDefCreateNativeOperation = "create_native"
const CorbaValueDefCreateOperationOperation = "create_operation"
const CorbaValueDefCreateStructOperation = "create_struct"
const CorbaValueDefCreateUnionOperation = "create_union"
const CorbaValueDefCreateValueOperation = "create_value"
const CorbaValueDefCreateValueBoxOperation = "create_value_box"
const CorbaValueDefCreateValueMemberOperation = "create_value_member"
const CorbaValueDefDescribeContentsOperation = "describe_contents"
const CorbaValueDefDescribeValueOperation = "describe_value"
const CorbaValueDefDestroyOperation = "destroy"
const CorbaValueDefIsAOperation = "is_a"
const CorbaValueDefLookupOperation = "lookup"
const CorbaValueDefLookupNameOperation = "lookup_name"
//noinspection GoSnakeCaseUsage
type CorbaValueDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaValueDefId_Const = "IDL:omg.org/CORBA/ValueDef:1.0"

func (self CorbaValueDef_Helper) Id() string {
	return CorbaValueDefId_Const
}

func (self CorbaValueDef_Helper) Read(stream __goidl__.IReadAny) (CorbaValueDef, error) {
	return nil, nil
}

func (self CorbaValueDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaValueDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaValueDefHelper = CorbaValueDef_Helper{}

func init() {
}
