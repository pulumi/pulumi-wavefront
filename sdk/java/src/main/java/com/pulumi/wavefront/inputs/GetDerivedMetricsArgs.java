// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDerivedMetricsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDerivedMetricsArgs Empty = new GetDerivedMetricsArgs();

    /**
     * Limit is the maximum number of results to be returned. Defaults to 100.
     * 
     */
    @Import(name="limit")
    private @Nullable Output<Integer> limit;

    /**
     * @return Limit is the maximum number of results to be returned. Defaults to 100.
     * 
     */
    public Optional<Output<Integer>> limit() {
        return Optional.ofNullable(this.limit);
    }

    /**
     * Offset is the offset from the first result to be returned. Defaults to 0.
     * 
     */
    @Import(name="offset")
    private @Nullable Output<Integer> offset;

    /**
     * @return Offset is the offset from the first result to be returned. Defaults to 0.
     * 
     */
    public Optional<Output<Integer>> offset() {
        return Optional.ofNullable(this.offset);
    }

    private GetDerivedMetricsArgs() {}

    private GetDerivedMetricsArgs(GetDerivedMetricsArgs $) {
        this.limit = $.limit;
        this.offset = $.offset;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDerivedMetricsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDerivedMetricsArgs $;

        public Builder() {
            $ = new GetDerivedMetricsArgs();
        }

        public Builder(GetDerivedMetricsArgs defaults) {
            $ = new GetDerivedMetricsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param limit Limit is the maximum number of results to be returned. Defaults to 100.
         * 
         * @return builder
         * 
         */
        public Builder limit(@Nullable Output<Integer> limit) {
            $.limit = limit;
            return this;
        }

        /**
         * @param limit Limit is the maximum number of results to be returned. Defaults to 100.
         * 
         * @return builder
         * 
         */
        public Builder limit(Integer limit) {
            return limit(Output.of(limit));
        }

        /**
         * @param offset Offset is the offset from the first result to be returned. Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder offset(@Nullable Output<Integer> offset) {
            $.offset = offset;
            return this;
        }

        /**
         * @param offset Offset is the offset from the first result to be returned. Defaults to 0.
         * 
         * @return builder
         * 
         */
        public Builder offset(Integer offset) {
            return offset(Output.of(offset));
        }

        public GetDerivedMetricsArgs build() {
            return $;
        }
    }

}
