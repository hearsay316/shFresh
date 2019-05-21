package controllers

import (
	"encoding/json"
	"fresh/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// 注册页面
func (C *UserController) ShowReg() {
	C.TplName = "register.html"
}

// 处理注册数据
func (C *UserController) HandleReg() {
	// 1.获取数据
	userName := C.GetString("user_name")
	pwd := C.GetString("pwd")
	cpwd := C.GetString("cpwd")
	email := C.GetString("email")
	//allow:=C.GetString("allow")
	if userName == "" || pwd == "" || cpwd == "" || email == "" {
		C.Data["error"] = "数据输错"
		C.TplName = "register.html"
	}

	// 2 校验数据

	// 3 梳理数据

	//  返回视图
}

type user struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Remember bool   `json:"remember"`
}

// Post:登录
func (C *UserController) HandleLogin() {
	var ob user
	var err error
	if err = json.Unmarshal(C.Ctx.Input.RequestBody, &ob); err == nil {
		C.Data["json"] = ob
	} else {
		C.Data["json"] = err.Error()
	}
	C.Ctx.SetCookie("userName", "", -1)
	C.SetSession("UserName", ob.UserName)
	C.ServeJSON()
	/*var user  = models.OrderInfo{Id:123,OrderId:"xsxs"}
	C.Data["json"] = user
	C.ServeJSON()*/
}

func (C *UserController) ShowUser() {
	var user = models.OrderInfo{Id: 123, OrderId: "xsxs"}
	C.Data["json"] = user
	C.ServeJSON()
}
func (C *UserController) HandleUser() {
	var user = models.OrderInfo{Id: 123, OrderId: "xsxs"}
	C.Data["json"] = user
	C.ServeJSON()
}
