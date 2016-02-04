package usermodules

import (
	"time"

	"github.com/bugisdev/helper"
	"github.com/jinzhu/gorm"
	"github.com/ngurajeka/ngurajeka.com/app"
)

// Login Function
func Login(f UserLoginForm, DB *gorm.DB) (user User, err []helper.ErrorMessage) {

	v := f.Validate()

	if v.HasErrors() {
		for _, value := range v.Errors {
			err = append(err, helper.ErrorMessage{
				Code:    409,
				Source:  helper.SourceErrors{Pointer: value.Key},
				Title:   "Insufficient Data",
				Details: value.Message,
			})
		}

		return user, err
	}

	_err := DB.Where(&User{Username: f.Data.Username}).First(&user).Error
	if _err != nil {
		err = append(err, helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{},
			Title:   "Failed Logging in User",
			Details: _err.Error(),
		})

		return user, err
	}

	loggedIn := ComparePassword(&user, f.Data.Password)
	if loggedIn == false {
		err = append(err, helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{},
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
func NewUser(f UserRegisterForm, DB *gorm.DB) (user User, err []helper.ErrorMessage) {

	v := f.Validate()

	if v.HasErrors() {
		for _, value := range v.Errors {
			err = append(err, helper.ErrorMessage{
				Code:    409,
				Source:  helper.SourceErrors{Pointer: value.Key},
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
		err = append(err, helper.ErrorMessage{
			Code:    409,
			Source:  helper.SourceErrors{},
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

// GetAll Function, Fetch All Users with Offset and Limit
func GetAll(limit, offset int, DB *gorm.DB) (users []User, err []helper.ErrorMessage) {

	_err := DB.Find(&users).Limit(limit).Offset(offset).Error
	if _err != nil {
		err = append(err, helper.ErrorMessage{
			Code:    400,
			Source:  helper.SourceErrors{},
			Title:   "Failed Fetch All User",
			Details: _err.Error(),
		})

		return users, err
	}

	return users, nil
}

// GetSingle Function, Fetch Single User and relation if true
func GetSingle(id int, DB *gorm.DB) (user User, err []helper.ErrorMessage) {

	_err := DB.First(&user, id).Error
	if _err != nil {
		err = append(err, helper.ErrorMessage{
			Code:    404,
			Source:  helper.SourceErrors{},
			Title:   "Failed Fetch Single user",
			Details: _err.Error(),
		})

		return user, err
	}

	var profile Profile
	DB.Model(&user).Related(&profile)
	user.Profile = profile
	return user, nil
}

// UpdateSingle Function, Updating User (Not Including Profile Relation)
func UpdateSingle(id int, f UserUpdateForm, DB *gorm.DB) (user User, err []helper.ErrorMessage) {

	user, _err := GetSingle(id, DB)
	if _err != nil {
		return user, _err
	}

	user.Fullname = f.Data.Fullname
	user.Profile.Gender = f.Data.Gender
	user.Profile.Address = f.Data.Address
	user.Profile.PhoneNumber = f.Data.PhoneNumber
	user.Profile.BirthPlace = f.Data.BirthPlace
	user.Profile.BirthDate, _ = time.Parse("01/02/2006", f.Data.BirthDate)

	_errSaving := app.DB.Save(&user).Error
	if _errSaving != nil {
		err = append(err, helper.ErrorMessage{
			Code:    400,
			Source:  helper.SourceErrors{},
			Title:   "Update User Failed",
			Details: _errSaving.Error(),
		})

		return user, err
	}

	return user, nil
}
