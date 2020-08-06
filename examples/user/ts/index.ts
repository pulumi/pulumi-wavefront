import * as pulumi from "@pulumi/pulumi";
import * as wavefront from "@pulumi/wavefront";

const user = new wavefront.User("demo-ts", {
    email: "test+ts@mycompany.io"
});
