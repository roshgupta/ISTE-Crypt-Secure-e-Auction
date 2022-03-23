package controllers

// import (
// 	models "Secure-e-Auc/models"
// 	"fmt"

// 	"github.com/astaxie/beego"
// )

// type RegistrationBidderController struct {
// 	beego.Controller
// }

// func (c *RegistrationBidderController) Get() {
// 	c.TplName = "registration-bidder.tpl"
// }

// func (c *RegistrationBidderController) Post() {
// 	var user_email = c.GetString("email")
// 	var user_password = c.GetString("password")
// 	fmt.Print(user_email)
// 	models.NewBidder(user_email, user_password)
// 	c.Redirect("/login-bidder", 302)
// }
