package controllers

import (
	_ "Secure-e-Auc/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "landing.tpl"
}
