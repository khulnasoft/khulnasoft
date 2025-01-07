package licensing

import (
	"reflect"
	"strings"

	"github.com/khulnasoft/khulnasoft/lib/errors"
)

// A Plan is a pricing plan, with an associated set of features that it offers.
type Plan string

// Details returns the name and features of the plan.
func (p Plan) Details() PlanDetails {
	return planDetails[p]
}

func (p Plan) IsFreePlan() bool {
	return p == PlanFree0 || p == PlanFree1
}

// HasFeature returns whether the plan has the given feature.
// If the target is a pointer, the plan's feature configuration will be
// set to the target.
func (p Plan) HasFeature(target Feature) bool {
	if target == nil {
		panic("licensing: target cannot be nil")
	}

	val := reflect.ValueOf(target)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		panic("licensing: target cannot be a nil pointer")
	}

	for _, f := range planDetails[p].Features {
		if target.FeatureName() == f.FeatureName() {
			if val.Kind() == reflect.Ptr {
				val.Elem().Set(reflect.ValueOf(f).Elem())
			}
			return true
		}
	}
	return false
}

const planTagPrefix = "plan:"

// Tag is the representation of the plan as a Tag in a license key.
func (p Plan) Tag() string { return planTagPrefix + string(p) }

// isKnown reports whether the plan is a known plan.
func (p Plan) isKnown() bool {
	for _, plan := range AllPlans {
		if p == plan {
			return true
		}
	}
	return false
}

// Plan is the pricing plan of the license.
func (info *Info) Plan() Plan {
	return PlanFromTags(info.Tags)
}

// HasUnknownPlan returns an error if the plan is presented in the license tags
// but unrecognizable. It returns nil if there is no tags found for plans.
func (info *Info) HasUnknownPlan() error {
	for _, tag := range info.Tags {
		// A tag that begins with "plan:" indicates the license's plan.
		if !strings.HasPrefix(tag, planTagPrefix) {
			continue
		}

		plan := Plan(tag[len(planTagPrefix):])
		if !plan.isKnown() {
			return errors.Errorf("The license has an unrecognizable plan in tag %q, please contact Khulnasoft support.", tag)
		}
	}
	return nil
}

// PlanFromTags returns the pricing plan of the license, based on the given tags.
func PlanFromTags(tags []string) Plan {
	for _, tag := range tags {
		// A tag that begins with "plan:" indicates the license's plan.
		if strings.HasPrefix(tag, planTagPrefix) {
			plan := Plan(tag[len(planTagPrefix):])
			if plan.isKnown() {
				return plan
			}
		}
	}

	// Backcompat: no tags means it is the old "Enterprise" plan.
	// TODO: In the future, we will no longer allow this and instances without a
	// plan tag will be on the Free plan instead.
	return PlanOldEnterprise
}
