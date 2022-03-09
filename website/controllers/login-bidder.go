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

// func (c *LoginBidderController) Post() {
// 	var email = c.GetString("email")
// 	var password = c.GetString("password")
// 	o := orm.NewOrm()
// 	o.Using("default")

// 	bidder := models.Bidder{}
// 	email, password := o.insert(&bi)
// }
