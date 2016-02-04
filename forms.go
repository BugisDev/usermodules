package usermodules

import "github.com/revel/revel"

// UserLoginForm Handling User Login
type UserLoginForm struct {
	Data LoginData `json:"data"`
}

// LoginData Data post when Login
type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate User Login
func (f *UserLoginForm) Validate() (v revel.Validation) {

	// Validate Username
	v.Required(f.Data.Username).Key("Username")
	v.MinSize(f.Data.Username, 3).Key("Username")
	v.MaxSize(f.Data.Username, 32).Key("Username")

	// Validate Password
	v.Required(f.Data.Password).Key("Password")
	v.MinSize(f.Data.Password, 8).Key("Password")

	return v
}

// UserRegisterForm Handling User Registration
type UserRegisterForm struct {
	Data RegisterData `json:"data"`
}

// RegisterData Data post when Register
type RegisterData struct {
	Email           string `json:"email"`
	Fullname        string `json:"full_name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// Validate User Registration
func (f *UserRegisterForm) Validate() (v revel.Validation) {

	// Validate Email
	v.Required(f.Data.Email).Key("Email")
	v.MaxSize(f.Data.Email, 120).Key("Email")
	v.Email(f.Data.Email).Key("Email")

	// Validate Fullname
	v.Required(f.Data.Fullname).Key("Fullname")
	v.MinSize(f.Data.Fullname, 3).Key("Fullname")
	v.MaxSize(f.Data.Fullname, 120).Key("Fullname")

	// Validate Username
	v.Required(f.Data.Username).Key("Username")
	v.MinSize(f.Data.Username, 3).Key("Username")
	v.MaxSize(f.Data.Username, 32).Key("Username")

	// Validate Password
	v.Required(f.Data.Password).Key("Password")
	v.Required(f.Data.ConfirmPassword).Key("ConfirmPassword")
	v.MinSize(f.Data.Password, 8).Key("Password")
	v.MinSize(f.Data.ConfirmPassword, 8).Key("ConfirmPassword")
	v.Required(f.Data.Password == f.Data.ConfirmPassword).Message("Password didn't match.").Key("ConfirmPassword")

	return v
}

// UserUpdateForm Handling User Update
type UserUpdateForm struct {
	Data UserUpdateData `json:"data"`
}

// UserUpdateData Data post when Updating User
type UserUpdateData struct {
	Fullname    string `json:"full_name"`
	Gender      int8   `json:"gender,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	BirthPlace  string `json:"birth_place,omitempty"`
	BirthDate   string `json:"birth_date,omitempty"`
}
