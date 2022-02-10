package models

type UserSeller struct {
	modelImpl
	FullName string
	Email    string
	Password string
}

func NewUserSeller(fullName, email, password string) *UserSeller {
	u := &UserSeller{
		FullName: fullName,
		Email:    email,
		Password: password,
	}
	u.SetId(email)
	return u
}

func (u *UserSeller) GetId() string {
	return u.Email
}
