package main

import (
	_ "Heimdallr/routers"
	_ "Heimdallr/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
