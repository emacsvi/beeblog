package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{}) // 智能路由
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.Router("/topic", &controllers.TopicController{})
	os.Mkdir("attachment", os.ModePerm)
	// beego.SetStaticPath("/attachment", "attachment")
	beego.Router("/attachment/:all", &controllers.AttachController{})
}
