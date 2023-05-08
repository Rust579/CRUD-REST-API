package Routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	model "CRUD_API/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

func CreatePost(c *gin.Context) {

	var DB = database.ConnectDB()
	var userCollection = getcollection.GetCollection(DB, "Users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	user := new(model.Users)
	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	userPayload := model.Users{
		ID:   primitive.NewObjectID(),
		Name: user.Name,
		Age:  user.Age,
	}
	startTime := time.Now()
	result, err := userCollection.InsertOne(ctx, userPayload)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Request create processed in %s\n", elapsedTime)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	createdUser := model.Users{}
	err = userCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&createdUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": createdUser})

}
