package main

import (
	"github.com/astaxie/beego"
	_ "github.com/henly2/hello-go/myproject/routers"
)

func main() {
	beego.Run()
}
