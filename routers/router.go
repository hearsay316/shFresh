package routers

import (
	"fresh/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func init() {
	beego.InsertFilter("/Article/*", beego.BeforeRouter, FiltFunc)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.UserController{}, "get:ShowUser;post:HandleUser")
	beego.Router("/Article/user", &controllers.UserController{}, "get:ShowUser;post:HandleUser")
}

var FiltFunc = func(ctx *context.Context) {
	logs.Info("app")
	Username := ctx.Input.Session("UserName")
	logs.Info(ctx.GetCookie("UserName"), "wdwdwd")
	if Username == nil {
		//ctx.Redirect(302, "/")
		ctx.ResponseWriter.WriteHeader(404)
		ctx.WriteString("'_xsrf' argument missing from POST")
	}
}
