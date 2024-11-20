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
from . import _utilities
from . import outputs

__all__ = [
    'SsisEnvironmentReferenceResponse',
    'SsisEnvironmentResponse',
    'SsisFolderResponse',
    'SsisPackageResponse',
    'SsisParameterResponse',
    'SsisProjectResponse',
    'SsisVariableResponse',
    'StorageAccountKeyResponse',
]

@khulnasoft.output_type
class SsisEnvironmentReferenceResponse(dict):
    """
    Ssis environment reference.
    """
    def __init__(__self__, *,
                 environment_folder_name: Optional[str] = None,
                 environment_name: Optional[str] = None,
                 id: Optional[float] = None,
                 reference_type: Optional[str] = None):
        """
        Ssis environment reference.
        :param str environment_folder_name: Environment folder name.
        :param str environment_name: Environment name.
        :param float id: Environment reference id.
        :param str reference_type: Reference type
        """
        if environment_folder_name is not None:
            khulnasoft.set(__self__, "environment_folder_name", environment_folder_name)
        if environment_name is not None:
            khulnasoft.set(__self__, "environment_name", environment_name)
        if id is not None:
            khulnasoft.set(__self__, "id", id)
        if reference_type is not None:
            khulnasoft.set(__self__, "reference_type", reference_type)

    @property
    @khulnasoft.getter(name="environmentFolderName")
    def environment_folder_name(self) -> Optional[str]:
        """
        Environment folder name.
        """
        return khulnasoft.get(self, "environment_folder_name")

    @property
    @khulnasoft.getter(name="environmentName")
    def environment_name(self) -> Optional[str]:
        """
        Environment name.
        """
        return khulnasoft.get(self, "environment_name")

    @property
    @khulnasoft.getter
    def id(self) -> Optional[float]:
        """
        Environment reference id.
        """
        return khulnasoft.get(self, "id")

    @property
    @khulnasoft.getter(name="referenceType")
    def reference_type(self) -> Optional[str]:
        """
        Reference type
        """
        return khulnasoft.get(self, "reference_type")


@khulnasoft.output_type
class SsisEnvironmentResponse(dict):
    """
    Ssis environment.
    """
    def __init__(__self__, *,
                 type: str,
                 description: Optional[str] = None,
                 folder_id: Optional[float] = None,
                 id: Optional[float] = None,
                 name: Optional[str] = None,
                 variables: Optional[Sequence['outputs.SsisVariableResponse']] = None):
        """
        Ssis environment.
        :param str type: The type of SSIS object metadata.
               Expected value is 'Environment'.
        :param str description: Metadata description.
        :param float folder_id: Folder id which contains environment.
        :param float id: Metadata id.
        :param str name: Metadata name.
        :param Sequence['SsisVariableResponse'] variables: Variable in environment
        """
        khulnasoft.set(__self__, "type", 'Environment')
        if description is not None:
            khulnasoft.set(__self__, "description", description)
        if folder_id is not None:
            khulnasoft.set(__self__, "folder_id", folder_id)
        if id is not None:
            khulnasoft.set(__self__, "id", id)
        if name is not None:
            khulnasoft.set(__self__, "name", name)
        if variables is not None:
            khulnasoft.set(__self__, "variables", variables)

    @property
    @khulnasoft.getter
    def type(self) -> str:
        """
        The type of SSIS object metadata.
        Expected value is 'Environment'.
        """
        return khulnasoft.get(self, "type")

    @property
    @khulnasoft.getter
    def description(self) -> Optional[str]:
        """
        Metadata description.
        """
        return khulnasoft.get(self, "description")

    @property
    @khulnasoft.getter(name="folderId")
    def folder_id(self) -> Optional[float]:
        """
        Folder id which contains environment.
        """
        return khulnasoft.get(self, "folder_id")

    @property
    @khulnasoft.getter
    def id(self) -> Optional[float]:
        """
        Metadata id.
        """
        return khulnasoft.get(self, "id")

    @property
    @khulnasoft.getter
    def name(self) -> Optional[str]:
        """
        Metadata name.
        """
        return khulnasoft.get(self, "name")

    @property
    @khulnasoft.getter
    def variables(self) -> Optional[Sequence['outputs.SsisVariableResponse']]:
        """
        Variable in environment
        """
        return khulnasoft.get(self, "variables")


