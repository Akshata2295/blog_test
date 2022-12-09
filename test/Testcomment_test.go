package test

import (
	"blog_test/controller"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateComment(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	body := bytes.NewBuffer([]byte("{\"ApiTest\":\"PostReq\"}"))
	req, err := http.NewRequest("POST", "/api/v1/article/1", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "form-data")

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := gin.Default()
	router.POST("/api/v1/comment/:article_id", controller.CreateComment)

	// Create Response Recorder
	w := httptest.NewRecorder()

	fmt.Println(w.Body)
}

func TestCommentOnComment(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	body := bytes.NewBuffer([]byte("{\"ApiTest\":\"PostReq\"}"))
	req, err := http.NewRequest("POST", "api/v1/comments/1/1", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")


	router := gin.Default()
	router.GET("api/v1/comments/:article_id/:comment_id", controller.CommentOnComment)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	fmt.Println(w.Body)
}
