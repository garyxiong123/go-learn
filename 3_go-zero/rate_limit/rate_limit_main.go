package main

import (
	"flag"
	"fmt"
	"github.com/garyxiong123/go-learn/web/go-zero/common/errorx"
	"github.com/garyxiong123/go-learn/web/go-zero/rate_limit/internal/config"
	"github.com/garyxiong123/go-learn/web/go-zero/rate_limit/internal/handler"
	"github.com/garyxiong123/go-learn/web/go-zero/rate_limit/internal/svc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis/redistest"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"net/http"
)

const (
	PERIOD_LIMIT = "periodLimit"
	TOKEN_LIMIT  = "tokenLimit"
)

var configFile = flag.String("f", "./etc/greet-api.yaml", "the config file")

var periodLimit *limit.PeriodLimit
var tokenLimit *limit.TokenLimiter

var rateLimitType = "periodLimit"

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	// 全局中间件
	server.Use(RateLimitHander)

	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("swagger-ui-dist"))))

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	InitRateLimitControl()

	server.Start()
}

func RateLimitHander(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		requestPath := request.URL.Path
		userId := request.Form.Get("userId")

		if err := RateLimitControlGlobal(requestPath); err != nil {
			http.Error(writer, err.Error(), http.StatusTooManyRequests)
			return
		}

		if err := RateLimitControlByUserId(requestPath, userId); err != nil {
			http.Error(writer, err.Error(), http.StatusTooManyRequests)
			return
		}

		next(writer, request)
	}
}

func RateLimitControlByUserId(path, userid string) error {
	limitKey := fmt.Sprintf("limit:%s:%s", path, userid)
	return RateLimitControl(limitKey)
}

func RateLimitControlGlobal(path string) error {
	limitKey := fmt.Sprintf("limit:%s", path)
	return RateLimitControl(limitKey)
}

func RateLimitControl(limitKey string) error {
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
