// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.wavefront.CloudIntegrationCloudWatchArgs;
import com.pulumi.wavefront.Utilities;
import com.pulumi.wavefront.inputs.CloudIntegrationCloudWatchState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Wavefront Cloud Integration for CloudWatch. This allows CloudWatch cloud integrations to be created,
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
 * import com.pulumi.wavefront.CloudIntegrationAwsExternalId;
 * import com.pulumi.wavefront.CloudIntegrationCloudWatch;
 * import com.pulumi.wavefront.CloudIntegrationCloudWatchArgs;
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
 *         var extId = new CloudIntegrationAwsExternalId("extId");
 * 
 *         var cloudwatch = new CloudIntegrationCloudWatch("cloudwatch", CloudIntegrationCloudWatchArgs.builder()
 *             .name("Test Integration")
 *             .forceSave(true)
 *             .roleArn("arn:aws::1234567:role/example-arn")
 *             .externalId(extId.id())
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
 * CloudWatch Cloud Integrations can be imported by using the `id`, e.g.:
 * 
 * ```sh
 * $ pulumi import wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch cloudwatch a411c16b-3cf7-4f03-bf11-8ca05aab898d
 * ```
 * 
 */
@ResourceType(type="wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch")
public class CloudIntegrationCloudWatch extends com.pulumi.resources.CustomResource {
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
     * The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
     * 
     */
    @Export(name="externalId", refs={String.class}, tree="[0]")
    private Output<String> externalId;

    /**
     * @return The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
     * 
     */
    public Output<String> externalId() {
        return this.externalId;
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
     * A string-&gt;string map allow list of instance tag-value pairs (in AWS).
     * If the instance&#39;s AWS tags match this allow list, CloudWatch data about this instance is ingested.
     * Multiple entries are OR&#39;ed.
     * 
     */
    @Export(name="instanceSelectionTags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> instanceSelectionTags;

    /**
     * @return A string-&gt;string map allow list of instance tag-value pairs (in AWS).
     * If the instance&#39;s AWS tags match this allow list, CloudWatch data about this instance is ingested.
     * Multiple entries are OR&#39;ed.
     * 
     */
    public Output<Optional<Map<String,String>>> instanceSelectionTags() {
        return Codegen.optional(this.instanceSelectionTags);
    }
    /**
     * A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
     * 
     */
    @Export(name="metricFilterRegex", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> metricFilterRegex;

    /**
     * @return A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.
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
     * A list of namespaces that limit what we query from CloudWatch.
     * 
     */
    @Export(name="namespaces", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> namespaces;

    /**
     * @return A list of namespaces that limit what we query from CloudWatch.
     * 
     */
    public Output<Optional<List<String>>> namespaces() {
        return Codegen.optional(this.namespaces);
    }
    /**
     * A regular expression that AWS tag key name must match (case-insensitively)
     * in order to be ingested.
     * 
     */
    @Export(name="pointTagFilterRegex", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pointTagFilterRegex;

    /**
     * @return A regular expression that AWS tag key name must match (case-insensitively)
     * in order to be ingested.
     * 
     */
    public Output<Optional<String>> pointTagFilterRegex() {
        return Codegen.optional(this.pointTagFilterRegex);
    }
    /**
     * The external ID corresponding to the Role ARN.
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output<String> roleArn;

    /**
     * @return The external ID corresponding to the Role ARN.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
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
     * A string-&gt;string map of allow list of volume tag-value pairs (in AWS).
     * If the volume&#39;s AWS tags match this allow list, CloudWatch data about this volume is ingested.
     * Multiple entries are OR&#39;ed.
     * 
     */
    @Export(name="volumeSelectionTags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> volumeSelectionTags;

    /**
     * @return A string-&gt;string map of allow list of volume tag-value pairs (in AWS).
     * If the volume&#39;s AWS tags match this allow list, CloudWatch data about this volume is ingested.
     * Multiple entries are OR&#39;ed.
     * 
     */
    public Output<Optional<Map<String,String>>> volumeSelectionTags() {
        return Codegen.optional(this.volumeSelectionTags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudIntegrationCloudWatch(java.lang.String name) {
        this(name, CloudIntegrationCloudWatchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudIntegrationCloudWatch(java.lang.String name, CloudIntegrationCloudWatchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudIntegrationCloudWatch(java.lang.String name, CloudIntegrationCloudWatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CloudIntegrationCloudWatch(java.lang.String name, Output<java.lang.String> id, @Nullable CloudIntegrationCloudWatchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/cloudIntegrationCloudWatch:CloudIntegrationCloudWatch", name, state, makeResourceOptions(options, id), false);
    }

    private static CloudIntegrationCloudWatchArgs makeArgs(CloudIntegrationCloudWatchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CloudIntegrationCloudWatchArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static CloudIntegrationCloudWatch get(java.lang.String name, Output<java.lang.String> id, @Nullable CloudIntegrationCloudWatchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudIntegrationCloudWatch(name, id, state, options);
    }
}
