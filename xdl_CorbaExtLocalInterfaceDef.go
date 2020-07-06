// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::ExtLocalInterfaceDef", generatedBy by: "WriteInterface"
type CorbaExtLocalInterfaceDef interface {
	// Original name: "contents"
	Contents(params CorbaExtLocalInterfaceDefContentsIn) (CorbaExtLocalInterfaceDefContentsOut, error)
	// Original name: "create_abstract_interface"
	CreateAbstractInterface(params CorbaExtLocalInterfaceDefCreateAbstractInterfaceIn) (CorbaExtLocalInterfaceDefCreateAbstractInterfaceOut, error)
	// Original name: "create_alias"
	CreateAlias(params CorbaExtLocalInterfaceDefCreateAliasIn) (CorbaExtLocalInterfaceDefCreateAliasOut, error)
	// Original name: "create_attribute"
	CreateAttribute(params CorbaExtLocalInterfaceDefCreateAttributeIn) (CorbaExtLocalInterfaceDefCreateAttributeOut, error)
	// Original name: "create_constant"
	CreateConstant(params CorbaExtLocalInterfaceDefCreateConstantIn) (CorbaExtLocalInterfaceDefCreateConstantOut, error)
	// Original name: "create_enum"
	CreateEnum(params CorbaExtLocalInterfaceDefCreateEnumIn) (CorbaExtLocalInterfaceDefCreateEnumOut, error)
	// Original name: "create_exception"
	CreateException(params CorbaExtLocalInterfaceDefCreateExceptionIn) (CorbaExtLocalInterfaceDefCreateExceptionOut, error)
	// Original name: "create_ext_attribute"
	CreateExtAttribute(params CorbaExtLocalInterfaceDefCreateExtAttributeIn) (CorbaExtLocalInterfaceDefCreateExtAttributeOut, error)
	// Original name: "create_ext_value"
	CreateExtValue(params CorbaExtLocalInterfaceDefCreateExtValueIn) (CorbaExtLocalInterfaceDefCreateExtValueOut, error)
	// Original name: "create_interface"
	CreateInterface(params CorbaExtLocalInterfaceDefCreateInterfaceIn) (CorbaExtLocalInterfaceDefCreateInterfaceOut, error)
	// Original name: "create_local_interface"
	CreateLocalInterface(params CorbaExtLocalInterfaceDefCreateLocalInterfaceIn) (CorbaExtLocalInterfaceDefCreateLocalInterfaceOut, error)
	// Original name: "create_module"
	CreateModule(params CorbaExtLocalInterfaceDefCreateModuleIn) (CorbaExtLocalInterfaceDefCreateModuleOut, error)
	// Original name: "create_native"
	CreateNative(params CorbaExtLocalInterfaceDefCreateNativeIn) (CorbaExtLocalInterfaceDefCreateNativeOut, error)
	// Original name: "create_operation"
	CreateOperation(params CorbaExtLocalInterfaceDefCreateOperationIn) (CorbaExtLocalInterfaceDefCreateOperationOut, error)
	// Original name: "create_struct"
	CreateStruct(params CorbaExtLocalInterfaceDefCreateStructIn) (CorbaExtLocalInterfaceDefCreateStructOut, error)
	// Original name: "create_union"
	CreateUnion(params CorbaExtLocalInterfaceDefCreateUnionIn) (CorbaExtLocalInterfaceDefCreateUnionOut, error)
	// Original name: "create_value"
	CreateValue(params CorbaExtLocalInterfaceDefCreateValueIn) (CorbaExtLocalInterfaceDefCreateValueOut, error)
	// Original name: "create_value_box"
	CreateValueBox(params CorbaExtLocalInterfaceDefCreateValueBoxIn) (CorbaExtLocalInterfaceDefCreateValueBoxOut, error)
	// Original name: "describe_contents"
	DescribeContents(params CorbaExtLocalInterfaceDefDescribeContentsIn) (CorbaExtLocalInterfaceDefDescribeContentsOut, error)
	// Original name: "describe_ext_interface"
	DescribeExtInterface(params CorbaExtLocalInterfaceDefDescribeExtInterfaceIn) (CorbaExtLocalInterfaceDefDescribeExtInterfaceOut, error)
	// Original name: "describe_interface"
	DescribeInterface(params CorbaExtLocalInterfaceDefDescribeInterfaceIn) (CorbaExtLocalInterfaceDefDescribeInterfaceOut, error)
	// Original name: "destroy"
	Destroy(params CorbaExtLocalInterfaceDefDestroyIn) (CorbaExtLocalInterfaceDefDestroyOut, error)
	// Original name: "is_a"
	IsA(params CorbaExtLocalInterfaceDefIsAIn) (CorbaExtLocalInterfaceDefIsAOut, error)
	// Original name: "lookup"
	Lookup(params CorbaExtLocalInterfaceDefLookupIn) (CorbaExtLocalInterfaceDefLookupOut, error)
	// Original name: "lookup_name"
	LookupName(params CorbaExtLocalInterfaceDefLookupNameIn) (CorbaExtLocalInterfaceDefLookupNameOut, error)
}

