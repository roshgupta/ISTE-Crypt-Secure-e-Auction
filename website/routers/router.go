package routers

import (
	"Secure-e-Auc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login-seller", &controllers.LoginSellerController{})
	beego.Router("/login-bidder", &controllers.LoginBidderController{})
	beego.Router("/registration-seller", &controllers.RegistrationSellerController{})
	beego.Router("/registration-bidder", &controllers.RegistrationBidderController{})

}
