package controllers

import (
	_ "Secure-e-Auc/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type Seller struct {
	beego.Controller
}

type Bidder struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "landing.tpl"

}
func (c *Seller) Get() {
	c.TplName = "seller.tpl"
}
func (c *Bidder) Get() {
	c.TplName = "bidder.tpl"
}
