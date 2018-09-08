package parser

import (
	"study.com/0504-parser-14/engine"
	"study.com/0504-parser-14/model"
	"regexp"
	"strconv"
	"fmt"
)

/*
type UserInfo struct {
	Name        string //姓名
	Age         int    //年龄
	Height      int    //身高
	Weight      int    //体重
	Sex         string //性别
	Salary      string //月收入
	Marr        string //婚姻状况
	Ocuppattion string //职业
	WorkAdrress string //工作地
	JianGuan    string //籍贯
	Education   string //学历
	House       string //房
	Car         string //车
}

*/

var nameRegexp = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
var ageRegexp = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marrRegexp = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>(\d+)CM</td>`)
var sexRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var salaryRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">(\d+)KG</span></td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var jiGuanRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var workAdrressRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var guessRe = regexp.MustCompile(`<a class="exp-user-name" target="_blank"[^=]+="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)



func UserInfoParser(contents []byte) (result engine.RequestResult) {

	userInfo := model.UserInfo{}

	// 年龄
	age, err := strconv.Atoi(UserInfoMatch(contents, ageRegexp))
	if err == nil {
		userInfo.Age = age
	}
	// 姓名
	userInfo.Name = UserInfoMatch(contents, nameRegexp)
	// 婚姻
	userInfo.Marr = UserInfoMatch(contents, marrRegexp)
	// 身高
	height, err := strconv.Atoi(UserInfoMatch(contents, heightRe))
	if err == nil {
		userInfo.Height = height
	}
	//体重
	weight, err := strconv.Atoi(UserInfoMatch(contents, weightRe))
	if err == nil {
		userInfo.Weight = weight
	}
	//职业
	userInfo.Ocuppattion = UserInfoMatch(contents, occupationRe)
	//工作地
	userInfo.WorkAdrress = UserInfoMatch(contents, workAdrressRe)
	//有没有房
	userInfo.House = UserInfoMatch(contents, houseRe)
	//有没有车
	userInfo.Car = UserInfoMatch(contents, carRe)
	//性别
	userInfo.Sex = UserInfoMatch(contents, sexRe)
	//教育
	userInfo.Education = UserInfoMatch(contents, educationRe)
	//籍贯
	userInfo.JianGuan = UserInfoMatch(contents, jiGuanRe)
	//月收入
	userInfo.Salary = UserInfoMatch(contents, salaryRe)

	fmt.Printf("%s\n", userInfo)
	result.Items = append(result.Items, userInfo)


	guessData := guessRe.FindAllSubmatch(contents,-1)

	for _,userData := range guessData {
		//fmt.Printf("--------%s\n",userData)
		result.R = append(result.R,engine.Request{
			Url: string(userData[1]),
			ParserFunc: UserInfoParser,
		})
	}

	return
}

/*
[[<a class="exp-user-name" target="_blank"
												href="http://album.zhenai.com/u/1662466794">麦芽糖</a> http://album.zhenai.com/u/1662466794 麦芽糖] [<a class="exp-user-name" target="_blank"
												href="http://album.zhenai.com/u/1388137348">缘来就是你</a> http://album.zhenai.com/u/1388137348 缘来就是你] [<a class="exp-user-name" target="_blank"
												href="http://album.zhenai.com/u/107905291">Dream</a> http://album.zhenai.com/u/107905291 Dream]]
*/


func UserInfoMatch(c []byte, rege *regexp.Regexp) string {
	data := rege.FindSubmatch(c)
	if len(data) >= 2 {
		return string(data[1])
	}
	return ""
}
