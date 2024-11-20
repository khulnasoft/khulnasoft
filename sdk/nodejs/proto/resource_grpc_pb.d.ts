// package: khulnasoftrpc
// file: khulnasoft/resource.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as khulnasoft_resource_pb from "./resource_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";
import * as khulnasoft_provider_pb from "./provider_pb";
import * as khulnasoft_alias_pb from "./alias_pb";
import * as khulnasoft_source_pb from "./source_pb";
import * as khulnasoft_callback_pb from "./callback_pb";

interface IResourceMonitorService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    supportsFeature: IResourceMonitorService_ISupportsFeature;
    invoke: IResourceMonitorService_IInvoke;
    streamInvoke: IResourceMonitorService_IStreamInvoke;
    call: IResourceMonitorService_ICall;
    readResource: IResourceMonitorService_IReadResource;
    registerResource: IResourceMonitorService_IRegisterResource;
    registerResourceOutputs: IResourceMonitorService_IRegisterResourceOutputs;
    registerStackTransform: IResourceMonitorService_IRegisterStackTransform;
    registerStackInvokeTransform: IResourceMonitorService_IRegisterStackInvokeTransform;
    registerPackage: IResourceMonitorService_IRegisterPackage;
}

interface IResourceMonitorService_ISupportsFeature extends grpc.MethodDefinition<khulnasoft_resource_pb.SupportsFeatureRequest, khulnasoft_resource_pb.SupportsFeatureResponse> {
    path: "/khulnasoftrpc.ResourceMonitor/SupportsFeature";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_resource_pb.SupportsFeatureRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_resource_pb.SupportsFeatureRequest>;
    responseSerialize: grpc.serialize<khulnasoft_resource_pb.SupportsFeatureResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_resource_pb.SupportsFeatureResponse>;
}
interface IResourceMonitorService_IInvoke extends grpc.MethodDefinition<khulnasoft_resource_pb.ResourceInvokeRequest, khulnasoft_provider_pb.InvokeResponse> {
    path: "/khulnasoftrpc.ResourceMonitor/Invoke";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_resource_pb.ResourceInvokeRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_resource_pb.ResourceInvokeRequest>;
    responseSerialize: grpc.serialize<khulnasoft_provider_pb.InvokeResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_provider_pb.InvokeResponse>;
}
interface IResourceMonitorService_IStreamInvoke extends grpc.MethodDefinition<khulnasoft_resource_pb.ResourceInvokeRequest, khulnasoft_provider_pb.InvokeResponse> {
    path: "/khulnasoftrpc.ResourceMonitor/StreamInvoke";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<khulnasoft_resource_pb.ResourceInvokeRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_resource_pb.ResourceInvokeRequest>;
    responseSerialize: grpc.serialize<khulnasoft_provider_pb.InvokeResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_provider_pb.InvokeResponse>;
}
interface IResourceMonitorService_ICall extends grpc.MethodDefinition<khulnasoft_resource_pb.ResourceCallRequest, khulnasoft_provider_pb.CallResponse> {
    path: "/khulnasoftrpc.ResourceMonitor/Call";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_resource_pb.ResourceCallRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_resource_pb.ResourceCallRequest>;
    responseSerialize: grpc.serialize<khulnasoft_provider_pb.CallResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_provider_pb.CallResponse>;
}
interface IResourceMonitorService_IReadResource extends grpc.MethodDefinition<khulnasoft_resource_pb.ReadResourceRequest, khulnasoft_resource_pb.ReadResourceResponse> {
    path: "/khulnasoftrpc.ResourceMonitor/ReadResource";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_resource_pb.ReadResourceRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_resource_pb.ReadResourceRequest>;
    responseSerialize: grpc.serialize<khulnasoft_resource_pb.ReadResourceResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_resource_pb.ReadResourceResponse>;
}
interface IResourceMonitorService_IRegisterResource extends grpc.MethodDefinition<khulnasoft_resource_pb.RegisterResourceRequest, khulnasoft_resource_pb.RegisterResourceResponse> {
    path: "/khulnasoftrpc.ResourceMonitor/RegisterResource";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_resource_pb.RegisterResourceRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_resource_pb.RegisterResourceRequest>;
    responseSerialize: grpc.serialize<khulnasoft_resource_pb.RegisterResourceResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_resource_pb.RegisterResourceResponse>;
}
interface IResourceMonitorService_IRegisterResourceOutputs extends grpc.MethodDefinition<khulnasoft_resource_pb.RegisterResourceOutputsRequest, google_protobuf_empty_pb.Empty> {
    path: "/khulnasoftrpc.ResourceMonitor/RegisterResourceOutputs";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_resource_pb.RegisterResourceOutputsRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_resource_pb.RegisterResourceOutputsRequest>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}
