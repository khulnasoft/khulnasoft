# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import khulnasoft
import khulnasoft.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from ... import meta as _meta

__all__ = [
    'ConfigMapArgs',
    'ConfigMapArgsDict',
]

MYPY = False

if not MYPY:
    class ConfigMapArgsDict(TypedDict):
        """
        ConfigMap holds configuration data for pods to consume.
        """
        api_version: NotRequired[khulnasoft.Input[str]]
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        binary_data: NotRequired[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]]
        """
        BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
        """
        data: NotRequired[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]]
        """
        Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
        """
        immutable: NotRequired[khulnasoft.Input[bool]]
        """
        Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
        """
        kind: NotRequired[khulnasoft.Input[str]]
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        metadata: NotRequired[khulnasoft.Input['_meta.v1.ObjectMetaArgsDict']]
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
elif False:
    ConfigMapArgsDict: TypeAlias = Mapping[str, Any]

@khulnasoft.input_type
class ConfigMapArgs:
    def __init__(__self__, *,
                 api_version: Optional[khulnasoft.Input[str]] = None,
                 binary_data: Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]] = None,
                 data: Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]] = None,
                 immutable: Optional[khulnasoft.Input[bool]] = None,
                 kind: Optional[khulnasoft.Input[str]] = None,
                 metadata: Optional[khulnasoft.Input['_meta.v1.ObjectMetaArgs']] = None):
        """
        ConfigMap holds configuration data for pods to consume.
        :param khulnasoft.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]] binary_data: BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
        :param khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]] data: Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
        :param khulnasoft.Input[bool] immutable: Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
        :param khulnasoft.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param khulnasoft.Input['_meta.v1.ObjectMetaArgs'] metadata: Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        if api_version is not None:
            khulnasoft.set(__self__, "api_version", 'v1')
        if binary_data is not None:
            khulnasoft.set(__self__, "binary_data", binary_data)
        if data is not None:
            khulnasoft.set(__self__, "data", data)
        if immutable is not None:
            khulnasoft.set(__self__, "immutable", immutable)
        if kind is not None:
            khulnasoft.set(__self__, "kind", 'ConfigMap')
        if metadata is not None:
            khulnasoft.set(__self__, "metadata", metadata)

    @property
    @khulnasoft.getter(name="apiVersion")
    def api_version(self) -> Optional[khulnasoft.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        return khulnasoft.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[khulnasoft.Input[str]]):
        khulnasoft.set(self, "api_version", value)

    @property
    @khulnasoft.getter(name="binaryData")
    def binary_data(self) -> Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]]:
        """
        BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet.
        """
        return khulnasoft.get(self, "binary_data")

    @binary_data.setter
    def binary_data(self, value: Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]]):
        khulnasoft.set(self, "binary_data", value)

    @property
    @khulnasoft.getter
    def data(self) -> Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]]:
        """
        Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.
        """
        return khulnasoft.get(self, "data")

    @data.setter
    def data(self, value: Optional[khulnasoft.Input[Mapping[str, khulnasoft.Input[str]]]]):
        khulnasoft.set(self, "data", value)

    @property
    @khulnasoft.getter
    def immutable(self) -> Optional[khulnasoft.Input[bool]]:
        """
        Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated (only object metadata can be modified). If not set to true, the field can be modified at any time. Defaulted to nil.
        """
        return khulnasoft.get(self, "immutable")

    @immutable.setter
    def immutable(self, value: Optional[khulnasoft.Input[bool]]):
        khulnasoft.set(self, "immutable", value)

    @property
    @khulnasoft.getter
    def kind(self) -> Optional[khulnasoft.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        return khulnasoft.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[khulnasoft.Input[str]]):
        khulnasoft.set(self, "kind", value)

    @property
    @khulnasoft.getter
    def metadata(self) -> Optional[khulnasoft.Input['_meta.v1.ObjectMetaArgs']]:
        """
        Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        """
        return khulnasoft.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[khulnasoft.Input['_meta.v1.ObjectMetaArgs']]):
        khulnasoft.set(self, "metadata", value)

