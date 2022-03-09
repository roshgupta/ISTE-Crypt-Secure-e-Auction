package controllers

import (
	models "Secure-e-Auc/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type RegistrationBidderController struct {
	beego.Controller
}

func (c *RegistrationBidderController) Get() {
	c.TplName = "registration-bidder.tpl"
}

func (c *RegistrationBidderController) Post() {
	var user_email = c.GetString("email")
	var user_password = c.GetString("password")
	fmt.Print(user_email)
	o := orm.NewOrm()
	o.Using("default")
	bidder := models.Bidder{}
	bidder.Email = user_email
	bidder.Password = user_password
	bidder.Insert()
	c.Redirect("/main", 302)
}
