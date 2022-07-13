package repository

import (
	"errors"
	"fmt"
	"log"
	"time"

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
func (usr *UserRepository) AddRecentlyView(userId string, articleId int) error {
	zkey := fmt.Sprintf("viewed:%s", userId)

	err := usr.rcl.ZAdd(ctx, zkey, redis.Z{
		Member: fmt.Sprintf("article:%d", articleId),
		Score:  float64(time.Now().Unix()),
	}).Err()
	return err

}

// zrange viewed:user:1 0 -1
// return list of articles:id
func (usr *UserRepository) GetRecentlyViews(userId string) ([]string, error) {
	zkey := fmt.Sprintf("viewed:%s", userId)

	start := int64(0)
	end := int64(-1)

	articleIds, err := usr.rcl.ZRange(ctx, zkey, start, end).Result()
	if err != nil {
		return nil, err
	}
	return articleIds, nil
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
