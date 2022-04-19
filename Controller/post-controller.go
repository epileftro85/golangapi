package controller

import (
	"encoding/json"
	"net/http"

	entity "github.com/epileftro85/goapi/Entity"
	service "github.com/epileftro85/goapi/Service"
	utils "github.com/epileftro85/goapi/Utils"
)

var (
	postService service.PostService = service.NewPostService()
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := utils.ServiceError{Message: "Error geting the post array"}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := utils.ServiceError{Message: "Error geting the post array"}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := utils.ServiceError{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	result, err1 := postService.Create(&post)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := utils.ServiceError{Message: "Error saving the post"}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
