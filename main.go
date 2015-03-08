package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true

	beego.Run()
}
