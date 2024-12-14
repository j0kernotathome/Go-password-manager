package Models

import (
	"errors"
	"unicode"

	handler "github.com/j0kernotathome/Go-password-manager/UserService/internal/singleton"
)

type User struct {
	Id       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	APIkey   string `json:"apikey"`
}

func (u *User) GetUserByLogin(login string) {
	db := handler.ConnectToDb()
	db.Where("login = ?", login).First(&u)
}
func CreateUserInDb(u User) error {
	db := handler.ConnectToDb()
	if u.Login == "" || u.Password == "" || u.Email == "" {
		return errors.New("nil value")
	}
	if len(u.Login) > 16 || len(u.Password) > 16 {
		return errors.New("login or password too long")
	}
	if len(u.Login) < 4 || len(u.Password) < 8 {
		return errors.New("login or password too short")
	}
	if isASCII(u.Login) == false || isASCII(u.Password) == false {
		return errors.New("login or password have non ascii symbols")
	}
	if isHaveBannedSybols(u.Login) || isHaveBannedSybols(u.Password) {
		return errors.New("login or password have banned symbols")
	}
	db.Create(&u)
	return nil
}
func isASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func isHaveBannedSybols(s string) bool {
	bannedList := `"'!@#$%^&*()_+-=[]\{}|;:,.<>/?`
	for _, char := range s {
		for _, bannedChar := range bannedList {
			if char == bannedChar {
				return true
			}
		}
	}
	return false
}
