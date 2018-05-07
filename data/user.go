package data

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	username string `json:"username"`
	password string `json:"password"`
	salt     string
}

// SetUsername sets the username
func (u *User) SetUsername(username string) {
	u.username = username
}

// GetUsername returns the username
func (u *User) GetUsername() string {
	return u.username
}

func (u *User) SetHashedPassword(hashedPw string) {
	u.password = hashedPw
}

// SetPassword hashes/salts and sets the hashed password
func (u *User) SetPassword(password string) error {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.password = string(hashedPw)
	return nil
}

func (u *User) GetPassword() string {
	return u.password
}

// VerifyPassword verifies a given password against a hashed password
func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
}

func MakeUser(username string, password string) (*User, error) {
	user := &User{}
	user.SetUsername(username)
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
