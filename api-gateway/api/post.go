package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/grpc-load-balancer-example/api-gateway/api/models"
	"github.com/golanguzb70/grpc-load-balancer-example/api-gateway/genproto/post_service"
)

// @Router /post [post]
// @Summary post Endpoint
// @Description API for create Post, `Authorization` required
// @Tags post
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param post body post_service.CreateRequest true "post"
// @Success 200 {object} post_service.Post
// @Failure default {object} models.Response
func (h *Handler) CreatePost(c *gin.Context) {
	var (
		body *post_service.CreateRequest
	)

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"messages": "Invalid request body",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	res, err := h.services.PostService().Create(ctx, body)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Error while create post",
		})
		return
	}

	c.JSON(200, res)
}

// @Router /post/{key} [put]
// @Summary post Endpoint
// @Description API for update post, `Authorization` required
// @Tags post
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param key path string true "key"
// @Param post body post_service.UpdateRequest true "post"
// @Success 200 {object} post_service.Post
// @Failure default {object} models.Response
func (h *Handler) UpdatePost(c *gin.Context) {
	var (
		body *post_service.UpdateRequest
	)

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"messages": "Invalid request body",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	res, err := h.services.PostService().Update(ctx, body)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Error while update post",
		})
		return
	}

	c.JSON(200, res)
}

// @Router /post/{key} [get]
// @Summary post Endpoint
// @Description API for get post with id, `Authorization` required
// @Tags post
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param key path string true "key"
// @Success 200 {object} post_service.Post
// @Failure default {object} models.Response
func (h *Handler) GetPost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	response, err := h.services.PostService().GetByKey(ctx, &post_service.KeyRequest{
		Id: c.Param("key"),
	})
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Error while get post",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Router /post/{key} [delete]
// @Summary post Endpoint
// @Description API for delete post with id, `Authorization` required
// @Tags post
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param key path string true "key"
// @Success 200 {object} models.Response
// @Failure default {object} models.Response
func (h *Handler) DeletePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	_, err := h.services.PostService().Delete(ctx, &post_service.KeyRequest{
		Id: c.Param("key"),
	})
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Error while delete post",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Messages: "Post deleted",
	})
}

// @Router /post [get]
// @Security BearerAuth
// @Summary Post List
// @Description API for get list a of posts
// @Tags post
// @Accept  json
// @Produce  json
// @Param filters query post_service.Filter true "filters"
// @Param search query string false "search"
// @Success 200 {object} post_service.Posts
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
func (h *Handler) ListPost(c *gin.Context) {
	var (
		request = &post_service.Filter{}
		err     error
	)

	page, _ := strconv.Atoi(c.Query("page"))
	request.Page = int32(page)
	limit, _ := strconv.Atoi(c.Query("limit"))
	request.Limit = int32(limit)
	request.Search = c.Query("search")
	request.SortAsc, _ = strconv.ParseBool(c.Query("sort_asc"))

	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	response, err := h.services.PostService().Find(ctx, request)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Error while get list post",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
