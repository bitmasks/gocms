package controllers

import (
	"github.com/astaxie/beego"
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

	c.Data["Block1"] = service.GetList(1)
	c.Data["Block2"] = service.GetList(4)
	c.TplName = "index.html"
}

