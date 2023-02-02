package ratelimit

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var GlobalTokenRateLimitMap map[string]*limit.TokenLimiter
var SingleTokenRateLimitMap map[string]*limit.TokenLimiter

func InitRateLimitControlByToken(redisInstance *redis.Redis) {
	GlobalTokenRateLimitMap = make(map[string]*limit.TokenLimiter)
	SingleTokenRateLimitMap = make(map[string]*limit.TokenLimiter)

	const (
		globalRate  = 200
		globalBurst = 20
	)
	GlobalTokenRateLimitMap["/greet"] = limit.NewTokenLimiter(globalRate, globalBurst, redisInstance, "GlobalTokenLimiter")

	const (
		singleRate  = 100
		singleBurst = 10
	)
	SingleTokenRateLimitMap["/greet"] = limit.NewTokenLimiter(singleRate, singleBurst, redisInstance, "SingleTokenLimiter")
}

func StartRateLimitControlByTokenLimit(param *LimitParam) error {
	fmt.Println("Start Token Rate Limit Control!!")
	err := LimitControlGlobalByToken(param, LimitControlSingleByToken)
	fmt.Println("End Token Rate Limit Control!!")
	return err
}
