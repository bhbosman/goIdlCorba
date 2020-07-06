// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::InterfaceDef", generatedBy by: "WriteInterface"
type CorbaInterfaceDef interface {
	// Original name: "contents"
	Contents(params CorbaInterfaceDefContentsIn) (CorbaInterfaceDefContentsOut, error)
	// Original name: "create_abstract_interface"
	CreateAbstractInterface(params CorbaInterfaceDefCreateAbstractInterfaceIn) (CorbaInterfaceDefCreateAbstractInterfaceOut, error)
	// Original name: "create_alias"
	CreateAlias(params CorbaInterfaceDefCreateAliasIn) (CorbaInterfaceDefCreateAliasOut, error)
	// Original name: "create_attribute"
	CreateAttribute(params CorbaInterfaceDefCreateAttributeIn) (CorbaInterfaceDefCreateAttributeOut, error)
	// Original name: "create_constant"
	CreateConstant(params CorbaInterfaceDefCreateConstantIn) (CorbaInterfaceDefCreateConstantOut, error)
	// Original name: "create_enum"
	CreateEnum(params CorbaInterfaceDefCreateEnumIn) (CorbaInterfaceDefCreateEnumOut, error)
	// Original name: "create_exception"
	CreateException(params CorbaInterfaceDefCreateExceptionIn) (CorbaInterfaceDefCreateExceptionOut, error)
	// Original name: "create_ext_value"
	CreateExtValue(params CorbaInterfaceDefCreateExtValueIn) (CorbaInterfaceDefCreateExtValueOut, error)
	// Original name: "create_interface"
	CreateInterface(params CorbaInterfaceDefCreateInterfaceIn) (CorbaInterfaceDefCreateInterfaceOut, error)
	// Original name: "create_local_interface"
	CreateLocalInterface(params CorbaInterfaceDefCreateLocalInterfaceIn) (CorbaInterfaceDefCreateLocalInterfaceOut, error)
	// Original name: "create_module"
	CreateModule(params CorbaInterfaceDefCreateModuleIn) (CorbaInterfaceDefCreateModuleOut, error)
	// Original name: "create_native"
	CreateNative(params CorbaInterfaceDefCreateNativeIn) (CorbaInterfaceDefCreateNativeOut, error)
	// Original name: "create_operation"
	CreateOperation(params CorbaInterfaceDefCreateOperationIn) (CorbaInterfaceDefCreateOperationOut, error)
	// Original name: "create_struct"
	CreateStruct(params CorbaInterfaceDefCreateStructIn) (CorbaInterfaceDefCreateStructOut, error)
	// Original name: "create_union"
	CreateUnion(params CorbaInterfaceDefCreateUnionIn) (CorbaInterfaceDefCreateUnionOut, error)
	// Original name: "create_value"
	CreateValue(params CorbaInterfaceDefCreateValueIn) (CorbaInterfaceDefCreateValueOut, error)
	// Original name: "create_value_box"
	CreateValueBox(params CorbaInterfaceDefCreateValueBoxIn) (CorbaInterfaceDefCreateValueBoxOut, error)
	// Original name: "describe_contents"
	DescribeContents(params CorbaInterfaceDefDescribeContentsIn) (CorbaInterfaceDefDescribeContentsOut, error)
	// Original name: "describe_interface"
	DescribeInterface(params CorbaInterfaceDefDescribeInterfaceIn) (CorbaInterfaceDefDescribeInterfaceOut, error)
	// Original name: "destroy"
	Destroy(params CorbaInterfaceDefDestroyIn) (CorbaInterfaceDefDestroyOut, error)
	// Original name: "is_a"
	IsA(params CorbaInterfaceDefIsAIn) (CorbaInterfaceDefIsAOut, error)
	// Original name: "lookup"
	Lookup(params CorbaInterfaceDefLookupIn) (CorbaInterfaceDefLookupOut, error)
	// Original name: "lookup_name"
	LookupName(params CorbaInterfaceDefLookupNameIn) (CorbaInterfaceDefLookupNameOut, error)
	// Property BaseInterfaces
	// Get Property BaseInterfaces
	GetBaseInterfaces() (CorbaInterfaceDefSeq, error)
	// Set Property BaseInterfaces
	SetBaseInterfaces(value CorbaInterfaceDefSeq) error
	// Property IsAbstract
	// Get Property IsAbstract
	GetIsAbstract() (bool, error)
	// Set Property IsAbstract
	SetIsAbstract(value bool) error
}

const CorbaInterfaceDefContentsOperation = "contents"
const CorbaInterfaceDefCreateAbstractInterfaceOperation = "create_abstract_interface"
const CorbaInterfaceDefCreateAliasOperation = "create_alias"
const CorbaInterfaceDefCreateAttributeOperation = "create_attribute"
const CorbaInterfaceDefCreateConstantOperation = "create_constant"
const CorbaInterfaceDefCreateEnumOperation = "create_enum"
const CorbaInterfaceDefCreateExceptionOperation = "create_exception"
const CorbaInterfaceDefCreateExtValueOperation = "create_ext_value"
const CorbaInterfaceDefCreateInterfaceOperation = "create_interface"
const CorbaInterfaceDefCreateLocalInterfaceOperation = "create_local_interface"
const CorbaInterfaceDefCreateModuleOperation = "create_module"
const CorbaInterfaceDefCreateNativeOperation = "create_native"
const CorbaInterfaceDefCreateOperationOperation = "create_operation"
const CorbaInterfaceDefCreateStructOperation = "create_struct"
const CorbaInterfaceDefCreateUnionOperation = "create_union"
const CorbaInterfaceDefCreateValueOperation = "create_value"
const CorbaInterfaceDefCreateValueBoxOperation = "create_value_box"
const CorbaInterfaceDefDescribeContentsOperation = "describe_contents"
const CorbaInterfaceDefDescribeInterfaceOperation = "describe_interface"
const CorbaInterfaceDefDestroyOperation = "destroy"
const CorbaInterfaceDefIsAOperation = "is_a"
const CorbaInterfaceDefLookupOperation = "lookup"
const CorbaInterfaceDefLookupNameOperation = "lookup_name"
//noinspection GoSnakeCaseUsage
type CorbaInterfaceDef_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaInterfaceDefId_Const = "IDL:omg.org/CORBA/InterfaceDef:1.0"

func (self CorbaInterfaceDef_Helper) Id() string {
	return CorbaInterfaceDefId_Const
}

func (self CorbaInterfaceDef_Helper) Read(stream __goidl__.IReadAny) (CorbaInterfaceDef, error) {
	return nil, nil
}

func (self CorbaInterfaceDef_Helper) Write(stream __goidl__.IWriteAny, v CorbaInterfaceDef) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaInterfaceDefHelper = CorbaInterfaceDef_Helper{}

func init() {
}
