package controllers

import (
	"github.com/astaxie/beego"
)

type LoginSellerController struct {
	beego.Controller
}

func (c *LoginSellerController) Get() {
	c.TplName = "login-seller.tpl"
}
