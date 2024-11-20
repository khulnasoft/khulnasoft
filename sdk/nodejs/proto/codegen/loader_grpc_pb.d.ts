// package: codegen
// file: khulnasoft/codegen/loader.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as khulnasoft_codegen_loader_pb from "../codegen/loader_pb";

interface ILoaderService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getSchema: ILoaderService_IGetSchema;
}

interface ILoaderService_IGetSchema extends grpc.MethodDefinition<khulnasoft_codegen_loader_pb.GetSchemaRequest, khulnasoft_codegen_loader_pb.GetSchemaResponse> {
    path: "/codegen.Loader/GetSchema";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_codegen_loader_pb.GetSchemaRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_codegen_loader_pb.GetSchemaRequest>;
    responseSerialize: grpc.serialize<khulnasoft_codegen_loader_pb.GetSchemaResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_codegen_loader_pb.GetSchemaResponse>;
}

export const LoaderService: ILoaderService;

export interface ILoaderServer extends grpc.UntypedServiceImplementation {
    getSchema: grpc.handleUnaryCall<khulnasoft_codegen_loader_pb.GetSchemaRequest, khulnasoft_codegen_loader_pb.GetSchemaResponse>;
}

export interface ILoaderClient {
    getSchema(request: khulnasoft_codegen_loader_pb.GetSchemaRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
    getSchema(request: khulnasoft_codegen_loader_pb.GetSchemaRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
    getSchema(request: khulnasoft_codegen_loader_pb.GetSchemaRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
}

export class LoaderClient extends grpc.Client implements ILoaderClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getSchema(request: khulnasoft_codegen_loader_pb.GetSchemaRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
    public getSchema(request: khulnasoft_codegen_loader_pb.GetSchemaRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
    public getSchema(request: khulnasoft_codegen_loader_pb.GetSchemaRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_loader_pb.GetSchemaResponse) => void): grpc.ClientUnaryCall;
}
