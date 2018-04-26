package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (r *ReplyController) Add() {
	tid := r.Input().Get("tid")
	name := r.Input().Get("nickname")
	content := r.Input().Get("content")
	if err := models.AddComment(tid, name, content); err != nil {
		beego.Error(err)
	}

	r.Redirect("/topic/view/"+tid, 302)
	return
}

func (r *ReplyController) Delete() {
	tid := r.Input().Get("tid")
	rid := r.Input().Get("rid")
	if err := models.DelComment(rid); err != nil {
		beego.Error(err)
	}

	r.Redirect("/topic/view/"+tid, 302)
	return
}
