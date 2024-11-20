"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
Copyright 2016-2018, Pulumi Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class PluginInfo(google.protobuf.message.Message):
    """PluginInfo is meta-information about a plugin that is used by the system."""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    VERSION_FIELD_NUMBER: builtins.int
    version: builtins.str
    """the semver for this plugin."""
    def __init__(
        self,
        *,
        version: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["version", b"version"]) -> None: ...

global___PluginInfo = PluginInfo

@typing_extensions.final
class PluginDependency(google.protobuf.message.Message):
    """PluginDependency is information about a plugin that a program may depend upon."""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    @typing_extensions.final
    class ChecksumsEntry(google.protobuf.message.Message):
        DESCRIPTOR: google.protobuf.descriptor.Descriptor

        KEY_FIELD_NUMBER: builtins.int
        VALUE_FIELD_NUMBER: builtins.int
        key: builtins.str
        value: builtins.bytes
        def __init__(
            self,
            *,
            key: builtins.str = ...,
            value: builtins.bytes = ...,
        ) -> None: ...
        def ClearField(self, field_name: typing_extensions.Literal["key", b"key", "value", b"value"]) -> None: ...

    NAME_FIELD_NUMBER: builtins.int
    KIND_FIELD_NUMBER: builtins.int
    VERSION_FIELD_NUMBER: builtins.int
    SERVER_FIELD_NUMBER: builtins.int
    CHECKSUMS_FIELD_NUMBER: builtins.int
    name: builtins.str
    """the name of the plugin."""
    kind: builtins.str
    """the kind of plugin (e.g., language, etc)."""
    version: builtins.str
    """the semver for this plugin."""
    server: builtins.str
    """the URL of a server that can be used to download this plugin, if needed."""
    @property
    def checksums(self) -> google.protobuf.internal.containers.ScalarMap[builtins.str, builtins.bytes]:
        """a map of the checksums for the plugin, will be empty from old language runtimes. The keys should match
        the os and architecture names used in pulumi releases, e.g. "darwin-amd64", "windows-arm64".
        """
    def __init__(
        self,
        *,
        name: builtins.str = ...,
        kind: builtins.str = ...,
        version: builtins.str = ...,
        server: builtins.str = ...,
        checksums: collections.abc.Mapping[builtins.str, builtins.bytes] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["checksums", b"checksums", "kind", b"kind", "name", b"name", "server", b"server", "version", b"version"]) -> None: ...

global___PluginDependency = PluginDependency

@typing_extensions.final
class PluginAttach(google.protobuf.message.Message):
    """PluginAttach is used to attach an already running plugin to the engine.

    Normally the engine starts the plugin process itself and passes the engine address as the first argumnent.
    But when debugging it can be useful to have an already running provider that the engine instead attaches
    to, this message is used so the provider can still be passed the engine address to communicate with.
    """

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ADDRESS_FIELD_NUMBER: builtins.int
    address: builtins.str
    """the grpc address for the engine"""
    def __init__(
        self,
        *,
        address: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["address", b"address"]) -> None: ...

global___PluginAttach = PluginAttach