package htpserver

import (
	"net/http"
	"testing"
	"time"
)

func Test_Connection_Storm_No_Time_Out(t *testing.T) {

	srv := &http.Server{
		Addr: "127.0.0.1:8899",
	}

	srv.ListenAndServe()

	//check connection

}

func Test_Connection_Storm(t *testing.T) {

	//handler := newHandler(http.HandlerFunc(Handle))

	srv := &http.Server{
		Addr:         "127.0.0.1:8899",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv.ListenAndServe()

	println(srv)

	//http.ListenAndServe("localhost:8899", handler)
}
