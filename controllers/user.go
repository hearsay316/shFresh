package controllers

import (
	"encoding/json"
	"fresh/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"regexp"
)

type UserController struct {
	beego.Controller
}
type user struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Remember bool   `json:"remember"`
}

/*user_name:''  ,
pwd:'',
cpwd:"",
email:""*/
type userReg struct {
	UserName  string `json:"username"`
	PassWord  string `json:"pwd"`
	CPassWord string `json:"cpwd"`
	Email     string `json:"email"`
}

// 注册页面
func (C *UserController) ShowReg() {
	C.TplName = "register.html"
}

// 处理注册数据
func (C *UserController) HandleReg() {
	// 1.获取数据
	/*	userName := C.GetString("user_name")
		pwd := C.GetString("pwd")
		cpwd := C.GetString("cpwd")
		email := C.GetString("email")
		//allow:=C.GetString("allow")
		if userName == "" || pwd == "" || cpwd == "" || email == "" {
			C.Data["error"] = "数据输错"
			C.TplName = "register.html"
		}*/
	defer C.ServeJSON()
	var ob userReg
	var err error
	if err = json.Unmarshal(C.Ctx.Input.RequestBody, &ob); err == nil {
		if ob.PassWord == "" || ob.UserName == "" || ob.CPassWord == "" || ob.Email == "" {
			C.Data["json"] = "数据不完整,请重新填写"
			return
		}
		if ob.PassWord != ob.CPassWord {
			C.Data["json"] = "两次输入密码有问题,请重新填写"
			return
		}
		reg, _ := regexp.Compile("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
		res := reg.FindString(ob.Email)
		if res == "" {
			C.Data["json"] = "邮箱格式不正确"
			return
		}

	} else {
		C.Data["json"] = err.Error()
		return
	}

	o := orm.NewOrm()
	var user models.User
	user.Name = ob.UserName
	user.Email = ob.Email
	user.PassWord = ob.PassWord
	user.Active = true
	_, err = o.Insert(&user)
	if err != nil {
		logs.Info("err", err)
		C.Data["json"] = "数据库写入错误"
		return
	}

	// 2 校验数据

	// 3 梳理数据

	//  返回视图

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
