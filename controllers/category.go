package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		if err := models.AddCategory(name); err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}

		if err := models.DelCategory(id); err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 302)
		return
	}
	c.Data["IsCategory"] = true
	c.TplName = "category.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
