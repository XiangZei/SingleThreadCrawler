package parser

import (
	"SingleThreadCrawler/singleThreadCrawler/model"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

//citylist 的 parseFunc func([]byte) ParseResult
//解析种子页面 -获取城市列表
func ParseCityList(contents []byte) model.ParseResult {
	result := model.ParseResult{}
	//正则表达式：（）用于提取
	rg := regexp.MustCompile(cityListRe)
	allSubmatch := rg.FindAllSubmatch(contents, -1)

	//便利每一个城市的匹配字段（城市Url和城市名），并且将Url和城市解析器封装为一个Request
	//最后将该Request添加到ParseResult中

	for _, m := range allSubmatch {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, model.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
