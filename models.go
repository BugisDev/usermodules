package usermodules

import "time"

// User Model
type User struct {
	ID                int        `json:"id"`
	Username          string     `json:"username" sql:"not null;unique;type:varchar(32);unique_index"`
	Password          []byte     `json:"password" sql:"not null;"`
	Email             string     `json:"email" sql:"not null;unique;type:varchar(120);unique_index"`
	Fullname          string     `json:"full_name" sql:"not null;type:varchar(120);"`
	Profile           Profile    `json:"profile,omitempty"`
	CreatedAt         time.Time  `json:"-"`
	UpdatedAt         time.Time  `json:"-"`
	DeletedAt         *time.Time `json:"-"`
	LastLoginAt       time.Time  `json:"-"`
	PasswordUpdatedAt time.Time  `json:"-"`
}

// Profile Model
// Related with User Model
type Profile struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"userId,omitempty" sql:"index"`
	Gender      int8      `json:"gender,omitempty"`
	Address     string    `json:"address,omitempty" sql:"null;type:varchar(255)"`
	PhoneNumber string    `json:"phone_number,omitempty" sql:"null;type:varchar(64)"`
	BirthPlace  string    `json:"birth_place,omitempty" sql:"null;type:varchar(64)"`
	BirthDate   time.Time `json:"birth_date,omitempty" sql:"null"`
}
