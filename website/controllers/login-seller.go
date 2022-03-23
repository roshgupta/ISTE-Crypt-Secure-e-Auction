package controllers

// import (
// 	"Secure-e-Auc/models"
// 	"fmt"

// 	"github.com/astaxie/beego"
// )

// type LoginSellerController struct {
// 	beego.Controller
// }

// func (c *LoginSellerController) Get() {
// 	c.TplName = "login-seller.tpl"
// }

// func (c *LoginSellerController) Post() {
// 	// c.ServeJSON()
// 	var user_email = c.GetString("email")
// 	var user_password = c.GetString("password")

// 	user, err := models.LoginSeller(user_email, user_password)
// 	fmt.Print(err != nil)
// 	if user.Password == user_password {
// 		c.Redirect("/seller", 302)
// 	}
// 	c.Redirect("/login-seller", 302)
// }
