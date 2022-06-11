package result

type Result struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Count   uint64      `json:"count"`

	// 错误
	Error string `json:"error"`
	// 错误码
	Code string `json:"code"`
}

/*
	统一返回接口
*/
func result(success bool) *Result {
	return &Result{
		Success: success,
		Data:    make(map[string]interface{}, 0),
	}
}

//WithData 设置数据
func (r *Result) WithData(data interface{}) *Result {
	r.Data = data
	return r
}
