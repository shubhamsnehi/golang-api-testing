package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shubhamsnehi/golang-api-testing/cache"
	"github.com/shubhamsnehi/golang-api-testing/entity"
	"github.com/shubhamsnehi/golang-api-testing/repository"
	"github.com/shubhamsnehi/golang-api-testing/service"
)

const (
	ID   string = "123"
	Name string = "Controller Test"
)

var (
	postRepo       repository.PostRepository = repository.NewMySQLRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postCh         cache.PostCache           = cache.NewRedisCache("localhost:6379", 0, 10)
	postController PostController            = NewPostController(postSrv, postCh)
)

func TestAddPost(t *testing.T) {
	// Create new HTTP request
	var jsonStr = []byte(`{"Name":"` + Name + `","text":"` + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.AddPost)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var post entity.User
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.ID)
	assert.Equal(t, Name, post.Name)

	// Cleanup database
	tearDown(post.Name)
}

func setup() {
	var post entity.User = entity.User{
		ID:   ID,
		Name: Name,
	}
	postRepo.Save(&post)
}

func tearDown(postID string) {
	var post entity.User = entity.User{
		ID: postID,
	}
	postRepo.Delete(&post)
}

func TestGetPosts(t *testing.T) {

	// Insert new post
	setup()

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.GetPosts)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var posts []entity.User
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP response
	assert.Equal(t, ID, posts[0].ID)
	assert.Equal(t, Name, posts[0].Name)

	// Cleanup database
	tearDown(ID)
}

func TestGetPostByID(t *testing.T) {

	// Insert new post
	setup()

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/posts/123", nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.GetPostByID)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var post entity.User
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.Equal(t, ID, post.ID)
	assert.Equal(t, Name, post.Name)

	// Cleanup database
	tearDown(ID)
}
