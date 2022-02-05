package controllers

import (
	"github.com/astaxie/beego"
)

type LoginBidderController struct {
	beego.Controller
}

func (c *LoginBidderController) Get() {
	c.TplName = "login-bidder.tpl"
}