@khulnasoft.output_type
class SsisFolderResponse(dict):
    """
    Ssis folder.
    """
    def __init__(__self__, *,
                 type: str,
                 description: Optional[str] = None,
                 id: Optional[float] = None,
                 name: Optional[str] = None):
        """
        Ssis folder.
        :param str type: The type of SSIS object metadata.
               Expected value is 'Folder'.
        :param str description: Metadata description.
        :param float id: Metadata id.
        :param str name: Metadata name.
        """
        khulnasoft.set(__self__, "type", 'Folder')
        if description is not None:
            khulnasoft.set(__self__, "description", description)
        if id is not None:
            khulnasoft.set(__self__, "id", id)
        if name is not None:
            khulnasoft.set(__self__, "name", name)

    @property
    @khulnasoft.getter
    def type(self) -> str:
        """
        The type of SSIS object metadata.
        Expected value is 'Folder'.
        """
        return khulnasoft.get(self, "type")

    @property
    @khulnasoft.getter
    def description(self) -> Optional[str]:
        """
        Metadata description.
        """
        return khulnasoft.get(self, "description")

    @property
    @khulnasoft.getter
    def id(self) -> Optional[float]:
        """
        Metadata id.
        """
        return khulnasoft.get(self, "id")

    @property
    @khulnasoft.getter
    def name(self) -> Optional[str]:
        """
        Metadata name.
        """
        return khulnasoft.get(self, "name")


@khulnasoft.output_type
class SsisPackageResponse(dict):
    """
    Ssis Package.
    """
    def __init__(__self__, *,
                 type: str,
                 description: Optional[str] = None,
                 folder_id: Optional[float] = None,
                 id: Optional[float] = None,
                 name: Optional[str] = None,
                 parameters: Optional[Sequence['outputs.SsisParameterResponse']] = None,
                 project_id: Optional[float] = None,
                 project_version: Optional[float] = None):
        """
        Ssis Package.
        :param str type: The type of SSIS object metadata.
               Expected value is 'Package'.
        :param str description: Metadata description.
        :param float folder_id: Folder id which contains package.
        :param float id: Metadata id.
        :param str name: Metadata name.
        :param Sequence['SsisParameterResponse'] parameters: Parameters in package
        :param float project_id: Project id which contains package.
        :param float project_version: Project version which contains package.
        """
        khulnasoft.set(__self__, "type", 'Package')
        if description is not None:
            khulnasoft.set(__self__, "description", description)
        if folder_id is not None:
            khulnasoft.set(__self__, "folder_id", folder_id)
        if id is not None:
            khulnasoft.set(__self__, "id", id)
        if name is not None:
            khulnasoft.set(__self__, "name", name)
        if parameters is not None:
            khulnasoft.set(__self__, "parameters", parameters)
        if project_id is not None:
            khulnasoft.set(__self__, "project_id", project_id)
        if project_version is not None:
            khulnasoft.set(__self__, "project_version", project_version)

    @property
    @khulnasoft.getter
    def type(self) -> str:
        """
        The type of SSIS object metadata.
        Expected value is 'Package'.
        """
        return khulnasoft.get(self, "type")

    @property
    @khulnasoft.getter
    def description(self) -> Optional[str]:
        """
        Metadata description.
        """
        return khulnasoft.get(self, "description")

    @property
    @khulnasoft.getter(name="folderId")
    def folder_id(self) -> Optional[float]:
        """
        Folder id which contains package.
        """
        return khulnasoft.get(self, "folder_id")

    @property
    @khulnasoft.getter
    def id(self) -> Optional[float]:
        """
        Metadata id.
        """
        return khulnasoft.get(self, "id")

    @property
    @khulnasoft.getter
    def name(self) -> Optional[str]:
        """
        Metadata name.
        """
        return khulnasoft.get(self, "name")

    @property
    @khulnasoft.getter
    def parameters(self) -> Optional[Sequence['outputs.SsisParameterResponse']]:
        """
        Parameters in package
        """
        return khulnasoft.get(self, "parameters")

    @property
    @khulnasoft.getter(name="projectId")
    def project_id(self) -> Optional[float]:
        """
        Project id which contains package.
        """
        return khulnasoft.get(self, "project_id")

    @property
    @khulnasoft.getter(name="projectVersion")
    def project_version(self) -> Optional[float]:
        """
        Project version which contains package.
        """
        return khulnasoft.get(self, "project_version")


