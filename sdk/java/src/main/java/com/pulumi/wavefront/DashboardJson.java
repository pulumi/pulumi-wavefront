// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.wavefront.DashboardJsonArgs;
import com.pulumi.wavefront.Utilities;
import com.pulumi.wavefront.inputs.DashboardJsonState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Wavefront Dashboard JSON resource. This allows dashboards to be created, updated, and deleted.
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
 * import com.pulumi.wavefront.DashboardJson;
 * import com.pulumi.wavefront.DashboardJsonArgs;
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
 *         var testDashboardJson = new DashboardJson(&#34;testDashboardJson&#34;, DashboardJsonArgs.builder()        
 *             .dashboardJson(&#34;&#34;&#34;
 *   {
 *     &#34;acl&#34;: {
 *       &#34;canModify&#34;: [
 *         &#34;group-uuid&#34;,
 *         &#34;role-uuid&#34;
 *       ],
 *       &#34;canView&#34;: [
 *         &#34;group-uuid&#34;,
 *         &#34;role-uuid&#34;
 *       ]
 *     },
 *     &#34;name&#34;: &#34;Terraform Test Dashboard Json&#34;,
 *     &#34;description&#34;: &#34;a&#34;,
 *     &#34;eventFilterType&#34;: &#34;BYCHART&#34;,
 *     &#34;eventQuery&#34;: &#34;&#34;,
 *     &#34;defaultTimeWindow&#34;: &#34;&#34;,
 *     &#34;url&#34;: &#34;tftestimport&#34;,
 *     &#34;displayDescription&#34;: false,
 *     &#34;displaySectionTableOfContents&#34;: true,
 *     &#34;displayQueryParameters&#34;: false,
 *     &#34;sections&#34;: [
 *       {
 *         &#34;name&#34;: &#34;section 1&#34;,
 *         &#34;rows&#34;: [
 *           {
 *             &#34;charts&#34;: [
 *               {
 *                 &#34;name&#34;: &#34;chart 1&#34;,
 *                 &#34;sources&#34;: [
 *                   {
 *                     &#34;name&#34;: &#34;source 1&#34;,
 *                     &#34;query&#34;: &#34;ts()&#34;,
 *                     &#34;scatterPlotSource&#34;: &#34;Y&#34;,
 *                     &#34;querybuilderEnabled&#34;: false,
 *                     &#34;sourceDescription&#34;: &#34;&#34;
 *                   }
 *                 ],
 *                 &#34;units&#34;: &#34;someunit&#34;,
 *                 &#34;base&#34;: 0,
 *                 &#34;noDefaultEvents&#34;: false,
 *                 &#34;interpolatePoints&#34;: false,
 *                 &#34;includeObsoleteMetrics&#34;: false,
 *                 &#34;description&#34;: &#34;This is chart 1, showing something&#34;,
 *                 &#34;chartSettings&#34;: {
 *                   &#34;type&#34;: &#34;markdown-widget&#34;,
 *                   &#34;max&#34;: 100,
 *                   &#34;expectedDataSpacing&#34;: 120,
 *                   &#34;windowing&#34;: &#34;full&#34;,
 *                   &#34;windowSize&#34;: 10,
 *                   &#34;autoColumnTags&#34;: false,
 *                   &#34;columnTags&#34;: &#34;deprecated&#34;,
 *                   &#34;tagMode&#34;: &#34;all&#34;,
 *                   &#34;numTags&#34;: 2,
 *                   &#34;customTags&#34;: [
 *                     &#34;tag1&#34;,
 *                     &#34;tag2&#34;
 *                   ],
 *                   &#34;groupBySource&#34;: true,
 *                   &#34;y1Max&#34;: 100,
 *                   &#34;y1Units&#34;: &#34;units&#34;,
 *                   &#34;y0ScaleSIBy1024&#34;: true,
 *                   &#34;y1ScaleSIBy1024&#34;: true,
 *                   &#34;y0UnitAutoscaling&#34;: true,
 *                   &#34;y1UnitAutoscaling&#34;: true,
 *                   &#34;fixedLegendEnabled&#34;: true,
 *                   &#34;fixedLegendUseRawStats&#34;: true,
 *                   &#34;fixedLegendPosition&#34;: &#34;RIGHT&#34;,
 *                   &#34;fixedLegendDisplayStats&#34;: [
 *                     &#34;stat1&#34;,
 *                     &#34;stat2&#34;
 *                   ],
 *                   &#34;fixedLegendFilterSort&#34;: &#34;TOP&#34;,
 *                   &#34;fixedLegendFilterLimit&#34;: 1,
 *                   &#34;fixedLegendFilterField&#34;: &#34;CURRENT&#34;,
 *                   &#34;plainMarkdownContent&#34;: &#34;markdown content&#34;
 *                 },
 *                 &#34;chartAttributes&#34;: {
 *                   &#34;dashboardLinks&#34;: {
 *                     &#34;*&#34;: {
 *                       &#34;variables&#34;: {
 *                         &#34;xxx&#34;: &#34;xxx&#34;
 *                       },
 *                       &#34;destination&#34;: &#34;/dashboards/xxxx&#34;
 *                     }
 *                   }
 *                 },
 *                 &#34;summarization&#34;: &#34;MEAN&#34;
 *               }
 *             ],
 *             &#34;heightFactor&#34;: 50
 *           }
 *         ]
 *       }
 *     ],
 *     &#34;parameterDetails&#34;: {
 *       &#34;param&#34;: {
 *         &#34;hideFromView&#34;: false,
 *         &#34;description&#34;: null,
 *         &#34;allowAll&#34;: null,
 *         &#34;tagKey&#34;: null,
 *         &#34;queryValue&#34;: null,
 *         &#34;dynamicFieldType&#34;: null,
 *         &#34;reverseDynSort&#34;: null,
 *         &#34;parameterType&#34;: &#34;SIMPLE&#34;,
 *         &#34;label&#34;: &#34;test&#34;,
 *         &#34;defaultValue&#34;: &#34;Label&#34;,
 *         &#34;valuesToReadableStrings&#34;: {
 *           &#34;Label&#34;: &#34;test&#34;
 *         },
 *         &#34;selectedLabel&#34;: &#34;Label&#34;,
 *         &#34;value&#34;: &#34;test&#34;
 *       }
 *     },
 *     &#34;tags&#34;: {
 *       &#34;customerTags&#34;: [
 *         &#34;terraform&#34;
 *       ]
 *     }
 *   }
 * 
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * * 
 * *Note:
 * ** If there are dynamic variables in the Wavefront dashboard json, then these variables must be present in a separate file as mentioned in the section below.
 * 
 * ## Import
 * 
 * Dashboard JSON can be imported by using the `id`, e.g.:
 * 
 * ```sh
 * $ pulumi import wavefront:index/dashboardJson:DashboardJson dashboard_json tftestimport
 * ```
 * 
 */
