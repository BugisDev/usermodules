package usermodules

import "github.com/jinzhu/gorm"

// ErrorMessage struct
type ErrorMessage struct {
	Code    int          `json:"code"`
	Source  SourceErrors `json:"source,omitempty"`
	Title   string       `json:"title"`
	Details string       `json:"details,omitempty"`
}

// SourceErrors Struct
type SourceErrors struct {
	Pointer string `json:"pointer,omitempty"`
}

// Login Function
func Login(f UserLoginForm, DB *gorm.DB) (user User, err []ErrorMessage) {

	v := f.Validate()

	if v.HasErrors() {
		for _, value := range v.Errors {
			err = append(err, ErrorMessage{
				Code:    409,
				Source:  SourceErrors{Pointer: value.Key},
				Title:   value.Message,
				Details: value.Message,
			})
		}

		return user, err
	}

	_err := DB.Where(&User{Username: f.Data.Username}).First(&user).Error
	if _err != nil {
		err = append(err, ErrorMessage{
			Code:    409,
			Source:  SourceErrors{},
			Title:   "Failed Logging in User",
			Details: _err.Error(),
		})

		return user, err
	}

	loggedIn := ComparePassword(&user, f.Data.Password)
	if loggedIn == false {
		err = append(err, ErrorMessage{
			Code:    409,
			Source:  SourceErrors{},
			Title:   "Failed Logging in User",
			Details: "Wrong Username / Password Combination",
		})

		return user, err
	}

	var profile Profile
	DB.Model(&user).Related(&profile)
	user.Profile = profile
	return user, nil
}

// NewUser Function
func NewUser(f UserRegisterForm, DB *gorm.DB) (user User, err []ErrorMessage) {

	v := f.Validate()

	if v.HasErrors() {
		for _, value := range v.Errors {
			err = append(err, ErrorMessage{
				Code:    409,
				Source:  SourceErrors{Pointer: value.Key},
				Title:   value.Message,
				Details: value.Message,
			})
		}

		return user, err
	}

	user.Username = f.Data.Username
	user.Fullname = f.Data.Fullname
	user.Email = f.Data.Email
	user.Password = GeneratePassword(f.Data.Password)

	_err := DB.Create(&user).Error
	if _err != nil {
		err = append(err, ErrorMessage{
			Code:    409,
			Source:  SourceErrors{},
			Title:   "Failed Creating New User",
			Details: _err.Error(),
		})

		return user, err
	}

	var profile Profile
	DB.Model(&user).Related(&profile)
	user.Profile = profile
	return user, nil
}
