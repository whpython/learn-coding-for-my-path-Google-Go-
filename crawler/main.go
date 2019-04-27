package main
import (
	"personal1/crawler/engine"
	"personal1/crawler/zhenai/parser"
	
)//non-main function must loacal in main function fprward (in vs code)

func main()  {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
	
}
