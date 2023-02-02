package ratelimit

import (
	"fmt"
)

// LimitControlSingleByPeroid Rate Limit Control in single instance dimension
func LimitControlSingleByPeroid(param *LimitParam, control LimitControl) error {
	rateLimitKey := fmt.Sprintf("limit:single:%s:%s", param.RequestPath, HostID)
	periodLimit := SinglePeriodRateLimitMap[param.RequestPath]
	if err := LimitControlByPeriod(rateLimitKey, periodLimit); err != nil {
		fmt.Printf("LimitControlSingle hit Period Limit, path:%s!\n", param.RequestPath)
		return err
	}
	return control(param, nil)
}
