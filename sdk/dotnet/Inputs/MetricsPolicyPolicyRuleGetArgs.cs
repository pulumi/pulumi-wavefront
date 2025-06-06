// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Wavefront.Inputs
{

    public sealed class MetricsPolicyPolicyRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid options are `ALLOW` and `BLOCK`.
        /// </summary>
        [Input("accessType", required: true)]
        public Input<string> AccessType { get; set; } = null!;

        [Input("accountIds")]
        private InputList<string>? _accountIds;

        /// <summary>
        /// List of account ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
        /// </summary>
        public InputList<string> AccountIds
        {
            get => _accountIds ?? (_accountIds = new InputList<string>());
            set => _accountIds = value;
        }

        /// <summary>
        /// A detailed description of the Metrics Policy. The description is visible only when you edit the rule.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The unique name identifier for a Metrics Policy. The name is visible on the Metrics Security Policy page.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("prefixes", required: true)]
        private InputList<string>? _prefixes;

        /// <summary>
        /// List of prefixes to match metrics on. You can specify the full metric name or use a wildcard character in metric names, sources, or point tags. The wildcard character alone (*) means all metrics.
        /// </summary>
        public InputList<string> Prefixes
        {
            get => _prefixes ?? (_prefixes = new InputList<string>());
            set => _prefixes = value;
        }

        [Input("roleIds")]
        private InputList<string>? _roleIds;

        /// <summary>
        /// List of role ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        [Input("tags")]
        private InputList<Inputs.MetricsPolicyPolicyRuleTagGetArgs>? _tags;

        /// <summary>
        /// List of Key/Value tags to select target metrics for policy.
        /// </summary>
        public InputList<Inputs.MetricsPolicyPolicyRuleTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.MetricsPolicyPolicyRuleTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Bool where `true` require all tags are met by selected metrics, else `false` select metrics that match any give tag.
        /// </summary>
        [Input("tagsAnded", required: true)]
        public Input<bool> TagsAnded { get; set; } = null!;

        [Input("userGroupIds")]
        private InputList<string>? _userGroupIds;

        /// <summary>
        /// List of user group ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.
        /// </summary>
        public InputList<string> UserGroupIds
        {
            get => _userGroupIds ?? (_userGroupIds = new InputList<string>());
            set => _userGroupIds = value;
        }

        public MetricsPolicyPolicyRuleGetArgs()
        {
        }
        public static new MetricsPolicyPolicyRuleGetArgs Empty => new MetricsPolicyPolicyRuleGetArgs();
    }
}
