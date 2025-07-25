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

__all__ = ['EventArgs', 'Event']

@pulumi.input_type
class EventArgs:
    def __init__(__self__, *,
                 annotations: pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]],
                 endtime_key: Optional[pulumi.Input[_builtins.int]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 start_time: Optional[pulumi.Input[_builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        The set of arguments for constructing a Event resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] annotations: The annotations associated with the event.
        :param pulumi.Input[_builtins.str] name: The name of the event as it is displayed in Wavefront.
        :param pulumi.Input[_builtins.int] start_time: The start time of the event in epoch milliseconds.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] tags: A set of tags to assign to this resource.
        """
        pulumi.set(__self__, "annotations", annotations)
        if endtime_key is not None:
            pulumi.set(__self__, "endtime_key", endtime_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @_builtins.property
    @pulumi.getter
    def annotations(self) -> pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]:
        """
        The annotations associated with the event.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]):
        pulumi.set(self, "annotations", value)

    @_builtins.property
    @pulumi.getter(name="endtimeKey")
    def endtime_key(self) -> Optional[pulumi.Input[_builtins.int]]:
        return pulumi.get(self, "endtime_key")

    @endtime_key.setter
    def endtime_key(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "endtime_key", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the event as it is displayed in Wavefront.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The start time of the event in epoch milliseconds.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "start_time", value)

    @_builtins.property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A set of tags to assign to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _EventState:
    def __init__(__self__, *,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 endtime_key: Optional[pulumi.Input[_builtins.int]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 start_time: Optional[pulumi.Input[_builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering Event resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] annotations: The annotations associated with the event.
        :param pulumi.Input[_builtins.str] name: The name of the event as it is displayed in Wavefront.
        :param pulumi.Input[_builtins.int] start_time: The start time of the event in epoch milliseconds.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] tags: A set of tags to assign to this resource.
        """
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if endtime_key is not None:
            pulumi.set(__self__, "endtime_key", endtime_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @_builtins.property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        The annotations associated with the event.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "annotations", value)

    @_builtins.property
    @pulumi.getter(name="endtimeKey")
    def endtime_key(self) -> Optional[pulumi.Input[_builtins.int]]:
        return pulumi.get(self, "endtime_key")

    @endtime_key.setter
    def endtime_key(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "endtime_key", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the event as it is displayed in Wavefront.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The start time of the event in epoch milliseconds.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "start_time", value)

    @_builtins.property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A set of tags to assign to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.type_token("wavefront:index/event:Event")
class Event(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 endtime_key: Optional[pulumi.Input[_builtins.int]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 start_time: Optional[pulumi.Input[_builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        """
        Provides a Wavefront event resource. This allows events to be created, updated, and deleted.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_wavefront as wavefront

        event = wavefront.Event("event",
            name="terraform-test",
            annotations={
                "severity": "info",
                "type": "event type",
                "details": "description",
            },
            tags=["eventTag1"])
        ```

        ## Import

        You can import events by using the id, for example:

        ```sh
        $ pulumi import wavefront:index/event:Event event 1479868728473
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] annotations: The annotations associated with the event.
        :param pulumi.Input[_builtins.str] name: The name of the event as it is displayed in Wavefront.
        :param pulumi.Input[_builtins.int] start_time: The start time of the event in epoch milliseconds.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] tags: A set of tags to assign to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Wavefront event resource. This allows events to be created, updated, and deleted.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_wavefront as wavefront

        event = wavefront.Event("event",
            name="terraform-test",
            annotations={
                "severity": "info",
                "type": "event type",
                "details": "description",
            },
            tags=["eventTag1"])
        ```

        ## Import

        You can import events by using the id, for example:

        ```sh
        $ pulumi import wavefront:index/event:Event event 1479868728473
        ```

        :param str resource_name: The name of the resource.
        :param EventArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 endtime_key: Optional[pulumi.Input[_builtins.int]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 start_time: Optional[pulumi.Input[_builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventArgs.__new__(EventArgs)

            if annotations is None and not opts.urn:
                raise TypeError("Missing required property 'annotations'")
            __props__.__dict__["annotations"] = annotations
            __props__.__dict__["endtime_key"] = endtime_key
            __props__.__dict__["name"] = name
            __props__.__dict__["start_time"] = start_time
            __props__.__dict__["tags"] = tags
        super(Event, __self__).__init__(
            'wavefront:index/event:Event',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            endtime_key: Optional[pulumi.Input[_builtins.int]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            start_time: Optional[pulumi.Input[_builtins.int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None) -> 'Event':
        """
        Get an existing Event resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] annotations: The annotations associated with the event.
        :param pulumi.Input[_builtins.str] name: The name of the event as it is displayed in Wavefront.
        :param pulumi.Input[_builtins.int] start_time: The start time of the event in epoch milliseconds.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] tags: A set of tags to assign to this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventState.__new__(_EventState)

        __props__.__dict__["annotations"] = annotations
        __props__.__dict__["endtime_key"] = endtime_key
        __props__.__dict__["name"] = name
        __props__.__dict__["start_time"] = start_time
        __props__.__dict__["tags"] = tags
        return Event(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Mapping[str, _builtins.str]]:
        """
        The annotations associated with the event.
        """
        return pulumi.get(self, "annotations")

    @_builtins.property
    @pulumi.getter(name="endtimeKey")
    def endtime_key(self) -> pulumi.Output[Optional[_builtins.int]]:
        return pulumi.get(self, "endtime_key")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the event as it is displayed in Wavefront.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        The start time of the event in epoch milliseconds.
        """
        return pulumi.get(self, "start_time")

    @_builtins.property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        A set of tags to assign to this resource.
        """
        return pulumi.get(self, "tags")

