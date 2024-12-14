package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	Models "github.com/j0kernotathome/Go-password-manager/PasswordService/internal/models/Users"
	"github.com/j0kernotathome/Go-password-manager/PasswordService/internal/singleton"
)

type password struct {
	UserId int64  `json:"userId"`
	Text   string `json:"password"`
}

func getPasswords(ctx *fiber.Ctx) error {
	var passwords []password
	var u Models.User
	var err error

	db := singleton.ConnectToDb()
	u.GetUserByLogin(ctx.Cookies("login"))
	db.Find(&passwords, "user_Id=?", u.Id)

	p := make([]string, len(passwords))
	for i, password := range passwords {
		p[i], err = decrypt(singleton.GetPasswordKey(), password.Text)
		if err != nil {
			fmt.Println(err)
		}
	}
	jsonPassword, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	ctx.Write(jsonPassword)
	return ctx.SendStatus(200)
}

func addPassword(ctx *fiber.Ctx) error {
	db := singleton.ConnectToDb()
	db.AutoMigrate(password{})
	var password password
	err := json.Unmarshal(ctx.Body(), &password)

	if err != nil {
		log.Println(err)
	}

	if len(password.Text) > 16 {
		return errors.New("password too long")
	}
	if len(password.Text) < 4 {
		return errors.New("password too short")
	}

	password.Text, _ = encrypt(singleton.GetPasswordKey(), password.Text)
	var u Models.User
	u.GetUserByLogin(ctx.Cookies("login"))
	password.UserId = u.Id
	db.Create(password)
	return ctx.SendStatus(200)
}

func encrypt(secret, value string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	plainText := []byte(value)

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(secret, value string) (string, error) {
	ciphertext, err := base64.RawStdEncoding.DecodeString(value)
	if err != nil {
		return "", fmt.Errorf("decoding base64: %w", err)
	}

	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
