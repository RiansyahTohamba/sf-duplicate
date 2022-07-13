package repository

import (
	"errors"

	"github.com/go-redis/redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	Email        string
	PasswordHash string
}

type UserRepository struct {
	rcl *redis.Client
	orm *gorm.DB
}

type EmailNotExistsError struct{}
type PasswordMissMatchError struct{}

func (*EmailNotExistsError) Error() string {
	return "email is not exists"
}

func (*PasswordMissMatchError) Error() string {
	return "password is invalid"
}

func NewUserRepo(rcl *redis.Client, orm *gorm.DB) *UserRepository {
	return &UserRepository{rcl, orm}
}

func (usr *UserRepository) CheckToken(token string) error {
	hkey := "login:"
	err := usr.rcl.HGet(ctx, hkey, token).Err()
	return err
}

// func (u *UserRepository) Login(email string, password string) (*int, error) {
// 	statement := "SELECT id, password FROM users WHERE email = ?"
// 	res := u.db.QueryRow(statement, email, password)
// 	var hashedPassword string
// 	var id int
// 	res.Scan(&id, &hashedPassword)
// 	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
// 		return nil, errors.New("Login Failed")
// 	}
// 	return &id, nil
// }

func SignUp() {

}

// TODO: skip
func (usr *UserRepository) WriteRecentlyView(userId string) error {
	//   - HSet("myhash", "key1", "value1", "key2", "value2")

	// conn.hset()
	hkey := "login:"

	// token menjadi key, valuenya adalah user
	token := ""

	// bisa diambil dari session
	err := usr.rcl.HSet(ctx, hkey, token, userId).Err()
	return err

}

func (usr *UserRepository) Login(email, password string) (*User, error) {
	var user User

	err := usr.orm.Find(&user, &User{Email: email}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &EmailNotExistsError{}
	}

	// compare hash password in DB vs password yg datang dari request.Body
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return nil, &PasswordMissMatchError{}
	}

	return &user, nil
}
