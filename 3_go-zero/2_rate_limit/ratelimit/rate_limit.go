package ratelimit

import (
	"github.com/shirou/gopsutil/host"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

var HostID string

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

		if err := StartRateLimitControlByPeroidLimit(limitParam); err != nil {
			http.Error(writer, err.Error(), http.StatusTooManyRequests)
			return
		}

		//if err := StartRateLimitControlByTokenLimit(limitParam); err != nil {
		//	http.Error(writer, err.Error(), http.StatusTooManyRequests)
		//	return
		//}

		next(writer, request)
	}
}

func InitRateLimitControl() {

	redisInstance := redis.New("127.0.0.1:6379")

	localHostID, _ := host.HostID()
	HostID = localHostID

	InitRateLimitControlByPeriod(redisInstance)
	InitRateLimitControlByToken(redisInstance)
}
