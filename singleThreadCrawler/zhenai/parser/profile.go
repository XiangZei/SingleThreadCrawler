package parser

import (
	"SingleThreadCrawler/singleThreadCrawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div data-v-bff6f798="" class="m-btn purple">([\d])+岁</div>`)

var incomeRe = regexp.MustCompile(`<div data-v-bff6f798="" class="m-btn purple">月收入:(.*?)</div>`)

// 解析单个人的主页
func ParseProfile(contents []byte, name string) model.ParseResult {
	profile := model.Profile{}

	// 1. 年龄
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	// 2. 月收入
	profile.Income = extractString(contents, incomeRe)

	// 3. 姓名
	profile.Name = name

	result := model.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(body []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(body) // 只找到第一个match的
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
