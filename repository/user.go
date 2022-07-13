package repository

import (
	"errors"
	"log"

	"github.com/go-redis/redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primary_key"`
	Username     string
	Email        string
	PasswordHash string `gorm:"column:password"`
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

func SignUp() {

}

// 1. Create via addRecentlyView(user_id, article_id)
// 2. Retrieve via getRecentlyViews(user_id)

// zadd viewed:user:1 10 'article:1'
func (usr *UserRepository) AddRecentlyView(userId string) error {
	//   - HSet("myhash", "key1", "value1", "key2", "value2")

	// conn.hset()
	hkey := "login:"

	// token menjadi key, valuenya adalah user
	token := ""

	// bisa diambil dari session
	err := usr.rcl.HSet(ctx, hkey, token, userId).Err()
	return err

}

// zrange viewed:user:1 0 -1
func (usr *UserRepository) GetRecentlyViews(userId string) {

}

func (usr *UserRepository) Login(email, password string) (*User, error) {
	var user User

	dbRresult := usr.orm.Where("email = ?", email).First(&user)

	if errors.Is(dbRresult.Error, gorm.ErrRecordNotFound) {
		// handle record not found
		return nil, &EmailNotExistsError{}
	}

	// compare hash password in DB vs password yg datang dari request.Body
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		log.Println(err)
		return nil, &PasswordMissMatchError{}
	}

	return &user, nil
}
