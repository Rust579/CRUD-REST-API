package Routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/databases"
	"CRUD_API/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func ReadAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var DB = database.ConnectDB()
	var userCollection = getcollection.GetCollection(DB, "Users")

	var results []model.Users

	startTime := time.Now()
	cursor, err := userCollection.Find(ctx, bson.M{})
	elapsedTime := time.Since(startTime)
	fmt.Printf("Request ReadAll processed in %s\n", elapsedTime)
	for cursor.Next(ctx) {
		// обработка данных
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result model.Users
		err := cursor.Decode(&result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{"message": "success!", "data": results})
}
