package controllers

import (
	"encoding/json"
	"fresh/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

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
func (C *UserController) ShowUser() {
	var user = models.OrderInfo{Id: 123, OrderId: "xsxs"}
	C.Data["json"] = user
	C.ServeJSON()
}

type user struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func (C *UserController) HandleUser() {
	var ob user
	var err error
	if err = json.Unmarshal(C.Ctx.Input.RequestBody, &ob); err == nil {
		C.Data["json"] = ob
	} else {
		C.Data["json"] = err.Error()
	}
	C.Ctx.SetCookie("userName", "", -1)
	C.SetSession("UserName", ob.UserName)
	logs.Info(ob)
	C.ServeJSON()
	/*var user  = models.OrderInfo{Id:123,OrderId:"xsxs"}
	C.Data["json"] = user
	C.ServeJSON()*/
}
