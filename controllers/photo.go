package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) Get() {
	//p := this.GetString("photo")
	//fmt.Printf("p -> %#v \n", p)
	//b := this.Ctx.Input.RequestBody
	//b := this.Ctx.Request.RequestURI
	//fmt.Printf("b -> %#v \n", b)
	//str := strings.Split(this.GetString("photo"), ".jpg")
	url := "http://wx.qlogo.cn/mmopen/" + strings.Split(this.GetString("photo"), ".jpg")[0]
	//url := "http://wx.qlogo.cn/mmopen/" + this.GetString("photo")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err -> %v, url -> %s\n", err, url)
	}
	if resp == nil {
		this.Ctx.Output.Body([]byte{})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err -> %v\n", err)
	}
	this.Ctx.Output.Body(body)
}
