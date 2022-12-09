package controller

import (

	"blog_test/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateArticle(c *gin.Context) {
	var article models.Article

	err := c.ShouldBindWith(&article, binding.Form)
		if err != nil {
			c.JSON(400, gin.H{"message": "could not parse the details"})
			return
	    }
	
		err = models.DB.Where("title = ?", article.Title).First(&article).Error
		if err == nil {
			c.JSON(400, gin.H{"error": "article already exists."})
		 return
		 }

		 now := time.Now()
		date := now.Format("Jan 2, 2006")
	
	 article.CreatedAt = date
	
	 err = models.DB.Select("title", "created_at", "nickname", "content").Create(&article).Error
	if err != nil {
	    c.JSON(400, gin.H{"message": err.Error()})
	   return
	}
	
	c.JSON(200, gin.H{
			"id": article.ID,
			"title": article.Title,
			"nickname": article.Nickname,
			"content": article.Content,
			"created_at": article.CreatedAt,
			"message": "Succesfully Created Article",

	 })
	
	}
	

func ListAllArticle(c *gin.Context) {
		var articles []models.Article

		page, _ := strconv.Atoi(c.Query("page"))
		size, _ := strconv.Atoi(c.Query("size"))

		err := models.DB.Table("articles").Limit(size).Offset((page - 1) * size).Find(&articles).Error
		if err != nil {
			c.JSON(400, gin.H{
				"message": "error in fetching from database",
			})
			return
		}
		
		c.JSON(http.StatusOK, articles)

}


func GetArticle(c *gin.Context) {
	var article models.Article

	// Check if the category already exists.
	err := models.DB.Where("id = ?", c.Param("id")).First(&article).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "article does not exists."})
		return
	}

	// GET FROM CACHE FIRST
	c.JSON(http.StatusOK, gin.H{"content": article.Content})
	
}