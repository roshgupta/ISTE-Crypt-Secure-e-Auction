package main

import (
	_ "Secure-e-Auc/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(localhost)/mysql")
	// force := true // Drop old table and create new
	// err := orm.RunSyncdb("default", force, beego.BConfig.RunMode == "dev")
	// if err != nil {
	// 	fmt.Printf("An Error Occurred: %v", err)
	// }
}

func main() {
	beego.Run()
}