interface IResourceMonitorService_IRegisterStackTransform extends grpc.MethodDefinition<khulnasoft_callback_pb.Callback, google_protobuf_empty_pb.Empty> {
    path: "/khulnasoftrpc.ResourceMonitor/RegisterStackTransform";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_callback_pb.Callback>;
    requestDeserialize: grpc.deserialize<khulnasoft_callback_pb.Callback>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}
interface IResourceMonitorService_IRegisterStackInvokeTransform extends grpc.MethodDefinition<khulnasoft_callback_pb.Callback, google_protobuf_empty_pb.Empty> {
    path: "/khulnasoftrpc.ResourceMonitor/RegisterStackInvokeTransform";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_callback_pb.Callback>;
    requestDeserialize: grpc.deserialize<khulnasoft_callback_pb.Callback>;
    responseSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    responseDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
}
interface IResourceMonitorService_IRegisterPackage extends grpc.MethodDefinition<khulnasoft_resource_pb.RegisterPackageRequest, khulnasoft_resource_pb.RegisterPackageResponse> {
    path: "/khulnasoftrpc.ResourceMonitor/RegisterPackage";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_resource_pb.RegisterPackageRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_resource_pb.RegisterPackageRequest>;
    responseSerialize: grpc.serialize<khulnasoft_resource_pb.RegisterPackageResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_resource_pb.RegisterPackageResponse>;
}

export const ResourceMonitorService: IResourceMonitorService;

export interface IResourceMonitorServer extends grpc.UntypedServiceImplementation {
    supportsFeature: grpc.handleUnaryCall<khulnasoft_resource_pb.SupportsFeatureRequest, khulnasoft_resource_pb.SupportsFeatureResponse>;
    invoke: grpc.handleUnaryCall<khulnasoft_resource_pb.ResourceInvokeRequest, khulnasoft_provider_pb.InvokeResponse>;
    streamInvoke: grpc.handleServerStreamingCall<khulnasoft_resource_pb.ResourceInvokeRequest, khulnasoft_provider_pb.InvokeResponse>;
    call: grpc.handleUnaryCall<khulnasoft_resource_pb.ResourceCallRequest, khulnasoft_provider_pb.CallResponse>;
    readResource: grpc.handleUnaryCall<khulnasoft_resource_pb.ReadResourceRequest, khulnasoft_resource_pb.ReadResourceResponse>;
    registerResource: grpc.handleUnaryCall<khulnasoft_resource_pb.RegisterResourceRequest, khulnasoft_resource_pb.RegisterResourceResponse>;
    registerResourceOutputs: grpc.handleUnaryCall<khulnasoft_resource_pb.RegisterResourceOutputsRequest, google_protobuf_empty_pb.Empty>;
    registerStackTransform: grpc.handleUnaryCall<khulnasoft_callback_pb.Callback, google_protobuf_empty_pb.Empty>;
    registerStackInvokeTransform: grpc.handleUnaryCall<khulnasoft_callback_pb.Callback, google_protobuf_empty_pb.Empty>;
    registerPackage: grpc.handleUnaryCall<khulnasoft_resource_pb.RegisterPackageRequest, khulnasoft_resource_pb.RegisterPackageResponse>;
}

