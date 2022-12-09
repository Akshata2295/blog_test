package models

type Article struct {
	ID        uint   `json:"id"`
	Title     string `form:"title" json:"title" binding:"required"`
	Nickname  string `form:"nickname" json:"nick_name" binding:"required"`
	Content   string `form:"content" json:"content" binding:"required"`
	CreatedAt string `json:"created_at"`
}

type Comment struct {
	ID        uint   `json:"id"`
	ArticleID uint   `json:"article_id"`
	Nickname  string `form:"nickname" json:"nickname" binding:"required"`
	Content   string `form:"content" json:"content" binding:"required"`
	CommentID uint   `json:"comment_id"`
	CreatedAt string `json:"created_at"`
}
