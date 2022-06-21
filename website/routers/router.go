package routers

import (
	"Secure-e-Auc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login-seller", &controllers.LoginSellerController{})
	beego.Router("/login-bidder", &controllers.LoginBidderController{})
	beego.Router("/registration-seller", &controllers.RegistrationSellerController{})
	beego.Router("/registration-bidder", &controllers.RegistrationBidderController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/new-auction", &controllers.NewAuctionController{})
	beego.Router("/bidder", &controllers.Bidder{})
	beego.Router("/seller", &controllers.Seller{})
	beego.Router("/add", &controllers.NewAuctionController{})
	beego.Router("/bid-details", &controllers.BidDetailsController{})
	beego.Router("/bid", &controllers.BidController{})
	beego.Router("/seller/:id", &controllers.SellerAuctionController{})
}
