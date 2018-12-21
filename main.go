package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	models "sitepointgoapp/models"
	_ "sitepointgoapp/routers"
)

func init() {
	// This is a dummy change to test Hound
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
	orm.RegisterModel(new(models.Article))
}

func main() {
	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
			ctx.Request.Method = ctx.Input.Query("_method")
		}
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)
	beego.Run()
}
