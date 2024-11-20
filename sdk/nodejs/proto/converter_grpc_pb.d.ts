// package: khulnasoftrpc
// file: khulnasoft/converter.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as khulnasoft_converter_pb from "./converter_pb";
import * as khulnasoft_codegen_hcl_pb from "./codegen/hcl_pb";

interface IConverterService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    convertState: IConverterService_IConvertState;
    convertProgram: IConverterService_IConvertProgram;
}

interface IConverterService_IConvertState extends grpc.MethodDefinition<khulnasoft_converter_pb.ConvertStateRequest, khulnasoft_converter_pb.ConvertStateResponse> {
    path: "/khulnasoftrpc.Converter/ConvertState";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_converter_pb.ConvertStateRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_converter_pb.ConvertStateRequest>;
    responseSerialize: grpc.serialize<khulnasoft_converter_pb.ConvertStateResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_converter_pb.ConvertStateResponse>;
}
interface IConverterService_IConvertProgram extends grpc.MethodDefinition<khulnasoft_converter_pb.ConvertProgramRequest, khulnasoft_converter_pb.ConvertProgramResponse> {
    path: "/khulnasoftrpc.Converter/ConvertProgram";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_converter_pb.ConvertProgramRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_converter_pb.ConvertProgramRequest>;
    responseSerialize: grpc.serialize<khulnasoft_converter_pb.ConvertProgramResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_converter_pb.ConvertProgramResponse>;
}

export const ConverterService: IConverterService;

export interface IConverterServer extends grpc.UntypedServiceImplementation {
    convertState: grpc.handleUnaryCall<khulnasoft_converter_pb.ConvertStateRequest, khulnasoft_converter_pb.ConvertStateResponse>;
    convertProgram: grpc.handleUnaryCall<khulnasoft_converter_pb.ConvertProgramRequest, khulnasoft_converter_pb.ConvertProgramResponse>;
}

export interface IConverterClient {
    convertState(request: khulnasoft_converter_pb.ConvertStateRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertStateResponse) => void): grpc.ClientUnaryCall;
    convertState(request: khulnasoft_converter_pb.ConvertStateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertStateResponse) => void): grpc.ClientUnaryCall;
    convertState(request: khulnasoft_converter_pb.ConvertStateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertStateResponse) => void): grpc.ClientUnaryCall;
    convertProgram(request: khulnasoft_converter_pb.ConvertProgramRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertProgramResponse) => void): grpc.ClientUnaryCall;
    convertProgram(request: khulnasoft_converter_pb.ConvertProgramRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertProgramResponse) => void): grpc.ClientUnaryCall;
    convertProgram(request: khulnasoft_converter_pb.ConvertProgramRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertProgramResponse) => void): grpc.ClientUnaryCall;
}

export class ConverterClient extends grpc.Client implements IConverterClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public convertState(request: khulnasoft_converter_pb.ConvertStateRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertStateResponse) => void): grpc.ClientUnaryCall;
    public convertState(request: khulnasoft_converter_pb.ConvertStateRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertStateResponse) => void): grpc.ClientUnaryCall;
    public convertState(request: khulnasoft_converter_pb.ConvertStateRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertStateResponse) => void): grpc.ClientUnaryCall;
    public convertProgram(request: khulnasoft_converter_pb.ConvertProgramRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertProgramResponse) => void): grpc.ClientUnaryCall;
    public convertProgram(request: khulnasoft_converter_pb.ConvertProgramRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertProgramResponse) => void): grpc.ClientUnaryCall;
    public convertProgram(request: khulnasoft_converter_pb.ConvertProgramRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_converter_pb.ConvertProgramResponse) => void): grpc.ClientUnaryCall;
}