@khulnasoft.output_type
class SsisParameterResponse(dict):
    """
    Ssis parameter.
    """
    def __init__(__self__, *,
                 data_type: Optional[str] = None,
                 default_value: Optional[str] = None,
                 description: Optional[str] = None,
                 design_default_value: Optional[str] = None,
                 id: Optional[float] = None,
                 name: Optional[str] = None,
                 required: Optional[bool] = None,
                 sensitive: Optional[bool] = None,
                 sensitive_default_value: Optional[str] = None,
                 value_set: Optional[bool] = None,
                 value_type: Optional[str] = None,
                 variable: Optional[str] = None):
        """
        Ssis parameter.
        :param str data_type: Parameter type.
        :param str default_value: Default value of parameter.
        :param str description: Parameter description.
        :param str design_default_value: Design default value of parameter.
        :param float id: Parameter id.
        :param str name: Parameter name.
        :param bool required: Whether parameter is required.
        :param bool sensitive: Whether parameter is sensitive.
        :param str sensitive_default_value: Default sensitive value of parameter.
        :param bool value_set: Parameter value set.
        :param str value_type: Parameter value type.
        :param str variable: Parameter reference variable.
        """
        if data_type is not None:
            khulnasoft.set(__self__, "data_type", data_type)
        if default_value is not None:
            khulnasoft.set(__self__, "default_value", default_value)
        if description is not None:
            khulnasoft.set(__self__, "description", description)
        if design_default_value is not None:
            khulnasoft.set(__self__, "design_default_value", design_default_value)
        if id is not None:
            khulnasoft.set(__self__, "id", id)
        if name is not None:
            khulnasoft.set(__self__, "name", name)
        if required is not None:
            khulnasoft.set(__self__, "required", required)
        if sensitive is not None:
            khulnasoft.set(__self__, "sensitive", sensitive)
        if sensitive_default_value is not None:
            khulnasoft.set(__self__, "sensitive_default_value", sensitive_default_value)
        if value_set is not None:
            khulnasoft.set(__self__, "value_set", value_set)
        if value_type is not None:
            khulnasoft.set(__self__, "value_type", value_type)
        if variable is not None:
            khulnasoft.set(__self__, "variable", variable)

    @property
    @khulnasoft.getter(name="dataType")
    def data_type(self) -> Optional[str]:
        """
        Parameter type.
        """
        return khulnasoft.get(self, "data_type")

    @property
    @khulnasoft.getter(name="defaultValue")
    def default_value(self) -> Optional[str]:
        """
        Default value of parameter.
        """
        return khulnasoft.get(self, "default_value")

    @property
    @khulnasoft.getter
    def description(self) -> Optional[str]:
        """
        Parameter description.
        """
        return khulnasoft.get(self, "description")

    @property
    @khulnasoft.getter(name="designDefaultValue")
    def design_default_value(self) -> Optional[str]:
        """
        Design default value of parameter.
        """
        return khulnasoft.get(self, "design_default_value")

    @property
    @khulnasoft.getter
    def id(self) -> Optional[float]:
        """
        Parameter id.
        """
        return khulnasoft.get(self, "id")

    @property
    @khulnasoft.getter
    def name(self) -> Optional[str]:
        """
        Parameter name.
        """
        return khulnasoft.get(self, "name")

    @property
    @khulnasoft.getter
    def required(self) -> Optional[bool]:
        """
        Whether parameter is required.
        """
        return khulnasoft.get(self, "required")

    @property
    @khulnasoft.getter
    def sensitive(self) -> Optional[bool]:
        """
        Whether parameter is sensitive.
        """
        return khulnasoft.get(self, "sensitive")

    @property
    @khulnasoft.getter(name="sensitiveDefaultValue")
    def sensitive_default_value(self) -> Optional[str]:
        """
        Default sensitive value of parameter.
        """
        return khulnasoft.get(self, "sensitive_default_value")

    @property
    @khulnasoft.getter(name="valueSet")
    def value_set(self) -> Optional[bool]:
        """
        Parameter value set.
        """
        return khulnasoft.get(self, "value_set")

    @property
    @khulnasoft.getter(name="valueType")
    def value_type(self) -> Optional[str]:
        """
        Parameter value type.
        """
        return khulnasoft.get(self, "value_type")

    @property
    @khulnasoft.getter
    def variable(self) -> Optional[str]:
        """
        Parameter reference variable.
        """
        return khulnasoft.get(self, "variable")


