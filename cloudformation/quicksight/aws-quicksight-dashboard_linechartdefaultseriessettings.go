// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_LineChartDefaultSeriesSettings AWS CloudFormation Resource (AWS::QuickSight::Dashboard.LineChartDefaultSeriesSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartdefaultseriessettings.html
type Dashboard_LineChartDefaultSeriesSettings struct {

	// AxisBinding AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartdefaultseriessettings.html#cfn-quicksight-dashboard-linechartdefaultseriessettings-axisbinding
	AxisBinding *string `json:"AxisBinding,omitempty"`

	// LineStyleSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartdefaultseriessettings.html#cfn-quicksight-dashboard-linechartdefaultseriessettings-linestylesettings
	LineStyleSettings *Dashboard_LineChartLineStyleSettings `json:"LineStyleSettings,omitempty"`

	// MarkerStyleSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-linechartdefaultseriessettings.html#cfn-quicksight-dashboard-linechartdefaultseriessettings-markerstylesettings
	MarkerStyleSettings *Dashboard_LineChartMarkerStyleSettings `json:"MarkerStyleSettings,omitempty"`

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
func (r *Dashboard_LineChartDefaultSeriesSettings) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.LineChartDefaultSeriesSettings"
}