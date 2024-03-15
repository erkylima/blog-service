package services

import (
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/ports"
)

type CommentService struct {
	repo ports.Repository[domains.Post]
}

func NewCommentService(repo ports.Repository[domains.Post]) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment domains.Comment) (*domains.Comment, error) {
	// return s.repo.Create(comment)
	return nil, nil
}

func (s *CommentService) ReadComment(id string) (*domains.Comment, error) {
	// return s.repo.ReadComment(id)
	return nil, nil
}

func (s *CommentService) UpdateComment(comment domains.Comment) (*domains.Comment, error) {
	// return s.repo.UpdateComment(comment)
	return nil, nil
}

func (s *CommentService) DeleteComment(id string) error {
	// return s.repo.DeleteComment(id)
	return nil
}

func (s *CommentService) ListComments() ([]domains.Comment, error) {
	// return s.repo.ListComments()
	return nil, nil
}

func (s *CommentService) ListCommentsByPost(postId string) ([]domains.Comment, error) {
	// return s.repo.ListCommentsByPost(postId)
	return nil, nil
}

func (s *CommentService) ListCommentsByAuthor(author string) ([]domains.Comment, error) {
	// return s.repo.ListCommentsByAuthor(author)
	return nil, nil
}

func (s *CommentService) ListCommentsByTag(tag string) ([]domains.Comment, error) {
	// return s.repo.ListCommentsByTag(tag)
	return nil, nil
}
