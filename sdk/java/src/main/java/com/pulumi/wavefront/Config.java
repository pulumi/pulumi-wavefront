// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.wavefront;

import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("wavefront");
    public String address() {
        return Codegen.stringProp("address").config(config).require();
    }
    public Optional<String> httpProxy() {
        return Codegen.stringProp("httpProxy").config(config).get();
    }
    public String token() {
        return Codegen.stringProp("token").config(config).require();
    }
}