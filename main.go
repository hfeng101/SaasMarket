package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {

		fmt.Println("Get  testtest123123 db is  %+v ", db)



		user := c.Params.ByName("name")

		fmt.Println("Get  testtest123123 user is  %+v ", user)
		value, ok := db[user]
		if ok {

			fmt.Println("Get OK testtest123123 value is %+v ", value)
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			fmt.Println("Get NO OK testtest123123 value is %+v ", value)
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	fmt.Println("testtest123123 %+v ", db)


	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		print("testtest user is %+v", user)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		fmt.Println("testtest, json value is %+v", json)
		fmt.Println("testtest123123.................. %+v ", db)

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
