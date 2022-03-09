package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type Bidder struct {
	Id       int64  `json:"id"`
	Email    string `json:"email" orm:"unique;index;size(191)"`
	Password string `json:"-"`
}

type Seller struct {
	Id       int64  `json:"id"`
	Email    string `json:"email" orm:"unique;index;size(191)"`
	Password string `json:"-"`
}

// func (m *Bidder) Read(fields ...string) error {
// 	if err := orm.NewOrm().Read(m, fields...); err != nil {
// 		return err
// 	}
// 	return nil
// }

func init() {
	orm.RegisterModel(new(Bidder), new(Seller))
}

func NewBidder(email, password string) (id int64, err error) {
	o := orm.NewOrm()
	user := Bidder{}
	user.Email = email
	user.Password = password

	uId, insertErr := o.Insert(&user)
	if insertErr != nil {
		return -1, errors.New("failed to insert user to database")
	}

	return uId, nil
}

func NewSeller(email, password string) (id int64, err error) {
	o := orm.NewOrm()
	user := Seller{}
	user.Email = email
	user.Password = password

	uId, insertErr := o.Insert(&user)
	if insertErr != nil {
		return -1, errors.New("failed to insert user to database")
	}

	return uId, nil
}

func FindBidderbyEmail(email string) (user *Bidder, err error) {
	o := orm.NewOrm()
	u := Bidder{Email: email}
	e := o.Read(&u, "Email")
	if e == orm.ErrNoRows {
		return nil, errors.New("user not found")
	} else if e == nil {
		return &u, nil
	} else {
		return nil, errors.New("unknown error occurred")
	}
}

func FindSellerbyEmail(email string) (user *Seller, err error) {
	o := orm.NewOrm()
	s := Seller{Email: email}
	e := o.Read(&s, "Email")
	if e == orm.ErrNoRows {
		return nil, errors.New("user not found")
	} else if e == nil {
		return &s, nil
	} else {
		return nil, errors.New("unknown error occurred")
	}
}

func LoginBidder(email, password string) (user *Bidder, err error) {
	u, e := FindBidderbyEmail(email)
	if e == nil {
		return u, e
	} else {
		return nil, e
	}
}

func LoginSeller(email, password string) (user *Seller, err error) {
	u, e := FindSellerbyEmail(email)
	if e == nil {
		return u, e
	} else {
		return nil, e
	}
}
