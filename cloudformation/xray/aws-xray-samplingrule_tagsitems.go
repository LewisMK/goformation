// Code generated by "go generate". Please don't change this file directly.

package xray

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// SamplingRule_TagsItems AWS CloudFormation Resource (AWS::XRay::SamplingRule.TagsItems)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-tagsitems.html
type SamplingRule_TagsItems struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-tagsitems.html#cfn-xray-samplingrule-tagsitems-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-xray-samplingrule-tagsitems.html#cfn-xray-samplingrule-tagsitems-value
	Value string `json:"Value"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *SamplingRule_TagsItems) AWSCloudFormationType() string {
	return "AWS::XRay::SamplingRule.TagsItems"
}