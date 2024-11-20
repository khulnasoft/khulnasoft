// package: khulnasoftrpc
// file: khulnasoft/language.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as khulnasoft_language_pb from "./language_pb";
import * as khulnasoft_codegen_hcl_pb from "./codegen/hcl_pb";
import * as khulnasoft_plugin_pb from "./plugin_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_struct_pb from "google-protobuf/google/protobuf/struct_pb";

interface ILanguageRuntimeService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getRequiredPlugins: ILanguageRuntimeService_IGetRequiredPlugins;
    run: ILanguageRuntimeService_IRun;
    getPluginInfo: ILanguageRuntimeService_IGetPluginInfo;
    installDependencies: ILanguageRuntimeService_IInstallDependencies;
    runtimeOptionsPrompts: ILanguageRuntimeService_IRuntimeOptionsPrompts;
    about: ILanguageRuntimeService_IAbout;
    getProgramDependencies: ILanguageRuntimeService_IGetProgramDependencies;
    runPlugin: ILanguageRuntimeService_IRunPlugin;
    generateProgram: ILanguageRuntimeService_IGenerateProgram;
    generateProject: ILanguageRuntimeService_IGenerateProject;
    generatePackage: ILanguageRuntimeService_IGeneratePackage;
    pack: ILanguageRuntimeService_IPack;
}

interface ILanguageRuntimeService_IGetRequiredPlugins extends grpc.MethodDefinition<khulnasoft_language_pb.GetRequiredPluginsRequest, khulnasoft_language_pb.GetRequiredPluginsResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/GetRequiredPlugins";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.GetRequiredPluginsRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.GetRequiredPluginsRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.GetRequiredPluginsResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.GetRequiredPluginsResponse>;
}
interface ILanguageRuntimeService_IRun extends grpc.MethodDefinition<khulnasoft_language_pb.RunRequest, khulnasoft_language_pb.RunResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/Run";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.RunRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.RunRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.RunResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.RunResponse>;
}
interface ILanguageRuntimeService_IGetPluginInfo extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, khulnasoft_plugin_pb.PluginInfo> {
    path: "/khulnasoftrpc.LanguageRuntime/GetPluginInfo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<khulnasoft_plugin_pb.PluginInfo>;
    responseDeserialize: grpc.deserialize<khulnasoft_plugin_pb.PluginInfo>;
}
interface ILanguageRuntimeService_IInstallDependencies extends grpc.MethodDefinition<khulnasoft_language_pb.InstallDependenciesRequest, khulnasoft_language_pb.InstallDependenciesResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/InstallDependencies";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.InstallDependenciesRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.InstallDependenciesRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.InstallDependenciesResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.InstallDependenciesResponse>;
}
interface ILanguageRuntimeService_IRuntimeOptionsPrompts extends grpc.MethodDefinition<khulnasoft_language_pb.RuntimeOptionsRequest, khulnasoft_language_pb.RuntimeOptionsResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/RuntimeOptionsPrompts";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.RuntimeOptionsRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.RuntimeOptionsRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.RuntimeOptionsResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.RuntimeOptionsResponse>;
}
interface ILanguageRuntimeService_IAbout extends grpc.MethodDefinition<khulnasoft_language_pb.AboutRequest, khulnasoft_language_pb.AboutResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/About";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.AboutRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.AboutRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.AboutResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.AboutResponse>;
}
interface ILanguageRuntimeService_IGetProgramDependencies extends grpc.MethodDefinition<khulnasoft_language_pb.GetProgramDependenciesRequest, khulnasoft_language_pb.GetProgramDependenciesResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/GetProgramDependencies";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.GetProgramDependenciesRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.GetProgramDependenciesRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.GetProgramDependenciesResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.GetProgramDependenciesResponse>;
}
interface ILanguageRuntimeService_IRunPlugin extends grpc.MethodDefinition<khulnasoft_language_pb.RunPluginRequest, khulnasoft_language_pb.RunPluginResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/RunPlugin";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.RunPluginRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.RunPluginRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.RunPluginResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.RunPluginResponse>;
}
interface ILanguageRuntimeService_IGenerateProgram extends grpc.MethodDefinition<khulnasoft_language_pb.GenerateProgramRequest, khulnasoft_language_pb.GenerateProgramResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/GenerateProgram";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.GenerateProgramRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.GenerateProgramRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.GenerateProgramResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.GenerateProgramResponse>;
}
interface ILanguageRuntimeService_IGenerateProject extends grpc.MethodDefinition<khulnasoft_language_pb.GenerateProjectRequest, khulnasoft_language_pb.GenerateProjectResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/GenerateProject";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.GenerateProjectRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.GenerateProjectRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.GenerateProjectResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.GenerateProjectResponse>;
}
interface ILanguageRuntimeService_IGeneratePackage extends grpc.MethodDefinition<khulnasoft_language_pb.GeneratePackageRequest, khulnasoft_language_pb.GeneratePackageResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/GeneratePackage";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.GeneratePackageRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.GeneratePackageRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.GeneratePackageResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.GeneratePackageResponse>;
}
interface ILanguageRuntimeService_IPack extends grpc.MethodDefinition<khulnasoft_language_pb.PackRequest, khulnasoft_language_pb.PackResponse> {
    path: "/khulnasoftrpc.LanguageRuntime/Pack";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<khulnasoft_language_pb.PackRequest>;
    requestDeserialize: grpc.deserialize<khulnasoft_language_pb.PackRequest>;
    responseSerialize: grpc.serialize<khulnasoft_language_pb.PackResponse>;
    responseDeserialize: grpc.deserialize<khulnasoft_language_pb.PackResponse>;
}

