package controllers

import (
	"github.com/astaxie/beego"
)

type RegistrationSellerController struct {
	beego.Controller
}

func (c *RegistrationSellerController) Get() {
	c.TplName = "registration-seller.tpl"
}
