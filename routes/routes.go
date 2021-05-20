package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/kellemnegasi/restapi-with-mux/entity"
	"github.com/kellemnegasi/restapi-with-mux/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"Error":"Error getting the psots"}`))
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	var post entity.Post
	res.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"Error":"Error unmarshaling the request"}`))
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)

}
