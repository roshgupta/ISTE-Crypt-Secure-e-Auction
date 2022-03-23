package controllers

// import (
// 	"Secure-e-Auc/models"
// 	"fmt"

// 	"github.com/astaxie/beego"
// )

// var user_id int = -1

// type LoginBidderController struct {
// 	beego.Controller
// }

// func (c *LoginBidderController) Get() {
// 	c.TplName = "login-bidder.tpl"
// }

// func (c *LoginBidderController) Post() {
// 	// c.ServeJSON()
// 	var user_email = c.GetString("email")
// 	var user_password = c.GetString("password")

// 	user, err := models.LoginBidder(user_email, user_password)
// 	fmt.Print(err != nil)
// 	if user.Password == user_password {
// 		user_id = int(user.Id)
// 		c.Redirect("/bidder", 302)
// 	}
// 	c.Redirect("/login-bidder", 302)
// }
