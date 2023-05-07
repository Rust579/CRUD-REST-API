package main

import (
	routes "CRUD_API/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/", routes.CreatePost)

	// called as localhost:3000/getOne/{id}
	router.GET("getOne/:userId", routes.ReadOnePost)

	// called as localhost:3000/update/{id}
	router.PUT("/update/:userId", routes.UpdatePost)

	// called as localhost:3000/delete/{id}
	router.DELETE("/delete/:userId", routes.DeletePost)

	// called as localhost:3000/getall
	router.GET("get/", routes.ReadAllUsers)

	router.Run(":3000")
}
