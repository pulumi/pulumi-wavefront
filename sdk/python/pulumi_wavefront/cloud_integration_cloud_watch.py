# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['CloudIntegrationCloudWatchArgs', 'CloudIntegrationCloudWatch']

@pulumi.input_type
class CloudIntegrationCloudWatchArgs:
    def __init__(__self__, *,
                 external_id: pulumi.Input[_builtins.str],
                 role_arn: pulumi.Input[_builtins.str],
                 additional_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 force_save: Optional[pulumi.Input[_builtins.bool]] = None,
                 instance_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 metric_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespaces: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 point_tag_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
                 service: Optional[pulumi.Input[_builtins.str]] = None,
                 service_refresh_rate_in_minutes: Optional[pulumi.Input[_builtins.int]] = None,
                 volume_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None):
        """
        The set of arguments for constructing a CloudIntegrationCloudWatch resource.
        :param pulumi.Input[_builtins.str] external_id: The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
        :param pulumi.Input[_builtins.str] role_arn: The external ID corresponding to the Role ARN.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] additional_tags: A list of point tag key-values to add to every point ingested using this integration.
        :param pulumi.Input[_builtins.bool] force_save: Forces this resource to save, even if errors are present.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] instance_selection_tags: A string->string map allow list of instance tag-value pairs (in AWS).
               If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
               Multiple entries are OR'ed.
        :param pulumi.Input[_builtins.str] metric_filter_regex: A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
        :param pulumi.Input[_builtins.str] name: The human-readable name of this integration.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] namespaces: A list of namespaces that limit what we query from CloudWatch.
        :param pulumi.Input[_builtins.str] point_tag_filter_regex: A regular expression that AWS tag key name must match (case-insensitively)
               in order to be ingested.
        :param pulumi.Input[_builtins.str] service: A value denoting which cloud service this service integrates with.
        :param pulumi.Input[_builtins.int] service_refresh_rate_in_minutes: How often, in minutes, to refresh the service.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] volume_selection_tags: A string->string map of allow list of volume tag-value pairs (in AWS).
               If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
               Multiple entries are OR'ed.
        """
        pulumi.set(__self__, "external_id", external_id)
        pulumi.set(__self__, "role_arn", role_arn)
        if additional_tags is not None:
            pulumi.set(__self__, "additional_tags", additional_tags)
        if force_save is not None:
            pulumi.set(__self__, "force_save", force_save)
        if instance_selection_tags is not None:
            pulumi.set(__self__, "instance_selection_tags", instance_selection_tags)
        if metric_filter_regex is not None:
            pulumi.set(__self__, "metric_filter_regex", metric_filter_regex)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespaces is not None:
            pulumi.set(__self__, "namespaces", namespaces)
        if point_tag_filter_regex is not None:
            pulumi.set(__self__, "point_tag_filter_regex", point_tag_filter_regex)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if service_refresh_rate_in_minutes is not None:
            pulumi.set(__self__, "service_refresh_rate_in_minutes", service_refresh_rate_in_minutes)
        if volume_selection_tags is not None:
            pulumi.set(__self__, "volume_selection_tags", volume_selection_tags)

    @_builtins.property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Input[_builtins.str]:
        """
        The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
        """
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "external_id", value)

    @_builtins.property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[_builtins.str]:
        """
        The external ID corresponding to the Role ARN.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "role_arn", value)

    @_builtins.property
    @pulumi.getter(name="additionalTags")
    def additional_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        A list of point tag key-values to add to every point ingested using this integration.
        """
        return pulumi.get(self, "additional_tags")

    @additional_tags.setter
    def additional_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "additional_tags", value)

    @_builtins.property
    @pulumi.getter(name="forceSave")
    def force_save(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Forces this resource to save, even if errors are present.
        """
        return pulumi.get(self, "force_save")

    @force_save.setter
    def force_save(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "force_save", value)

    @_builtins.property
    @pulumi.getter(name="instanceSelectionTags")
    def instance_selection_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        A string->string map allow list of instance tag-value pairs (in AWS).
        If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
        Multiple entries are OR'ed.
        """
        return pulumi.get(self, "instance_selection_tags")

    @instance_selection_tags.setter
    def instance_selection_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "instance_selection_tags", value)

    @_builtins.property
    @pulumi.getter(name="metricFilterRegex")
    def metric_filter_regex(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
        """
        return pulumi.get(self, "metric_filter_regex")

    @metric_filter_regex.setter
    def metric_filter_regex(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "metric_filter_regex", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The human-readable name of this integration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def namespaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A list of namespaces that limit what we query from CloudWatch.
        """
        return pulumi.get(self, "namespaces")

    @namespaces.setter
    def namespaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "namespaces", value)

    @_builtins.property
    @pulumi.getter(name="pointTagFilterRegex")
    def point_tag_filter_regex(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A regular expression that AWS tag key name must match (case-insensitively)
        in order to be ingested.
        """
        return pulumi.get(self, "point_tag_filter_regex")

    @point_tag_filter_regex.setter
    def point_tag_filter_regex(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "point_tag_filter_regex", value)

    @_builtins.property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A value denoting which cloud service this service integrates with.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "service", value)

    @_builtins.property
    @pulumi.getter(name="serviceRefreshRateInMinutes")
    def service_refresh_rate_in_minutes(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        How often, in minutes, to refresh the service.
        """
        return pulumi.get(self, "service_refresh_rate_in_minutes")

    @service_refresh_rate_in_minutes.setter
    def service_refresh_rate_in_minutes(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "service_refresh_rate_in_minutes", value)

    @_builtins.property
    @pulumi.getter(name="volumeSelectionTags")
    def volume_selection_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        A string->string map of allow list of volume tag-value pairs (in AWS).
        If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
        Multiple entries are OR'ed.
        """
        return pulumi.get(self, "volume_selection_tags")

    @volume_selection_tags.setter
    def volume_selection_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "volume_selection_tags", value)


@pulumi.input_type
class _CloudIntegrationCloudWatchState:
    def __init__(__self__, *,
                 additional_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 external_id: Optional[pulumi.Input[_builtins.str]] = None,
                 force_save: Optional[pulumi.Input[_builtins.bool]] = None,
                 instance_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 metric_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespaces: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 point_tag_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
                 role_arn: Optional[pulumi.Input[_builtins.str]] = None,
                 service: Optional[pulumi.Input[_builtins.str]] = None,
                 service_refresh_rate_in_minutes: Optional[pulumi.Input[_builtins.int]] = None,
                 volume_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering CloudIntegrationCloudWatch resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] additional_tags: A list of point tag key-values to add to every point ingested using this integration.
        :param pulumi.Input[_builtins.str] external_id: The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
        :param pulumi.Input[_builtins.bool] force_save: Forces this resource to save, even if errors are present.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] instance_selection_tags: A string->string map allow list of instance tag-value pairs (in AWS).
               If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
               Multiple entries are OR'ed.
        :param pulumi.Input[_builtins.str] metric_filter_regex: A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
        :param pulumi.Input[_builtins.str] name: The human-readable name of this integration.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] namespaces: A list of namespaces that limit what we query from CloudWatch.
        :param pulumi.Input[_builtins.str] point_tag_filter_regex: A regular expression that AWS tag key name must match (case-insensitively)
               in order to be ingested.
        :param pulumi.Input[_builtins.str] role_arn: The external ID corresponding to the Role ARN.
        :param pulumi.Input[_builtins.str] service: A value denoting which cloud service this service integrates with.
        :param pulumi.Input[_builtins.int] service_refresh_rate_in_minutes: How often, in minutes, to refresh the service.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] volume_selection_tags: A string->string map of allow list of volume tag-value pairs (in AWS).
               If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
               Multiple entries are OR'ed.
        """
        if additional_tags is not None:
            pulumi.set(__self__, "additional_tags", additional_tags)
        if external_id is not None:
            pulumi.set(__self__, "external_id", external_id)
        if force_save is not None:
            pulumi.set(__self__, "force_save", force_save)
        if instance_selection_tags is not None:
            pulumi.set(__self__, "instance_selection_tags", instance_selection_tags)
        if metric_filter_regex is not None:
            pulumi.set(__self__, "metric_filter_regex", metric_filter_regex)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespaces is not None:
            pulumi.set(__self__, "namespaces", namespaces)
        if point_tag_filter_regex is not None:
            pulumi.set(__self__, "point_tag_filter_regex", point_tag_filter_regex)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if service_refresh_rate_in_minutes is not None:
            pulumi.set(__self__, "service_refresh_rate_in_minutes", service_refresh_rate_in_minutes)
        if volume_selection_tags is not None:
            pulumi.set(__self__, "volume_selection_tags", volume_selection_tags)

    @_builtins.property
    @pulumi.getter(name="additionalTags")
    def additional_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        A list of point tag key-values to add to every point ingested using this integration.
        """
        return pulumi.get(self, "additional_tags")

    @additional_tags.setter
    def additional_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "additional_tags", value)

    @_builtins.property
    @pulumi.getter(name="externalId")
    def external_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
        """
        return pulumi.get(self, "external_id")

    @external_id.setter
    def external_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "external_id", value)

    @_builtins.property
    @pulumi.getter(name="forceSave")
    def force_save(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Forces this resource to save, even if errors are present.
        """
        return pulumi.get(self, "force_save")

    @force_save.setter
    def force_save(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "force_save", value)

    @_builtins.property
    @pulumi.getter(name="instanceSelectionTags")
    def instance_selection_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        A string->string map allow list of instance tag-value pairs (in AWS).
        If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
        Multiple entries are OR'ed.
        """
        return pulumi.get(self, "instance_selection_tags")

    @instance_selection_tags.setter
    def instance_selection_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "instance_selection_tags", value)

    @_builtins.property
    @pulumi.getter(name="metricFilterRegex")
    def metric_filter_regex(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
        """
        return pulumi.get(self, "metric_filter_regex")

    @metric_filter_regex.setter
    def metric_filter_regex(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "metric_filter_regex", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The human-readable name of this integration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def namespaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A list of namespaces that limit what we query from CloudWatch.
        """
        return pulumi.get(self, "namespaces")

    @namespaces.setter
    def namespaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "namespaces", value)

    @_builtins.property
    @pulumi.getter(name="pointTagFilterRegex")
    def point_tag_filter_regex(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A regular expression that AWS tag key name must match (case-insensitively)
        in order to be ingested.
        """
        return pulumi.get(self, "point_tag_filter_regex")

    @point_tag_filter_regex.setter
    def point_tag_filter_regex(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "point_tag_filter_regex", value)

    @_builtins.property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The external ID corresponding to the Role ARN.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "role_arn", value)

    @_builtins.property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A value denoting which cloud service this service integrates with.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "service", value)

    @_builtins.property
    @pulumi.getter(name="serviceRefreshRateInMinutes")
    def service_refresh_rate_in_minutes(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        How often, in minutes, to refresh the service.
        """
        return pulumi.get(self, "service_refresh_rate_in_minutes")

    @service_refresh_rate_in_minutes.setter
    def service_refresh_rate_in_minutes(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "service_refresh_rate_in_minutes", value)

    @_builtins.property
    @pulumi.getter(name="volumeSelectionTags")
    def volume_selection_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        A string->string map of allow list of volume tag-value pairs (in AWS).
        If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
        Multiple entries are OR'ed.
        """
        return pulumi.get(self, "volume_selection_tags")

    @volume_selection_tags.setter
    def volume_selection_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "volume_selection_tags", value)


@pulumi.type_token("wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch")
class CloudIntegrationCloudWatch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 external_id: Optional[pulumi.Input[_builtins.str]] = None,
                 force_save: Optional[pulumi.Input[_builtins.bool]] = None,
                 instance_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 metric_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespaces: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 point_tag_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
                 role_arn: Optional[pulumi.Input[_builtins.str]] = None,
                 service: Optional[pulumi.Input[_builtins.str]] = None,
                 service_refresh_rate_in_minutes: Optional[pulumi.Input[_builtins.int]] = None,
                 volume_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        """
        Provides a Wavefront Cloud Integration for CloudWatch. This allows CloudWatch cloud integrations to be created,
        updated, and deleted.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_wavefront as wavefront

        ext_id = wavefront.CloudIntegrationAwsExternalId("ext_id")
        cloudwatch = wavefront.CloudIntegrationCloudWatch("cloudwatch",
            name="Test Integration",
            force_save=True,
            role_arn="arn:aws::1234567:role/example-arn",
            external_id=ext_id.id)
        ```

        ## Import

        CloudWatch Cloud Integrations can be imported by using the `id`, e.g.:

        ```sh
        $ pulumi import wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch cloudwatch a411c16b-3cf7-4f03-bf11-8ca05aab898d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] additional_tags: A list of point tag key-values to add to every point ingested using this integration.
        :param pulumi.Input[_builtins.str] external_id: The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
        :param pulumi.Input[_builtins.bool] force_save: Forces this resource to save, even if errors are present.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] instance_selection_tags: A string->string map allow list of instance tag-value pairs (in AWS).
               If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
               Multiple entries are OR'ed.
        :param pulumi.Input[_builtins.str] metric_filter_regex: A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
        :param pulumi.Input[_builtins.str] name: The human-readable name of this integration.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] namespaces: A list of namespaces that limit what we query from CloudWatch.
        :param pulumi.Input[_builtins.str] point_tag_filter_regex: A regular expression that AWS tag key name must match (case-insensitively)
               in order to be ingested.
        :param pulumi.Input[_builtins.str] role_arn: The external ID corresponding to the Role ARN.
        :param pulumi.Input[_builtins.str] service: A value denoting which cloud service this service integrates with.
        :param pulumi.Input[_builtins.int] service_refresh_rate_in_minutes: How often, in minutes, to refresh the service.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] volume_selection_tags: A string->string map of allow list of volume tag-value pairs (in AWS).
               If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
               Multiple entries are OR'ed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudIntegrationCloudWatchArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Wavefront Cloud Integration for CloudWatch. This allows CloudWatch cloud integrations to be created,
        updated, and deleted.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_wavefront as wavefront

        ext_id = wavefront.CloudIntegrationAwsExternalId("ext_id")
        cloudwatch = wavefront.CloudIntegrationCloudWatch("cloudwatch",
            name="Test Integration",
            force_save=True,
            role_arn="arn:aws::1234567:role/example-arn",
            external_id=ext_id.id)
        ```

        ## Import

        CloudWatch Cloud Integrations can be imported by using the `id`, e.g.:

        ```sh
        $ pulumi import wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch cloudwatch a411c16b-3cf7-4f03-bf11-8ca05aab898d
        ```

        :param str resource_name: The name of the resource.
        :param CloudIntegrationCloudWatchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudIntegrationCloudWatchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 external_id: Optional[pulumi.Input[_builtins.str]] = None,
                 force_save: Optional[pulumi.Input[_builtins.bool]] = None,
                 instance_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 metric_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespaces: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 point_tag_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
                 role_arn: Optional[pulumi.Input[_builtins.str]] = None,
                 service: Optional[pulumi.Input[_builtins.str]] = None,
                 service_refresh_rate_in_minutes: Optional[pulumi.Input[_builtins.int]] = None,
                 volume_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudIntegrationCloudWatchArgs.__new__(CloudIntegrationCloudWatchArgs)

            __props__.__dict__["additional_tags"] = additional_tags
            if external_id is None and not opts.urn:
                raise TypeError("Missing required property 'external_id'")
            __props__.__dict__["external_id"] = external_id
            __props__.__dict__["force_save"] = force_save
            __props__.__dict__["instance_selection_tags"] = instance_selection_tags
            __props__.__dict__["metric_filter_regex"] = metric_filter_regex
            __props__.__dict__["name"] = name
            __props__.__dict__["namespaces"] = namespaces
            __props__.__dict__["point_tag_filter_regex"] = point_tag_filter_regex
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            __props__.__dict__["service"] = service
            __props__.__dict__["service_refresh_rate_in_minutes"] = service_refresh_rate_in_minutes
            __props__.__dict__["volume_selection_tags"] = volume_selection_tags
        super(CloudIntegrationCloudWatch, __self__).__init__(
            'wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            external_id: Optional[pulumi.Input[_builtins.str]] = None,
            force_save: Optional[pulumi.Input[_builtins.bool]] = None,
            instance_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            metric_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            namespaces: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            point_tag_filter_regex: Optional[pulumi.Input[_builtins.str]] = None,
            role_arn: Optional[pulumi.Input[_builtins.str]] = None,
            service: Optional[pulumi.Input[_builtins.str]] = None,
            service_refresh_rate_in_minutes: Optional[pulumi.Input[_builtins.int]] = None,
            volume_selection_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None) -> 'CloudIntegrationCloudWatch':
        """
        Get an existing CloudIntegrationCloudWatch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] additional_tags: A list of point tag key-values to add to every point ingested using this integration.
        :param pulumi.Input[_builtins.str] external_id: The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
        :param pulumi.Input[_builtins.bool] force_save: Forces this resource to save, even if errors are present.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] instance_selection_tags: A string->string map allow list of instance tag-value pairs (in AWS).
               If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
               Multiple entries are OR'ed.
        :param pulumi.Input[_builtins.str] metric_filter_regex: A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
        :param pulumi.Input[_builtins.str] name: The human-readable name of this integration.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] namespaces: A list of namespaces that limit what we query from CloudWatch.
        :param pulumi.Input[_builtins.str] point_tag_filter_regex: A regular expression that AWS tag key name must match (case-insensitively)
               in order to be ingested.
        :param pulumi.Input[_builtins.str] role_arn: The external ID corresponding to the Role ARN.
        :param pulumi.Input[_builtins.str] service: A value denoting which cloud service this service integrates with.
        :param pulumi.Input[_builtins.int] service_refresh_rate_in_minutes: How often, in minutes, to refresh the service.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] volume_selection_tags: A string->string map of allow list of volume tag-value pairs (in AWS).
               If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
               Multiple entries are OR'ed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudIntegrationCloudWatchState.__new__(_CloudIntegrationCloudWatchState)

        __props__.__dict__["additional_tags"] = additional_tags
        __props__.__dict__["external_id"] = external_id
        __props__.__dict__["force_save"] = force_save
        __props__.__dict__["instance_selection_tags"] = instance_selection_tags
        __props__.__dict__["metric_filter_regex"] = metric_filter_regex
        __props__.__dict__["name"] = name
        __props__.__dict__["namespaces"] = namespaces
        __props__.__dict__["point_tag_filter_regex"] = point_tag_filter_regex
        __props__.__dict__["role_arn"] = role_arn
        __props__.__dict__["service"] = service
        __props__.__dict__["service_refresh_rate_in_minutes"] = service_refresh_rate_in_minutes
        __props__.__dict__["volume_selection_tags"] = volume_selection_tags
        return CloudIntegrationCloudWatch(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="additionalTags")
    def additional_tags(self) -> pulumi.Output[Optional[Mapping[str, _builtins.str]]]:
        """
        A list of point tag key-values to add to every point ingested using this integration.
        """
        return pulumi.get(self, "additional_tags")

    @_builtins.property
    @pulumi.getter(name="externalId")
    def external_id(self) -> pulumi.Output[_builtins.str]:
        """
        The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
        """
        return pulumi.get(self, "external_id")

    @_builtins.property
    @pulumi.getter(name="forceSave")
    def force_save(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Forces this resource to save, even if errors are present.
        """
        return pulumi.get(self, "force_save")

    @_builtins.property
    @pulumi.getter(name="instanceSelectionTags")
    def instance_selection_tags(self) -> pulumi.Output[Optional[Mapping[str, _builtins.str]]]:
        """
        A string->string map allow list of instance tag-value pairs (in AWS).
        If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.
        Multiple entries are OR'ed.
        """
        return pulumi.get(self, "instance_selection_tags")

    @_builtins.property
    @pulumi.getter(name="metricFilterRegex")
    def metric_filter_regex(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
        """
        return pulumi.get(self, "metric_filter_regex")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The human-readable name of this integration.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def namespaces(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        A list of namespaces that limit what we query from CloudWatch.
        """
        return pulumi.get(self, "namespaces")

    @_builtins.property
    @pulumi.getter(name="pointTagFilterRegex")
    def point_tag_filter_regex(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        A regular expression that AWS tag key name must match (case-insensitively)
        in order to be ingested.
        """
        return pulumi.get(self, "point_tag_filter_regex")

    @_builtins.property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[_builtins.str]:
        """
        The external ID corresponding to the Role ARN.
        """
        return pulumi.get(self, "role_arn")

    @_builtins.property
    @pulumi.getter
    def service(self) -> pulumi.Output[_builtins.str]:
        """
        A value denoting which cloud service this service integrates with.
        """
        return pulumi.get(self, "service")

    @_builtins.property
    @pulumi.getter(name="serviceRefreshRateInMinutes")
    def service_refresh_rate_in_minutes(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        How often, in minutes, to refresh the service.
        """
        return pulumi.get(self, "service_refresh_rate_in_minutes")

    @_builtins.property
    @pulumi.getter(name="volumeSelectionTags")
    def volume_selection_tags(self) -> pulumi.Output[Optional[Mapping[str, _builtins.str]]]:
        """
        A string->string map of allow list of volume tag-value pairs (in AWS).
        If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.
        Multiple entries are OR'ed.
        """
        return pulumi.get(self, "volume_selection_tags")

