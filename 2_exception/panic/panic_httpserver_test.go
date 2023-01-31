package panic

import (
	"net/http"
	"testing"
)

func newHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			// 调用 HTTP 处理函数
			next.ServeHTTP(rw, req)
		},
	)
}

func Handle(rw http.ResponseWriter, req *http.Request) {
	panic("sss")
}

// http://localhost:8899/

func Test_Get_Http_Id(t *testing.T) {

	handler := newHandler(http.HandlerFunc(Handle))
	http.ListenAndServe("localhost:8899", handler)
}
