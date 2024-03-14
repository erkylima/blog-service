package services

import (
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/ports"
)

type PostService struct {
	repo ports.Repository
}

func NewPostService(repo ports.Repository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post domains.Post) (*domains.Post, error) {
	// return s.repo.CreatePost(post)
	return nil, nil
}

func (s *PostService) ReadPost(slug string) (*domains.Post, error) {
	// return s.repo.ReadPost(slug)
	return nil, nil
}

func (s *PostService) UpdatePost(post domains.Post) (*domains.Post, error) {
	// return s.repo.UpdatePost(post)
	return nil, nil
}

func (s *PostService) DeletePost(slug string) error {
	// return s.repo.DeletePost(slug)
	return nil
}

func (s *PostService) ListPosts() ([]domains.Post, error) {
	// return s.repo.ListPosts()
	return nil, nil
}

func (s *PostService) ListPostsByTag(tag string) ([]domains.Post, error) {
	// return s.repo.ListPostsByTag(tag)
	return nil, nil
}

func (s *PostService) ListPostsByCategory(category string) ([]domains.Post, error) {
	// return s.repo.ListPostsByCategory(category)
	return nil, nil
}

func (s *PostService) ListPostsByAuthor(author string) ([]domains.Post, error) {
	// return s.repo.ListPostsByAuthor(author)
	return nil, nil
}

func (s *PostService) ListPostsByDate(date string) ([]domains.Post, error) {
	// return s.repo.ListPostsByDate(date)
	return nil, nil
}
