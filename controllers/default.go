package controllers

import "github.com/astaxie/beego"

// MainController struct
type MainController struct {
	beego.Controller
}

// Get -> route to index
func (c *MainController) Get() {
	c.TplName = "index.tpl"
}
