package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (main *MainController) Get() {
	main.Data["Website"] = "beego.me"
	main.Data["Email"] = "astaxie@gmail.com"
	main.TplName = "index.tpl"
}

func (main *MainController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.TplName = "default/hello-sitepoint.tpl"
}
