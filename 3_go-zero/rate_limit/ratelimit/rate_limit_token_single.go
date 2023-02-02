package ratelimit

import (
	"fmt"
	"github.com/pkg/errors"
)

// LimitControlSingleByToken Rate Limit Control in single instance dimension
func LimitControlSingleByToken(param *LimitParam, control LimitControl) error {
	tokenLimiter := SingleTokenRateLimitMap[param.RequestPath]
	if tokenLimiter.Allow() {
		return nil
	} else {
		fmt.Printf("LimitControlSingle hit Token Limit, path:%s!\n", param.RequestPath)
		return errors.New("Too Many Request!")
	}
}
