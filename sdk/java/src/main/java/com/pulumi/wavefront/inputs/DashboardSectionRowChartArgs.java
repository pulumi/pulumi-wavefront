// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.wavefront.inputs.DashboardSectionRowChartChartSettingArgs;
import com.pulumi.wavefront.inputs.DashboardSectionRowChartSourceArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardSectionRowChartArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardSectionRowChartArgs Empty = new DashboardSectionRowChartArgs();

    /**
     * The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
     * 
     */
    @Import(name="base")
    private @Nullable Output<Integer> base;

    /**
     * @return The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
     * 
     */
    public Optional<Output<Integer>> base() {
        return Optional.ofNullable(this.base);
    }

    @Import(name="chartAttribute")
    private @Nullable Output<String> chartAttribute;

    public Optional<Output<String>> chartAttribute() {
        return Optional.ofNullable(this.chartAttribute);
    }

    /**
     * Chart settings. See chart settings.
     * 
     */
    @Import(name="chartSetting", required=true)
    private Output<DashboardSectionRowChartChartSettingArgs> chartSetting;

    /**
     * @return Chart settings. See chart settings.
     * 
     */
    public Output<DashboardSectionRowChartChartSettingArgs> chartSetting() {
        return this.chartSetting;
    }

    /**
     * Description of the chart.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the chart.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name of the source.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the source.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Show events related to the sources included in queries
     * 
     */
    @Import(name="noDefaultEvents")
    private @Nullable Output<Boolean> noDefaultEvents;

    /**
     * @return Show events related to the sources included in queries
     * 
     */
    public Optional<Output<Boolean>> noDefaultEvents() {
        return Optional.ofNullable(this.noDefaultEvents);
    }

    /**
     * Query expression to plot on the chart. See chart source queries.
     * 
     */
    @Import(name="sources", required=true)
    private Output<List<DashboardSectionRowChartSourceArgs>> sources;

    /**
     * @return Query expression to plot on the chart. See chart source queries.
     * 
     */
    public Output<List<DashboardSectionRowChartSourceArgs>> sources() {
        return this.sources;
    }

    /**
     * Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
     * `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
     * 
     */
    @Import(name="summarization", required=true)
    private Output<String> summarization;

    /**
     * @return Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
     * `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
     * 
     */
    public Output<String> summarization() {
        return this.summarization;
    }

    /**
     * String to label the units of the chart on the Y-Axis.
     * 
     */
    @Import(name="units", required=true)
    private Output<String> units;

    /**
     * @return String to label the units of the chart on the Y-Axis.
     * 
     */
    public Output<String> units() {
        return this.units;
    }

    private DashboardSectionRowChartArgs() {}

    private DashboardSectionRowChartArgs(DashboardSectionRowChartArgs $) {
        this.base = $.base;
        this.chartAttribute = $.chartAttribute;
        this.chartSetting = $.chartSetting;
        this.description = $.description;
        this.name = $.name;
        this.noDefaultEvents = $.noDefaultEvents;
        this.sources = $.sources;
        this.summarization = $.summarization;
        this.units = $.units;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardSectionRowChartArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardSectionRowChartArgs $;

        public Builder() {
            $ = new DashboardSectionRowChartArgs();
        }

        public Builder(DashboardSectionRowChartArgs defaults) {
            $ = new DashboardSectionRowChartArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param base The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
         * 
         * @return builder
         * 
         */
        public Builder base(@Nullable Output<Integer> base) {
            $.base = base;
            return this;
        }

        /**
         * @param base The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
         * 
         * @return builder
         * 
         */
        public Builder base(Integer base) {
            return base(Output.of(base));
        }

        public Builder chartAttribute(@Nullable Output<String> chartAttribute) {
            $.chartAttribute = chartAttribute;
            return this;
        }

        public Builder chartAttribute(String chartAttribute) {
            return chartAttribute(Output.of(chartAttribute));
        }

        /**
         * @param chartSetting Chart settings. See chart settings.
         * 
         * @return builder
         * 
         */
        public Builder chartSetting(Output<DashboardSectionRowChartChartSettingArgs> chartSetting) {
            $.chartSetting = chartSetting;
            return this;
        }

        /**
         * @param chartSetting Chart settings. See chart settings.
         * 
         * @return builder
         * 
         */
        public Builder chartSetting(DashboardSectionRowChartChartSettingArgs chartSetting) {
            return chartSetting(Output.of(chartSetting));
        }

        /**
         * @param description Description of the chart.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the chart.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Name of the source.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the source.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param noDefaultEvents Show events related to the sources included in queries
         * 
         * @return builder
         * 
         */
        public Builder noDefaultEvents(@Nullable Output<Boolean> noDefaultEvents) {
            $.noDefaultEvents = noDefaultEvents;
            return this;
        }

        /**
         * @param noDefaultEvents Show events related to the sources included in queries
         * 
         * @return builder
         * 
         */
        public Builder noDefaultEvents(Boolean noDefaultEvents) {
            return noDefaultEvents(Output.of(noDefaultEvents));
        }

        /**
         * @param sources Query expression to plot on the chart. See chart source queries.
         * 
         * @return builder
         * 
         */
        public Builder sources(Output<List<DashboardSectionRowChartSourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        /**
         * @param sources Query expression to plot on the chart. See chart source queries.
         * 
         * @return builder
         * 
         */
        public Builder sources(List<DashboardSectionRowChartSourceArgs> sources) {
            return sources(Output.of(sources));
        }

        /**
         * @param sources Query expression to plot on the chart. See chart source queries.
         * 
         * @return builder
         * 
         */
        public Builder sources(DashboardSectionRowChartSourceArgs... sources) {
            return sources(List.of(sources));
        }

        /**
         * @param summarization Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
         * `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
         * 
         * @return builder
         * 
         */
        public Builder summarization(Output<String> summarization) {
            $.summarization = summarization;
            return this;
        }

        /**
         * @param summarization Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
         * `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
         * 
         * @return builder
         * 
         */
        public Builder summarization(String summarization) {
            return summarization(Output.of(summarization));
        }

        /**
         * @param units String to label the units of the chart on the Y-Axis.
         * 
         * @return builder
         * 
         */
        public Builder units(Output<String> units) {
            $.units = units;
            return this;
        }

        /**
         * @param units String to label the units of the chart on the Y-Axis.
         * 
         * @return builder
         * 
         */
        public Builder units(String units) {
            return units(Output.of(units));
        }

        public DashboardSectionRowChartArgs build() {
            if ($.chartSetting == null) {
                throw new MissingRequiredPropertyException("DashboardSectionRowChartArgs", "chartSetting");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("DashboardSectionRowChartArgs", "name");
            }
            if ($.sources == null) {
                throw new MissingRequiredPropertyException("DashboardSectionRowChartArgs", "sources");
            }
            if ($.summarization == null) {
                throw new MissingRequiredPropertyException("DashboardSectionRowChartArgs", "summarization");
            }
            if ($.units == null) {
                throw new MissingRequiredPropertyException("DashboardSectionRowChartArgs", "units");
            }
            return $;
        }
    }

}
