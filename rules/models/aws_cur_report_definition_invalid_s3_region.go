// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCurReportDefinitionInvalidS3RegionRule checks the pattern is valid
type AwsCurReportDefinitionInvalidS3RegionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCurReportDefinitionInvalidS3RegionRule returns new rule with default attributes
func NewAwsCurReportDefinitionInvalidS3RegionRule() *AwsCurReportDefinitionInvalidS3RegionRule {
	return &AwsCurReportDefinitionInvalidS3RegionRule{
		resourceType:  "aws_cur_report_definition",
		attributeName: "s3_region",
		enum: []string{
			"af-south-1",
			"ap-east-1",
			"ap-south-1",
			"ap-southeast-1",
			"ap-southeast-2",
			"ap-northeast-1",
			"ap-northeast-2",
			"ap-northeast-3",
			"ca-central-1",
			"eu-central-1",
			"eu-west-1",
			"eu-west-2",
			"eu-west-3",
			"eu-north-1",
			"eu-south-1",
			"me-south-1",
			"sa-east-1",
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"cn-north-1",
			"cn-northwest-1",
		},
	}
}

// Name returns the rule name
func (r *AwsCurReportDefinitionInvalidS3RegionRule) Name() string {
	return "aws_cur_report_definition_invalid_s3_region"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCurReportDefinitionInvalidS3RegionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCurReportDefinitionInvalidS3RegionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCurReportDefinitionInvalidS3RegionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCurReportDefinitionInvalidS3RegionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as s3_region`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}