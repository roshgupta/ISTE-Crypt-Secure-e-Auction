package controllers

import (
	"Secure-e-Auc/models"
	"fmt"

	"github.com/astaxie/beego"
)

type LoginBidderController struct {
	beego.Controller
}

func (c *LoginBidderController) Get() {
	c.TplName = "login-bidder.tpl"
}

func (c *LoginBidderController) Post() {
	// c.ServeJSON()
	var user_email = c.GetString("email")
	var user_password = c.GetString("password")

	user, err := models.LoginBidder(user_email, user_password)
	fmt.Print(err != nil)
	if user.Email == user_email {
		c.Redirect("/bidder", 302)
	}
	c.Redirect("/login-bidder", 302)
}
