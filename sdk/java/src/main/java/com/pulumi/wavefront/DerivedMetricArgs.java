// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DerivedMetricArgs extends com.pulumi.resources.ResourceArgs {

    public static final DerivedMetricArgs Empty = new DerivedMetricArgs();

    /**
     * User-supplied additional explanatory information for the derived metric.
     * 
     */
    @Import(name="additionalInformation")
    private @Nullable Output<String> additionalInformation;

    /**
     * @return User-supplied additional explanatory information for the derived metric.
     * 
     */
    public Optional<Output<String>> additionalInformation() {
        return Optional.ofNullable(this.additionalInformation);
    }

    /**
     * How frequently the query generating the derived metric is run.
     * 
     */
    @Import(name="minutes", required=true)
    private Output<Integer> minutes;

    /**
     * @return How frequently the query generating the derived metric is run.
     * 
     */
    public Output<Integer> minutes() {
        return this.minutes;
    }

    /**
     * The name of the Derived Metric in Wavefront.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Derived Metric in Wavefront.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A Wavefront query that is evaluated at regular intervals (default is 1 minute).
     * 
     */
    @Import(name="query", required=true)
    private Output<String> query;

    /**
     * @return A Wavefront query that is evaluated at regular intervals (default is 1 minute).
     * 
     */
    public Output<String> query() {
        return this.query;
    }

    /**
     * A set of tags to assign to this resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return A set of tags to assign to this resource.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private DerivedMetricArgs() {}

    private DerivedMetricArgs(DerivedMetricArgs $) {
        this.additionalInformation = $.additionalInformation;
        this.minutes = $.minutes;
        this.name = $.name;
        this.query = $.query;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DerivedMetricArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DerivedMetricArgs $;

        public Builder() {
            $ = new DerivedMetricArgs();
        }

        public Builder(DerivedMetricArgs defaults) {
            $ = new DerivedMetricArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param additionalInformation User-supplied additional explanatory information for the derived metric.
         * 
         * @return builder
         * 
         */
        public Builder additionalInformation(@Nullable Output<String> additionalInformation) {
            $.additionalInformation = additionalInformation;
            return this;
        }

        /**
         * @param additionalInformation User-supplied additional explanatory information for the derived metric.
         * 
         * @return builder
         * 
         */
        public Builder additionalInformation(String additionalInformation) {
            return additionalInformation(Output.of(additionalInformation));
        }

        /**
         * @param minutes How frequently the query generating the derived metric is run.
         * 
         * @return builder
         * 
         */
        public Builder minutes(Output<Integer> minutes) {
            $.minutes = minutes;
            return this;
        }

        /**
         * @param minutes How frequently the query generating the derived metric is run.
         * 
         * @return builder
         * 
         */
        public Builder minutes(Integer minutes) {
            return minutes(Output.of(minutes));
        }

        /**
         * @param name The name of the Derived Metric in Wavefront.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Derived Metric in Wavefront.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param query A Wavefront query that is evaluated at regular intervals (default is 1 minute).
         * 
         * @return builder
         * 
         */
        public Builder query(Output<String> query) {
            $.query = query;
            return this;
        }

        /**
         * @param query A Wavefront query that is evaluated at regular intervals (default is 1 minute).
         * 
         * @return builder
         * 
         */
        public Builder query(String query) {
            return query(Output.of(query));
        }

        /**
         * @param tags A set of tags to assign to this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A set of tags to assign to this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags A set of tags to assign to this resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public DerivedMetricArgs build() {
            $.minutes = Objects.requireNonNull($.minutes, "expected parameter 'minutes' to be non-null");
            $.query = Objects.requireNonNull($.query, "expected parameter 'query' to be non-null");
            return $;
        }
    }

}