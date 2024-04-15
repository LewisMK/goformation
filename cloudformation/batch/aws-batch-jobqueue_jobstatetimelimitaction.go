// Code generated by "go generate". Please don't change this file directly.

package batch

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// JobQueue_JobStateTimeLimitAction AWS CloudFormation Resource (AWS::Batch::JobQueue.JobStateTimeLimitAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html
type JobQueue_JobStateTimeLimitAction struct {

	// Action AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html#cfn-batch-jobqueue-jobstatetimelimitaction-action
	Action string `json:"Action"`

	// MaxTimeSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html#cfn-batch-jobqueue-jobstatetimelimitaction-maxtimeseconds
	MaxTimeSeconds int `json:"MaxTimeSeconds"`

	// Reason AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html#cfn-batch-jobqueue-jobstatetimelimitaction-reason
	Reason string `json:"Reason"`

	// State AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html#cfn-batch-jobqueue-jobstatetimelimitaction-state
	State string `json:"State"`

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
func (r *JobQueue_JobStateTimeLimitAction) AWSCloudFormationType() string {
	return "AWS::Batch::JobQueue.JobStateTimeLimitAction"
}