export interface IResourceMonitorClient {
    supportsFeature(request: khulnasoft_resource_pb.SupportsFeatureRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.SupportsFeatureResponse) => void): grpc.ClientUnaryCall;
    supportsFeature(request: khulnasoft_resource_pb.SupportsFeatureRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.SupportsFeatureResponse) => void): grpc.ClientUnaryCall;
    supportsFeature(request: khulnasoft_resource_pb.SupportsFeatureRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.SupportsFeatureResponse) => void): grpc.ClientUnaryCall;
    invoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.InvokeResponse) => void): grpc.ClientUnaryCall;
    invoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.InvokeResponse) => void): grpc.ClientUnaryCall;
    invoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.InvokeResponse) => void): grpc.ClientUnaryCall;
    streamInvoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_provider_pb.InvokeResponse>;
    streamInvoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_provider_pb.InvokeResponse>;
    call(request: khulnasoft_resource_pb.ResourceCallRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.CallResponse) => void): grpc.ClientUnaryCall;
    call(request: khulnasoft_resource_pb.ResourceCallRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.CallResponse) => void): grpc.ClientUnaryCall;
    call(request: khulnasoft_resource_pb.ResourceCallRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.CallResponse) => void): grpc.ClientUnaryCall;
    readResource(request: khulnasoft_resource_pb.ReadResourceRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.ReadResourceResponse) => void): grpc.ClientUnaryCall;
    readResource(request: khulnasoft_resource_pb.ReadResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.ReadResourceResponse) => void): grpc.ClientUnaryCall;
    readResource(request: khulnasoft_resource_pb.ReadResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.ReadResourceResponse) => void): grpc.ClientUnaryCall;
    registerResource(request: khulnasoft_resource_pb.RegisterResourceRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterResourceResponse) => void): grpc.ClientUnaryCall;
    registerResource(request: khulnasoft_resource_pb.RegisterResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterResourceResponse) => void): grpc.ClientUnaryCall;
    registerResource(request: khulnasoft_resource_pb.RegisterResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterResourceResponse) => void): grpc.ClientUnaryCall;
    registerResourceOutputs(request: khulnasoft_resource_pb.RegisterResourceOutputsRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerResourceOutputs(request: khulnasoft_resource_pb.RegisterResourceOutputsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerResourceOutputs(request: khulnasoft_resource_pb.RegisterResourceOutputsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerStackTransform(request: khulnasoft_callback_pb.Callback, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerStackTransform(request: khulnasoft_callback_pb.Callback, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerStackTransform(request: khulnasoft_callback_pb.Callback, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerStackInvokeTransform(request: khulnasoft_callback_pb.Callback, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerStackInvokeTransform(request: khulnasoft_callback_pb.Callback, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerStackInvokeTransform(request: khulnasoft_callback_pb.Callback, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    registerPackage(request: khulnasoft_resource_pb.RegisterPackageRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterPackageResponse) => void): grpc.ClientUnaryCall;
    registerPackage(request: khulnasoft_resource_pb.RegisterPackageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterPackageResponse) => void): grpc.ClientUnaryCall;
    registerPackage(request: khulnasoft_resource_pb.RegisterPackageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterPackageResponse) => void): grpc.ClientUnaryCall;
}

export class ResourceMonitorClient extends grpc.Client implements IResourceMonitorClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public supportsFeature(request: khulnasoft_resource_pb.SupportsFeatureRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.SupportsFeatureResponse) => void): grpc.ClientUnaryCall;
    public supportsFeature(request: khulnasoft_resource_pb.SupportsFeatureRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.SupportsFeatureResponse) => void): grpc.ClientUnaryCall;
    public supportsFeature(request: khulnasoft_resource_pb.SupportsFeatureRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.SupportsFeatureResponse) => void): grpc.ClientUnaryCall;
    public invoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.InvokeResponse) => void): grpc.ClientUnaryCall;
    public invoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.InvokeResponse) => void): grpc.ClientUnaryCall;
    public invoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.InvokeResponse) => void): grpc.ClientUnaryCall;
    public streamInvoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_provider_pb.InvokeResponse>;
    public streamInvoke(request: khulnasoft_resource_pb.ResourceInvokeRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_provider_pb.InvokeResponse>;
    public call(request: khulnasoft_resource_pb.ResourceCallRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.CallResponse) => void): grpc.ClientUnaryCall;
    public call(request: khulnasoft_resource_pb.ResourceCallRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.CallResponse) => void): grpc.ClientUnaryCall;
    public call(request: khulnasoft_resource_pb.ResourceCallRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_provider_pb.CallResponse) => void): grpc.ClientUnaryCall;
    public readResource(request: khulnasoft_resource_pb.ReadResourceRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.ReadResourceResponse) => void): grpc.ClientUnaryCall;
    public readResource(request: khulnasoft_resource_pb.ReadResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.ReadResourceResponse) => void): grpc.ClientUnaryCall;
    public readResource(request: khulnasoft_resource_pb.ReadResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.ReadResourceResponse) => void): grpc.ClientUnaryCall;
    public registerResource(request: khulnasoft_resource_pb.RegisterResourceRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterResourceResponse) => void): grpc.ClientUnaryCall;
    public registerResource(request: khulnasoft_resource_pb.RegisterResourceRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterResourceResponse) => void): grpc.ClientUnaryCall;
    public registerResource(request: khulnasoft_resource_pb.RegisterResourceRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterResourceResponse) => void): grpc.ClientUnaryCall;
    public registerResourceOutputs(request: khulnasoft_resource_pb.RegisterResourceOutputsRequest, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerResourceOutputs(request: khulnasoft_resource_pb.RegisterResourceOutputsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerResourceOutputs(request: khulnasoft_resource_pb.RegisterResourceOutputsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerStackTransform(request: khulnasoft_callback_pb.Callback, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerStackTransform(request: khulnasoft_callback_pb.Callback, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerStackTransform(request: khulnasoft_callback_pb.Callback, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerStackInvokeTransform(request: khulnasoft_callback_pb.Callback, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerStackInvokeTransform(request: khulnasoft_callback_pb.Callback, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerStackInvokeTransform(request: khulnasoft_callback_pb.Callback, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: google_protobuf_empty_pb.Empty) => void): grpc.ClientUnaryCall;
    public registerPackage(request: khulnasoft_resource_pb.RegisterPackageRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterPackageResponse) => void): grpc.ClientUnaryCall;
    public registerPackage(request: khulnasoft_resource_pb.RegisterPackageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterPackageResponse) => void): grpc.ClientUnaryCall;
    public registerPackage(request: khulnasoft_resource_pb.RegisterPackageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_resource_pb.RegisterPackageResponse) => void): grpc.ClientUnaryCall;
}
