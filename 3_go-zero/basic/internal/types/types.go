// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me|error|panic"`
}

type Response struct {
	Message string `json:"message"`
}