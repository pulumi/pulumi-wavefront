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
    'GetAlertsResult',
    'AwaitableGetAlertsResult',
    'get_alerts',
    'get_alerts_output',
]

@pulumi.output_type
class GetAlertsResult:
    """
    A collection of values returned by getAlerts.
    """
    def __init__(__self__, alerts=None, id=None, limit=None, offset=None):
        if alerts and not isinstance(alerts, list):
            raise TypeError("Expected argument 'alerts' to be a list")
        pulumi.set(__self__, "alerts", alerts)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if limit and not isinstance(limit, int):
            raise TypeError("Expected argument 'limit' to be a int")
        pulumi.set(__self__, "limit", limit)
        if offset and not isinstance(offset, int):
            raise TypeError("Expected argument 'offset' to be a int")
        pulumi.set(__self__, "offset", offset)

    @property
    @pulumi.getter
    def alerts(self) -> Sequence['outputs.GetAlertsAlertResult']:
        """
        List of all alerts in Wavefront. For each alert you will see a list of attributes.
        """
        return pulumi.get(self, "alerts")

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


class AwaitableGetAlertsResult(GetAlertsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlertsResult(
            alerts=self.alerts,
            id=self.id,
            limit=self.limit,
            offset=self.offset)


def get_alerts(limit: Optional[int] = None,
               offset: Optional[int] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlertsResult:
    """
    Use this data source to get information about all Wavefront alerts.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_wavefront as wavefront

    example = wavefront.get_alerts(limit=10,
        offset=0)
    ```


    :param int limit: Limit is the maximum number of results to be returned. Defaults to 100.
    :param int offset: Offset is the offset from the first result to be returned. Defaults to 0.
    """
    __args__ = dict()
    __args__['limit'] = limit
    __args__['offset'] = offset
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('wavefront:index/getAlerts:getAlerts', __args__, opts=opts, typ=GetAlertsResult).value

    return AwaitableGetAlertsResult(
        alerts=__ret__.alerts,
        id=__ret__.id,
        limit=__ret__.limit,
        offset=__ret__.offset)


@_utilities.lift_output_func(get_alerts)
def get_alerts_output(limit: Optional[pulumi.Input[Optional[int]]] = None,
                      offset: Optional[pulumi.Input[Optional[int]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAlertsResult]:
    """
    Use this data source to get information about all Wavefront alerts.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_wavefront as wavefront

    example = wavefront.get_alerts(limit=10,
        offset=0)
    ```


    :param int limit: Limit is the maximum number of results to be returned. Defaults to 100.
    :param int offset: Offset is the offset from the first result to be returned. Defaults to 0.
    """
    ...