package main

import (
	"github.com/astaxie/beego"
	_ "shFresh/models"
	_ "shFresh/routers"
)

func main() {
	beego.Run()
}
