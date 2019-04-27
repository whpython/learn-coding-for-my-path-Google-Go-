package fetcher
import (
	
	"fmt"
	"net/http"
	"bufio"
	
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	
	
  	"log"
	"io/ioutil" 
)
func determineEncoding(r *bufio.Reader)encoding.Encoding{
	bytes,err:=r.Peek(1024)
	if err!=nil{
		log.Printf("Fetcher error: %v",err)
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}
//Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36
func Fetch(url string) ([]byte, error){
	/*resp, err :=http.Get(url)
	if err!=nil {
		return nil,err
	}
	defer resp.Body.Close()*/


	//come from solution 
	//https://blog.csdn.net/qq_36183935/article/details/80499183
	client :=&http.Client{}
	rep,err:=http.NewRequest("GET",url,nil)
	if err!=nil {
		log.Fatalln(err)
	}
	rep.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36")

	resp,err :=client.Do(rep)
	if err!=nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()


	if resp.StatusCode!=http.StatusOK {
		return nil,
		fmt.Errorf("wrong status code:%d ",resp.StatusCode)
	}
	bodyReader :=bufio.NewReader(resp.Body)
	e :=determineEncoding(bodyReader)
	utf8Reader :=transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
