package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"os"
	"qiniupkg.com/x/url.v7"
)

type AttachController struct {
	beego.Controller
}

func (a *AttachController) Get() {
	// URI=/attachment/%2342%223%222.avi
	filePath, err := url.QueryUnescape(a.Ctx.Request.RequestURI[1:])
	if err != nil {
		beego.Error(err)
		a.Ctx.WriteString(err.Error())
		return
	}
	f, err := os.Open(filePath)
	if err != nil {
		beego.Error(err)
		a.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(a.Ctx.ResponseWriter, f)
	if err != nil {
		beego.Error(err)
		a.Ctx.WriteString(err.Error())
		return
	}
}
