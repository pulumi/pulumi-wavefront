// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.wavefront.outputs.GetDashboardsDashboardSectionRow;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDashboardsDashboardSection {
    /**
     * @return The name of the parameters.
     * 
     */
    private String name;
    private List<GetDashboardsDashboardSectionRow> rows;

    private GetDashboardsDashboardSection() {}
    /**
     * @return The name of the parameters.
     * 
     */
    public String name() {
        return this.name;
    }
    public List<GetDashboardsDashboardSectionRow> rows() {
        return this.rows;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDashboardsDashboardSection defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private List<GetDashboardsDashboardSectionRow> rows;
        public Builder() {}
        public Builder(GetDashboardsDashboardSection defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.rows = defaults.rows;
        }

        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder rows(List<GetDashboardsDashboardSectionRow> rows) {
            this.rows = Objects.requireNonNull(rows);
            return this;
        }
        public Builder rows(GetDashboardsDashboardSectionRow... rows) {
            return rows(List.of(rows));
        }
        public GetDashboardsDashboardSection build() {
            final var o = new GetDashboardsDashboardSection();
            o.name = name;
            o.rows = rows;
            return o;
        }
    }
}