package controllers

import (
	"test-sharing-vision/app/models"
	"test-sharing-vision/platform/database"
	"test-sharing-vision/utils"

	"github.com/gin-gonic/gin"
)

// get articles with pagination
func GetArticles(c *gin.Context) {
	db, err := database.OpenDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	limit := utils.StringToInt(c.Param("limit"))
	offset := utils.StringToInt(c.Param("offset"))
	articles, err := db.GetArticleWithPagination(offset, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)

}

// get article by id
func GetArticleById(c *gin.Context) {
	db, err := database.OpenDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	articles, err := db.GetArticleById(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)

}

// delete article by id
func DeleteArticleById(c *gin.Context) {
	db, err := database.OpenDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = db.DeleteArticleById(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "article deleted"})

}

// update article by id
func UpdateArticleById(c *gin.Context) {
	var article models.Article
	db, err := database.OpenDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.BindJSON(&article)
	article_valid, hasil := utils.ValidateArticle(article)
	if !article_valid {
		c.JSON(400, gin.H{"error": hasil})
		return
	}

	_, err = db.UpdateArticleById(article)
	if err != nil {
		c.JSON(404, gin.H{"error": "article not found" + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "article updated"})

}

// create article
func CreateArticle(c *gin.Context) {
	var article models.Article
	db, err := database.OpenDBConnection()
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.BindJSON(&article)
	article_valid, hasil := utils.ValidateArticle(article)
	if !article_valid {
		c.JSON(400, gin.H{"error": hasil})
		return
	}

	_, err = db.CreateArticle(article)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "article created"})

}
