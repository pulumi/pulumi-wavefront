// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ExternalLinkState extends com.pulumi.resources.ResourceArgs {

    public static final ExternalLinkState Empty = new ExternalLinkState();

    /**
     * Human-readable description for this link.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Human-readable description for this link.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Whether this is a &#34;Log Integration&#34; subType of external link.
     * 
     */
    @Import(name="isLogIntegration")
    private @Nullable Output<Boolean> isLogIntegration;

    /**
     * @return Whether this is a &#34;Log Integration&#34; subType of external link.
     * 
     */
    public Optional<Output<Boolean>> isLogIntegration() {
        return Optional.ofNullable(this.isLogIntegration);
    }

    /**
     * Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
     * 
     */
    @Import(name="metricFilterRegex")
    private @Nullable Output<String> metricFilterRegex;

    /**
     * @return Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
     * 
     */
    public Optional<Output<String>> metricFilterRegex() {
        return Optional.ofNullable(this.metricFilterRegex);
    }

    /**
     * The name of the external link.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the external link.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Controls whether a link is displayed in the context menu of a highlighted
     * series. This is a map from string to regular expression. The highlighted series must contain point tags whose
     * keys are present in the keys of this map and whose values match the regular expressions associated with those
     * keys in order for the link to be displayed.
     * 
     */
    @Import(name="pointTagFilterRegexes")
    private @Nullable Output<Map<String,String>> pointTagFilterRegexes;

    /**
     * @return Controls whether a link is displayed in the context menu of a highlighted
     * series. This is a map from string to regular expression. The highlighted series must contain point tags whose
     * keys are present in the keys of this map and whose values match the regular expressions associated with those
     * keys in order for the link to be displayed.
     * 
     */
    public Optional<Output<Map<String,String>>> pointTagFilterRegexes() {
        return Optional.ofNullable(this.pointTagFilterRegexes);
    }

    /**
     * Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
     * 
     */
    @Import(name="sourceFilterRegex")
    private @Nullable Output<String> sourceFilterRegex;

    /**
     * @return Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
     * 
     */
    public Optional<Output<String>> sourceFilterRegex() {
        return Optional.ofNullable(this.sourceFilterRegex);
    }

    /**
     * The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
     * 
     */
    @Import(name="template")
    private @Nullable Output<String> template;

    /**
     * @return The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
     * 
     */
    public Optional<Output<String>> template() {
        return Optional.ofNullable(this.template);
    }

    private ExternalLinkState() {}

    private ExternalLinkState(ExternalLinkState $) {
        this.description = $.description;
        this.isLogIntegration = $.isLogIntegration;
        this.metricFilterRegex = $.metricFilterRegex;
        this.name = $.name;
        this.pointTagFilterRegexes = $.pointTagFilterRegexes;
        this.sourceFilterRegex = $.sourceFilterRegex;
        this.template = $.template;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExternalLinkState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExternalLinkState $;

        public Builder() {
            $ = new ExternalLinkState();
        }

        public Builder(ExternalLinkState defaults) {
            $ = new ExternalLinkState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Human-readable description for this link.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Human-readable description for this link.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param isLogIntegration Whether this is a &#34;Log Integration&#34; subType of external link.
         * 
         * @return builder
         * 
         */
        public Builder isLogIntegration(@Nullable Output<Boolean> isLogIntegration) {
            $.isLogIntegration = isLogIntegration;
            return this;
        }

        /**
         * @param isLogIntegration Whether this is a &#34;Log Integration&#34; subType of external link.
         * 
         * @return builder
         * 
         */
        public Builder isLogIntegration(Boolean isLogIntegration) {
            return isLogIntegration(Output.of(isLogIntegration));
        }

        /**
         * @param metricFilterRegex Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
         * 
         * @return builder
         * 
         */
        public Builder metricFilterRegex(@Nullable Output<String> metricFilterRegex) {
            $.metricFilterRegex = metricFilterRegex;
            return this;
        }

        /**
         * @param metricFilterRegex Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.
         * 
         * @return builder
         * 
         */
        public Builder metricFilterRegex(String metricFilterRegex) {
            return metricFilterRegex(Output.of(metricFilterRegex));
        }

        /**
         * @param name The name of the external link.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the external link.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param pointTagFilterRegexes Controls whether a link is displayed in the context menu of a highlighted
         * series. This is a map from string to regular expression. The highlighted series must contain point tags whose
         * keys are present in the keys of this map and whose values match the regular expressions associated with those
         * keys in order for the link to be displayed.
         * 
         * @return builder
         * 
         */
        public Builder pointTagFilterRegexes(@Nullable Output<Map<String,String>> pointTagFilterRegexes) {
            $.pointTagFilterRegexes = pointTagFilterRegexes;
            return this;
        }

        /**
         * @param pointTagFilterRegexes Controls whether a link is displayed in the context menu of a highlighted
         * series. This is a map from string to regular expression. The highlighted series must contain point tags whose
         * keys are present in the keys of this map and whose values match the regular expressions associated with those
         * keys in order for the link to be displayed.
         * 
         * @return builder
         * 
         */
        public Builder pointTagFilterRegexes(Map<String,String> pointTagFilterRegexes) {
            return pointTagFilterRegexes(Output.of(pointTagFilterRegexes));
        }

        /**
         * @param sourceFilterRegex Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
         * 
         * @return builder
         * 
         */
        public Builder sourceFilterRegex(@Nullable Output<String> sourceFilterRegex) {
            $.sourceFilterRegex = sourceFilterRegex;
            return this;
        }

        /**
         * @param sourceFilterRegex Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.
         * 
         * @return builder
         * 
         */
        public Builder sourceFilterRegex(String sourceFilterRegex) {
            return sourceFilterRegex(Output.of(sourceFilterRegex));
        }

        /**
         * @param template The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
         * 
         * @return builder
         * 
         */
        public Builder template(@Nullable Output<String> template) {
            $.template = template;
            return this;
        }

        /**
         * @param template The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.
         * 
         * @return builder
         * 
         */
        public Builder template(String template) {
            return template(Output.of(template));
        }

        public ExternalLinkState build() {
            return $;
        }
    }

}
