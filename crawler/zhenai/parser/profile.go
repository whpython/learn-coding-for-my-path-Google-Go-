package parser
import (
	"personal1/crawler/engine"
	"personal1/crawler/model"
	"regexp"
	"strconv"
	//"fmt"
)
// do it by myself
var agere=regexp.MustCompile(
	`<span class="grayL">年龄：</span>([^<]+)</td>`)
var marriagere=regexp.MustCompile(
	`<span class="grayL">婚况：</span>([^<]+)</td>`)
var namere=regexp.MustCompile(
	`<a href="http://album.zhenai.com/u/[0-9]+" [^>]*>([^<]+)</a>`)
var genderre=regexp.MustCompile(//chang the gender for id
	`<span class="grayL">性别：</span>([^<]+)</td>`)
var heightre=regexp.MustCompile(
	`<span class="grayL">身   高：</span>(\d+)</td>`)
var weightre=regexp.MustCompile(
	`<span class="grayL">体   重：</span>(\d+)</td>`)	
var incomere=regexp.MustCompile(
	`<span class="grayL">月   薪：</span>([^<]+)</td>`)
var educationre=regexp.MustCompile(
	`<span class="grayL">学   历：</span>([^<]+)</td>`)
var occupationre=regexp.MustCompile(
	`<span class="grayL">职   业：</span>([^<]+)</td>`)
var hokoure=regexp.MustCompile(
	`<span class="grayL">居住地：</span>([^<]+)</td>`)
var xinzuore=regexp.MustCompile(
	`<span class="grayL">星   座：</span>([^<]+)</td>`)
var housere=regexp.MustCompile(
	`<span class="grayL">住房条件：</span>([^<]+)</td>`)
var carre=regexp.MustCompile(
	`<span class="grayL">是否购车：</span>([^<]+)</td>`)


/* copy form github.com
var agere = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightre = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightre = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomere = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var genderre = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuore = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriagere = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationre = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationre = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokoure = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var housere = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carre = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
*/

func ExtracString(contents []byte,re *regexp.Regexp)string{
	
	match :=re.FindSubmatch(contents)	

	if len(match)>=2 {
		return string(match[1])
	}else{
		return ""
	}

}
func ParseProfile(contents []byte,name string)engine.ParseResult{
	profile :=model.Profile{}
	profile.Name=name

	age,err :=strconv.Atoi(ExtracString(contents,agere))//age is a int
	if err!=nil{
		profile.Age=age
	}
	//fmt.Println(profile.Name)//test
	profile.Marriage=ExtracString(contents,marriagere)//marriage is a string
	//profile.Name=ExtracString(contents,namere)//above have a example
	profile.Gender=ExtracString(contents,genderre)
	height,err :=strconv.Atoi(ExtracString(contents,heightre))
	if err!=nil{
		profile.Height=height
	}
	weight,err :=strconv.Atoi(ExtracString(contents,weightre))
	if err!=nil{
		profile.Weight=weight
	}
	profile.Income=ExtracString(contents,incomere)
	profile.Education=ExtracString(contents,educationre)
	profile.Occupation=ExtracString(contents,occupationre)
	profile.Hokou=ExtracString(contents,hokoure)
	profile.Xinzuo=ExtracString(contents,xinzuore)
	profile.House=ExtracString(contents,housere)
	profile.Car=ExtracString(contents,carre)
	result :=engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
	
}