export const LanguageRuntimeService: ILanguageRuntimeService;

export interface ILanguageRuntimeServer extends grpc.UntypedServiceImplementation {
    getRequiredPlugins: grpc.handleUnaryCall<khulnasoft_language_pb.GetRequiredPluginsRequest, khulnasoft_language_pb.GetRequiredPluginsResponse>;
    run: grpc.handleUnaryCall<khulnasoft_language_pb.RunRequest, khulnasoft_language_pb.RunResponse>;
    getPluginInfo: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, khulnasoft_plugin_pb.PluginInfo>;
    installDependencies: grpc.handleServerStreamingCall<khulnasoft_language_pb.InstallDependenciesRequest, khulnasoft_language_pb.InstallDependenciesResponse>;
    runtimeOptionsPrompts: grpc.handleUnaryCall<khulnasoft_language_pb.RuntimeOptionsRequest, khulnasoft_language_pb.RuntimeOptionsResponse>;
    about: grpc.handleUnaryCall<khulnasoft_language_pb.AboutRequest, khulnasoft_language_pb.AboutResponse>;
    getProgramDependencies: grpc.handleUnaryCall<khulnasoft_language_pb.GetProgramDependenciesRequest, khulnasoft_language_pb.GetProgramDependenciesResponse>;
    runPlugin: grpc.handleServerStreamingCall<khulnasoft_language_pb.RunPluginRequest, khulnasoft_language_pb.RunPluginResponse>;
    generateProgram: grpc.handleUnaryCall<khulnasoft_language_pb.GenerateProgramRequest, khulnasoft_language_pb.GenerateProgramResponse>;
    generateProject: grpc.handleUnaryCall<khulnasoft_language_pb.GenerateProjectRequest, khulnasoft_language_pb.GenerateProjectResponse>;
    generatePackage: grpc.handleUnaryCall<khulnasoft_language_pb.GeneratePackageRequest, khulnasoft_language_pb.GeneratePackageResponse>;
    pack: grpc.handleUnaryCall<khulnasoft_language_pb.PackRequest, khulnasoft_language_pb.PackResponse>;
}

