package handlers

import (
	"net/http"

	"github.com/erkylima/posts-service/internal/pages/domains"
	"github.com/erkylima/posts-service/internal/pages/ports"
	"github.com/erkylima/posts-service/internal/shared/entities"
	"github.com/gin-gonic/gin"
)

type pagesHandler struct {
	svc ports.PageService
}

func NewPagesHandler(svc ports.PageService) *pagesHandler {
	return &pagesHandler{svc: svc}
}

func (h *pagesHandler) CreateHandler(ctx *gin.Context) {

	var page domains.Page
	if err := ctx.ShouldBindJSON(&page); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	_, err := h.svc.CreatePage(&page)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Header("X-Slug", page.Slug)
	ctx.Status(http.StatusCreated)
}

func (h *pagesHandler) ListHandler(ctx *gin.Context) {
	var filters []entities.Filter
	if err := ctx.ShouldBindJSON(&filters); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	pages, err := h.svc.ListPages(filters)
	if err != nil {
		HandleError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, pages)
}
