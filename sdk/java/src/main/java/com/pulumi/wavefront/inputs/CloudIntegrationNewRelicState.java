// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.wavefront.inputs.CloudIntegrationNewRelicMetricFilterArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudIntegrationNewRelicState extends com.pulumi.resources.ResourceArgs {

    public static final CloudIntegrationNewRelicState Empty = new CloudIntegrationNewRelicState();

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
     * New Relic REST API key.
     * 
     */
    @Import(name="apiKey")
    private @Nullable Output<String> apiKey;

    /**
     * @return New Relic REST API key.
     * 
     */
    public Optional<Output<String>> apiKey() {
        return Optional.ofNullable(this.apiKey);
    }

    /**
     * A regular expression that an application name must match (case-insensitively) in order to collect metrics.
     * 
     */
    @Import(name="appFilterRegex")
    private @Nullable Output<String> appFilterRegex;

    /**
     * @return A regular expression that an application name must match (case-insensitively) in order to collect metrics.
     * 
     */
    public Optional<Output<String>> appFilterRegex() {
        return Optional.ofNullable(this.appFilterRegex);
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
     * A regular expression that a host name must match (case-insensitively) in order to collect metrics.
     * 
     */
    @Import(name="hostFilterRegex")
    private @Nullable Output<String> hostFilterRegex;

    /**
     * @return A regular expression that a host name must match (case-insensitively) in order to collect metrics.
     * 
     */
    public Optional<Output<String>> hostFilterRegex() {
        return Optional.ofNullable(this.hostFilterRegex);
    }

    /**
     * See Metric Filter.
     * 
     */
    @Import(name="metricFilters")
    private @Nullable Output<List<CloudIntegrationNewRelicMetricFilterArgs>> metricFilters;

    /**
     * @return See Metric Filter.
     * 
     */
    public Optional<Output<List<CloudIntegrationNewRelicMetricFilterArgs>>> metricFilters() {
        return Optional.ofNullable(this.metricFilters);
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
     * A value denoting which cloud service this service integrates with.
     * 
     */
    @Import(name="service")
    private @Nullable Output<String> service;

    /**
     * @return A value denoting which cloud service this service integrates with.
     * 
     */
    public Optional<Output<String>> service() {
        return Optional.ofNullable(this.service);
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

    private CloudIntegrationNewRelicState() {}

    private CloudIntegrationNewRelicState(CloudIntegrationNewRelicState $) {
        this.additionalTags = $.additionalTags;
        this.apiKey = $.apiKey;
        this.appFilterRegex = $.appFilterRegex;
        this.forceSave = $.forceSave;
        this.hostFilterRegex = $.hostFilterRegex;
        this.metricFilters = $.metricFilters;
        this.name = $.name;
        this.service = $.service;
        this.serviceRefreshRateInMinutes = $.serviceRefreshRateInMinutes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudIntegrationNewRelicState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudIntegrationNewRelicState $;

        public Builder() {
            $ = new CloudIntegrationNewRelicState();
        }

        public Builder(CloudIntegrationNewRelicState defaults) {
            $ = new CloudIntegrationNewRelicState(Objects.requireNonNull(defaults));
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
         * @param apiKey New Relic REST API key.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(@Nullable Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey New Relic REST API key.
         * 
         * @return builder
         * 
         */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
        }

        /**
         * @param appFilterRegex A regular expression that an application name must match (case-insensitively) in order to collect metrics.
         * 
         * @return builder
         * 
         */
        public Builder appFilterRegex(@Nullable Output<String> appFilterRegex) {
            $.appFilterRegex = appFilterRegex;
            return this;
        }

        /**
         * @param appFilterRegex A regular expression that an application name must match (case-insensitively) in order to collect metrics.
         * 
         * @return builder
         * 
         */
        public Builder appFilterRegex(String appFilterRegex) {
            return appFilterRegex(Output.of(appFilterRegex));
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
         * @param hostFilterRegex A regular expression that a host name must match (case-insensitively) in order to collect metrics.
         * 
         * @return builder
         * 
         */
        public Builder hostFilterRegex(@Nullable Output<String> hostFilterRegex) {
            $.hostFilterRegex = hostFilterRegex;
            return this;
        }

        /**
         * @param hostFilterRegex A regular expression that a host name must match (case-insensitively) in order to collect metrics.
         * 
         * @return builder
         * 
         */
        public Builder hostFilterRegex(String hostFilterRegex) {
            return hostFilterRegex(Output.of(hostFilterRegex));
        }

        /**
         * @param metricFilters See Metric Filter.
         * 
         * @return builder
         * 
         */
        public Builder metricFilters(@Nullable Output<List<CloudIntegrationNewRelicMetricFilterArgs>> metricFilters) {
            $.metricFilters = metricFilters;
            return this;
        }

        /**
         * @param metricFilters See Metric Filter.
         * 
         * @return builder
         * 
         */
        public Builder metricFilters(List<CloudIntegrationNewRelicMetricFilterArgs> metricFilters) {
            return metricFilters(Output.of(metricFilters));
        }

        /**
         * @param metricFilters See Metric Filter.
         * 
         * @return builder
         * 
         */
        public Builder metricFilters(CloudIntegrationNewRelicMetricFilterArgs... metricFilters) {
            return metricFilters(List.of(metricFilters));
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
         * @param service A value denoting which cloud service this service integrates with.
         * 
         * @return builder
         * 
         */
        public Builder service(@Nullable Output<String> service) {
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

        public CloudIntegrationNewRelicState build() {
            return $;
        }
    }

}