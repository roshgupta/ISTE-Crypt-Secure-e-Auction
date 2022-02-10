package models

type UserBidder struct {
	modelImpl
	FullName string
	Email    string
	Password string
}

func NewUserBidder(fullName, email, password string) *UserBidder {
	u := &UserBidder{
		FullName: fullName,
		Email:    email,
		Password: password,
	}
	u.SetId(email)
	return u
}

func (u *UserBidder) GetId() string {
	return u.Email
}