@khulnasoft.output_type
class SsisProjectResponse(dict):
    """
    Ssis project.
    """
    def __init__(__self__, *,
                 type: str,
                 description: Optional[str] = None,
                 environment_refs: Optional[Sequence['outputs.SsisEnvironmentReferenceResponse']] = None,
                 folder_id: Optional[float] = None,
                 id: Optional[float] = None,
                 name: Optional[str] = None,
                 parameters: Optional[Sequence['outputs.SsisParameterResponse']] = None,
                 version: Optional[float] = None):
        """
        Ssis project.
        :param str type: The type of SSIS object metadata.
               Expected value is 'Project'.
        :param str description: Metadata description.
        :param Sequence['SsisEnvironmentReferenceResponse'] environment_refs: Environment reference in project
        :param float folder_id: Folder id which contains project.
        :param float id: Metadata id.
        :param str name: Metadata name.
        :param Sequence['SsisParameterResponse'] parameters: Parameters in project
        :param float version: Project version.
        """
        khulnasoft.set(__self__, "type", 'Project')
        if description is not None:
            khulnasoft.set(__self__, "description", description)
        if environment_refs is not None:
            khulnasoft.set(__self__, "environment_refs", environment_refs)
        if folder_id is not None:
            khulnasoft.set(__self__, "folder_id", folder_id)
        if id is not None:
            khulnasoft.set(__self__, "id", id)
        if name is not None:
            khulnasoft.set(__self__, "name", name)
        if parameters is not None:
            khulnasoft.set(__self__, "parameters", parameters)
        if version is not None:
            khulnasoft.set(__self__, "version", version)

    @property
    @khulnasoft.getter
    def type(self) -> str:
        """
        The type of SSIS object metadata.
        Expected value is 'Project'.
        """
        return khulnasoft.get(self, "type")

    @property
    @khulnasoft.getter
    def description(self) -> Optional[str]:
        """
        Metadata description.
        """
        return khulnasoft.get(self, "description")

    @property
    @khulnasoft.getter(name="environmentRefs")
    def environment_refs(self) -> Optional[Sequence['outputs.SsisEnvironmentReferenceResponse']]:
        """
        Environment reference in project
        """
        return khulnasoft.get(self, "environment_refs")

    @property
    @khulnasoft.getter(name="folderId")
    def folder_id(self) -> Optional[float]:
        """
        Folder id which contains project.
        """
        return khulnasoft.get(self, "folder_id")

    @property
    @khulnasoft.getter
    def id(self) -> Optional[float]:
        """
        Metadata id.
        """
        return khulnasoft.get(self, "id")

    @property
    @khulnasoft.getter
    def name(self) -> Optional[str]:
        """
        Metadata name.
        """
        return khulnasoft.get(self, "name")

    @property
    @khulnasoft.getter
    def parameters(self) -> Optional[Sequence['outputs.SsisParameterResponse']]:
        """
        Parameters in project
        """
        return khulnasoft.get(self, "parameters")

    @property
    @khulnasoft.getter
    def version(self) -> Optional[float]:
        """
        Project version.
        """
        return khulnasoft.get(self, "version")


