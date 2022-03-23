package controllers

import (
	"Secure-e-Auc/models"

	"github.com/astaxie/beego"
)

type NewAuctionController struct {
	beego.Controller
}

func (c *NewAuctionController) Get() {
	c.TplName = "new-auction.tpl"
}

func (c *NewAuctionController) Post() {
	// c.ServeJSON()
	var user_product = c.GetString("productTitle")
	var user_desc = c.GetString("productDesc")
	var seller_id = user_id

	models.NewAuction(user_product, user_desc, -1, int64(seller_id), false)
	c.Redirect("/seller", 302)
}
