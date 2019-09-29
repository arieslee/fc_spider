package main
import (
	"bytes"
	"github.com/gogf/gf/frame/g"
	_ "github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
)
var intChan = make(chan int)
func GetHtml() (string,error){
	var codeStr string
	url := g.Config().GetString("app.url")
	if response, err := ghttp.Get(url); err != nil {
		return codeStr, err
	} else {
		defer response.Close()
		data := response.ReadAllString()
		strings := gstr.Explode("/images/info/public/ball/ball_red.gif",data)
		codeStack := gstr.Explode("<!-- -->", strings[2])
		match, _ := gregex.MatchAllString(`<li class="ball_red">(\d+)<\/li>`, codeStack[0])
		var buffer bytes.Buffer
		for _,val := range match{
			buffer.WriteString(val[1])
			buffer.WriteString(" ")
		}
		blue, _ := gregex.MatchString(`<li class="ball_blue">(\d+)<\/li>`, codeStack[0])
		buffer.WriteString(blue[1])
		codeStr = buffer.String()
		return codeStr,nil
	}
}
func appRun()  {
	s := g.Server()
	s.BindHandler("/win", func(r *ghttp.Request) {
		code, _ := GetHtml()
		r.Response.Writeln(code)
	})
	s.SetPort(8199)
	s.Run()
}

func main() {
    appRun()
}