// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.wavefront.outputs.GetDashboardParameterDetail;
import com.pulumi.wavefront.outputs.GetDashboardSection;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetDashboardResult {
    /**
     * @return A list of users that have modify ACL access to the dashboard.
     * 
     */
    private List<String> canModifies;
    /**
     * @return A list of users that have view ACL access to the dashboard.
     * 
     */
    private List<String> canViews;
    private String chartTitleBgColor;
    private String chartTitleColor;
    private Integer chartTitleScalar;
    private Integer createdEpochMillis;
    private String creatorId;
    private String customer;
    private Integer defaultEndTime;
    private Integer defaultStartTime;
    private String defaultTimeWindow;
    private Boolean deleted;
    /**
     * @return Description of the chart.
     * 
     */
    private String description;
    private Boolean displayDescription;
    /**
     * @return Whether the dashboard parameters section is opened by default when the dashboard
     * is shown.
     * 
     */
    private Boolean displayQueryParameters;
    /**
     * @return Whether the &#34;pills&#34; quick-linked the sections of the dashboard are
     * displayed by default when the dashboard is shown.
     * 
     */
    private Boolean displaySectionTableOfContents;
    /**
     * @return How charts belonging to this dashboard should display events. `BYCHART` is default if
     * unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.
     * 
     */
    private String eventFilterType;
    private String eventQuery;
    private Boolean favorite;
    private Boolean hidden;
    private String id;
    /**
     * @return The name of the parameters.
     * 
     */
    private String name;
    private Integer numCharts;
    private Integer numFavorites;
    /**
     * @return The current JSON representation of dashboard parameters. See parameter details.
     * 
     */
    private List<GetDashboardParameterDetail> parameterDetails;
    private Map<String,String> parameters;
    private List<GetDashboardSection> sections;
    private Boolean systemOwned;
    /**
     * @return A set of tags to assign to this resource.
     * 
     */
    private List<String> tags;
    private Integer updatedEpochMillis;
    private String updaterId;
    /**
     * @return Unique identifier, also a URL slug of the dashboard.
     * 
     */
    private String url;
    private Integer viewsLastDay;
    private Integer viewsLastMonth;
    private Integer viewsLastWeek;

    private GetDashboardResult() {}
    /**
     * @return A list of users that have modify ACL access to the dashboard.
     * 
     */
    public List<String> canModifies() {
        return this.canModifies;
    }
    /**
     * @return A list of users that have view ACL access to the dashboard.
     * 
     */
    public List<String> canViews() {
        return this.canViews;
    }
    public String chartTitleBgColor() {
        return this.chartTitleBgColor;
    }
    public String chartTitleColor() {
        return this.chartTitleColor;
    }
    public Integer chartTitleScalar() {
        return this.chartTitleScalar;
    }
    public Integer createdEpochMillis() {
        return this.createdEpochMillis;
    }
    public String creatorId() {
        return this.creatorId;
    }
    public String customer() {
        return this.customer;
    }
    public Integer defaultEndTime() {
        return this.defaultEndTime;
    }
    public Integer defaultStartTime() {
        return this.defaultStartTime;
    }
    public String defaultTimeWindow() {
        return this.defaultTimeWindow;
    }
    public Boolean deleted() {
        return this.deleted;
    }
    /**
     * @return Description of the chart.
     * 
     */
    public String description() {
        return this.description;
    }
    public Boolean displayDescription() {
        return this.displayDescription;
    }
    /**
     * @return Whether the dashboard parameters section is opened by default when the dashboard
     * is shown.
     * 
     */
    public Boolean displayQueryParameters() {
        return this.displayQueryParameters;
    }
    /**
     * @return Whether the &#34;pills&#34; quick-linked the sections of the dashboard are
     * displayed by default when the dashboard is shown.
     * 
     */
    public Boolean displaySectionTableOfContents() {
        return this.displaySectionTableOfContents;
    }
    /**
     * @return How charts belonging to this dashboard should display events. `BYCHART` is default if
     * unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.
     * 
     */
    public String eventFilterType() {
        return this.eventFilterType;
    }
    public String eventQuery() {
        return this.eventQuery;
    }
    public Boolean favorite() {
        return this.favorite;
    }
    public Boolean hidden() {
        return this.hidden;
    }
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the parameters.
     * 
     */
    public String name() {
        return this.name;
    }
    public Integer numCharts() {
        return this.numCharts;
    }
    public Integer numFavorites() {
        return this.numFavorites;
    }
    /**
     * @return The current JSON representation of dashboard parameters. See parameter details.
     * 
     */
    public List<GetDashboardParameterDetail> parameterDetails() {
        return this.parameterDetails;
    }
    public Map<String,String> parameters() {
        return this.parameters;
    }
    public List<GetDashboardSection> sections() {
        return this.sections;
    }
    public Boolean systemOwned() {
        return this.systemOwned;
    }
    /**
     * @return A set of tags to assign to this resource.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    public Integer updatedEpochMillis() {
        return this.updatedEpochMillis;
    }
    public String updaterId() {
        return this.updaterId;
    }
    /**
     * @return Unique identifier, also a URL slug of the dashboard.
     * 
     */
    public String url() {
        return this.url;
    }
    public Integer viewsLastDay() {
        return this.viewsLastDay;
    }
    public Integer viewsLastMonth() {
        return this.viewsLastMonth;
    }
    public Integer viewsLastWeek() {
        return this.viewsLastWeek;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDashboardResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> canModifies;
        private List<String> canViews;
        private String chartTitleBgColor;
        private String chartTitleColor;
        private Integer chartTitleScalar;
        private Integer createdEpochMillis;
        private String creatorId;
        private String customer;
        private Integer defaultEndTime;
        private Integer defaultStartTime;
        private String defaultTimeWindow;
        private Boolean deleted;
        private String description;
        private Boolean displayDescription;
        private Boolean displayQueryParameters;
        private Boolean displaySectionTableOfContents;
        private String eventFilterType;
        private String eventQuery;
        private Boolean favorite;
        private Boolean hidden;
        private String id;
        private String name;
        private Integer numCharts;
        private Integer numFavorites;
        private List<GetDashboardParameterDetail> parameterDetails;
        private Map<String,String> parameters;
        private List<GetDashboardSection> sections;
        private Boolean systemOwned;
        private List<String> tags;
        private Integer updatedEpochMillis;
        private String updaterId;
        private String url;
        private Integer viewsLastDay;
        private Integer viewsLastMonth;
        private Integer viewsLastWeek;
        public Builder() {}
        public Builder(GetDashboardResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.canModifies = defaults.canModifies;
    	      this.canViews = defaults.canViews;
    	      this.chartTitleBgColor = defaults.chartTitleBgColor;
    	      this.chartTitleColor = defaults.chartTitleColor;
    	      this.chartTitleScalar = defaults.chartTitleScalar;
    	      this.createdEpochMillis = defaults.createdEpochMillis;
    	      this.creatorId = defaults.creatorId;
    	      this.customer = defaults.customer;
    	      this.defaultEndTime = defaults.defaultEndTime;
    	      this.defaultStartTime = defaults.defaultStartTime;
    	      this.defaultTimeWindow = defaults.defaultTimeWindow;
    	      this.deleted = defaults.deleted;
    	      this.description = defaults.description;
    	      this.displayDescription = defaults.displayDescription;
    	      this.displayQueryParameters = defaults.displayQueryParameters;
    	      this.displaySectionTableOfContents = defaults.displaySectionTableOfContents;
    	      this.eventFilterType = defaults.eventFilterType;
    	      this.eventQuery = defaults.eventQuery;
    	      this.favorite = defaults.favorite;
    	      this.hidden = defaults.hidden;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.numCharts = defaults.numCharts;
    	      this.numFavorites = defaults.numFavorites;
    	      this.parameterDetails = defaults.parameterDetails;
    	      this.parameters = defaults.parameters;
    	      this.sections = defaults.sections;
    	      this.systemOwned = defaults.systemOwned;
    	      this.tags = defaults.tags;
    	      this.updatedEpochMillis = defaults.updatedEpochMillis;
    	      this.updaterId = defaults.updaterId;
    	      this.url = defaults.url;
    	      this.viewsLastDay = defaults.viewsLastDay;
    	      this.viewsLastMonth = defaults.viewsLastMonth;
    	      this.viewsLastWeek = defaults.viewsLastWeek;
        }

        @CustomType.Setter
        public Builder canModifies(List<String> canModifies) {
            if (canModifies == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "canModifies");
            }
            this.canModifies = canModifies;
            return this;
        }
        public Builder canModifies(String... canModifies) {
            return canModifies(List.of(canModifies));
        }
        @CustomType.Setter
        public Builder canViews(List<String> canViews) {
            if (canViews == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "canViews");
            }
            this.canViews = canViews;
            return this;
        }
        public Builder canViews(String... canViews) {
            return canViews(List.of(canViews));
        }
        @CustomType.Setter
        public Builder chartTitleBgColor(String chartTitleBgColor) {
            if (chartTitleBgColor == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "chartTitleBgColor");
            }
            this.chartTitleBgColor = chartTitleBgColor;
            return this;
        }
        @CustomType.Setter
        public Builder chartTitleColor(String chartTitleColor) {
            if (chartTitleColor == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "chartTitleColor");
            }
            this.chartTitleColor = chartTitleColor;
            return this;
        }
        @CustomType.Setter
        public Builder chartTitleScalar(Integer chartTitleScalar) {
            if (chartTitleScalar == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "chartTitleScalar");
            }
            this.chartTitleScalar = chartTitleScalar;
            return this;
        }
        @CustomType.Setter
        public Builder createdEpochMillis(Integer createdEpochMillis) {
            if (createdEpochMillis == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "createdEpochMillis");
            }
            this.createdEpochMillis = createdEpochMillis;
            return this;
        }
        @CustomType.Setter
        public Builder creatorId(String creatorId) {
            if (creatorId == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "creatorId");
            }
            this.creatorId = creatorId;
            return this;
        }
        @CustomType.Setter
        public Builder customer(String customer) {
            if (customer == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "customer");
            }
            this.customer = customer;
            return this;
        }
        @CustomType.Setter
        public Builder defaultEndTime(Integer defaultEndTime) {
            if (defaultEndTime == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "defaultEndTime");
            }
            this.defaultEndTime = defaultEndTime;
            return this;
        }
        @CustomType.Setter
        public Builder defaultStartTime(Integer defaultStartTime) {
            if (defaultStartTime == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "defaultStartTime");
            }
            this.defaultStartTime = defaultStartTime;
            return this;
        }
        @CustomType.Setter
        public Builder defaultTimeWindow(String defaultTimeWindow) {
            if (defaultTimeWindow == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "defaultTimeWindow");
            }
            this.defaultTimeWindow = defaultTimeWindow;
            return this;
        }
        @CustomType.Setter
        public Builder deleted(Boolean deleted) {
            if (deleted == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "deleted");
            }
            this.deleted = deleted;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder displayDescription(Boolean displayDescription) {
            if (displayDescription == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "displayDescription");
            }
            this.displayDescription = displayDescription;
            return this;
        }
        @CustomType.Setter
        public Builder displayQueryParameters(Boolean displayQueryParameters) {
            if (displayQueryParameters == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "displayQueryParameters");
            }
            this.displayQueryParameters = displayQueryParameters;
            return this;
        }
        @CustomType.Setter
        public Builder displaySectionTableOfContents(Boolean displaySectionTableOfContents) {
            if (displaySectionTableOfContents == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "displaySectionTableOfContents");
            }
            this.displaySectionTableOfContents = displaySectionTableOfContents;
            return this;
        }
        @CustomType.Setter
        public Builder eventFilterType(String eventFilterType) {
            if (eventFilterType == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "eventFilterType");
            }
            this.eventFilterType = eventFilterType;
            return this;
        }
        @CustomType.Setter
        public Builder eventQuery(String eventQuery) {
            if (eventQuery == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "eventQuery");
            }
            this.eventQuery = eventQuery;
            return this;
        }
        @CustomType.Setter
        public Builder favorite(Boolean favorite) {
            if (favorite == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "favorite");
            }
            this.favorite = favorite;
            return this;
        }
        @CustomType.Setter
        public Builder hidden(Boolean hidden) {
            if (hidden == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "hidden");
            }
            this.hidden = hidden;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder numCharts(Integer numCharts) {
            if (numCharts == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "numCharts");
            }
            this.numCharts = numCharts;
            return this;
        }
        @CustomType.Setter
        public Builder numFavorites(Integer numFavorites) {
            if (numFavorites == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "numFavorites");
            }
            this.numFavorites = numFavorites;
            return this;
        }
        @CustomType.Setter
        public Builder parameterDetails(List<GetDashboardParameterDetail> parameterDetails) {
            if (parameterDetails == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "parameterDetails");
            }
            this.parameterDetails = parameterDetails;
            return this;
        }
        public Builder parameterDetails(GetDashboardParameterDetail... parameterDetails) {
            return parameterDetails(List.of(parameterDetails));
        }
        @CustomType.Setter
        public Builder parameters(Map<String,String> parameters) {
            if (parameters == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "parameters");
            }
            this.parameters = parameters;
            return this;
        }
        @CustomType.Setter
        public Builder sections(List<GetDashboardSection> sections) {
            if (sections == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "sections");
            }
            this.sections = sections;
            return this;
        }
        public Builder sections(GetDashboardSection... sections) {
            return sections(List.of(sections));
        }
        @CustomType.Setter
        public Builder systemOwned(Boolean systemOwned) {
            if (systemOwned == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "systemOwned");
            }
            this.systemOwned = systemOwned;
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "tags");
            }
            this.tags = tags;
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder updatedEpochMillis(Integer updatedEpochMillis) {
            if (updatedEpochMillis == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "updatedEpochMillis");
            }
            this.updatedEpochMillis = updatedEpochMillis;
            return this;
        }
        @CustomType.Setter
        public Builder updaterId(String updaterId) {
            if (updaterId == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "updaterId");
            }
            this.updaterId = updaterId;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "url");
            }
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder viewsLastDay(Integer viewsLastDay) {
            if (viewsLastDay == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "viewsLastDay");
            }
            this.viewsLastDay = viewsLastDay;
            return this;
        }
        @CustomType.Setter
        public Builder viewsLastMonth(Integer viewsLastMonth) {
            if (viewsLastMonth == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "viewsLastMonth");
            }
            this.viewsLastMonth = viewsLastMonth;
            return this;
        }
        @CustomType.Setter
        public Builder viewsLastWeek(Integer viewsLastWeek) {
            if (viewsLastWeek == null) {
              throw new MissingRequiredPropertyException("GetDashboardResult", "viewsLastWeek");
            }
            this.viewsLastWeek = viewsLastWeek;
            return this;
        }
        public GetDashboardResult build() {
            final var _resultValue = new GetDashboardResult();
            _resultValue.canModifies = canModifies;
            _resultValue.canViews = canViews;
            _resultValue.chartTitleBgColor = chartTitleBgColor;
            _resultValue.chartTitleColor = chartTitleColor;
            _resultValue.chartTitleScalar = chartTitleScalar;
            _resultValue.createdEpochMillis = createdEpochMillis;
            _resultValue.creatorId = creatorId;
            _resultValue.customer = customer;
            _resultValue.defaultEndTime = defaultEndTime;
            _resultValue.defaultStartTime = defaultStartTime;
            _resultValue.defaultTimeWindow = defaultTimeWindow;
            _resultValue.deleted = deleted;
            _resultValue.description = description;
            _resultValue.displayDescription = displayDescription;
            _resultValue.displayQueryParameters = displayQueryParameters;
            _resultValue.displaySectionTableOfContents = displaySectionTableOfContents;
            _resultValue.eventFilterType = eventFilterType;
            _resultValue.eventQuery = eventQuery;
            _resultValue.favorite = favorite;
            _resultValue.hidden = hidden;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.numCharts = numCharts;
            _resultValue.numFavorites = numFavorites;
            _resultValue.parameterDetails = parameterDetails;
            _resultValue.parameters = parameters;
            _resultValue.sections = sections;
            _resultValue.systemOwned = systemOwned;
            _resultValue.tags = tags;
            _resultValue.updatedEpochMillis = updatedEpochMillis;
            _resultValue.updaterId = updaterId;
            _resultValue.url = url;
            _resultValue.viewsLastDay = viewsLastDay;
            _resultValue.viewsLastMonth = viewsLastMonth;
            _resultValue.viewsLastWeek = viewsLastWeek;
            return _resultValue;
        }
    }
}
