package controllers

import (
	"Secure-e-Auc/models"
	"fmt"

	"github.com/astaxie/beego"
)

var bidder_user_id int = -1
var seller_user_id int = -1
var auctions_id = -1

type Seller struct {
	beego.Controller
}

type LoginBidderController struct {
	beego.Controller
}

type LoginSellerController struct {
	beego.Controller
}

type RegistrationBidderController struct {
	beego.Controller
}

type RegistrationSellerController struct {
	beego.Controller
}

type NewAuctionController struct {
	beego.Controller
}

type Bidder struct {
	beego.Controller
}

type BidDetailsController struct {
	beego.Controller
}
type BidController struct {
	beego.Controller
}
type MainController struct {
	beego.Controller
}

type SellerAuctionController struct {
	beego.Controller
}

func (c *LoginBidderController) Get() {
	c.TplName = "login-bidder.tpl"
}

func (c *LoginBidderController) Post() {
	// c.ServeJSON()
	var user_email = c.GetString("email")
	var user_password = c.GetString("password")

	user, err := models.LoginBidder(user_email, user_password)
	fmt.Print(err != nil)
	if user.Password == user_password {
		bidder_user_id = int(user.Id)
		c.Redirect("/bidder", 302)
	}
	c.Redirect("/login-bidder", 302)
}

func (c *LoginSellerController) Get() {
	c.TplName = "login-seller.tpl"

}

// Seller Login

func (c *LoginSellerController) Post() {
	// c.ServeJSON()
	var user_email = c.GetString("email")
	var user_password = c.GetString("password")

	user, err := models.LoginSeller(user_email, user_password)
	fmt.Print(err != nil)
	if user.Password == user_password {
		seller_user_id = int(user.Id)
		c.Redirect("/seller", 302)
	}
	c.Redirect("/login-seller", 302)
}

//Bidder Registration
func (c *RegistrationBidderController) Get() {
	c.TplName = "registration-bidder.tpl"
}

func (c *RegistrationBidderController) Post() {
	var user_email = c.GetString("email")
	var user_password = c.GetString("password")
	fmt.Print(user_email)
	models.NewBidder(user_email, user_password)
	c.Redirect("/login-bidder", 302)
}

//Seller Registrations

func (c *RegistrationSellerController) Get() {
	c.TplName = "registration-seller.tpl"
}

func (c *RegistrationSellerController) Post() {
	var user_email = c.GetString("email")
	var user_password = c.GetString("password")
	fmt.Print(user_email)
	models.NewSeller(user_email, user_password)
	c.Redirect("/login-seller", 302)
}

// Add a product to the auction

func (c *NewAuctionController) Get() {
	c.TplName = "new-auction.tpl"
}

func (c *NewAuctionController) Post() {
	if seller_user_id == -1 {
		c.Redirect("/seller", 302)
	}
	// c.Redirect("/", 404)

	// }
	// c.ServeJSON()
	var user_product = c.GetString("productTitle")
	var user_desc = c.GetString("productDesc")
	var seller_id = seller_user_id

	models.NewAuction(user_product, user_desc, int64(seller_id), false)
	c.Redirect("/seller", 302)

}

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

func (c *Bidder) Get() {
	if bidder_user_id == -1 {
		c.Redirect("/seller", 302)
	}
	c.TplName = "bidder.tpl"
	auctions, err := models.AuctionListBidder()
	fmt.Print(err != nil)
	c.Data["auctions"] = auctions
}

// func (c *Bidder) Post() {
// 	if bidder_user_id == -1 {
// 		c.Redirect("/", 302)
// 	}
// 	var bid_amount, err = c.GetInt("bidAmount")
// 	var bidder_id = bidder_user_id

// 	models.Bidder_List(bid_amount, int64(bidder_id))
// 	c.Redirect("/bidder", 302)
// }that

func (c *BidDetailsController) Get() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Print(err)
	}
	auctions_id = id
	auctions, err := models.AuctionDetails(id)
	bidderlist, err := models.AuctionBidList(int64(id))
	c.Data["auctions"] = auctions
	c.Data["bidders"] = bidderlist
	c.TplName = "bid-details.tpl"
}

func (c *BidDetailsController) Post() {
	models.EndBid(int64(auctions_id))
	c.Redirect("/seller", 302)
}

func (c *Seller) Get() {
	if seller_user_id == -1 {
		c.Redirect("/", 302)
	}

	c.TplName = "seller.tpl"
	auctions, err := models.AuctionListSeller(int64(seller_user_id))
	fmt.Print(err != nil)
	c.Data["auctions"] = auctions
}

func (c *MainController) Get() {
	c.TplName = "landing.tpl"
}

func (c *BidController) Get() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Print(err)
	}
	auctions_id = id
	auctions, err := models.AuctionDetails(auctions_id)
	c.Data["auctions"] = auctions
	c.TplName = "bid.tpl"
}

func (c *BidController) Post() {
	amount, err := c.GetInt("bidAmount")
	fmt.Print(err)
	fmt.Print(amount)
	bidId, err := models.AddBid(int64(auctions_id), int64(bidder_user_id), int64(amount))
	fmt.Println(bidId)
	c.Redirect("/bidder", 302)
}
