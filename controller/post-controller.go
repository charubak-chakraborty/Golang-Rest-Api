package controller

import (
	"encoding/json"
	"net/http"

	"../entity"
	"../errors"
	"../service"
)

var (
	postService service.PostService
)

type controller struct{}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
}

func (*controller) GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		// res.Write([]byte(`{"error":"error getting the post"}`))
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "error getting the post"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func (*controller) AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var post entity.Post

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		// res.Write([]byte(`{"error":"error unmarshalling the request"}`))
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "error unmarshalling the request"})
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	result, err := postService.Create(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "error saving the post"})
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}
