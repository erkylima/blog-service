package services

import (
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/ports"
)

type CommentService struct {
	repo ports.CommentRepository
}

func NewCommentService(repo ports.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment domains.Comment) (*domains.Comment, error) {
	return s.repo.CreateComment(comment)
}

func (s *CommentService) ReadComment(id string) (*domains.Comment, error) {
	return s.repo.ReadComment(id)
}

func (s *CommentService) UpdateComment(comment domains.Comment) (*domains.Comment, error) {
	return s.repo.UpdateComment(comment)
}

func (s *CommentService) DeleteComment(id string) error {
	return s.repo.DeleteComment(id)
}

func (s *CommentService) ListComments() ([]domains.Comment, error) {
	return s.repo.ListComments()
}

func (s *CommentService) ListCommentsByPost(postId string) ([]domains.Comment, error) {
	return s.repo.ListCommentsByPost(postId)
}

func (s *CommentService) ListCommentsByAuthor(author string) ([]domains.Comment, error) {
	return s.repo.ListCommentsByAuthor(author)
}

func (s *CommentService) ListCommentsByTag(tag string) ([]domains.Comment, error) {
	return s.repo.ListCommentsByTag(tag)
}
