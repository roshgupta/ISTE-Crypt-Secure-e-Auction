package controllers

import (
	_ "Secure-e-Auc/models"

	"github.com/astaxie/beego"
)

type Bidder struct {
	beego.Controller
}

type Auction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Auctions []Auction

var auctions []Auction

func init() {
	auctions = Auctions{
		Auction{Name: "Dildo", Description: "You can suck it or sit on it"},
		Auction{Name: "Viagra", Description: "If your PP can't handle sex"},
	}
}

func (c *Bidder) Get() {
	c.TplName = "bidder.tpl"
	c.Data["auctions"] = auctions
}
