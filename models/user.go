package models

import (
	// "fmt"
	"errors"
	"fmt"
	// "time"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v4"

	"github.com/adelberteng/reservation_system/db"
	"github.com/adelberteng/reservation_system/utils"
)

var (
	logger = utils.GetLogger()
	engine = db.GetSQLClient()
)

type User struct {
	Uid          string `xorm:"not null pk autoincr INT(11)"`
	Name         string `xorm:"not null unique VARCHAR(32)"`
	PasswordHash string `xorm:"not null VARCHAR(128)"`
	Phone        string `xorm:"not null unique CHAR(10)"`
	Email        string `xorm:"not null unique VARCHAR(255)"`
}

func init() {
	err := engine.Sync2(new(User))
	if err != nil {
		logger.Error(err)
	}
}

func (u *User) TableName() string {
	return "user_tbl"
}

func GeneratePasswordHash(password string) (string, error) {
	pw := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(pw, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func verifyPassword(password, passwordHash string) bool {
	pw := []byte(password)
	hash := []byte(passwordHash)

	err := bcrypt.CompareHashAndPassword(hash, pw)
	if err != nil {
		return false
	}

	return true
}

func Register(name, password, phone, email string) error {
	nameResult, err := engine.Table("user_tbl").Where("name = ?", name).QueryString()
	if err != nil {
		return err
	} else if nameResult != nil {
		return errors.New("This user name had been registered")
	}

	phoneResult, err := engine.Table("user_tbl").Where("phone = ?", phone).QueryString()
	if err != nil {
		return err
	} else if phoneResult != nil {
		return errors.New("This phone number had been registered")
	}

	emailResult, err := engine.Table("user_tbl").Where("email = ?", email).QueryString()
	if err != nil {
		return err
	} else if emailResult != nil {
		return errors.New("This email address had been registered")
	}

	// generate the password hash with salt.
	passwordHash, err := GeneratePasswordHash(password)
	if err != nil {
		return err
	}
	user := User{Name: name, PasswordHash: passwordHash, Phone: phone, Email: email}

	_, err = engine.Insert(&user)
	if err != nil {
		return err
	}

	return nil
}

func Login(name, password string) (User, error) {
	res, err := engine.Table("user_tbl").Where("name = ?", name).QueryString()
	passwordHash := res[0]["password_hash"]
	if err != nil {
		fmt.Println(err)
	}

	isCorrect := verifyPassword(password, passwordHash)
	if !isCorrect {
		return User{}, errors.New("password is incorrect, please try again.")
	}

	user := User{
		Uid:   res[0]["uid"],
		Name:  res[0]["name"],
		Phone: res[0]["phone"],
		Email: res[0]["email"],
	}

	return user, nil
}

func ApplyJWT() (string, error) {
	Claims := jwt.MapClaims{
		"foo": "bar",
		// "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	}

	var hmacSecret []byte

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	fmt.Println(token)
	tokenString, err := token.SignedString(hmacSecret)
	fmt.Println(tokenString, err)
	return tokenString, err
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	var hmacSecret []byte

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