const CorbaExtLocalInterfaceDefContentsOperation = "contents"
const CorbaExtLocalInterfaceDefCreateAbstractInterfaceOperation = "create_abstract_interface"
const CorbaExtLocalInterfaceDefCreateAliasOperation = "create_alias"
const CorbaExtLocalInterfaceDefCreateAttributeOperation = "create_attribute"
const CorbaExtLocalInterfaceDefCreateConstantOperation = "create_constant"
const CorbaExtLocalInterfaceDefCreateEnumOperation = "create_enum"
const CorbaExtLocalInterfaceDefCreateExceptionOperation = "create_exception"
const CorbaExtLocalInterfaceDefCreateExtAttributeOperation = "create_ext_attribute"
const CorbaExtLocalInterfaceDefCreateExtValueOperation = "create_ext_value"
const CorbaExtLocalInterfaceDefCreateInterfaceOperation = "create_interface"
const CorbaExtLocalInterfaceDefCreateLocalInterfaceOperation = "create_local_interface"
const CorbaExtLocalInterfaceDefCreateModuleOperation = "create_module"
const CorbaExtLocalInterfaceDefCreateNativeOperation = "create_native"
const CorbaExtLocalInterfaceDefCreateOperationOperation = "create_operation"
const CorbaExtLocalInterfaceDefCreateStructOperation = "create_struct"
const CorbaExtLocalInterfaceDefCreateUnionOperation = "create_union"
const CorbaExtLocalInterfaceDefCreateValueOperation = "create_value"
const CorbaExtLocalInterfaceDefCreateValueBoxOperation = "create_value_box"
const CorbaExtLocalInterfaceDefDescribeContentsOperation = "describe_contents"
const CorbaExtLocalInterfaceDefDescribeExtInterfaceOperation = "describe_ext_interface"
const CorbaExtLocalInterfaceDefDescribeInterfaceOperation = "describe_interface"
const CorbaExtLocalInterfaceDefDestroyOperation = "destroy"
const CorbaExtLocalInterfaceDefIsAOperation = "is_a"
const CorbaExtLocalInterfaceDefLookupOperation = "lookup"
const CorbaExtLocalInterfaceDefLookupNameOperation = "lookup_name"
//noinspection GoSnakeCaseUsage
type CorbaExtLocalInterfaceDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaExtLocalInterfaceDefId_Const = "IDL:omg.org/CORBA/ExtLocalInterfaceDef:1.0"

func (self CorbaExtLocalInterfaceDef_Helper) Id() string {
	return CorbaExtLocalInterfaceDefId_Const
}

func (self CorbaExtLocalInterfaceDef_Helper) Read(stream __goidl__.IReadAny) (CorbaExtLocalInterfaceDef, error) {
	return nil, nil
}

func (self CorbaExtLocalInterfaceDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaExtLocalInterfaceDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaExtLocalInterfaceDefHelper = CorbaExtLocalInterfaceDef_Helper{}

func init() {
}
