package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (t *TopicController) Get() {
	t.Data["IsTopic"] = true
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	t.TplName = "topic.html"
	topics, err := models.GetAllTopics("", "", false)
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
	category := t.Input().Get("category")
	label := t.Input().Get("label")
	tid := t.Input().Get("tid")

	// 获取附件
	_, fh, err := t.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		attachment = fh.Filename
		beego.Info(attachment)
		err = t.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err := models.AddTopic(title, category, label, content, attachment)
		if err != nil {
			beego.Error(err)
		}
	} else {
		err := models.ModifyTopic(tid, title, category, label, content, attachment)
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
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	/*
	reqUrl := t.Ctx.Request.RequestURI
	i := strings.LastIndex(reqUrl, "/")
	tid := reqUrl[i+1:]
	*/

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
	t.Data["Labels"] = strings.Split(topic.Labels, " ")
	t.TplName = "topic_view.html"

	repies, err := models.GetAllComments()
	if err != nil {
		beego.Error(err)
		return
	}

	t.Data["Replies"] = repies
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
