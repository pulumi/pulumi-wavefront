# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetRoleResult',
    'AwaitableGetRoleResult',
    'get_role',
    'get_role_output',
]

@pulumi.output_type
class GetRoleResult:
    """
    A collection of values returned by getRole.
    """
    def __init__(__self__, description=None, id=None, name=None, permissions=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Human-readable description of the role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the role in Wavefront.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the role in Wavefront.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> Sequence[str]:
        """
        The list of permissions associated with role.
        """
        return pulumi.get(self, "permissions")


class AwaitableGetRoleResult(GetRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleResult(
            description=self.description,
            id=self.id,
            name=self.name,
            permissions=self.permissions)


def get_role(id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleResult:
    """
    Use this data source to get information about a Wavefront role by its ID.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_wavefront as wavefront

    example = wavefront.get_role(id="role-id")
    ```


    :param str id: The ID associated with the role data to be fetched.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('wavefront:index/getRole:getRole', __args__, opts=opts, typ=GetRoleResult).value

    return AwaitableGetRoleResult(
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        permissions=__ret__.permissions)


@_utilities.lift_output_func(get_role)
def get_role_output(id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRoleResult]:
    """
    Use this data source to get information about a Wavefront role by its ID.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_wavefront as wavefront

    example = wavefront.get_role(id="role-id")
    ```


    :param str id: The ID associated with the role data to be fetched.
    """
    ...