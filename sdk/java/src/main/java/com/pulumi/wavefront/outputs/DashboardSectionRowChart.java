// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.wavefront.outputs.DashboardSectionRowChartChartSetting;
import com.pulumi.wavefront.outputs.DashboardSectionRowChartSource;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardSectionRowChart {
    /**
     * @return The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
     * 
     */
    private @Nullable Integer base;
    private @Nullable String chartAttribute;
    /**
     * @return Chart settings. See chart settings.
     * 
     */
    private DashboardSectionRowChartChartSetting chartSetting;
    /**
     * @return Description of the chart.
     * 
     */
    private @Nullable String description;
    /**
     * @return Name of the source.
     * 
     */
    private String name;
    /**
     * @return Query expression to plot on the chart. See chart source queries.
     * 
     */
    private List<DashboardSectionRowChartSource> sources;
    /**
     * @return Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
     * `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
     * 
     */
    private String summarization;
    /**
     * @return String to label the units of the chart on the Y-Axis.
     * 
     */
    private String units;

    private DashboardSectionRowChart() {}
    /**
     * @return The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale.
     * 
     */
    public Optional<Integer> base() {
        return Optional.ofNullable(this.base);
    }
    public Optional<String> chartAttribute() {
        return Optional.ofNullable(this.chartAttribute);
    }
    /**
     * @return Chart settings. See chart settings.
     * 
     */
    public DashboardSectionRowChartChartSetting chartSetting() {
        return this.chartSetting;
    }
    /**
     * @return Description of the chart.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return Name of the source.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Query expression to plot on the chart. See chart source queries.
     * 
     */
    public List<DashboardSectionRowChartSource> sources() {
        return this.sources;
    }
    /**
     * @return Summarization strategy for the chart. MEAN is default. Valid options are, `MEAN`,
     * `MEDIAN`, `MIN`, `MAX`, `SUM`, `COUNT`, `LAST`, `FIRST`.
     * 
     */
    public String summarization() {
        return this.summarization;
    }
    /**
     * @return String to label the units of the chart on the Y-Axis.
     * 
     */
    public String units() {
        return this.units;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardSectionRowChart defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer base;
        private @Nullable String chartAttribute;
        private DashboardSectionRowChartChartSetting chartSetting;
        private @Nullable String description;
        private String name;
        private List<DashboardSectionRowChartSource> sources;
        private String summarization;
        private String units;
        public Builder() {}
        public Builder(DashboardSectionRowChart defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.base = defaults.base;
    	      this.chartAttribute = defaults.chartAttribute;
    	      this.chartSetting = defaults.chartSetting;
    	      this.description = defaults.description;
    	      this.name = defaults.name;
    	      this.sources = defaults.sources;
    	      this.summarization = defaults.summarization;
    	      this.units = defaults.units;
        }

        @CustomType.Setter
        public Builder base(@Nullable Integer base) {
            this.base = base;
            return this;
        }
        @CustomType.Setter
        public Builder chartAttribute(@Nullable String chartAttribute) {
            this.chartAttribute = chartAttribute;
            return this;
        }
        @CustomType.Setter
        public Builder chartSetting(DashboardSectionRowChartChartSetting chartSetting) {
            this.chartSetting = Objects.requireNonNull(chartSetting);
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder sources(List<DashboardSectionRowChartSource> sources) {
            this.sources = Objects.requireNonNull(sources);
            return this;
        }
        public Builder sources(DashboardSectionRowChartSource... sources) {
            return sources(List.of(sources));
        }
        @CustomType.Setter
        public Builder summarization(String summarization) {
            this.summarization = Objects.requireNonNull(summarization);
            return this;
        }
        @CustomType.Setter
        public Builder units(String units) {
            this.units = Objects.requireNonNull(units);
            return this;
        }
        public DashboardSectionRowChart build() {
            final var o = new DashboardSectionRowChart();
            o.base = base;
            o.chartAttribute = chartAttribute;
            o.chartSetting = chartSetting;
            o.description = description;
            o.name = name;
            o.sources = sources;
            o.summarization = summarization;
            o.units = units;
            return o;
        }
    }
}