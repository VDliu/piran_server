package main

import (
	_ "pirain_server/routers"
	"github.com/astaxie/beego"
	"fmt"
)




func main() {
	fmt.Println("main begin")
	beego.Run()
}