@ResourceType(type="wavefront:index/dashboardJson:DashboardJson")
public class DashboardJson extends com.pulumi.resources.CustomResource {
    /**
     * See the [Wavefront API Documentation](https://docs.wavefront.com/wavefront_api.html#api-documentation-wavefront-instance)
     * for instructions on how to get to your API documentation for more details.
     * 
     */
    @Export(name="dashboardJson", refs={String.class}, tree="[0]")
    private Output<String> dashboardJson;

    /**
     * @return See the [Wavefront API Documentation](https://docs.wavefront.com/wavefront_api.html#api-documentation-wavefront-instance)
     * for instructions on how to get to your API documentation for more details.
     * 
     */
    public Output<String> dashboardJson() {
        return this.dashboardJson;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DashboardJson(String name) {
        this(name, DashboardJsonArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DashboardJson(String name, DashboardJsonArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DashboardJson(String name, DashboardJsonArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/dashboardJson:DashboardJson", name, args == null ? DashboardJsonArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DashboardJson(String name, Output<String> id, @Nullable DashboardJsonState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/dashboardJson:DashboardJson", name, state, makeResourceOptions(options, id));
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
    public static DashboardJson get(String name, Output<String> id, @Nullable DashboardJsonState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DashboardJson(name, id, state, options);
    }
}
