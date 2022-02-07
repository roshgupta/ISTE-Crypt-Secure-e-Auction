package routers

import (
	"Secure-e-Auc/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/main", &controllers.Landing{})
	beego.Router("/bidder", &controllers.Bidder{})
	beego.Router("/seller", &controllers.Seller{})
}
