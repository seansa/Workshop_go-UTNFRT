package main

import (
	"net/http"
	"os"
	"workshop/config"
	"workshop/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	routes()
	// var db = db.DBPurchases{}
	// for k, val := range db.GetAll() {
	// 	fmt.Printf("Key: %v - Value: %#v \n", k, val)
	// }

	if port := os.Getenv("PORT"); port != "" {
		router.Run(":" + port)
	} else {
		router.Run(":8080")
	}
}

func routes() {
	router.POST("/purchases", challenge, controllers.CreatePurchase)
	router.GET("/purchases", controllers.GetPurchases)
	router.GET("/purchases/:id", controllers.ReadPurchases)
	router.PUT("/purchases/:id", challenge, controllers.UpdatePurchase)
	router.DELETE("/purchases/:id", challenge, controllers.DeletePurchase)

	router.POST("/users")
	router.GET("/users")
	router.GET("/users/:id")
	router.PUT("/users/:id")
	router.DELETE("/users/:id")
}

func checkQueryParams(c *gin.Context) {
	if userID := c.Query("user_id"); userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Query params user_id required.")
	}
}

func onlyAdmin(c *gin.Context) {
	if role := c.GetHeader("role"); role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized!")
	}
}

func challenge(c *gin.Context) {
	if config.IsProduction() {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, "Only gets allowed, for challenge purpose! - visit github.com/seansa/Workshop_go-UTNFRT")
		return
	}
}
