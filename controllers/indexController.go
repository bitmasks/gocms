package controllers

import (
	"github.com/astaxie/beego"
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

	//todo  分类列表

	//根据分类的文章列表

	c.TplName = "index.html"
}
