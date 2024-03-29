// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class AlertTargetRoute {
    /**
     * @return (Required) String that filters the route. Space delimited. Currently only allows a single key value pair.
     * (e.g. `env prod`)
     * 
     */
    private @Nullable Map<String,String> filter;
    /**
     * @return The notification method used for notification target. One of `WEBHOOK`, `EMAIL`, `PAGERDUTY`.
     * 
     */
    private String method;
    /**
     * @return (Required) The endpoint for the alert route. `EMAIL`: email address. `PAGERDUTY`: PagerDuty routing
     * key. `WEBHOOK`: URL endpoint.
     * 
     */
    private String target;

    private AlertTargetRoute() {}
    /**
     * @return (Required) String that filters the route. Space delimited. Currently only allows a single key value pair.
     * (e.g. `env prod`)
     * 
     */
    public Map<String,String> filter() {
        return this.filter == null ? Map.of() : this.filter;
    }
    /**
     * @return The notification method used for notification target. One of `WEBHOOK`, `EMAIL`, `PAGERDUTY`.
     * 
     */
    public String method() {
        return this.method;
    }
    /**
     * @return (Required) The endpoint for the alert route. `EMAIL`: email address. `PAGERDUTY`: PagerDuty routing
     * key. `WEBHOOK`: URL endpoint.
     * 
     */
    public String target() {
        return this.target;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AlertTargetRoute defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> filter;
        private String method;
        private String target;
        public Builder() {}
        public Builder(AlertTargetRoute defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filter = defaults.filter;
    	      this.method = defaults.method;
    	      this.target = defaults.target;
        }

        @CustomType.Setter
        public Builder filter(@Nullable Map<String,String> filter) {

            this.filter = filter;
            return this;
        }
        @CustomType.Setter
        public Builder method(String method) {
            if (method == null) {
              throw new MissingRequiredPropertyException("AlertTargetRoute", "method");
            }
            this.method = method;
            return this;
        }
        @CustomType.Setter
        public Builder target(String target) {
            if (target == null) {
              throw new MissingRequiredPropertyException("AlertTargetRoute", "target");
            }
            this.target = target;
            return this;
        }
        public AlertTargetRoute build() {
            final var _resultValue = new AlertTargetRoute();
            _resultValue.filter = filter;
            _resultValue.method = method;
            _resultValue.target = target;
            return _resultValue;
        }
    }
}
