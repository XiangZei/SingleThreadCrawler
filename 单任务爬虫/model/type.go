package model

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	//请求出来的多个 Request 任务
	Requests []Request
	//解析出来的实体 （例如城市名），是任意类别 （interface{} 类比java Object）
	Items []interface{}
}
