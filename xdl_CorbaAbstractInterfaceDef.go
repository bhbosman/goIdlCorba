// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::AbstractInterfaceDef", generatedBy by: "WriteInterface"
type CorbaAbstractInterfaceDef interface {
	// Original name: "contents"
	Contents(params CorbaAbstractInterfaceDefContentsIn) (CorbaAbstractInterfaceDefContentsOut, error)
	// Original name: "create_abstract_interface"
	CreateAbstractInterface(params CorbaAbstractInterfaceDefCreateAbstractInterfaceIn) (CorbaAbstractInterfaceDefCreateAbstractInterfaceOut, error)
	// Original name: "create_alias"
	CreateAlias(params CorbaAbstractInterfaceDefCreateAliasIn) (CorbaAbstractInterfaceDefCreateAliasOut, error)
	// Original name: "create_attribute"
	CreateAttribute(params CorbaAbstractInterfaceDefCreateAttributeIn) (CorbaAbstractInterfaceDefCreateAttributeOut, error)
	// Original name: "create_constant"
	CreateConstant(params CorbaAbstractInterfaceDefCreateConstantIn) (CorbaAbstractInterfaceDefCreateConstantOut, error)
	// Original name: "create_enum"
	CreateEnum(params CorbaAbstractInterfaceDefCreateEnumIn) (CorbaAbstractInterfaceDefCreateEnumOut, error)
	// Original name: "create_exception"
	CreateException(params CorbaAbstractInterfaceDefCreateExceptionIn) (CorbaAbstractInterfaceDefCreateExceptionOut, error)
	// Original name: "create_ext_value"
	CreateExtValue(params CorbaAbstractInterfaceDefCreateExtValueIn) (CorbaAbstractInterfaceDefCreateExtValueOut, error)
	// Original name: "create_interface"
	CreateInterface(params CorbaAbstractInterfaceDefCreateInterfaceIn) (CorbaAbstractInterfaceDefCreateInterfaceOut, error)
	// Original name: "create_local_interface"
	CreateLocalInterface(params CorbaAbstractInterfaceDefCreateLocalInterfaceIn) (CorbaAbstractInterfaceDefCreateLocalInterfaceOut, error)
	// Original name: "create_module"
	CreateModule(params CorbaAbstractInterfaceDefCreateModuleIn) (CorbaAbstractInterfaceDefCreateModuleOut, error)
	// Original name: "create_native"
	CreateNative(params CorbaAbstractInterfaceDefCreateNativeIn) (CorbaAbstractInterfaceDefCreateNativeOut, error)
	// Original name: "create_operation"
	CreateOperation(params CorbaAbstractInterfaceDefCreateOperationIn) (CorbaAbstractInterfaceDefCreateOperationOut, error)
	// Original name: "create_struct"
	CreateStruct(params CorbaAbstractInterfaceDefCreateStructIn) (CorbaAbstractInterfaceDefCreateStructOut, error)
	// Original name: "create_union"
	CreateUnion(params CorbaAbstractInterfaceDefCreateUnionIn) (CorbaAbstractInterfaceDefCreateUnionOut, error)
	// Original name: "create_value"
	CreateValue(params CorbaAbstractInterfaceDefCreateValueIn) (CorbaAbstractInterfaceDefCreateValueOut, error)
	// Original name: "create_value_box"
	CreateValueBox(params CorbaAbstractInterfaceDefCreateValueBoxIn) (CorbaAbstractInterfaceDefCreateValueBoxOut, error)
	// Original name: "describe_contents"
	DescribeContents(params CorbaAbstractInterfaceDefDescribeContentsIn) (CorbaAbstractInterfaceDefDescribeContentsOut, error)
	// Original name: "describe_interface"
	DescribeInterface(params CorbaAbstractInterfaceDefDescribeInterfaceIn) (CorbaAbstractInterfaceDefDescribeInterfaceOut, error)
	// Original name: "destroy"
	Destroy(params CorbaAbstractInterfaceDefDestroyIn) (CorbaAbstractInterfaceDefDestroyOut, error)
	// Original name: "is_a"
	IsA(params CorbaAbstractInterfaceDefIsAIn) (CorbaAbstractInterfaceDefIsAOut, error)
	// Original name: "lookup"
	Lookup(params CorbaAbstractInterfaceDefLookupIn) (CorbaAbstractInterfaceDefLookupOut, error)
	// Original name: "lookup_name"
	LookupName(params CorbaAbstractInterfaceDefLookupNameIn) (CorbaAbstractInterfaceDefLookupNameOut, error)
}

const CorbaAbstractInterfaceDefContentsOperation = "contents"
const CorbaAbstractInterfaceDefCreateAbstractInterfaceOperation = "create_abstract_interface"
const CorbaAbstractInterfaceDefCreateAliasOperation = "create_alias"
const CorbaAbstractInterfaceDefCreateAttributeOperation = "create_attribute"
const CorbaAbstractInterfaceDefCreateConstantOperation = "create_constant"
const CorbaAbstractInterfaceDefCreateEnumOperation = "create_enum"
const CorbaAbstractInterfaceDefCreateExceptionOperation = "create_exception"
const CorbaAbstractInterfaceDefCreateExtValueOperation = "create_ext_value"
const CorbaAbstractInterfaceDefCreateInterfaceOperation = "create_interface"
const CorbaAbstractInterfaceDefCreateLocalInterfaceOperation = "create_local_interface"
const CorbaAbstractInterfaceDefCreateModuleOperation = "create_module"
const CorbaAbstractInterfaceDefCreateNativeOperation = "create_native"
const CorbaAbstractInterfaceDefCreateOperationOperation = "create_operation"
const CorbaAbstractInterfaceDefCreateStructOperation = "create_struct"
const CorbaAbstractInterfaceDefCreateUnionOperation = "create_union"
const CorbaAbstractInterfaceDefCreateValueOperation = "create_value"
const CorbaAbstractInterfaceDefCreateValueBoxOperation = "create_value_box"
const CorbaAbstractInterfaceDefDescribeContentsOperation = "describe_contents"
const CorbaAbstractInterfaceDefDescribeInterfaceOperation = "describe_interface"
const CorbaAbstractInterfaceDefDestroyOperation = "destroy"
const CorbaAbstractInterfaceDefIsAOperation = "is_a"
const CorbaAbstractInterfaceDefLookupOperation = "lookup"
const CorbaAbstractInterfaceDefLookupNameOperation = "lookup_name"
//noinspection GoSnakeCaseUsage
type CorbaAbstractInterfaceDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaAbstractInterfaceDefId_Const = "IDL:omg.org/CORBA/AbstractInterfaceDef:1.0"

func (self CorbaAbstractInterfaceDef_Helper) Id() string {
	return CorbaAbstractInterfaceDefId_Const
}

func (self CorbaAbstractInterfaceDef_Helper) Read(stream __goidl__.IReadAny) (CorbaAbstractInterfaceDef, error) {
	return nil, nil
}

func (self CorbaAbstractInterfaceDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaAbstractInterfaceDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaAbstractInterfaceDefHelper = CorbaAbstractInterfaceDef_Helper{}

func init() {
}
