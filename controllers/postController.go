package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/jsonCrudGorm/initializers"
	"github.com/lordofthemind/jsonCrudGorm/models"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, gin.H{"data": post, "message": "Post created successfully"})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": posts})
}

func GetPost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
	var post models.Post

	// Retrieve the post with the given ID
	thatPost := initializers.DB.First(&post, c.Param("id"))
	if thatPost.Error != nil {
		c.JSON(500, gin.H{"error": thatPost.Error.Error()})
		return
	}

	// Update the post with the data from the JSON payload
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Save the updated post
	result := initializers.DB.Save(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"data": post, "message": "Post updated successfully"})
}

// func UpdatePost(c *gin.Context) {
// 	var post models.Post
// 	if err := c.ShouldBindJSON(&post); err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
// 		return
// 	}

// 	result := initializers.DB.Save(&post)
// 	if result.Error != nil {
// 		c.JSON(500, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"data": post, "message": "Post updated successfully"})
// }

func DeletePost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.Delete(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
