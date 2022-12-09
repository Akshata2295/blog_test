package controller

import (
	"fmt"
	"blog_test/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
   var article models.Article

	err := c.ShouldBindWith(&comment, binding.Form)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{"message": "could not parse the details"})
			return
	    }
		 now := time.Now()
		date := now.Format("Jan 2, 2006")
	
	 comment.CreatedAt = date

	 ID :=c.Param("article_id")
	 str, _:= strconv.Atoi(ID)

	 comment.ArticleID = uint(str)
     


    
	 err = models.DB.Where("id = ?", comment.ArticleID).First(&article).Error
		if err != nil {
			c.JSON(400, gin.H{"error": "article already exists."})
		 return
		 }


	 err = models.DB.Select("nickname", "created_at", "article_id", "content").Create(&comment).Error
	if err != nil {
	    fmt.Println(err)
	}
	
	c.JSON(200, gin.H{
			"id": comment.ID,
			"nickname": comment.Nickname,
			"content": comment.Content,
			"created_at": comment.CreatedAt,
			"article_id": comment.ArticleID,
			"message": "Succesfully Created comment",

	 })
}

func CommentOnComment(c *gin.Context) {
    var comment models.Comment
    var article models.Article

    err := c.ShouldBindWith(&comment, binding.Form)
    if err != nil {
        fmt.Println(err)
        c.JSON(400, gin.H{"message": "could not parse the details"})
        return
    }
    now := time.Now()
    date := now.Format("Jan 2, 2006")

    comment.CreatedAt = date

    article_id := c.Param("article_id")
    id, _ := strconv.Atoi(article_id)

    comment.ArticleID = uint(id)

    if err = models.DB.Where("id = ?", comment.ArticleID).First(&article).Error; err != nil {
        c.JSON(400, gin.H{"message": "Article Does not exists"})
        return
    }

    Comment_id := c.Param("comment_id")
    id1, _ := strconv.Atoi(Comment_id)

    comment.CommentID = uint(id1)
    fmt.Println("id", comment.CommentID)

    if err = models.DB.Where("id = ?", comment.CommentID).First(&comment).Error; err != nil {
        c.JSON(400, gin.H{"message": "Comment Does not exists"})
        return
    }
    comment.CommentID = uint(id1)
    fmt.Println("id", comment.CommentID)

    err = models.DB.Select("created_at", "nickname", "content", "article_id", "comment_id").Create(&comment).Error
    if err != nil {
        // fmt.Println(err)
        c.JSON(400, gin.H{"message": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "id":                comment.ID,
        "article_id":        comment.ArticleID,
        "nickname":          comment.Nickname,
        "content":           comment.Content,
        "created_at":        comment.CreatedAt,
        "parent_comment_id": comment.CommentID,
        "message":           "Succesfully posted comment",
    })
}