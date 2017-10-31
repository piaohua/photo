package main

import (
	"photo/controllers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/mmopen", &controllers.PhotoController{})
	beego.Router("/domainurl", &controllers.UrlController{})
	//beego.AutoRouter(&controllers.PhotoController{})
	beego.Run()
}
