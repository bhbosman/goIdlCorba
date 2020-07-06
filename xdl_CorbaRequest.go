// Code generated by me. DO NOT EDIT.

package goIdlCorba

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CORBA::Request", generatedBy by: "WriteInterface"
type CorbaRequest interface {
	// Original name: "add_arg"
	AddArg(params CorbaRequestAddArgIn) (CorbaRequestAddArgOut, error)
	// Original name: "delete"
	Delete(params CorbaRequestDeleteIn) (CorbaRequestDeleteOut, error)
	//Exceptions for : GetResponse
	//	CorbaWrongTransaction
	// Original name: "get_response"
	GetResponse(params CorbaRequestGetResponseIn) (CorbaRequestGetResponseOut, error)
	// Original name: "invoke"
	Invoke(params CorbaRequestInvokeIn) (CorbaRequestInvokeOut, error)
	// Original name: "poll_response"
	PollResponse(params CorbaRequestPollResponseIn) (CorbaRequestPollResponseOut, error)
	// Original name: "prepare"
	Prepare(params CorbaRequestPrepareIn) (CorbaRequestPrepareOut, error)
	// Original name: "send"
	Send(params CorbaRequestSendIn) (CorbaRequestSendOut, error)
	// Original name: "sendc"
	Sendc(params CorbaRequestSendcIn) (CorbaRequestSendcOut, error)
	// Original name: "sendp"
	Sendp(params CorbaRequestSendpIn) (CorbaRequestSendpOut, error)
}

const CorbaRequestAddArgOperation = "add_arg"
const CorbaRequestDeleteOperation = "delete"
const CorbaRequestGetResponseOperation = "get_response"
const CorbaRequestInvokeOperation = "invoke"
const CorbaRequestPollResponseOperation = "poll_response"
const CorbaRequestPrepareOperation = "prepare"
const CorbaRequestSendOperation = "send"
const CorbaRequestSendcOperation = "sendc"
const CorbaRequestSendpOperation = "sendp"
//noinspection GoSnakeCaseUsage
type CorbaRequest_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CorbaRequestId_Const = "IDL:omg.org/CORBA/Request:1.0"

func (self CorbaRequest_Helper) Id() string {
	return CorbaRequestId_Const
}

func (self CorbaRequest_Helper) Read(stream __goidl__.IReadAny) (CorbaRequest, error) {
	return nil, nil
}

func (self CorbaRequest_Helper) Write(stream __goidl__.IWriteAny, v CorbaRequest) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CorbaRequestHelper = CorbaRequest_Helper{}

func init() {
}