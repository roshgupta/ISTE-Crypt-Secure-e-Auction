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
func (c *RegistrationBidderController) Post() {
	c.TplName = "thankyouTemplate.tpl"
	c.Data["name"] = c.GetString("userName")
	c.Data["email"] = c.GetString("email")
	c.Data["password"] = c.GetString("password")
}
