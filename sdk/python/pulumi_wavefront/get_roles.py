# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetRolesResult',
    'AwaitableGetRolesResult',
    'get_roles',
    'get_roles_output',
]

@pulumi.output_type
class GetRolesResult:
    """
    A collection of values returned by getRoles.
    """
    def __init__(__self__, id=None, limit=None, offset=None, roles=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)
        if offset and not isinstance(offset, int):
            raise TypeError("Expected argument 'offset' to be a int")
        pulumi.set(__self__, "offset", offset)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def limit(self) -> Optional[int]:
        return pulumi.get(self, "limit")

    @property
    @pulumi.getter
    def offset(self) -> Optional[int]:
        return pulumi.get(self, "offset")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetRolesRoleResult']:
        """
        List of Wavefront Roles.
        """
        return pulumi.get(self, "roles")


class AwaitableGetRolesResult(GetRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRolesResult(
            id=self.id,
            limit=self.limit,
            offset=self.offset,
            roles=self.roles)


def get_roles(limit: Optional[int] = None,
              offset: Optional[int] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRolesResult:
    """
    Use this data source to get all Roles in Wavefront.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_wavefront as wavefront

    roles = wavefront.get_roles(limit=10,
        offset=0)
    ```


    :param int limit: Limit is the maximum number of results to be returned. Defaults to 100.
    :param int offset: Offset is the offset from the first result to be returned. Defaults to 0.
    """
    __args__ = dict()
    __args__['limit'] = limit
    __args__['offset'] = offset
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('wavefront:index/getRoles:getRoles', __args__, opts=opts, typ=GetRolesResult).value

    return AwaitableGetRolesResult(
        id=__ret__.id,
        limit=__ret__.limit,
        offset=__ret__.offset,
        roles=__ret__.roles)


@_utilities.lift_output_func(get_roles)
def get_roles_output(limit: Optional[pulumi.Input[Optional[int]]] = None,
                     offset: Optional[pulumi.Input[Optional[int]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRolesResult]:
    """
    Use this data source to get all Roles in Wavefront.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_wavefront as wavefront

    roles = wavefront.get_roles(limit=10,
        offset=0)
    ```


    :param int limit: Limit is the maximum number of results to be returned. Defaults to 100.
    :param int offset: Offset is the offset from the first result to be returned. Defaults to 0.
    """
    ...