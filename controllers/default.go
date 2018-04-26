package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	cate := c.Input().Get("cate")
	label := c.Input().Get("label")
	topics, err := models.GetAllTopics(cate, label,  true)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}

	cates, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Categories"] = cates
}
