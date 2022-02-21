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

func (c *RegistrationSellerController) Post() {
	c.TplName = "thankyouTemplate.tpl"

	c.Data["name"] = c.GetString("userName")
	c.Data["email"] = c.GetString("email")
	c.Data["password"] = c.GetString("password")

}
