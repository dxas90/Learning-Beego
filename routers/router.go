package routers

import (
	"github.com/astaxie/beego"
	"sitepointgoapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitepoint")
	beego.Router("/manage/", &controllers.ManageController{}, "get:Index")
	beego.Router("/manage/new", &controllers.ManageController{}, "get:New;post:Post")
	beego.Router("/manage/show/?:id([0-9]+)", &controllers.ManageController{}, "get:Show")
	beego.Router("/manage/edit/?:id([0-9]+)", &controllers.ManageController{}, "get:Edit")
	beego.Router("/manage/?:id([0-9]+)", &controllers.ManageController{}, "delete:Delete")
	beego.Router("/manage/?:id([0-9]+)", &controllers.ManageController{}, "put:Put")
}
