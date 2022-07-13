package api

import (
	"sf-duplicate/api/handler"
	"sf-duplicate/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func StartRouter(arRepo *repository.ArticleRepo, usrRepo *repository.UserRepository) {
	arHandler := handler.NewArticleHandler(arRepo)
	usrHandler := handler.NewUserHandler(usrRepo)

	router := gin.Default()
	router.Use(sessions.Sessions("sfsession", getRedisStore()))

	router.GET("/", rootHandler)

	router.POST("/login", usrHandler.Login)

	router.POST("/logout", usrHandler.Logout)

	user := router.Group("/api/v1/user")
	user.Use(sessionAuth())

	{
		user.GET("/home", arHandler.ListArticles)
	}
	// playground
	router.Use(sessions.Sessions("counter", getRedisStore()))
	router.GET("/incr", incrementHandler)
	router.POST("/testmapping", testMapping)

	router.Run(":8080")
}

type UserTest struct {
	Password string `json:"password"`
}

func testMapping(ctx *gin.Context) {
	var utreq UserTest
	err := ctx.BindJSON(&utreq)

	hashedPassword := "$2a$10$faZx20KXiAIJjHH6MrBlKOipS.AbsPrIu1W55Nnk6CO6UDnDYZFO6"

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(utreq.Password))

	if err != nil {
		ctx.JSON(500, gin.H{
			"msg": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"msg": utreq.Password,
		})
	}

}
func rootHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "welcome API",
	})
}

func getRedisStore() redis.Store {
	pwd := []byte("secret")
	size := 10
	redisStore, _ := redis.NewStore(size, "tcp", "localhost:6379", "", pwd)
	return redisStore
}

// try redis session
func incrementHandler(ctx *gin.Context) {
	var counter int
	session := sessions.Default(ctx)

	token := session.Get("token")
	val := session.Get("counter")

	if val == nil {
		counter = 0
	} else {
		counter = val.(int)
		counter++
	}
	session.Set("counter", counter)
	session.Save()

	ctx.JSON(200, gin.H{
		"counter": counter,
		"token":   token,
	})
}
