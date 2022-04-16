package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	entity "github.com/epileftro85/goapi/Entity"
	repository "github.com/epileftro85/goapi/Repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Error geting the post array" }`))
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
		w.Write([]byte(`{ "error": "Error unmarshalling the request" }`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
