package controllers

import (
	"github.com/astaxie/beego"
)



type BidDetailsController struct {
	beego.Controller
}

func (c *BidDetailsController) Get() {
	c.TplName = "bid-details.tpl"
}

