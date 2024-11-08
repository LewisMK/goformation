// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_StringParameterDeclaration AWS CloudFormation Resource (AWS::QuickSight::Analysis.StringParameterDeclaration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameterdeclaration.html
type Analysis_StringParameterDeclaration struct {

	// DefaultValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameterdeclaration.html#cfn-quicksight-analysis-stringparameterdeclaration-defaultvalues
	DefaultValues *Analysis_StringDefaultValues `json:"DefaultValues,omitempty"`

	// MappedDataSetParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameterdeclaration.html#cfn-quicksight-analysis-stringparameterdeclaration-mappeddatasetparameters
	MappedDataSetParameters []Analysis_MappedDataSetParameter `json:"MappedDataSetParameters,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameterdeclaration.html#cfn-quicksight-analysis-stringparameterdeclaration-name
	Name string `json:"Name"`

	// ParameterValueType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameterdeclaration.html#cfn-quicksight-analysis-stringparameterdeclaration-parametervaluetype
	ParameterValueType string `json:"ParameterValueType"`

	// ValueWhenUnset AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-stringparameterdeclaration.html#cfn-quicksight-analysis-stringparameterdeclaration-valuewhenunset
	ValueWhenUnset *Analysis_StringValueWhenUnsetConfiguration `json:"ValueWhenUnset,omitempty"`

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
func (r *Analysis_StringParameterDeclaration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.StringParameterDeclaration"
}
