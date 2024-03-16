package services

import (
	"github.com/erkylima/posts-service/internal/posts/domains"
	"github.com/erkylima/posts-service/internal/shared/entities"
	"github.com/erkylima/posts-service/internal/shared/ports"
)

type PostService struct {
	repo ports.Repository[domains.Post]
}

func NewPostService(repo ports.Repository[domains.Post]) *PostService {
	return &PostService{repo: repo}
}

func (ps *PostService) CreatePost(post *domains.Post) (string, error) {
	result, err := ps.repo.Create(post)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (ps *PostService) ReadPost(slug string) (*domains.Post, error) {
	post := &domains.Post{}
	err := ps.repo.Read(slug, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *PostService) UpdatePost(slug string, post *domains.Post) (*domains.Post, error) {
	err := ps.repo.Update(slug, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (ps *PostService) DeletePost(slug string) error {
	err := ps.repo.Delete(slug)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostService) ListPosts(filter []entities.Filter) ([]domains.Post, error) {

	posts, err := ps.repo.List(filter)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
