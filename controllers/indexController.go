package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gocms/service"
)

type IndexController struct {
	beego.Controller
}

// @router /
func (c *IndexController) Get() {
	c.Data["WebsiteName"] = "网站名称"
	c.Data["WebsiteIndexName"] = "网站名称"
	c.Data["WebsiteKeywords"] = "关键词"
	c.Data["WebsiteDescription"] = "网站描述"
	c.Data["Nav"] = service.GetCate()
	c.Data["Block1"] = service.GetGlobalList()
	logs.SetLogger(logs.AdapterFile,`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	logs.Info(c.Data["Block1"])
	c.TplName = "index.html"
}

