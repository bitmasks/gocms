package main

import (
	_ "gocms/routers"
	"github.com/astaxie/beego"
)

/**
xorm reverse mysql "root:taorong2012@(192.81.218.34:3306)/gocms?charset=utf8" templates/goxorm
 */
func main() {
	beego.Run()
}

