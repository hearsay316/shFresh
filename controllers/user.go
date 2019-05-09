package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (C *UserController) ShowReg() {
	C.TplName = "register.html"
}
func (C *UserController) HandleRge() {
	// 1.获取数据
	user_name := C.GetString("user_name")
	pwd := C.GetString("pwd")
	cpwd := C.GetString("cpwd")
	email := C.GetString("email")
	//allow:=C.GetString("allow")
	if user_name == "" || pwd == "" || cpwd == "" || email == "" {
		C.Data["error"] = "数据输错"
		C.TplName = "register.html"
	}

	// 2 校验数据

	// 3 梳理数据

	//  返回视图
}
