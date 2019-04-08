package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["WebsiteName"] = "网站名称"
	c.Data["WebsiteKeywords"] = "关键词"
	c.Data["WebsiteDescription"] = "网站描述"
	c.TplName = "index.tpl"
}
