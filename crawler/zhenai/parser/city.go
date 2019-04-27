package parser
import (
	"personal1/crawler/engine"
	"regexp"
)

const cityre=`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
func ParseCity(contents []byte)engine.ParseResult {

	re :=regexp.MustCompile(cityre)
	matches :=re.FindAllSubmatch(contents,-1)

	result :=engine.ParseResult{}
	for _, m :=range matches {
		result.Items=append(result.Items,"User"+string(m[2]))
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:func(c []byte)engine.ParseResult{
				return ParseProfile(c,string(m[2]))
			},
			
		})
	}
	return result
}