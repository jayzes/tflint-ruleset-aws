// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloud9EnvironmentEc2InvalidNameRule checks the pattern is valid
type AwsCloud9EnvironmentEc2InvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloud9EnvironmentEc2InvalidNameRule returns new rule with default attributes
func NewAwsCloud9EnvironmentEc2InvalidNameRule() *AwsCloud9EnvironmentEc2InvalidNameRule {
	return &AwsCloud9EnvironmentEc2InvalidNameRule{
		resourceType:  "aws_cloud9_environment_ec2",
		attributeName: "name",
		max:           60,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Name() string {
	return "aws_cloud9_environment_ec2_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 60 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}