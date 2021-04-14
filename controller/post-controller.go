package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/shubhamsnehi/golang-api-testing/cache"
	"github.com/shubhamsnehi/golang-api-testing/entity"
	"github.com/shubhamsnehi/golang-api-testing/errors"
	"github.com/shubhamsnehi/golang-api-testing/service"
)

var (
	postService service.PostService
	postCache   cache.PostCache
)

type controller struct{}

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
	GetPostByID(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) GetPostByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	postID := strings.Split(request.URL.Path, "/")[2]
	var post *entity.User = postCache.Get(postID)
	if post == nil {
		post, err := postService.FindByID(postID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "No post found"})
			return
		}
		postCache.Set(postID, post)
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.User
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		//response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		//response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}

	result, err2 := postService.Create1(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		//response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
