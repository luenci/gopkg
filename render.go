package gopkg

import (
	"sync"
)

// Result Request the result structure.
type Result struct {
	Data interface{} `json:"data"` // 返回内容
	Msg  string      `json:"msg"`  // 请求结果
	Code int         `json:"code"` // 状态码
}

func (r *Result) reset() {
	r.Data = nil
	r.Msg = ""
	r.Code = 0
}

var resultPool = &sync.Pool{
	New: func() interface{} {
		return new(Result)
	},
}

// Response response method.
func Response(data *Result)*Result{
	result := resultPool.Get().(*Result)
	defer func() {
		result.reset()
		resultPool.Put(result)
	}()
	return &Result{
		Data: data.Data,
		Msg:  data.Msg,
		Code: data.Code,
	}
}
