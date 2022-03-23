package controllers

// import (
// 	// "Secure-e-Auc/models"

// 	"github.com/astaxie/beego"
// )

// type Bidder struct {
// 	beego.Controller
// }

// type Auction struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// type Auctions []Auction

// var auctions []Auction

// func init() {
// 	auctions = Auctions{
// 		Auction{Name: "name", Description: "You can suck it or sit on it"},
// 	}
// }

// func (c *Bidder) Get() {
// 	c.TplName = "bidder.tpl"
// 	c.Data["auctions"] = auctions
// }

// // func (c *Bidder) Post() {
// // 	var bid_amount = c.GetInt("bidAmount")
// // 	var bidder_id = user_id

// // 	models.Bidder_List(bid_amount, int64(bidder_id))
// // 	c.Redirect("/bidder", 302)
// // }