package controllers

import (
	"github.com/astaxie/beego"
)

type RegistrationBidderController struct {
	beego.Controller
}

func (c *RegistrationBidderController) Get() {
	c.TplName = "registration-bidder.tpl"
}
