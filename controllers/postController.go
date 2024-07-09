package controllers

import (
	"go-api-server/initializers"
	"go-api-server/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	// get data from req body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	var body struct {
		ID    int
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post

	initializers.DB.First(&post, body.ID)

	post.Body = body.Body
	post.Title = body.Title

	initializers.DB.Save(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Delete post successfully!",
	})
}
