package main

import (
	"app/user"
	"context"
	"time"

	//gin
	"github.com/gin-gonic/gin"
)

func main() {

	user.ConnectDB()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	defer user.Client.Disconnect(ctx)
	//RESTful API
	r := gin.Default()
	r.POST("/register", user.Register)
	r.Run() //change port here r.Run(:8011)

}
