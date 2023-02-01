package ratelimit

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis/redistest"
	"net/http"
)

var periodLimit *limit.PeriodLimit

var tokenLimit *limit.TokenLimiter

type LimitParam struct {
	RequestPath    string
	UserIdentifier string
}

type LimitControl func(*LimitParam, LimitControl) error

func LimitHander(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		requestPath := request.URL.Path
		userId := request.Form.Get("userId")

		limitParam := &LimitParam{
			RequestPath:    requestPath,
			UserIdentifier: userId,
		}

		if err := StartRateLimitControl(limitParam, LimitControlGlobal); err != nil {
			http.Error(writer, err.Error(), http.StatusTooManyRequests)
			return
		}

		next(writer, request)
	}
}

func StartRateLimitControl(param *LimitParam, control LimitControl) error {
	fmt.Println("Start Rate Limit Control!!")
	err := control(param, LimitControlSingle)
	fmt.Println("End Rate Limit Control!!")
	return err
}

// LimitControlGlobal Rate Limit Control in global dimension
func LimitControlGlobal(param *LimitParam, control LimitControl) error {
	rateLimitKey := fmt.Sprintf("limit:%s", param.RequestPath)
	if err := LimitControlByPeriod(rateLimitKey); err != nil {
		return err
	}
	return control(param, LimitControlSingle)
}

// LimitControlSingle Rate Limit Control in single instance dimension
func LimitControlSingle(param *LimitParam, control LimitControl) error {
	rateLimitKey := fmt.Sprintf("limit:%s", param.RequestPath)
	if err := LimitControlByPeriod(rateLimitKey); err != nil {
		return err
	}
	return control(param, LimitControlByUserId)
}

// LimitControlByUserId Rate Limit Control by user dimension
func LimitControlByUserId(param *LimitParam, control LimitControl) error {
	rateLimitKey := fmt.Sprintf("limit:%s:%s", param.RequestPath, param.UserIdentifier)
	if err := LimitControlByPeriod(rateLimitKey); err != nil {
		return err
	}
	return nil
}

func LimitControlByPeriod(limitKey string) error {
	status, err := periodLimit.Take(limitKey)
	if err != nil {
		return err
	}
	switch status {
	case limit.Allowed:
		return nil
	case limit.HitQuota:
		return errors.New(limitKey + " has hit quota!")
	case limit.OverQuota:
		return errors.New(limitKey + " has been over quota!")
	}
	return errors.New("Unknown rate limit status error!")
}

func InitRateLimitControl() {

	store, clean, err := redistest.CreateRedis()
	if err != nil {
		panic("Creating Redis Instance Fails")
	}
	defer clean()

	const (
		second = 1
		quota  = 2
	)

	periodLimit = limit.NewPeriodLimit(second, quota, store, "PeriodLimiter")

	const (
		rate  = 1
		burst = 2
	)
	tokenLimit = limit.NewTokenLimiter(rate, burst, store, "TokenLimiter")
}
