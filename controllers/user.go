package controllers

import (
	"encoding/json"
	"fresh/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strconv"
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
	HashPass, err := bcrypt.GenerateFromPassword([]byte(ob.PassWord), 4)
	if err != nil {
		C.Data["json"] = err.Error()
		return
	}
	user.PassWord = string(HashPass)
	_, err = o.Insert(&user)
	if err != nil {
		logs.Info("err", err)
		C.Data["json"] = "数据库写入错误"
		return
	}
	config := `{"username":"908388349@qq.com","password":"alopfzkexgrgbbeg","host":"smtp.qq.com","port":587}`
	email := utils.NewEMail(config)
	email.To = []string{ob.Email}
	email.Subject = "激活邮件"
	// 注意发送给用户的是清货请求地址
	email.Text = "http://localhost:8080/active?id=" + strconv.Itoa(user.Id)
	logs.Info(`这个是{{strconv.Itoa(user.Id)}}`)
	err = email.Send()
	if err != nil {
		logs.Info("错误", err)
		C.Data["json"] = "数据库写入错误"
		return
	}
	C.Data["json"] = "邮件发送成功"

	// 2 校验数据

	// 3 梳理数据

	//  返回视图

}

// Post:登录
func (C *UserController) HandleLogin() {
	// 获取数据
	var ob user
	var err error
	defer C.ServeJSON()
	if err = json.Unmarshal(C.Ctx.Input.RequestBody, &ob); err == nil {
		C.Data["json"] = ob
	} else {
		C.Data["json"] = err.Error()
	}
	logs.Info(ob)
	// orm
	o := orm.NewOrm()
	var user models.User
	user.Name = ob.UserName
	err = o.Read(&user, "Name")
	if err != nil {
		C.Data["json"] = "用户不存在"
		return
	}
	logs.Info(user.PassWord)
	if user.Active != true {
		C.Data["json"] = "没有激活"
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(ob.PassWord))
	if err != nil {
		C.Data["json"] = "密码错误"
		logs.Info(err.Error())
		return
	}

	C.Ctx.SetCookie("userName", "", -1)
	C.SetSession("UserName", ob.UserName)

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

// 用户激活
func (C *UserController) ShowActive() {
	// 获取数据
	id, err := C.GetInt("id")
	if err != nil {
		C.Redirect("http://localhost:8081/#/register?error=要激活的用户不存在", 302)
		return
	}
	// 更新数据
	o := orm.NewOrm()
	var user models.User
	user.Id = id
	err = o.Read(&user)
	if err != nil {
		C.Redirect("http://localhost:8081/#/register?error=要激活的用户不存在", 302)
		return
	}
	user.Active = true
	_, err = o.Update(&user)
	if err != nil {
		C.Redirect("http://localhost:8081/#/register?error=写出数据库错误", 302)
		return
	}

	C.Redirect("http://localhost:8081/#/?success=注册成功,登陆吧", 302)

}
