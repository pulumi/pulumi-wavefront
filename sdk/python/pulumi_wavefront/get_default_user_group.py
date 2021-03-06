# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetDefaultUserGroupResult',
    'AwaitableGetDefaultUserGroupResult',
    'get_default_user_group',
]

@pulumi.output_type
class GetDefaultUserGroupResult:
    """
    A collection of values returned by getDefaultUserGroup.
    """
    def __init__(__self__, group_id=None, id=None):
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        """
        Set to the Group ID of the `Everyone` group, suitable for referencing
        in other resources that support group memberships. s
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetDefaultUserGroupResult(GetDefaultUserGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultUserGroupResult(
            group_id=self.group_id,
            id=self.id)


def get_default_user_group(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultUserGroupResult:
    """
    Use this data source to get the Group ID of the `Everyone` group in Wavefront.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_wavefront as wavefront

    everyone_group = wavefront.get_default_user_group()
    ```
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('wavefront:index/getDefaultUserGroup:getDefaultUserGroup', __args__, opts=opts, typ=GetDefaultUserGroupResult).value

    return AwaitableGetDefaultUserGroupResult(
        group_id=__ret__.group_id,
        id=__ret__.id)
