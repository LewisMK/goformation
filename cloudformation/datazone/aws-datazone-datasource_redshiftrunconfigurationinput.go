// Code generated by "go generate". Please don't change this file directly.

package datazone

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DataSource_RedshiftRunConfigurationInput AWS CloudFormation Resource (AWS::DataZone::DataSource.RedshiftRunConfigurationInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html
type DataSource_RedshiftRunConfigurationInput struct {

	// DataAccessRole AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-dataaccessrole
	DataAccessRole *string `json:"DataAccessRole,omitempty"`

	// RedshiftCredentialConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-redshiftcredentialconfiguration
	RedshiftCredentialConfiguration *DataSource_RedshiftCredentialConfiguration `json:"RedshiftCredentialConfiguration"`

	// RedshiftStorage AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-redshiftstorage
	RedshiftStorage *DataSource_RedshiftStorage `json:"RedshiftStorage"`

	// RelationalFilterConfigurations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-datasource-redshiftrunconfigurationinput.html#cfn-datazone-datasource-redshiftrunconfigurationinput-relationalfilterconfigurations
	RelationalFilterConfigurations []DataSource_RelationalFilterConfiguration `json:"RelationalFilterConfigurations"`

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
func (r *DataSource_RedshiftRunConfigurationInput) AWSCloudFormationType() string {
	return "AWS::DataZone::DataSource.RedshiftRunConfigurationInput"
}