export interface ILanguageRuntimeClient {
    getRequiredPlugins(request: khulnasoft_language_pb.GetRequiredPluginsRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    getRequiredPlugins(request: khulnasoft_language_pb.GetRequiredPluginsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    getRequiredPlugins(request: khulnasoft_language_pb.GetRequiredPluginsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    run(request: khulnasoft_language_pb.RunRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    run(request: khulnasoft_language_pb.RunRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    run(request: khulnasoft_language_pb.RunRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: khulnasoft_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    installDependencies(request: khulnasoft_language_pb.InstallDependenciesRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_language_pb.InstallDependenciesResponse>;
    installDependencies(request: khulnasoft_language_pb.InstallDependenciesRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_language_pb.InstallDependenciesResponse>;
    runtimeOptionsPrompts(request: khulnasoft_language_pb.RuntimeOptionsRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    runtimeOptionsPrompts(request: khulnasoft_language_pb.RuntimeOptionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    runtimeOptionsPrompts(request: khulnasoft_language_pb.RuntimeOptionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    about(request: khulnasoft_language_pb.AboutRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    about(request: khulnasoft_language_pb.AboutRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    about(request: khulnasoft_language_pb.AboutRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    getProgramDependencies(request: khulnasoft_language_pb.GetProgramDependenciesRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    getProgramDependencies(request: khulnasoft_language_pb.GetProgramDependenciesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    getProgramDependencies(request: khulnasoft_language_pb.GetProgramDependenciesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    runPlugin(request: khulnasoft_language_pb.RunPluginRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_language_pb.RunPluginResponse>;
    runPlugin(request: khulnasoft_language_pb.RunPluginRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_language_pb.RunPluginResponse>;
    generateProgram(request: khulnasoft_language_pb.GenerateProgramRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    generateProgram(request: khulnasoft_language_pb.GenerateProgramRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    generateProgram(request: khulnasoft_language_pb.GenerateProgramRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    generateProject(request: khulnasoft_language_pb.GenerateProjectRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    generateProject(request: khulnasoft_language_pb.GenerateProjectRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    generateProject(request: khulnasoft_language_pb.GenerateProjectRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    generatePackage(request: khulnasoft_language_pb.GeneratePackageRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    generatePackage(request: khulnasoft_language_pb.GeneratePackageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    generatePackage(request: khulnasoft_language_pb.GeneratePackageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    pack(request: khulnasoft_language_pb.PackRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
    pack(request: khulnasoft_language_pb.PackRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
    pack(request: khulnasoft_language_pb.PackRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
}

export class LanguageRuntimeClient extends grpc.Client implements ILanguageRuntimeClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getRequiredPlugins(request: khulnasoft_language_pb.GetRequiredPluginsRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    public getRequiredPlugins(request: khulnasoft_language_pb.GetRequiredPluginsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    public getRequiredPlugins(request: khulnasoft_language_pb.GetRequiredPluginsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetRequiredPluginsResponse) => void): grpc.ClientUnaryCall;
    public run(request: khulnasoft_language_pb.RunRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    public run(request: khulnasoft_language_pb.RunRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    public run(request: khulnasoft_language_pb.RunRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RunResponse) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: khulnasoft_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public getPluginInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_plugin_pb.PluginInfo) => void): grpc.ClientUnaryCall;
    public installDependencies(request: khulnasoft_language_pb.InstallDependenciesRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_language_pb.InstallDependenciesResponse>;
    public installDependencies(request: khulnasoft_language_pb.InstallDependenciesRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_language_pb.InstallDependenciesResponse>;
    public runtimeOptionsPrompts(request: khulnasoft_language_pb.RuntimeOptionsRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    public runtimeOptionsPrompts(request: khulnasoft_language_pb.RuntimeOptionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    public runtimeOptionsPrompts(request: khulnasoft_language_pb.RuntimeOptionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.RuntimeOptionsResponse) => void): grpc.ClientUnaryCall;
    public about(request: khulnasoft_language_pb.AboutRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    public about(request: khulnasoft_language_pb.AboutRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    public about(request: khulnasoft_language_pb.AboutRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.AboutResponse) => void): grpc.ClientUnaryCall;
    public getProgramDependencies(request: khulnasoft_language_pb.GetProgramDependenciesRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    public getProgramDependencies(request: khulnasoft_language_pb.GetProgramDependenciesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    public getProgramDependencies(request: khulnasoft_language_pb.GetProgramDependenciesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GetProgramDependenciesResponse) => void): grpc.ClientUnaryCall;
    public runPlugin(request: khulnasoft_language_pb.RunPluginRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_language_pb.RunPluginResponse>;
    public runPlugin(request: khulnasoft_language_pb.RunPluginRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<khulnasoft_language_pb.RunPluginResponse>;
    public generateProgram(request: khulnasoft_language_pb.GenerateProgramRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    public generateProgram(request: khulnasoft_language_pb.GenerateProgramRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    public generateProgram(request: khulnasoft_language_pb.GenerateProgramRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProgramResponse) => void): grpc.ClientUnaryCall;
    public generateProject(request: khulnasoft_language_pb.GenerateProjectRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    public generateProject(request: khulnasoft_language_pb.GenerateProjectRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    public generateProject(request: khulnasoft_language_pb.GenerateProjectRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GenerateProjectResponse) => void): grpc.ClientUnaryCall;
    public generatePackage(request: khulnasoft_language_pb.GeneratePackageRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    public generatePackage(request: khulnasoft_language_pb.GeneratePackageRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    public generatePackage(request: khulnasoft_language_pb.GeneratePackageRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.GeneratePackageResponse) => void): grpc.ClientUnaryCall;
    public pack(request: khulnasoft_language_pb.PackRequest, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
    public pack(request: khulnasoft_language_pb.PackRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
    public pack(request: khulnasoft_language_pb.PackRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: khulnasoft_language_pb.PackResponse) => void): grpc.ClientUnaryCall;
}
