package models

import "github.com/astaxie/beego/orm"

type Bidder struct {
	Id       int    `orm:"pk" form:"-"`
	Email    string `orm:"size(256)"`
	Password string `orm:"size(128)"`
}

func (m *Bidder) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

// func (m *Bidder) Read(fields ...string) error {
// 	if err := orm.NewOrm().Read(m, fields...); err != nil {
// 		return err
// 	}
// 	return nil
// }

func init() {
	// orm.RegisterModel(new(Bidder))
}
