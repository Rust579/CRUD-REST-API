package Routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	model "CRUD_API/model"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func ReadOnePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var userCollection = getcollection.GetCollection(DB, "Users")

	userId := c.Param("userId")
	var result model.Users

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "data": res})
}
