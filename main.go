package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sesi1-crud/config"
	"sesi1-crud/modules/user"
	"sesi1-crud/utils/database"

	"github.com/gin-gonic/gin"
)

func init() {
	config.SetupConfiguration()
}

func main() {
	//cache.LearnRedis()
	fmt.Println(os.Getenv("DB_HOST"))
	router := gin.New()
	db := database.GetDatabase()

	userHandler := user.NewRouter(db)
	userHandler.Handle(router)

	//learn simple endpoint
	router.GET("/hello", helloHandler)
	router.GET("/say", getMessageHandler)

	err := router.Run(fmt.Sprintf(":%s", config.Config.Server.Port))
	if err != nil {
		log.Fatal("router can't run", err)
	}
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messange": "Hello World API",
	})
}

func getMessageHandler(c *gin.Context) {
	str := c.DefaultQuery("message", "Hello World API")
	c.JSON(http.StatusOK, gin.H{
		"messange": str,
	})
}
