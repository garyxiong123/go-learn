package ratelimit

import "fmt"

// LimitControlByUserIdByPeroid Rate Limit Control by user dimension
func LimitControlByUserIdByPeroid(param *LimitParam, control LimitControl) error {
	rateLimitKey := fmt.Sprintf("limit:userId:%s:%s", param.RequestPath, param.UserIdentifier)
	periodLimit := UserPeriodRateLimitMap[param.RequestPath]
	if err := LimitControlByPeriod(rateLimitKey, periodLimit); err != nil {
		fmt.Printf("LimitControlByUser hit Period Limit, path:%s, UserId:%s!\n",
			param.RequestPath, param.UserIdentifier)
		return err
	}
	return nil
}
