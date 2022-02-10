package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

type RegisterSellerController struct {
	beego.Controller
}

func (c *RegisterSellerController) RegisterUser(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
	fmt.Print(c)
	c.Ctx.ResponseWriter.WriteHeader(200)
	c.TplName = "login-seller.tpl"
}
