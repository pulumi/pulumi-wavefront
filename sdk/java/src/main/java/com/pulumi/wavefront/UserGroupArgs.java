// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserGroupArgs Empty = new UserGroupArgs();

    /**
     * A short description of the user group.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return A short description of the user group.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * The name of the user group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the user group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private UserGroupArgs() {}

    private UserGroupArgs(UserGroupArgs $) {
        this.description = $.description;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserGroupArgs $;

        public Builder() {
            $ = new UserGroupArgs();
        }

        public Builder(UserGroupArgs defaults) {
            $ = new UserGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A short description of the user group.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A short description of the user group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the user group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the user group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public UserGroupArgs build() {
            $.description = Objects.requireNonNull($.description, "expected parameter 'description' to be non-null");
            return $;
        }
    }

}