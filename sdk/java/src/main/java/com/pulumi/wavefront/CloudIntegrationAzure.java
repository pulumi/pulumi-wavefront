// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.wavefront.CloudIntegrationAzureArgs;
import com.pulumi.wavefront.Utilities;
import com.pulumi.wavefront.inputs.CloudIntegrationAzureState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Wavefront Cloud Integration for Microsoft Azure. This allows Azure cloud integrations to be created,
 * updated, and deleted.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.wavefront.CloudIntegrationAzureActivityLog;
 * import com.pulumi.wavefront.CloudIntegrationAzureActivityLogArgs;
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
 *         var azureActivityLog = new CloudIntegrationAzureActivityLog("azureActivityLog", CloudIntegrationAzureActivityLogArgs.builder()
 *             .name("Test Integration")
 *             .clientId("client-id2")
 *             .clientSecret("client-secret2")
 *             .tenant("my-tenant2")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Azure Cloud Integrations can be imported by using the `id`, e.g.:
 * 
 * ```sh
 * $ pulumi import wavefront:index/cloudIntegrationAzure:CloudIntegrationAzure azure a411c16b-3cf7-4f03-bf11-8ca05aab898d
 * ```
 * 
 */
@ResourceType(type="wavefront:index/cloudIntegrationAzure:CloudIntegrationAzure")
public class CloudIntegrationAzure extends com.pulumi.resources.CustomResource {
    /**
     * A list of point tag key-values to add to every point ingested using this integration.
     * 
     */
    @Export(name="additionalTags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> additionalTags;

    /**
     * @return A list of point tag key-values to add to every point ingested using this integration.
     * 
     */
    public Output<Optional<Map<String,String>>> additionalTags() {
        return Codegen.optional(this.additionalTags);
    }
    /**
     * A list of Azure Activity Log categories.
     * 
     */
    @Export(name="categoryFilters", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> categoryFilters;

    /**
     * @return A list of Azure Activity Log categories.
     * 
     */
    public Output<Optional<List<String>>> categoryFilters() {
        return Codegen.optional(this.categoryFilters);
    }
    /**
     * Client ID for an Azure service account within your project.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    /**
     * @return Client ID for an Azure service account within your project.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * Client secret for an Azure service account within your project.
     * 
     */
    @Export(name="clientSecret", refs={String.class}, tree="[0]")
    private Output<String> clientSecret;

    /**
     * @return Client secret for an Azure service account within your project.
     * 
     */
    public Output<String> clientSecret() {
        return this.clientSecret;
    }
    /**
     * Forces this resource to save, even if errors are present.
     * 
     */
    @Export(name="forceSave", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceSave;

    /**
     * @return Forces this resource to save, even if errors are present.
     * 
     */
    public Output<Optional<Boolean>> forceSave() {
        return Codegen.optional(this.forceSave);
    }
    /**
     * A regular expression that a metric name must match (case-insensitively) in order to be ingested.
     * 
     */
    @Export(name="metricFilterRegex", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> metricFilterRegex;

    /**
     * @return A regular expression that a metric name must match (case-insensitively) in order to be ingested.
     * 
     */
    public Output<Optional<String>> metricFilterRegex() {
        return Codegen.optional(this.metricFilterRegex);
    }
    /**
     * The human-readable name of this integration.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The human-readable name of this integration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A list of Azure resource groups from which to pull metrics.
     * 
     */
    @Export(name="resourceGroupFilters", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> resourceGroupFilters;

    /**
     * @return A list of Azure resource groups from which to pull metrics.
     * 
     */
    public Output<Optional<List<String>>> resourceGroupFilters() {
        return Codegen.optional(this.resourceGroupFilters);
    }
    /**
     * A value denoting which cloud service this service integrates with.
     * 
     */
    @Export(name="service", refs={String.class}, tree="[0]")
    private Output<String> service;

    /**
     * @return A value denoting which cloud service this service integrates with.
     * 
     */
    public Output<String> service() {
        return this.service;
    }
    /**
     * How often, in minutes, to refresh the service.
     * 
     */
    @Export(name="serviceRefreshRateInMinutes", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> serviceRefreshRateInMinutes;

    /**
     * @return How often, in minutes, to refresh the service.
     * 
     */
    public Output<Optional<Integer>> serviceRefreshRateInMinutes() {
        return Codegen.optional(this.serviceRefreshRateInMinutes);
    }
    /**
     * Tenant ID for an Azure service account within your project.
     * 
     */
    @Export(name="tenant", refs={String.class}, tree="[0]")
    private Output<String> tenant;

    /**
     * @return Tenant ID for an Azure service account within your project.
     * 
     */
    public Output<String> tenant() {
        return this.tenant;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudIntegrationAzure(java.lang.String name) {
        this(name, CloudIntegrationAzureArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudIntegrationAzure(java.lang.String name, CloudIntegrationAzureArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudIntegrationAzure(java.lang.String name, CloudIntegrationAzureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/cloudIntegrationAzure:CloudIntegrationAzure", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CloudIntegrationAzure(java.lang.String name, Output<java.lang.String> id, @Nullable CloudIntegrationAzureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/cloudIntegrationAzure:CloudIntegrationAzure", name, state, makeResourceOptions(options, id), false);
    }

    private static CloudIntegrationAzureArgs makeArgs(CloudIntegrationAzureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CloudIntegrationAzureArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "clientSecret"
            ))
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
    public static CloudIntegrationAzure get(java.lang.String name, Output<java.lang.String> id, @Nullable CloudIntegrationAzureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudIntegrationAzure(name, id, state, options);
    }
}
