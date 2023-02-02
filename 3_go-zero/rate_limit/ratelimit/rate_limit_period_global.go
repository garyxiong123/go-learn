package ratelimit

import "fmt"

// LimitControlGlobalByPeroid Rate Limit Control in global dimension
func LimitControlGlobalByPeroid(param *LimitParam, control LimitControl) error {
	rateLimitKey := fmt.Sprintf("limit:global:%s", param.RequestPath)
	periodLimit := GlobalPeriodRateLimitMap[param.RequestPath]
	if err := LimitControlByPeriod(rateLimitKey, periodLimit); err != nil {
		fmt.Printf("LimitControlGlobal hit Period Limit, path:%s!\n", param.RequestPath)
		return err
	}
	return control(param, LimitControlByUserIdByPeroid)
}
