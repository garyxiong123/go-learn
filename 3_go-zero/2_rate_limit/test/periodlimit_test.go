package test

import (
	"github.com/zeromicro/go-zero/core/limit"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/redis/redistest"
)

func TestPeriodLimit_Take(t *testing.T) {
	testPeriodLimit(t)
}

func TestPeriodLimit_TakeWithAlign(t *testing.T) {
	testPeriodLimit(t, limit.Align())
}

func TestPeriodLimit_RedisUnavailable(t *testing.T) {
	s, err := miniredis.Run()
	assert.Nil(t, err)

	const (
		seconds = 1
		quota   = 5
	)
	l := limit.NewPeriodLimit(seconds, quota, redis.New(s.Addr()), "periodlimit")
	//s.Close()
	val, err := l.Take("first")
	assert.NotNil(t, err)
	assert.Equal(t, 0, val)
}

func testPeriodLimit(t *testing.T, opts ...limit.PeriodOption) {
	store, clean, err := redistest.CreateRedis()
	assert.Nil(t, err)
	defer clean()

	const (
		seconds = 1
		total   = 100
		quota   = 5
	)
	l := limit.NewPeriodLimit(seconds, quota, store, "periodlimit", opts...)
	var allowed, hitQuota, overQuota int
	for i := 0; i < total; i++ {
		val, err := l.Take("sss")
		if err != nil {
			t.Error(err)
		}
		switch val {
		case limit.Allowed:
			allowed++
		case limit.HitQuota:
			hitQuota++
		case limit.OverQuota:
			overQuota++
		default:
			t.Error("unknown status")
		}
	}

	assert.Equal(t, quota-1, allowed)
	assert.Equal(t, 1, hitQuota)
	assert.Equal(t, total-quota, overQuota)
}

func TestQuotaFull(t *testing.T) {
	s, err := miniredis.Run()
	assert.Nil(t, err)

	l := limit.NewPeriodLimit(1, 1, redis.New(s.Addr()), "periodlimit")
	val, err := l.Take("first")
	assert.Nil(t, err)
	assert.Equal(t, limit.HitQuota, val)
}
