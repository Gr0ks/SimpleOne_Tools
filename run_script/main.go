package main
import (
    "fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"flag"
	"net/url"
)
type Data struct {
	Result string `json:"result"`
	Info string `json:"info"`
} 
type Timing struct {
	Before_echo float64 `json:"before_echo"`	
  }
type Respon struct {
	Status string `json:"status"`
	Data Data
	Messages []string `json:"messages"`
	Page_debug_id string `json:"page_debug_id"`
	Timing Timing
  }

type Head struct {
	Par string `json:"par"`	
	Val string `json:"val"`
  }
type InitData struct {
	Url string `json:"url"`	
	Headers []Head
  }

func main() {    
	wordPtr := flag.String("f", "", " = --f=[file.name] name of runing script")
	flag.Parse()
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
			fmt.Println(err)
		}
	initData := InitData{}
	json.Unmarshal([]byte(data), &initData)

	myUrl := initData.Url
	if *wordPtr != "" {
		scriptFromFile, err := ioutil.ReadFile(*wordPtr)
		if err != nil {
			fmt.Println(err)
		}
		script := url.Values{}
		script.Set("script", string(scriptFromFile))
		urlEncodedScript := bytes.NewBuffer([]byte(script.Encode()))
		
		req, err := http.NewRequest("POST", myUrl, urlEncodedScript)
		for _, value := range initData.Headers {
			req.Header.Set(value.Par, value.Val)
		}
		client := &http.Client{}		
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		
		body, _ := ioutil.ReadAll(resp.Body)
		respon := Respon{}
		json.Unmarshal([]byte(body), &respon)
		fmt.Println("RESPONSE: ", string(body))
		fmt.Println()
		fmt.Println("--------------------------------------")
		fmt.Println("SCRIPT SS.INFO: ", string(respon.Data.Info))
		fmt.Println()
	}
}