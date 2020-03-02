package main
import (
    "fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"runtime"
	"os/exec"	
)

type Respon struct {
	Status string `json:"status"`
	Data []string `json:"data"`	
	Error []string `json:"ERROR"`
  }
type Auth struct {
	Login string `json:"login"`	
	Pass string `json:"pass"`
  }
type Files struct {
	HtmlTemplate string `json:"htmlTemplate"`	
	Css string `json:"css"`
	ServerScript string `json:"serverScript"`	
	ClientScript string `json:"clientScript"`
  }  
type InitData struct {
	WidgetUrl string `json:"widgetRestUrl"`	
	WidgetInstanceUrl string `json:"widgetInstanceUrl"`
	Auth Auth
	Files Files
  }

type Widget struct {
	HtmlTemplate string `json:"template"`	
	Css string `json:"css"`
	ServerScript string `json:"server_script"`	
	ClientScript string `json:"client_script"`
  }

func main() {
	data, err := ioutil.ReadFile("conf.json")
	if err != nil {
			fmt.Println(err)
		}
	initData := InitData{}
	json.Unmarshal([]byte(data), &initData)
// input Widget data from files	
	htmlFile, err := ioutil.ReadFile(initData.Files.HtmlTemplate)
	if err != nil {
		fmt.Println(err)
	}
	cssFile, err := ioutil.ReadFile(initData.Files.Css)
	if err != nil {
		fmt.Println(err)
	}
	serverFile, err := ioutil.ReadFile(initData.Files.ServerScript)
	if err != nil {
		fmt.Println(err)
	}
	clientFile, err := ioutil.ReadFile(initData.Files.ClientScript)
	if err != nil {
		fmt.Println(err)
	}
// Create Widget structure to send it
	widgetJsonBody := &Widget{
		HtmlTemplate: string(htmlFile),
		Css: string(cssFile),
		ServerScript: string(serverFile),
		ClientScript: string(clientFile),
	}
// prepare request data
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(widgetJsonBody)
	//reqUrl := 
// create request	
	req, err := http.NewRequest("PUT", initData.WidgetUrl, reqBody)
	req.Header.Set("content-type", "application/json")
	req.SetBasicAuth(initData.Auth.Login, initData.Auth.Pass)
	client := &http.Client{}
// send request	
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
// parse response	
	body, _ := ioutil.ReadAll(resp.Body)
	respon := Respon{}
	json.Unmarshal([]byte(body), &respon)
// write result to console	
	if string(respon.Status) == "OK" {
		fmt.Println("Widget saving: ", respon.Status)
		if len(initData.WidgetInstanceUrl)>0{
			// open widget instance in new tab
			err := open(initData.WidgetInstanceUrl)
			if err != nil {
				panic(err)
			}
			fmt.Println("Opening page: ", initData.WidgetInstanceUrl)
		}		
	} else {
		fmt.Println("Widget saving: ERROR")
	}
}

func open(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}