package Routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func DeletePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	userId := c.Param("userId")

	var postCollection = getcollection.GetCollection(DB, "Users")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(userId)
	result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete " + userId})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User deleted successfully", "data": res})
}
