package main

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"



/* A function that generates Random Strings of Fixed Length */
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}



func main() {
	/* Connect to the database */
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database("lenk-cf").Collection("shorturls")



	/* Create a new gin router */
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the URL Shortener API",
			"status":  http.StatusOK,
			"version": "1.0.0",
			"author":  "Adithya",
			"GETendpoints": []string{
				"/",
				"/:id",
				"all",
			},
			"POSTendpoints": []string{
				"/",
			},

		})
	})




	 r.GET("/all", func(c *gin.Context) {
		var result bson.M
		cur, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"});return
		};for cur.Next(context.TODO()) {
			err := cur.Decode(&result)
			if err != nil {c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()});return};c.JSON(http.StatusOK, result)
		}})


	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var result bson.M
		err := collection.FindOne(context.TODO(), bson.M{"shortId": id}).Decode(&result)
		if err != nil {c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"});return
		};c.Redirect(http.StatusMovedPermanently, result["url"].(string))
	})

	r.POST("/", func(c *gin.Context) {
		/* extract the variable from the body */
		var json struct {
			Url     string `json:"url"`
			ShortId string `json:"shortId"`
		}
	 /* check if shortId is null */
		if json.ShortId == "" {json.ShortId = RandStringBytes(6)}
		c.Bind(&json)

		/* check if shortId is present in database  */
		var result bson.M
		collection.FindOne(context.TODO(), bson.M{"shortId": json.ShortId}).Decode(&result)
		if result != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "ShortId already exists!"});return
		}

		/* insert the url and shortId into database */
		_, err := collection.InsertOne(context.TODO(), bson.M{"url": json.Url, "shortId": json.ShortId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "shortId" : json.ShortId, "url" : json.Url, "message": "ShortUrl created successfully!"})
	})

	r.Run(":3000")

}