@khulnasoft.output_type
class SsisVariableResponse(dict):
    """
    Ssis variable.
    """
    def __init__(__self__, *,
                 data_type: Optional[str] = None,
                 description: Optional[str] = None,
                 id: Optional[float] = None,
                 name: Optional[str] = None,
                 sensitive: Optional[bool] = None,
                 sensitive_value: Optional[str] = None,
                 value: Optional[str] = None):
        """
        Ssis variable.
        :param str data_type: Variable type.
        :param str description: Variable description.
        :param float id: Variable id.
        :param str name: Variable name.
        :param bool sensitive: Whether variable is sensitive.
        :param str sensitive_value: Variable sensitive value.
        :param str value: Variable value.
        """
        if data_type is not None:
            khulnasoft.set(__self__, "data_type", data_type)
        if description is not None:
            khulnasoft.set(__self__, "description", description)
        if id is not None:
            khulnasoft.set(__self__, "id", id)
        if name is not None:
            khulnasoft.set(__self__, "name", name)
        if sensitive is not None:
            khulnasoft.set(__self__, "sensitive", sensitive)
        if sensitive_value is not None:
            khulnasoft.set(__self__, "sensitive_value", sensitive_value)
        if value is not None:
            khulnasoft.set(__self__, "value", value)

    @property
    @khulnasoft.getter(name="dataType")
    def data_type(self) -> Optional[str]:
        """
        Variable type.
        """
        return khulnasoft.get(self, "data_type")

    @property
    @khulnasoft.getter
    def description(self) -> Optional[str]:
        """
        Variable description.
        """
        return khulnasoft.get(self, "description")

    @property
    @khulnasoft.getter
    def id(self) -> Optional[float]:
        """
        Variable id.
        """
        return khulnasoft.get(self, "id")

    @property
    @khulnasoft.getter
    def name(self) -> Optional[str]:
        """
        Variable name.
        """
        return khulnasoft.get(self, "name")

    @property
    @khulnasoft.getter
    def sensitive(self) -> Optional[bool]:
        """
        Whether variable is sensitive.
        """
        return khulnasoft.get(self, "sensitive")

    @property
    @khulnasoft.getter(name="sensitiveValue")
    def sensitive_value(self) -> Optional[str]:
        """
        Variable sensitive value.
        """
        return khulnasoft.get(self, "sensitive_value")

    @property
    @khulnasoft.getter
    def value(self) -> Optional[str]:
        """
        Variable value.
        """
        return khulnasoft.get(self, "value")


@khulnasoft.output_type
class StorageAccountKeyResponse(dict):
    """
    An access key for the storage account.
    """
    def __init__(__self__, *,
                 creation_time: str,
                 key_name: str,
                 permissions: str,
                 value: str):
        """
        An access key for the storage account.
        :param str creation_time: Creation time of the key, in round trip date format.
        :param str key_name: Name of the key.
        :param str permissions: Permissions for the key -- read-only or full permissions.
        :param str value: Base 64-encoded value of the key.
        """
        khulnasoft.set(__self__, "creation_time", creation_time)
        khulnasoft.set(__self__, "key_name", key_name)
        khulnasoft.set(__self__, "permissions", permissions)
        khulnasoft.set(__self__, "value", value)

    @property
    @khulnasoft.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        Creation time of the key, in round trip date format.
        """
        return khulnasoft.get(self, "creation_time")

    @property
    @khulnasoft.getter(name="keyName")
    def key_name(self) -> str:
        """
        Name of the key.
        """
        return khulnasoft.get(self, "key_name")

    @property
    @khulnasoft.getter
    def permissions(self) -> str:
        """
        Permissions for the key -- read-only or full permissions.
        """
        return khulnasoft.get(self, "permissions")

    @property
    @khulnasoft.getter
    def value(self) -> str:
        """
        Base 64-encoded value of the key.
        """
        return khulnasoft.get(self, "value")


