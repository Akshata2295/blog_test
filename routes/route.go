package routes

import (
	"blog_test/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	// r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		//allow all
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			platform := c.GetHeader("User-agent")
			fmt.Println("platform ", platform)
		}
		c.Next()
	})

	v1 := r.Group("/api/v1")
	article := v1.Group("/")
	article.POST("/article", controller.CreateArticle)
	article.GET("/articles", controller.ListAllArticle)
	article.GET("article/:id", controller.GetArticle)

	comment := v1.Group("/")
	comment.POST("/comment/:article_id", controller.CreateComment)
	comment.POST("/comments/:article_id/:comment_id", controller.CommentOnComment)

	return r
}
