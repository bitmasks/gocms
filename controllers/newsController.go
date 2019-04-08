package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"gocms/service"
	"strconv"
)

type NewsController struct {
	beego.Controller
}


// @router /:nav/
func(c *NewsController) List()  {
	nvId := c.Ctx.Input.Param(":nav")
	info := service.GetCategoryInfo(nvId)
	c.Data["Title"] = info.CategoryName
	c.Data["datalist"] = service.GetDataList(nvId)
	c.TplName = "list.html"
}

// @router /:nav/:id
func (c *NewsController) Info()  {
	id,_ := strconv.ParseInt(c.Ctx.Input.Param(":id"),10,64)
	data,_ := service.GetDataInfo(id)
	fmt.Println(id)
	c.Data["Title"] = data.Title
	c.Data["Desc"] = data.Desc
	c.Data["Content"] = data.Content
	c.TplName = "info.html"
}