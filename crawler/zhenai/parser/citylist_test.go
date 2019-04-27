package parser
import(
	"testing"
	//"personal1/crawler/fetcher"
	//"fmt"
	"io/ioutil"

)
func TestParseCityList(t *testing.T)  {

	contents,err :=ioutil.ReadFile("citylist_test_data.html")

	if err!=nil{
		panic(err)
	}
	result :=ParseCityList(contents)
	const resultSize=470

	expectedUrls :=[]string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",

	}
	expectedCites :=[]string{
		"city阿坝",
		"city阿克苏",
		"city阿拉善盟",
	}

	if len(result.Requests)!=resultSize{
		t.Errorf("result should have %d "+"request;but had %d",
	resultSize,len(result.Requests))
	}

	for i,url :=range expectedUrls{
		if result.Requests[i].Url !=url{
			t.Errorf("epected url #%d: %s; but"+"was %s",
		i,url,result.Requests[i].Url)
		}
	}
	if len(result.Items)!=resultSize{
		t.Error("result should have %d "+"request;but had %d",
	resultSize,len(result.Items))
	}
	for i,city :=range expectedCites{
		if result.Items[i].(string) !=city{
			t.Errorf("epected city #%d: %s; but"+"was %s",
		i,city,result.Items[i].(string))
		}
	}
	
	/*if len(result.Items)!=resultSize{
		t.Error("result should have %d "+"request;but had %d",
	resultSize,len(result.Items))
	}*/

	/*fmt.Printf("%v\n",result)   
	
	contents,err :=fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s\n",contents)
	//ParseCityList(contents)
	*/
}
