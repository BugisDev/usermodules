package usermodules

import (
	"crypto/sha256"
	"crypto/subtle"
)

// GeneratePassword Function
func GeneratePassword(password string) []byte {
	_password := sha256.Sum256([]byte(password))
	return _password[:]
}

// ComparePassword Function
func ComparePassword(user *User, loginpassword string) bool {
	_loginpassword := sha256.Sum256([]byte(loginpassword))
	return subtle.ConstantTimeCompare(user.Password, _loginpassword[:]) == 1
}
