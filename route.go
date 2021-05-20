package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "title1", Text: "text1"}}
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	result, err := json.Marshal(posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"Error":"Error marshaling the psots"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)

}
