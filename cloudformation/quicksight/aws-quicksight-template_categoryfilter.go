// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_CategoryFilter AWS CloudFormation Resource (AWS::QuickSight::Template.CategoryFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html
type Template_CategoryFilter struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html#cfn-quicksight-template-categoryfilter-column
	Column *Template_ColumnIdentifier `json:"Column"`

	// Configuration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html#cfn-quicksight-template-categoryfilter-configuration
	Configuration *Template_CategoryFilterConfiguration `json:"Configuration"`

	// DefaultFilterControlConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html#cfn-quicksight-template-categoryfilter-defaultfiltercontrolconfiguration
	DefaultFilterControlConfiguration *Template_DefaultFilterControlConfiguration `json:"DefaultFilterControlConfiguration,omitempty"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html#cfn-quicksight-template-categoryfilter-filterid
	FilterId string `json:"FilterId"`

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
func (r *Template_CategoryFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.CategoryFilter"
}
