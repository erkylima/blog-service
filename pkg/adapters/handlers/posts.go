package handlers

import (
	"net/http"

	"github.com/erkylima/posts-service/internal/posts/domains"
	"github.com/erkylima/posts-service/internal/posts/ports"
	"github.com/erkylima/posts-service/internal/shared/entities"
	"github.com/gin-gonic/gin"
)

type postsHandler struct {
	svc ports.PostService
}

func NewPostsHandler(svc ports.PostService) *postsHandler {
	return &postsHandler{svc: svc}
}

func (h *postsHandler) CreateHandler(ctx *gin.Context) {

	var post domains.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	_, err := h.svc.CreatePost(&post)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Header("X-Slug", post.Slug)
	ctx.Status(http.StatusCreated)
}

func (h *postsHandler) ListHandler(ctx *gin.Context) {
	var filters []entities.Filter
	if err := ctx.ShouldBindJSON(&filters); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	posts, err := h.svc.ListPosts(filters)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, posts)
}
