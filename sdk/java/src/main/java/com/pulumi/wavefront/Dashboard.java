// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.wavefront.DashboardArgs;
import com.pulumi.wavefront.Utilities;
import com.pulumi.wavefront.inputs.DashboardState;
import com.pulumi.wavefront.outputs.DashboardParameterDetail;
import com.pulumi.wavefront.outputs.DashboardSection;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Wavefront Dashboard resource. This allows dashboards to be created, updated, and deleted.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.wavefront.User;
 * import com.pulumi.wavefront.UserArgs;
 * import com.pulumi.wavefront.Dashboard;
 * import com.pulumi.wavefront.DashboardArgs;
 * import com.pulumi.wavefront.inputs.DashboardSectionArgs;
 * import com.pulumi.wavefront.inputs.DashboardParameterDetailArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var basic = new User(&#34;basic&#34;, UserArgs.builder()        
 *             .email(&#34;test+tftesting@example.com&#34;)
 *             .groups(            
 *                 &#34;agent_management&#34;,
 *                 &#34;alerts_management&#34;)
 *             .build());
 * 
 *         var testDashboard = new Dashboard(&#34;testDashboard&#34;, DashboardArgs.builder()        
 *             .description(&#34;testing, testing&#34;)
 *             .url(&#34;tftestcreate&#34;)
 *             .displaySectionTableOfContents(true)
 *             .displayQueryParameters(true)
 *             .canViews(basic.id())
 *             .sections(DashboardSectionArgs.builder()
 *                 .name(&#34;section 1&#34;)
 *                 .rows(DashboardSectionRowArgs.builder()
 *                     .charts(DashboardSectionRowChartArgs.builder()
 *                         .name(&#34;chart 1&#34;)
 *                         .description(&#34;chart number 1&#34;)
 *                         .units(&#34;something per unit&#34;)
 *                         .sources(DashboardSectionRowChartSourceArgs.builder()
 *                             .name(&#34;source name&#34;)
 *                             .query(&#34;ts()&#34;)
 *                             .build())
 *                         .chartSetting(DashboardSectionRowChartChartSettingArgs.builder()
 *                             .type(&#34;linear&#34;)
 *                             .build())
 *                         .summarization(&#34;MEAN&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .parameterDetails(DashboardParameterDetailArgs.builder()
 *                 .name(&#34;param1&#34;)
 *                 .label(&#34;param1&#34;)
 *                 .defaultValue(&#34;Label&#34;)
 *                 .hideFromView(false)
 *                 .parameterType(&#34;SIMPLE&#34;)
 *                 .valuesToReadableStrings(Map.of(&#34;Label&#34;, &#34;test&#34;))
 *                 .build())
 *             .tags(            
 *                 &#34;b&#34;,
 *                 &#34;terraform&#34;,
 *                 &#34;a&#34;,
 *                 &#34;test&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Dashboards can be imported by using the `id`, e.g.:
 * 
 * ```sh
 * $ pulumi import wavefront:index/dashboard:Dashboard dashboard tftestimport
 * ```
 * 
 */
@ResourceType(type="wavefront:index/dashboard:Dashboard")
public class Dashboard extends com.pulumi.resources.CustomResource {
    /**
     * A list of users/groups/roles that can modify the dashboard.
     * 
     */
    @Export(name="canModifies", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> canModifies;

    /**
     * @return A list of users/groups/roles that can modify the dashboard.
     * 
     */
    public Output<List<String>> canModifies() {
        return this.canModifies;
    }
    /**
     * A list of users/groups/roles that can view the dashboard.
     * 
     */
    @Export(name="canViews", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> canViews;

    /**
     * @return A list of users/groups/roles that can view the dashboard.
     * 
     */
    public Output<Optional<List<String>>> canViews() {
        return Codegen.optional(this.canViews);
    }
    /**
     * Human-readable description of the dashboard.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Human-readable description of the dashboard.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Whether the dashboard parameters section is opened by default when the dashboard
     * is shown.
     * 
     */
    @Export(name="displayQueryParameters", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> displayQueryParameters;

    /**
     * @return Whether the dashboard parameters section is opened by default when the dashboard
     * is shown.
     * 
     */
    public Output<Optional<Boolean>> displayQueryParameters() {
        return Codegen.optional(this.displayQueryParameters);
    }
    /**
     * Whether the &#34;pills&#34; quick-linked the sections of the dashboard are
     * displayed by default when the dashboard is shown.
     * 
     */
    @Export(name="displaySectionTableOfContents", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> displaySectionTableOfContents;

    /**
     * @return Whether the &#34;pills&#34; quick-linked the sections of the dashboard are
     * displayed by default when the dashboard is shown.
     * 
     */
    public Output<Optional<Boolean>> displaySectionTableOfContents() {
        return Codegen.optional(this.displaySectionTableOfContents);
    }
    /**
     * How charts belonging to this dashboard should display events. `BYCHART` is default if
     * unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.
     * 
     */
    @Export(name="eventFilterType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> eventFilterType;

    /**
     * @return How charts belonging to this dashboard should display events. `BYCHART` is default if
     * unspecified. Valid options are: `BYCHART`, `AUTOMATIC`, `ALL`, `NONE`, `BYDASHBOARD`, and `BYCHARTANDDASHBOARD`.
     * 
     */
    public Output<Optional<String>> eventFilterType() {
        return Codegen.optional(this.eventFilterType);
    }
    /**
     * Name of the dashboard.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the dashboard.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The current JSON representation of dashboard parameters. See parameter details.
     * 
     */
    @Export(name="parameterDetails", refs={List.class,DashboardParameterDetail.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DashboardParameterDetail>> parameterDetails;

    /**
     * @return The current JSON representation of dashboard parameters. See parameter details.
     * 
     */
    public Output<Optional<List<DashboardParameterDetail>>> parameterDetails() {
        return Codegen.optional(this.parameterDetails);
    }
    /**
     * Dashboard chart sections. See dashboard sections.
     * 
     */
    @Export(name="sections", refs={List.class,DashboardSection.class}, tree="[0,1]")
    private Output<List<DashboardSection>> sections;

    /**
     * @return Dashboard chart sections. See dashboard sections.
     * 
     */
    public Output<List<DashboardSection>> sections() {
        return this.sections;
    }
    /**
     * A set of tags to assign to this resource.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tags;

    /**
     * @return A set of tags to assign to this resource.
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }
    /**
     * Unique identifier, also a URL slug of the dashboard.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return Unique identifier, also a URL slug of the dashboard.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Dashboard(String name) {
        this(name, DashboardArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Dashboard(String name, DashboardArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Dashboard(String name, DashboardArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/dashboard:Dashboard", name, args == null ? DashboardArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Dashboard(String name, Output<String> id, @Nullable DashboardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/dashboard:Dashboard", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Dashboard get(String name, Output<String> id, @Nullable DashboardState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Dashboard(name, id, state, options);
    }
}
