package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/viviyee/go-crud/initializers"
	"github.com/viviyee/go-crud/models"
)

var postReq struct {
	Title string
	Body  string
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func StorePost(c *gin.Context) {

	c.Bind(&postReq)

	post := models.Post{
		Title: postReq.Title,
		Body:  postReq.Body,
	}

	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(201, gin.H{
		"post": post,
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
	id := c.Param("id")
	c.Bind(&postReq)

	// post
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: postReq.Title,
		Body:  postReq.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.Delete(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}
