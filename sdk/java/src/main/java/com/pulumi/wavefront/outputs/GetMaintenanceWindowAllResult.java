// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.wavefront.outputs.GetMaintenanceWindowAllMaintenanceWindow;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetMaintenanceWindowAllResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable Integer limit;
    private List<GetMaintenanceWindowAllMaintenanceWindow> maintenanceWindows;
    private @Nullable Integer offset;

    private GetMaintenanceWindowAllResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<Integer> limit() {
        return Optional.ofNullable(this.limit);
    }
    public List<GetMaintenanceWindowAllMaintenanceWindow> maintenanceWindows() {
        return this.maintenanceWindows;
    }
    public Optional<Integer> offset() {
        return Optional.ofNullable(this.offset);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMaintenanceWindowAllResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable Integer limit;
        private List<GetMaintenanceWindowAllMaintenanceWindow> maintenanceWindows;
        private @Nullable Integer offset;
        public Builder() {}
        public Builder(GetMaintenanceWindowAllResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.limit = defaults.limit;
    	      this.maintenanceWindows = defaults.maintenanceWindows;
    	      this.offset = defaults.offset;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder limit(@Nullable Integer limit) {
            this.limit = limit;
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceWindows(List<GetMaintenanceWindowAllMaintenanceWindow> maintenanceWindows) {
            this.maintenanceWindows = Objects.requireNonNull(maintenanceWindows);
            return this;
        }
        public Builder maintenanceWindows(GetMaintenanceWindowAllMaintenanceWindow... maintenanceWindows) {
            return maintenanceWindows(List.of(maintenanceWindows));
        }
        @CustomType.Setter
        public Builder offset(@Nullable Integer offset) {
            this.offset = offset;
            return this;
        }
        public GetMaintenanceWindowAllResult build() {
            final var o = new GetMaintenanceWindowAllResult();
            o.id = id;
            o.limit = limit;
            o.maintenanceWindows = maintenanceWindows;
            o.offset = offset;
            return o;
        }
    }
}