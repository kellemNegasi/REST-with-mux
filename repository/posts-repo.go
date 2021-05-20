package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/kellemnegasi/restapi-with-mux/entity"
	"google.golang.org/api/iterator"
)

type PostRepository interface {
	Save(posts *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

const (
	projectId      = "golang-project-fa5f4"
	collectionName = "posts"
)

type repo struct {
}

func NewPostRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to create a firestore client %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("failed to add a new post %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to create a firestore client %v", err)
		return nil, err
	}
	defer client.Close()

	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			log.Fatalf("error in iterating through the posts %v", err)
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
