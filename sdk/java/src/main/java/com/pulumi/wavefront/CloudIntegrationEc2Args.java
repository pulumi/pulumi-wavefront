// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudIntegrationEc2Args extends com.pulumi.resources.ResourceArgs {

    public static final CloudIntegrationEc2Args Empty = new CloudIntegrationEc2Args();

    /**
     * A list of point tag key-values to add to every point ingested using this integration.
     * 
     */
    @Import(name="additionalTags")
    private @Nullable Output<Map<String,String>> additionalTags;

    /**
     * @return A list of point tag key-values to add to every point ingested using this integration.
     * 
     */
    public Optional<Output<Map<String,String>>> additionalTags() {
        return Optional.ofNullable(this.additionalTags);
    }

    /**
     * The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
     * 
     */
    @Import(name="externalId", required=true)
    private Output<String> externalId;

    /**
     * @return The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
     * 
     */
    public Output<String> externalId() {
        return this.externalId;
    }

    /**
     * Forces this resource to save, even if errors are present.
     * 
     */
    @Import(name="forceSave")
    private @Nullable Output<Boolean> forceSave;

    /**
     * @return Forces this resource to save, even if errors are present.
     * 
     */
    public Optional<Output<Boolean>> forceSave() {
        return Optional.ofNullable(this.forceSave);
    }

    /**
     * A list of AWS instance tags to use as the `source` name
     * in a series. Default is `[&#34;hostname&#34;, &#34;host&#34;, &#34;name&#34;]`. If no tag in the list is found, the series source
     * is set to the instance id.
     * 
     */
    @Import(name="hostnameTags")
    private @Nullable Output<List<String>> hostnameTags;

    /**
     * @return A list of AWS instance tags to use as the `source` name
     * in a series. Default is `[&#34;hostname&#34;, &#34;host&#34;, &#34;name&#34;]`. If no tag in the list is found, the series source
     * is set to the instance id.
     * 
     */
    public Optional<Output<List<String>>> hostnameTags() {
        return Optional.ofNullable(this.hostnameTags);
    }

    /**
     * The human-readable name of this integration.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The human-readable name of this integration.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The external ID corresponding to the Role ARN.
     * 
     */
    @Import(name="roleArn", required=true)
    private Output<String> roleArn;

    /**
     * @return The external ID corresponding to the Role ARN.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }

    /**
     * A value denoting which cloud service this service integrates with.
     * 
     */
    @Import(name="service", required=true)
    private Output<String> service;

    /**
     * @return A value denoting which cloud service this service integrates with.
     * 
     */
    public Output<String> service() {
        return this.service;
    }

    /**
     * How often, in minutes, to refresh the service.
     * 
     */
    @Import(name="serviceRefreshRateInMinutes")
    private @Nullable Output<Integer> serviceRefreshRateInMinutes;

    /**
     * @return How often, in minutes, to refresh the service.
     * 
     */
    public Optional<Output<Integer>> serviceRefreshRateInMinutes() {
        return Optional.ofNullable(this.serviceRefreshRateInMinutes);
    }

    private CloudIntegrationEc2Args() {}

    private CloudIntegrationEc2Args(CloudIntegrationEc2Args $) {
        this.additionalTags = $.additionalTags;
        this.externalId = $.externalId;
        this.forceSave = $.forceSave;
        this.hostnameTags = $.hostnameTags;
        this.name = $.name;
        this.roleArn = $.roleArn;
        this.service = $.service;
        this.serviceRefreshRateInMinutes = $.serviceRefreshRateInMinutes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudIntegrationEc2Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudIntegrationEc2Args $;

        public Builder() {
            $ = new CloudIntegrationEc2Args();
        }

        public Builder(CloudIntegrationEc2Args defaults) {
            $ = new CloudIntegrationEc2Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalTags A list of point tag key-values to add to every point ingested using this integration.
         * 
         * @return builder
         * 
         */
        public Builder additionalTags(@Nullable Output<Map<String,String>> additionalTags) {
            $.additionalTags = additionalTags;
            return this;
        }

        /**
         * @param additionalTags A list of point tag key-values to add to every point ingested using this integration.
         * 
         * @return builder
         * 
         */
        public Builder additionalTags(Map<String,String> additionalTags) {
            return additionalTags(Output.of(additionalTags));
        }

        /**
         * @param externalId The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
         * 
         * @return builder
         * 
         */
        public Builder externalId(Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param forceSave Forces this resource to save, even if errors are present.
         * 
         * @return builder
         * 
         */
        public Builder forceSave(@Nullable Output<Boolean> forceSave) {
            $.forceSave = forceSave;
            return this;
        }

        /**
         * @param forceSave Forces this resource to save, even if errors are present.
         * 
         * @return builder
         * 
         */
        public Builder forceSave(Boolean forceSave) {
            return forceSave(Output.of(forceSave));
        }

        /**
         * @param hostnameTags A list of AWS instance tags to use as the `source` name
         * in a series. Default is `[&#34;hostname&#34;, &#34;host&#34;, &#34;name&#34;]`. If no tag in the list is found, the series source
         * is set to the instance id.
         * 
         * @return builder
         * 
         */
        public Builder hostnameTags(@Nullable Output<List<String>> hostnameTags) {
            $.hostnameTags = hostnameTags;
            return this;
        }

        /**
         * @param hostnameTags A list of AWS instance tags to use as the `source` name
         * in a series. Default is `[&#34;hostname&#34;, &#34;host&#34;, &#34;name&#34;]`. If no tag in the list is found, the series source
         * is set to the instance id.
         * 
         * @return builder
         * 
         */
        public Builder hostnameTags(List<String> hostnameTags) {
            return hostnameTags(Output.of(hostnameTags));
        }

        /**
         * @param hostnameTags A list of AWS instance tags to use as the `source` name
         * in a series. Default is `[&#34;hostname&#34;, &#34;host&#34;, &#34;name&#34;]`. If no tag in the list is found, the series source
         * is set to the instance id.
         * 
         * @return builder
         * 
         */
        public Builder hostnameTags(String... hostnameTags) {
            return hostnameTags(List.of(hostnameTags));
        }

        /**
         * @param name The human-readable name of this integration.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The human-readable name of this integration.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param roleArn The external ID corresponding to the Role ARN.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn The external ID corresponding to the Role ARN.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param service A value denoting which cloud service this service integrates with.
         * 
         * @return builder
         * 
         */
        public Builder service(Output<String> service) {
            $.service = service;
            return this;
        }

        /**
         * @param service A value denoting which cloud service this service integrates with.
         * 
         * @return builder
         * 
         */
        public Builder service(String service) {
            return service(Output.of(service));
        }

        /**
         * @param serviceRefreshRateInMinutes How often, in minutes, to refresh the service.
         * 
         * @return builder
         * 
         */
        public Builder serviceRefreshRateInMinutes(@Nullable Output<Integer> serviceRefreshRateInMinutes) {
            $.serviceRefreshRateInMinutes = serviceRefreshRateInMinutes;
            return this;
        }

        /**
         * @param serviceRefreshRateInMinutes How often, in minutes, to refresh the service.
         * 
         * @return builder
         * 
         */
        public Builder serviceRefreshRateInMinutes(Integer serviceRefreshRateInMinutes) {
            return serviceRefreshRateInMinutes(Output.of(serviceRefreshRateInMinutes));
        }

        public CloudIntegrationEc2Args build() {
            $.externalId = Objects.requireNonNull($.externalId, "expected parameter 'externalId' to be non-null");
            $.roleArn = Objects.requireNonNull($.roleArn, "expected parameter 'roleArn' to be non-null");
            $.service = Objects.requireNonNull($.service, "expected parameter 'service' to be non-null");
            return $;
        }
    }

}