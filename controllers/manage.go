package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	models "sitepointgoapp/models"
	"strconv"
)

type ManageController struct {
	beego.Controller
}

// Prepare implements Prepare method for baseRouter.
func (this *ManageController) Prepare() {
	this.Layout = "basic-layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
}

func (this *ManageController) Index() {
	this.TplName = "manage/list.tpl"

	flash := beego.ReadFromRequest(&this.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		this.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		this.Data["notices"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	var articles []*models.Article
	num, err := o.QueryTable("articles").All(&articles)

	if err != orm.ErrNoRows && num > 0 {
		this.Data["records"] = articles
	}
}

func (this *ManageController) Show() {
	this.TplName = "manage/show.tpl"

	// convert the string value to an int
	articleId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	o.Using("default")

	article := models.Article{Id: articleId}

	err := o.Read(&article)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(article.Id, article.Name)
	}
	this.Data["record"] = article
}

func (this *ManageController) New() {
	this.Data["Form"] = &models.Article{}
	this.TplName = "manage/add.tpl"
}

func (this *ManageController) Edit() {

	// convert the string value to an int
	articleId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	o.Using("default")

	article := models.Article{Id: articleId}

	err := o.Read(&article)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(article.Id, article.Name)
	}
	this.Data["Article"] = article
	this.TplName = "manage/edit.tpl"
}

func (this *ManageController) Post() {

	flash := beego.ReadFromRequest(&this.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		this.Data["flash"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	article := models.Article{}
	if err := this.ParseForm(&article); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {
		valid := validation.Validation{}
		valid.Valid(article)

		if !valid.HasErrors() {
			searchArticle := models.Article{Name: article.Name}
			beego.Debug("Article name supplied:", article.Name)
			err = o.Read(&searchArticle)
			beego.Debug("Err:", err)
			flash := beego.NewFlash()

			if err == orm.ErrNoRows || err == orm.ErrMissPK {
				beego.Debug("No article found matching details supplied. Attempting to insert article: ", article)
				id, err := o.Insert(&article)
				if err == nil {
					msg := fmt.Sprintf("Article inserted with id: %d", id)
					beego.Debug(msg)
					flash.Notice(msg)
					flash.Store(&this.Controller)
					this.Redirect(beego.URLFor("ManageController.Show", ":id", id), 302)
				} else {
					msg := fmt.Sprintf("Couldn't insert new article. Reason: ", err)
					beego.Debug(msg)
					flash.Error(msg)
					flash.Store(&this.Controller)
				}
			} else {
				beego.Debug("Article found matching details supplied. Cannot insert")
			}
			this.Data["Article"] = article
			this.TplName = "manage/add.tpl"
		} else {
			this.Data["Errors"] = valid.ErrorsMap
			for _, err := range valid.Errors {
				fmt.Println(err.Key, err.Message)
			}
			beego.Error("Form didn't validate.")
			this.Data["Article"] = article
			this.TplName = "manage/add.tpl"
		}
	}
}

func (this *ManageController) Put() {
	o := orm.NewOrm()
	o.Using("default")
	flash := beego.NewFlash()
	id := this.Ctx.Input.Param(":id")
	var article models.Article
	// convert the string value to an int
	if articleId, err := strconv.Atoi(this.Ctx.Input.Param(":id")); err == nil {
		article = models.Article{Id: articleId}
		// attempt to load the record from the database
		if o.Read(&article) == nil {
			// set the Client and Url properties to arbitrary values
			if err := this.ParseForm(&article); err != nil {
				beego.Error("Couldn't parse the form. Reason: ", err)
			} else {
				valid := validation.Validation{}
				valid.Valid(article)

				if !valid.HasErrors() {
					if num, err := o.Update(&article); err == nil {
						flash.Notice("Record Was Updated.")
						flash.Store(&this.Controller)
						beego.Info("Record Was Updated. ", num)
						// redirect afterwards
						this.Redirect(beego.URLFor("ManageController.Show", ":id", id), 302)
					}
				}
			}
		} else {
			flash.Notice("Record Was NOT Updated.")
			flash.Store(&this.Controller)
			beego.Error("Couldn't find article matching id: ", articleId)
		}
	} else {
		flash.Notice("Record Was NOT Updated.")
		flash.Store(&this.Controller)
		beego.Error("Couldn't convert id from a string to a number. ", err)
	}
	this.Data["Article"] = article
	this.TplName = "manage/edit.tpl"
}

func (this *ManageController) Delete() {

	// convert the string value to an int
	articleId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))

	o := orm.NewOrm()
	o.Using("default")

	article := models.Article{}

	// Check if the article exists first
	if exist := o.QueryTable(article.TableName()).Filter("Id", articleId).Exist(); exist {
		if num, err := o.Delete(&models.Article{Id: articleId}); err == nil {
			beego.Info("Record Deleted. ", num)
		} else {
			beego.Error("Record couldn't be deleted. Reason: ", err)
		}
	} else {
		beego.Info("Record Doesn't exist.")
	}
	this.Redirect(beego.URLFor("ManageController.Index"), 302)
}
