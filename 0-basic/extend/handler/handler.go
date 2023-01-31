package handler

import "net/http"

type Handler interface {
	ServeHTTP( (w http.ResponseWriter, r *http.Request)
}

//类型方法级别的继承
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//方法级别的继承
func (h HandlerFunc) ServeHTTP() (w http.ResponseWriter, r *http.Request) {

	panic("implement me")
}


func (h HandlerFunc) ServeHTTP2() (w http.ResponseWriter, r *http.Request) {

	panic("implement me")
}


type gary struct {
	h  Handler
	h1 HandlerFunc
}
