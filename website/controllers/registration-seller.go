package controllers

import (
	models "Secure-e-Auc/models"
	"fmt"

	"github.com/astaxie/beego"
)

type RegistrationSellerController struct {
	beego.Controller
}

func (c *RegistrationSellerController) Get() {
	c.TplName = "registration-seller.tpl"
}

func (c *RegistrationSellerController) Post() {
	var user_email = c.GetString("email")
	var user_password = c.GetString("password")
	fmt.Print(user_email)
	models.NewSeller(user_email, user_password)
	c.Redirect("/login-seller", 302)
}
