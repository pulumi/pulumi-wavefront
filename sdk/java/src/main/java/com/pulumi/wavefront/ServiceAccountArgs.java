// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceAccountArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceAccountArgs Empty = new ServiceAccountArgs();

    /**
     * Whether or not the service account is active.
     * 
     */
    @Import(name="active")
    private @Nullable Output<Boolean> active;

    /**
     * @return Whether or not the service account is active.
     * 
     */
    public Optional<Output<Boolean>> active() {
        return Optional.ofNullable(this.active);
    }

    /**
     * The description of the service account.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the service account.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The unique identifier of the service account to create. Must have the prefix `sa::`.
     * 
     */
    @Import(name="identifier", required=true)
    private Output<String> identifier;

    /**
     * @return The unique identifier of the service account to create. Must have the prefix `sa::`.
     * 
     */
    public Output<String> identifier() {
        return this.identifier;
    }

    /**
     * ID of ingestion policy.
     * 
     */
    @Import(name="ingestionPolicy")
    private @Nullable Output<String> ingestionPolicy;

    /**
     * @return ID of ingestion policy.
     * 
     */
    public Optional<Output<String>> ingestionPolicy() {
        return Optional.ofNullable(this.ingestionPolicy);
    }

    /**
     * List of permission to grant to this service account. Valid options are
     * `agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
     * `host_tag_management`, `metrics_management`, and `user_management`.
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<List<String>> permissions;

    /**
     * @return List of permission to grant to this service account. Valid options are
     * `agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
     * `host_tag_management`, `metrics_management`, and `user_management`.
     * 
     */
    public Optional<Output<List<String>>> permissions() {
        return Optional.ofNullable(this.permissions);
    }

    /**
     * List of user groups for this service account.
     * 
     */
    @Import(name="userGroups")
    private @Nullable Output<List<String>> userGroups;

    /**
     * @return List of user groups for this service account.
     * 
     */
    public Optional<Output<List<String>>> userGroups() {
        return Optional.ofNullable(this.userGroups);
    }

    private ServiceAccountArgs() {}

    private ServiceAccountArgs(ServiceAccountArgs $) {
        this.active = $.active;
        this.description = $.description;
        this.identifier = $.identifier;
        this.ingestionPolicy = $.ingestionPolicy;
        this.permissions = $.permissions;
        this.userGroups = $.userGroups;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceAccountArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceAccountArgs $;

        public Builder() {
            $ = new ServiceAccountArgs();
        }

        public Builder(ServiceAccountArgs defaults) {
            $ = new ServiceAccountArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param active Whether or not the service account is active.
         * 
         * @return builder
         * 
         */
        public Builder active(@Nullable Output<Boolean> active) {
            $.active = active;
            return this;
        }

        /**
         * @param active Whether or not the service account is active.
         * 
         * @return builder
         * 
         */
        public Builder active(Boolean active) {
            return active(Output.of(active));
        }

        /**
         * @param description The description of the service account.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the service account.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param identifier The unique identifier of the service account to create. Must have the prefix `sa::`.
         * 
         * @return builder
         * 
         */
        public Builder identifier(Output<String> identifier) {
            $.identifier = identifier;
            return this;
        }

        /**
         * @param identifier The unique identifier of the service account to create. Must have the prefix `sa::`.
         * 
         * @return builder
         * 
         */
        public Builder identifier(String identifier) {
            return identifier(Output.of(identifier));
        }

        /**
         * @param ingestionPolicy ID of ingestion policy.
         * 
         * @return builder
         * 
         */
        public Builder ingestionPolicy(@Nullable Output<String> ingestionPolicy) {
            $.ingestionPolicy = ingestionPolicy;
            return this;
        }

        /**
         * @param ingestionPolicy ID of ingestion policy.
         * 
         * @return builder
         * 
         */
        public Builder ingestionPolicy(String ingestionPolicy) {
            return ingestionPolicy(Output.of(ingestionPolicy));
        }

        /**
         * @param permissions List of permission to grant to this service account. Valid options are
         * `agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
         * `host_tag_management`, `metrics_management`, and `user_management`.
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<List<String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions List of permission to grant to this service account. Valid options are
         * `agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
         * `host_tag_management`, `metrics_management`, and `user_management`.
         * 
         * @return builder
         * 
         */
        public Builder permissions(List<String> permissions) {
            return permissions(Output.of(permissions));
        }

        /**
         * @param permissions List of permission to grant to this service account. Valid options are
         * `agent_management`, `alerts_management`, `dashboard_management`, `embedded_charts`, `events_management`, `external_links_management`,
         * `host_tag_management`, `metrics_management`, and `user_management`.
         * 
         * @return builder
         * 
         */
        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }

        /**
         * @param userGroups List of user groups for this service account.
         * 
         * @return builder
         * 
         */
        public Builder userGroups(@Nullable Output<List<String>> userGroups) {
            $.userGroups = userGroups;
            return this;
        }

        /**
         * @param userGroups List of user groups for this service account.
         * 
         * @return builder
         * 
         */
        public Builder userGroups(List<String> userGroups) {
            return userGroups(Output.of(userGroups));
        }

        /**
         * @param userGroups List of user groups for this service account.
         * 
         * @return builder
         * 
         */
        public Builder userGroups(String... userGroups) {
            return userGroups(List.of(userGroups));
        }

        public ServiceAccountArgs build() {
            if ($.identifier == null) {
                throw new MissingRequiredPropertyException("ServiceAccountArgs", "identifier");
            }
            return $;
        }
    }

}
