// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetEventResult {
    /**
     * @return Annotations associated with the event.
     * 
     */
    private Map<String,String> annotations;
    /**
     * @return The description of the event.
     * 
     */
    private String details;
    private Integer endtimeKey;
    /**
     * @return The ID of the event in Wavefront.
     * 
     */
    private String id;
    /**
     * @return A Boolean flag. If set to `true`, creates a point-in-time event (i.e. with no duration).
     * 
     */
    private Boolean isEphemeral;
    /**
     * @return The name of the event in Wavefront.
     * 
     */
    private String name;
    /**
     * @return The severity category of the event.
     * 
     */
    private String severity;
    /**
     * @return The start time of the event in epoch milliseconds.
     * 
     */
    private Integer startTime;
    /**
     * @return A set of tags assigned to the event.
     * 
     */
    private List<String> tags;
    /**
     * @return The type of the event.
     * 
     */
    private String type;

    private GetEventResult() {}
    /**
     * @return Annotations associated with the event.
     * 
     */
    public Map<String,String> annotations() {
        return this.annotations;
    }
    /**
     * @return The description of the event.
     * 
     */
    public String details() {
        return this.details;
    }
    public Integer endtimeKey() {
        return this.endtimeKey;
    }
    /**
     * @return The ID of the event in Wavefront.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A Boolean flag. If set to `true`, creates a point-in-time event (i.e. with no duration).
     * 
     */
    public Boolean isEphemeral() {
        return this.isEphemeral;
    }
    /**
     * @return The name of the event in Wavefront.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The severity category of the event.
     * 
     */
    public String severity() {
        return this.severity;
    }
    /**
     * @return The start time of the event in epoch milliseconds.
     * 
     */
    public Integer startTime() {
        return this.startTime;
    }
    /**
     * @return A set of tags assigned to the event.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return The type of the event.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEventResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> annotations;
        private String details;
        private Integer endtimeKey;
        private String id;
        private Boolean isEphemeral;
        private String name;
        private String severity;
        private Integer startTime;
        private List<String> tags;
        private String type;
        public Builder() {}
        public Builder(GetEventResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.details = defaults.details;
    	      this.endtimeKey = defaults.endtimeKey;
    	      this.id = defaults.id;
    	      this.isEphemeral = defaults.isEphemeral;
    	      this.name = defaults.name;
    	      this.severity = defaults.severity;
    	      this.startTime = defaults.startTime;
    	      this.tags = defaults.tags;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder annotations(Map<String,String> annotations) {
            if (annotations == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "annotations");
            }
            this.annotations = annotations;
            return this;
        }
        @CustomType.Setter
        public Builder details(String details) {
            if (details == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "details");
            }
            this.details = details;
            return this;
        }
        @CustomType.Setter
        public Builder endtimeKey(Integer endtimeKey) {
            if (endtimeKey == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "endtimeKey");
            }
            this.endtimeKey = endtimeKey;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isEphemeral(Boolean isEphemeral) {
            if (isEphemeral == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "isEphemeral");
            }
            this.isEphemeral = isEphemeral;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder severity(String severity) {
            if (severity == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "severity");
            }
            this.severity = severity;
            return this;
        }
        @CustomType.Setter
        public Builder startTime(Integer startTime) {
            if (startTime == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "startTime");
            }
            this.startTime = startTime;
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "tags");
            }
            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetEventResult", "type");
            }
            this.type = type;
            return this;
        }
        public GetEventResult build() {
            final var _resultValue = new GetEventResult();
            _resultValue.annotations = annotations;
            _resultValue.details = details;
            _resultValue.endtimeKey = endtimeKey;
            _resultValue.id = id;
            _resultValue.isEphemeral = isEphemeral;
            _resultValue.name = name;
            _resultValue.severity = severity;
            _resultValue.startTime = startTime;
            _resultValue.tags = tags;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
