package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
)

type UrlController struct {
	beego.Controller
}

func (this *UrlController) Get() {
	urlStr := this.GetString("url")
	purl, err := url.Parse(urlStr)
	//fm.Println(purl.Query().Encode())
	if err != nil {
		fmt.Printf("err -> %v, qurl -> %s\n", err, purl)
		beego.Trace("err, qurl, ", err, purl)
		this.Ctx.Output.Body([]byte{})
		return
	}
	qurl := purl.String()
	fmt.Printf("url -> %s\n", qurl)
	beego.Trace("url, ", qurl)
	resp, err := http.Get(qurl)
	if err != nil {
		fmt.Printf("err -> %v, url -> %s\n", err, qurl)
		beego.Trace("err, url, ", err, qurl)
	}
	if resp == nil {
		this.Ctx.Output.Body([]byte{})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err -> %v\n", err)
		beego.Trace("err, ", err)
	}
	this.Ctx.Output.Body(body)
}

func (this *UrlController) Post() {
	urlStr := this.GetString("url")
	purl, err := url.Parse(urlStr)
	//fm.Println(purl.Query().Encode())
	if err != nil {
		fmt.Printf("err -> %v, qurl -> %s\n", err, purl)
		beego.Trace("err, qurl, ", err, purl)
		this.Ctx.Output.Body([]byte{})
		return
	}
	qurl := purl.String()
	fmt.Printf("url -> %s\n", qurl)
	beego.Trace("url, ", qurl)
	resp, err := http.Get(qurl)
	if err != nil {
		fmt.Printf("err -> %v, url -> %s\n", err, qurl)
		beego.Trace("err, url, ", err, qurl)
	}
	if resp == nil {
		this.Ctx.Output.Body([]byte{})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err -> %v\n", err)
		beego.Trace("err, ", err)
	}
	this.Ctx.Output.Body(body)
}
