// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudIntegrationGcpArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudIntegrationGcpArgs Empty = new CloudIntegrationGcpArgs();

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
     * A list of Google Cloud Platform (GCP) services. Valid values are `APPENGINE`,
     * `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
     * `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
     * `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
     * `TPU`, and `VPN`.
     * 
     */
    @Import(name="categories")
    private @Nullable Output<List<String>> categories;

    /**
     * @return A list of Google Cloud Platform (GCP) services. Valid values are `APPENGINE`,
     * `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
     * `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
     * `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
     * `TPU`, and `VPN`.
     * 
     */
    public Optional<Output<List<String>>> categories() {
        return Optional.ofNullable(this.categories);
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
     * Private key for a Google Cloud Platform (GCP) service account within your project.
     * The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
     * 
     */
    @Import(name="jsonKey", required=true)
    private Output<String> jsonKey;

    /**
     * @return Private key for a Google Cloud Platform (GCP) service account within your project.
     * The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
     * 
     */
    public Output<String> jsonKey() {
        return this.jsonKey;
    }

    /**
     * A regular expression that a metric name must match (case-insensitively) in order to be ingested.
     * 
     */
    @Import(name="metricFilterRegex")
    private @Nullable Output<String> metricFilterRegex;

    /**
     * @return A regular expression that a metric name must match (case-insensitively) in order to be ingested.
     * 
     */
    public Optional<Output<String>> metricFilterRegex() {
        return Optional.ofNullable(this.metricFilterRegex);
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
     * The Google Cloud Platform (GCP) Project ID.
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return The Google Cloud Platform (GCP) Project ID.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
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

    private CloudIntegrationGcpArgs() {}

    private CloudIntegrationGcpArgs(CloudIntegrationGcpArgs $) {
        this.additionalTags = $.additionalTags;
        this.categories = $.categories;
        this.forceSave = $.forceSave;
        this.jsonKey = $.jsonKey;
        this.metricFilterRegex = $.metricFilterRegex;
        this.name = $.name;
        this.projectId = $.projectId;
        this.service = $.service;
        this.serviceRefreshRateInMinutes = $.serviceRefreshRateInMinutes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudIntegrationGcpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudIntegrationGcpArgs $;

        public Builder() {
            $ = new CloudIntegrationGcpArgs();
        }

        public Builder(CloudIntegrationGcpArgs defaults) {
            $ = new CloudIntegrationGcpArgs(Objects.requireNonNull(defaults));
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
         * @param categories A list of Google Cloud Platform (GCP) services. Valid values are `APPENGINE`,
         * `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
         * `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
         * `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
         * `TPU`, and `VPN`.
         * 
         * @return builder
         * 
         */
        public Builder categories(@Nullable Output<List<String>> categories) {
            $.categories = categories;
            return this;
        }

        /**
         * @param categories A list of Google Cloud Platform (GCP) services. Valid values are `APPENGINE`,
         * `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
         * `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
         * `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
         * `TPU`, and `VPN`.
         * 
         * @return builder
         * 
         */
        public Builder categories(List<String> categories) {
            return categories(Output.of(categories));
        }

        /**
         * @param categories A list of Google Cloud Platform (GCP) services. Valid values are `APPENGINE`,
         * `BIGQUERY`, `BIGTABLE`, `CLOUDFUNCTIONS`, `CLOUDIOT`, `CLOUDSQL`, `CLOUDTASKS`, `COMPUTE`, `CONTAINER`,
         * `DATAFLOW`, `DATAPROC`, `DATASTORE`, `FIREBASEDATABASE`, `FIREBASEHOSTING`, `FIRESTORE`, `INTERCONNECT`,
         * `LOADBALANCING`, `LOGGING`, `ML`, `MONITORING`, `PUBSUB`, `REDIS`, `ROUTER`, `SERVICERUNTIME`, `SPANNER`, `STORAGE`,
         * `TPU`, and `VPN`.
         * 
         * @return builder
         * 
         */
        public Builder categories(String... categories) {
            return categories(List.of(categories));
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
         * @param jsonKey Private key for a Google Cloud Platform (GCP) service account within your project.
         * The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
         * 
         * @return builder
         * 
         */
        public Builder jsonKey(Output<String> jsonKey) {
            $.jsonKey = jsonKey;
            return this;
        }

        /**
         * @param jsonKey Private key for a Google Cloud Platform (GCP) service account within your project.
         * The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.
         * 
         * @return builder
         * 
         */
        public Builder jsonKey(String jsonKey) {
            return jsonKey(Output.of(jsonKey));
        }

        /**
         * @param metricFilterRegex A regular expression that a metric name must match (case-insensitively) in order to be ingested.
         * 
         * @return builder
         * 
         */
        public Builder metricFilterRegex(@Nullable Output<String> metricFilterRegex) {
            $.metricFilterRegex = metricFilterRegex;
            return this;
        }

        /**
         * @param metricFilterRegex A regular expression that a metric name must match (case-insensitively) in order to be ingested.
         * 
         * @return builder
         * 
         */
        public Builder metricFilterRegex(String metricFilterRegex) {
            return metricFilterRegex(Output.of(metricFilterRegex));
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
         * @param projectId The Google Cloud Platform (GCP) Project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The Google Cloud Platform (GCP) Project ID.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
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

        public CloudIntegrationGcpArgs build() {
            if ($.jsonKey == null) {
                throw new MissingRequiredPropertyException("CloudIntegrationGcpArgs", "jsonKey");
            }
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("CloudIntegrationGcpArgs", "projectId");
            }
            return $;
        }
    }

}
