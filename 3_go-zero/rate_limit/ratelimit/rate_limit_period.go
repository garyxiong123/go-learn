package ratelimit

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var GlobalPeriodRateLimitMap map[string]*limit.PeriodLimit
var SinglePeriodRateLimitMap map[string]*limit.PeriodLimit
var UserPeriodRateLimitMap map[string]*limit.PeriodLimit

func InitRateLimitControlByPeriod(redisInstance *redis.Redis) {
	GlobalPeriodRateLimitMap = make(map[string]*limit.PeriodLimit)
	SinglePeriodRateLimitMap = make(map[string]*limit.PeriodLimit)
	UserPeriodRateLimitMap = make(map[string]*limit.PeriodLimit)

	const (
		GlobalSecond = 1
		GlobalQuota  = 20
	)
	GlobalPeriodRateLimitMap["/greet"] = limit.NewPeriodLimit(GlobalSecond, GlobalQuota, redisInstance, "GlobalPeriodLimiter")
	const (
		SingleSecond = 1
		SingleQuota  = 10
	)
	SinglePeriodRateLimitMap["/greet"] = limit.NewPeriodLimit(SingleSecond, SingleQuota, redisInstance, "SinglePeriodLimiter")
	const (
		UserSecond  = 1
		UserleQuota = 5
	)
	UserPeriodRateLimitMap["/greet"] = limit.NewPeriodLimit(UserSecond, UserleQuota, redisInstance, "UserPeriodLimiter")
}

func StartRateLimitControlByPeroidLimit(param *LimitParam) error {
	fmt.Println("Start Period Rate Limit Control!!")
	err := LimitControlGlobalByPeroid(param, LimitControlSingleByPeroid)
	fmt.Println("End Period Rate Limit Control!!")
	return err
}

func LimitControlByPeriod(limitKey string, periodLimit *limit.PeriodLimit) error {
	status, err := periodLimit.Take(limitKey)
	if err != nil {
		return err
	}
	switch status {
	case limit.Allowed:
		return nil
	case limit.HitQuota:
		return errors.New("Too Many Request!")
	case limit.OverQuota:
		return errors.New("Too Many Request!")
	}
	return errors.New("Unknown rate limit status error!")
}
