package parser

import (
	"regexp"
	"study.com/Day15/zhenai/moudle"
	"study.com/Day15/zhenai/engine"
	"fmt"
	"strconv"
)

var nameRegexp = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
var ageRegexp = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marrRegexp = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var sexRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var salaryRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var jiGuanRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var workAdrressRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var guessRe = regexp.MustCompile(`<a class="exp-user-name" target="_blank"[^=]+="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)

/*
	用户信息解析器
 */
func UserInfoParser(str string) (r engine.RequestResult) {
	user := moudle.UserInfo{}

	name := regexMatch(str, nameRegexp)
	user.Name = name

	ageStr := regexMatch(str, ageRegexp)
	age, err := strconv.Atoi(ageStr)
	if err == nil {
		user.Age = age
	}

	marr := regexMatch(str, marrRegexp)
	user.Marr = marr

	height := regexMatch(str, heightRe)
	user.Height = height

	sex := regexMatch(str, sexRe)
	user.Sex = sex

	salary := regexMatch(sex, salaryRe)
	user.Salary = salary

	weight := regexMatch(str, weightRe)
	user.Weight = weight

	occupation := regexMatch(str, occupationRe)
	user.Ocuppattion = occupation

	jiGuan := regexMatch(str, jiGuanRe)
	user.JianGuan = jiGuan

	education := regexMatch(str, educationRe)
	user.Education = education

	house := regexMatch(str, houseRe)
	user.House = house

	car := regexMatch(str, carRe)
	user.Car = car

	workAdrress := regexMatch(str, workAdrressRe)
	user.WorkAdrress = workAdrress

	r.Items = append(r.Items, user)

	//猜你喜欢
	allGuess := guessRe.FindAllStringSubmatch(str, -1)
	//fmt.Println("allGuess: ",allGuess)
	for _, v := range allGuess {
		r.R = append(r.R, engine.Request{
			Url:   v[1],
			Parse: UserInfoParser,
		})
	}

	fmt.Println("user:", user)

	return r

}

func regexMatch(str string, regex *regexp.Regexp) string {
	subStrArr := regex.FindStringSubmatch(str)
	if len(subStrArr) >= 2 {
		return subStrArr[1]
	}
	return ""
}
