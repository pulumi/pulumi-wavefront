// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.wavefront.CloudIntegrationCloudTrailArgs;
import com.pulumi.wavefront.Utilities;
import com.pulumi.wavefront.inputs.CloudIntegrationCloudTrailState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Wavefront Cloud Integration for CloudTrail. This allows CloudTrail cloud integrations to be created,
 * updated, and deleted.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.wavefront.CloudIntegrationAwsExternalId;
 * import com.pulumi.wavefront.CloudIntegrationCloudTrail;
 * import com.pulumi.wavefront.CloudIntegrationCloudTrailArgs;
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
 *         var extId = new CloudIntegrationAwsExternalId(&#34;extId&#34;);
 * 
 *         var cloudtrail = new CloudIntegrationCloudTrail(&#34;cloudtrail&#34;, CloudIntegrationCloudTrailArgs.builder()        
 *             .roleArn(&#34;arn:aws::1234567:role/example-arn&#34;)
 *             .externalId(extId.id())
 *             .region(&#34;us-west-2&#34;)
 *             .bucketName(&#34;example-s3-bucket&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CloudTrail Cloud Integrations can be imported by using the `id`, e.g.
 * 
 * ```sh
 *  $ pulumi import wavefront:index/cloudIntegrationCloudTrail:CloudIntegrationCloudTrail cloudtrail a411c16b-3cf7-4f03-bf11-8ca05aab898d
 * ```
 * 
 */
@ResourceType(type="wavefront:index/cloudIntegrationCloudTrail:CloudIntegrationCloudTrail")
public class CloudIntegrationCloudTrail extends com.pulumi.resources.CustomResource {
    /**
     * A list of point tag key-values to add to every point ingested using this integration.
     * 
     */
    @Export(name="additionalTags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> additionalTags;

    /**
     * @return A list of point tag key-values to add to every point ingested using this integration.
     * 
     */
    public Output<Optional<Map<String,String>>> additionalTags() {
        return Codegen.optional(this.additionalTags);
    }
    /**
     * Name of the S3 bucket where CloudTrail logs are stored.
     * 
     */
    @Export(name="bucketName", type=String.class, parameters={})
    private Output<String> bucketName;

    /**
     * @return Name of the S3 bucket where CloudTrail logs are stored.
     * 
     */
    public Output<String> bucketName() {
        return this.bucketName;
    }
    /**
     * The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
     * 
     */
    @Export(name="externalId", type=String.class, parameters={})
    private Output<String> externalId;

    /**
     * @return The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.
     * 
     */
    public Output<String> externalId() {
        return this.externalId;
    }
    /**
     * Rule to filter CloudTrail log event.
     * 
     */
    @Export(name="filterRule", type=String.class, parameters={})
    private Output</* @Nullable */ String> filterRule;

    /**
     * @return Rule to filter CloudTrail log event.
     * 
     */
    public Output<Optional<String>> filterRule() {
        return Codegen.optional(this.filterRule);
    }
    /**
     * Forces this resource to save, even if errors are present.
     * 
     */
    @Export(name="forceSave", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> forceSave;

    /**
     * @return Forces this resource to save, even if errors are present.
     * 
     */
    public Output<Optional<Boolean>> forceSave() {
        return Codegen.optional(this.forceSave);
    }
    /**
     * The human-readable name of this integration.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The human-readable name of this integration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The common prefix, if any, appended to all CloudTrail log files.
     * 
     */
    @Export(name="prefix", type=String.class, parameters={})
    private Output</* @Nullable */ String> prefix;

    /**
     * @return The common prefix, if any, appended to all CloudTrail log files.
     * 
     */
    public Output<Optional<String>> prefix() {
        return Codegen.optional(this.prefix);
    }
    /**
     * The AWS region of the S3 bucket where CloudTrail logs are stored.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The AWS region of the S3 bucket where CloudTrail logs are stored.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The external ID corresponding to the Role ARN.
     * 
     */
    @Export(name="roleArn", type=String.class, parameters={})
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
    @Export(name="service", type=String.class, parameters={})
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
    @Export(name="serviceRefreshRateInMinutes", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> serviceRefreshRateInMinutes;

    /**
     * @return How often, in minutes, to refresh the service.
     * 
     */
    public Output<Optional<Integer>> serviceRefreshRateInMinutes() {
        return Codegen.optional(this.serviceRefreshRateInMinutes);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CloudIntegrationCloudTrail(String name) {
        this(name, CloudIntegrationCloudTrailArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CloudIntegrationCloudTrail(String name, CloudIntegrationCloudTrailArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CloudIntegrationCloudTrail(String name, CloudIntegrationCloudTrailArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/cloudIntegrationCloudTrail:CloudIntegrationCloudTrail", name, args == null ? CloudIntegrationCloudTrailArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CloudIntegrationCloudTrail(String name, Output<String> id, @Nullable CloudIntegrationCloudTrailState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("wavefront:index/cloudIntegrationCloudTrail:CloudIntegrationCloudTrail", name, state, makeResourceOptions(options, id));
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
    public static CloudIntegrationCloudTrail get(String name, Output<String> id, @Nullable CloudIntegrationCloudTrailState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CloudIntegrationCloudTrail(name, id, state, options);
    }
}