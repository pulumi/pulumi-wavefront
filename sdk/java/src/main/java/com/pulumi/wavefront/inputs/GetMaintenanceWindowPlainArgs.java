// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetMaintenanceWindowPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMaintenanceWindowPlainArgs Empty = new GetMaintenanceWindowPlainArgs();

    /**
     * The ID of the maintenance window.
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return The ID of the maintenance window.
     * 
     */
    public String id() {
        return this.id;
    }

    private GetMaintenanceWindowPlainArgs() {}

    private GetMaintenanceWindowPlainArgs(GetMaintenanceWindowPlainArgs $) {
        this.id = $.id;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMaintenanceWindowPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMaintenanceWindowPlainArgs $;

        public Builder() {
            $ = new GetMaintenanceWindowPlainArgs();
        }

        public Builder(GetMaintenanceWindowPlainArgs defaults) {
            $ = new GetMaintenanceWindowPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The ID of the maintenance window.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        public GetMaintenanceWindowPlainArgs build() {
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            return $;
        }
    }

}