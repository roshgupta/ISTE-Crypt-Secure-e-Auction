package controllers

import (
	"Secure-e-Auc/models"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

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
	user, err := models.LoginBidder(user_email, user_password)
	if err == nil {
		print(err)
	}
	app := "node"
	arg0 := "auction/auction-simple/application-javascript/registerEnrollUser.js"
	arg1 := "org2"
	s := strconv.FormatInt(user.Id, 10)
	arg2 := "bidder" + s
	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()
	fmt.Println(string(stdout), err)
	c.Redirect("/login-bidder", 302)
}

//Seller Registrations

func (c *RegistrationSellerController) Get() {
	c.TplName = "registration-seller.tpl"
}

func (c *RegistrationSellerController) Post() {
	var user_email = c.GetString("email")
	var user_password = c.GetString("password")
	// fmt.Print(user_email)
	models.NewSeller(user_email, user_password)
	user, err := models.LoginSeller(user_email, user_password)
	app := "node"
	arg0 := "auction/auction-simple/application-javascript/registerEnrollUser.js"
	arg1 := "org1"
	s := strconv.FormatInt(user.Id, 10)
	arg2 := "seller" + s
	print(arg2)
	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()
	fmt.Println(string(stdout), err)

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

	id, err := models.NewAuction(user_product, user_desc, int64(seller_id), false)
	auctions, err := models.AuctionDetails(int(id))
	app := "node"
	arg0 := "auction/auction-simple/application-javascript/createAuction.js"
	arg1 := "org1"
	s := strconv.Itoa(seller_user_id)
	arg2 := "seller" + s
	s = strconv.FormatInt(auctions.Id, 10)
	arg3 := "auction" + s
	arg4 := string(user_product)
	fmt.Println(app, arg0, arg1, arg2, arg3, arg4)
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	stdout, err := cmd.Output()
	fmt.Println(string(stdout))
	if err == nil {
	}
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
	if err != nil {
		fmt.Print(err)
	}
	bidderlist, err := models.AuctionBidList(int64(id))
	if err != nil {
		fmt.Print(err)
	}
	c.Data["auctions"] = auctions
	c.Data["bidders"] = bidderlist
	c.TplName = "bid-details.tpl"
}

func (c *BidDetailsController) Post() {
	models.EndBid(int64(auctions_id))
	app := "node"
	arg0 := "auction/auction-simple/application-javascript/closeAuction.js"
	arg1 := "org1"
	arg2 := "seller" + strconv.Itoa(seller_user_id)
	arg3 := "auction" + strconv.Itoa(auctions_id)
	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()
	fmt.Println(string(stdout), err)
	bidHash, err := models.BidHashList(int64(auctions_id))
	arg0 = "auction/auction-simple/application-javascript/revealBid.js"
	arg1 = "org2"
	arg3 = "auction" + strconv.Itoa(auctions_id)
	for i := 0; i < len(bidHash); i++ {
		bidder_id := bidHash[i].Bidder_id
		arg2 = "bidder" + strconv.FormatInt(bidder_id, 10)
		arg4 := bidHash[i].Hash
		cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
		stdout, err := cmd.Output()
		fmt.Println(string(stdout), err)
	}
	app = "node"
	arg0 = "auction/auction-simple/application-javascript/endAuction.js"
	arg1 = "org1"
	arg2 = "seller" + strconv.Itoa(seller_user_id)
	arg3 = "auction" + strconv.Itoa(auctions_id)
	cmd = exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err = cmd.Output()
	fmt.Println(string(stdout), err)
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
	if err != nil {
		fmt.Print(err)
	}
	c.Data["auctions"] = auctions
	c.TplName = "bid.tpl"
}

// func exportBidId(app string, auctions_id, bid_id string) (toprint string, err error) {
// 	fmt.Println(app, auctions_id+"_BID_ID="+bid_id)

// 	cmd := exec.Command(app, auctions_id+"_BID_ID="+bid_id)
// 	stdout, err := cmd.Output()
// 	arg0 := "$" + auctions_id + "_BID_ID"
// 	cmd = exec.Command("echo", arg0)
// 	stdout, err = cmd.Output()
// 	return string(stdout), err
// }

func addBid(app, arg0, arg1, arg2, arg3, arg4 string) (toprint string, err error) {
	fmt.Println(app, arg0, arg1, arg2, arg3, arg4)
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	stdout, err := cmd.Output()
	return string(stdout), err
}

func proposeBid(app, arg0, arg1, arg2, arg3, arg4 string) (toprint string, err error) {
	fmt.Println(app, arg0, arg1, arg2, arg3, arg4)
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	stdout, err := cmd.Output()
	return string(stdout), err
}

func (c *BidController) Post() {
	amount, err := c.GetInt("bidAmount")
	fmt.Print(err)
	fmt.Print(amount)

	bidId, err := models.AddBid(int64(auctions_id), int64(bidder_user_id), int64(amount))
	fmt.Println(bidId, "/n")
	app := "node"
	arg0 := "auction/auction-simple/application-javascript/bid_web.js"
	arg1 := "org2"
	arg2 := "bidder" + strconv.Itoa(bidder_user_id)
	arg3 := "auction" + strconv.Itoa(auctions_id)
	arg4 := strconv.Itoa(amount)
	fmt.Println(app, arg0, arg1, arg2, arg3, arg4)
	toprint, err := proposeBid(app, arg0, arg1, arg2, arg3, arg4)
	bidHash := strings.Split(toprint, "\n")
	bidHash = bidHash[:len(bidHash)-1]
	fmt.Println(toprint)
	fmt.Println(bidHash)
	fmt.Println(err)

	id, err := models.StoreBidderHash(int64(auctions_id), int64(bidder_user_id), bidHash[0])
	print(err)
	app = "node"
	arg0 = "auction/auction-simple/application-javascript/submitBid.js"
	arg1 = "org2"
	arg2 = "bidder" + strconv.Itoa(bidder_user_id)
	arg3 = "auction" + strconv.Itoa(auctions_id)
	arg4 = bidHash[0]
	toprint, err = addBid(app, arg0, arg1, arg2, arg3, arg4)
	print(err)
	print(toprint, id)
	c.Redirect("/bidder", 302)
}
