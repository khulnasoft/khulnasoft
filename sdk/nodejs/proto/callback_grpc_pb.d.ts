// package: khulnasoftrpc
// file: khulnasoft/callback.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as khulnasoft_callback_pb from "./callback_pb";

interface ICallbacksService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    invoke: ICallbacksService_IInvoke;
}

interface ICallbacksService_IInvoke extends grpc.MethodDefinition<khulnasoft_callback_pb.CallbackInvokeRequest, khulnasoft_callback_pb.CallbackInvokeResponse> {
    path: "/khulnasoftrpc.Callbacks/Invoke";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_callback_pb.CallbackInvokeRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_callback_pb.CallbackInvokeRequest>;
    responseSerialize: grpc.serialize<khulnasoft_callback_pb.CallbackInvokeResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_callback_pb.CallbackInvokeResponse>;
}

export const CallbacksService: ICallbacksService;

export interface ICallbacksServer extends grpc.UntypedServiceImplementation {
    invoke: grpc.handleUnaryCall<khulnasoft_callback_pb.CallbackInvokeRequest, khulnasoft_callback_pb.CallbackInvokeResponse>;
}

export interface ICallbacksClient {
    invoke(request: khulnasoft_callback_pb.CallbackInvokeRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
    invoke(request: khulnasoft_callback_pb.CallbackInvokeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
    invoke(request: khulnasoft_callback_pb.CallbackInvokeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
}

export class CallbacksClient extends grpc.Client implements ICallbacksClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public invoke(request: khulnasoft_callback_pb.CallbackInvokeRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
    public invoke(request: khulnasoft_callback_pb.CallbackInvokeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
    public invoke(request: khulnasoft_callback_pb.CallbackInvokeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_callback_pb.CallbackInvokeResponse) => void): grpc.ClientUnaryCall;
}
