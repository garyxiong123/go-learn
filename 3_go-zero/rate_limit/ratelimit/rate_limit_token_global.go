package ratelimit

import (
	"fmt"
	"github.com/pkg/errors"
)

// LimitControlGlobalByToken Rate Limit Control in global dimension
func LimitControlGlobalByToken(param *LimitParam, control LimitControl) error {
	tokenLimiter := GlobalTokenRateLimitMap[param.RequestPath]
	if tokenLimiter.Allow() {
		return control(param, nil)
	} else {
		fmt.Printf("LimitControlGlobal hit Token Limit, path:%s!\n", param.RequestPath)
		return errors.New("Too Many Request!")
	}
}
