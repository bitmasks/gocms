package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gocms/models"
	"gocms/service"
)

type IndexController struct {
	beego.Controller
}

// @router /
func (c *IndexController) Get() {

	//common
	c.Data["WebsiteName"] = beego.AppConfig.String("WebsiteName")
	c.Data["WebsiteIndexName"] = beego.AppConfig.String("WebsiteIndexName")
	c.Data["WebsiteKeywords"] = beego.AppConfig.String("WebsiteKeywords")
	c.Data["WebsiteDescription"] = beego.AppConfig.String("WebsiteDescription")
	logs.SetLogger(logs.AdapterFile, `{"filename":"debug.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)

	var category = service.GetCate()

	c.Data["Nav"] = category

	//first  block
	var firstScreen = service.GetGlobalList(33)
	c.Data["Block1"] = firstScreen[0:1]
	c.Data["Block2"] = firstScreen[2:6]
	c.Data["Block3"] = firstScreen[7:11]
	c.Data["Block4"] = firstScreen[12:17]
	c.Data["Block5"] = firstScreen[18:32]

	type IndexBlock struct {
		data         [5][]models.Data
		categoryId   string
		categoryName string
	}

	type indexScreenData struct {
		count  int
		screen [5]IndexBlock
	}
	var cat = indexScreenData{}
	for i, v := range category {

		//data from  per screen
		var currScreenData = service.GetList(string(v.CategoryId), 32)

		var block = IndexBlock{}
		block.categoryName = v.CategoryName
		block.categoryId = v.CategoryId
		block.data[0] = currScreenData[0:1]
		block.data[1] = currScreenData[2:6]
		block.data[2] = currScreenData[7:11]
		block.data[3] = currScreenData[12:17]
		block.data[4] = currScreenData[18:32]
		cat.screen[i] = block

		//count for  data of per screen
		cat.count = i

	}

	logs.Info(cat)

	c.Data["screen"] = cat

	//view
	c.TplName = "index.html"
}
