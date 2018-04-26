package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type TopicController struct {
	beego.Controller
}

func (t *TopicController) Get() {
	t.Data["IsTopic"] = true
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	t.TplName = "topic.html"
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		t.Data["Topics"] = topics
	}
}

func (t *TopicController) Post() {
	if !checkAccount(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}

	title := t.Input().Get("title")
	content := t.Input().Get("content")
	tid := t.Input().Get("tid")
	if len(tid) == 0 {
		err := models.AddTopic(title, content)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.ModifyTopic(tid, title, content)
		if err != nil {
			beego.Error(err)
		}
	}
	t.Redirect("/topic", 302)
	return
}

func (t *TopicController) Add() {
	if !checkAccount(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	t.TplName = "topic_add.html"
}

func (t *TopicController) View() {
	if !checkAccount(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	p := t.Ctx.Input.Params()
	var tid string
	if v, ok := p["0"]; !ok {
		t.Redirect("/", 302)
		return
	} else {
		tid = v
	}
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		t.Redirect("/", 302)
		return
	}
	t.Data["Topic"] = topic
	t.Data["Tid"] = tid
	t.TplName = "topic_view.html"
}

func (t *TopicController) Modify() {
	if !checkAccount(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	tid := t.Input().Get("tid")
	if len(tid) == 0 {
		t.Redirect("/topic", 302)
		return
	}
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		t.Redirect("/", 302)
		return
	}
	t.Data["Topic"] = topic
	t.Data["Tid"] = tid
	t.TplName = "topic_modify.html"
}

func (t *TopicController) Delete() {
	if !checkAccount(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	p := t.Ctx.Input.Params()
	var tid string
	if v, ok := p["0"]; !ok {
		t.Redirect("/", 302)
		return
	} else {
		tid = v
	}
	if err := models.DelTopic(tid); err != nil {
		beego.Error(err)
	}
	t.Redirect("/topic", 302)
	return
}
