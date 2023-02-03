//Package serializer ...
package types

// Result 基础序列化器
type Result struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	TrackID string `json:"track_id"`
}

func BuildOkResult(item interface{}) Result {
	return Result{}
}

// BuildListResponse 带有总数的列表构建器
func BuildListResult(items interface{}, total uint) Result {
	return Result{
		Status: 200,
		Msg:    "ok",
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
