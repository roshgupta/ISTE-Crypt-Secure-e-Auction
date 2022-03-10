package controllers

import (
	_ "Secure-e-Auc/models"

	"github.com/astaxie/beego"
)

type Seller struct {
	beego.Controller
}

func (c *Seller) Get() {
	c.TplName = "seller.tpl"
}
