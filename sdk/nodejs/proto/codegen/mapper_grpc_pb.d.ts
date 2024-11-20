// package: codegen
// file: khulnasoft/codegen/mapper.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as khulnasoft_codegen_mapper_pb from "../codegen/mapper_pb";

interface IMapperService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getMapping: IMapperService_IGetMapping;
}

interface IMapperService_IGetMapping extends grpc.MethodDefinition<khulnasoft_codegen_mapper_pb.GetMappingRequest, khulnasoft_codegen_mapper_pb.GetMappingResponse> {
    path: "/codegen.Mapper/GetMapping";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_codegen_mapper_pb.GetMappingRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_codegen_mapper_pb.GetMappingRequest>;
    responseSerialize: grpc.serialize<khulnasoft_codegen_mapper_pb.GetMappingResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_codegen_mapper_pb.GetMappingResponse>;
}

export const MapperService: IMapperService;

export interface IMapperServer extends grpc.UntypedServiceImplementation {
    getMapping: grpc.handleUnaryCall<khulnasoft_codegen_mapper_pb.GetMappingRequest, khulnasoft_codegen_mapper_pb.GetMappingResponse>;
}

export interface IMapperClient {
    getMapping(request: khulnasoft_codegen_mapper_pb.GetMappingRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_mapper_pb.GetMappingResponse) => void): grpc.ClientUnaryCall;
    getMapping(request: khulnasoft_codegen_mapper_pb.GetMappingRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_mapper_pb.GetMappingResponse) => void): grpc.ClientUnaryCall;
    getMapping(request: khulnasoft_codegen_mapper_pb.GetMappingRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_mapper_pb.GetMappingResponse) => void): grpc.ClientUnaryCall;
}

export class MapperClient extends grpc.Client implements IMapperClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getMapping(request: khulnasoft_codegen_mapper_pb.GetMappingRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_mapper_pb.GetMappingResponse) => void): grpc.ClientUnaryCall;
    public getMapping(request: khulnasoft_codegen_mapper_pb.GetMappingRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_mapper_pb.GetMappingResponse) => void): grpc.ClientUnaryCall;
    public getMapping(request: khulnasoft_codegen_mapper_pb.GetMappingRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_codegen_mapper_pb.GetMappingResponse) => void): grpc.ClientUnaryCall;
}
