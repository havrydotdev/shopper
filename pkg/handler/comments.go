package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopper"
	"strconv"
)

// @Summary AddComment
// @Security ApiKeyAuth
// @Tags comments
// @Description Create comment endpoint
// @ID addComment
// @Accept  json
// @Produce  json
// @Param input body shopper.Comment true "Comment text, user and item id"
// @Param item_id path int true  "Item id"
// @Success 200 {object} idJson
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/items/{item_id}/comments/ [post]
func (h *Handler) addComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input shopper.Comment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.AddComment(itemId, userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idJson{
		Id: id,
	})
}

// @Summary UpdateComment
// @Security ApiKeyAuth
// @Tags comments
// @Description Update comment endpoint
// @ID updateComment
// @Accept json
// @Param input body shopper.UpdateCommentInput true "Updated comment text"
// @Param id path int true  "Comment id to be updated"
// @Produce json
// @Success 200 {object} status
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/comments/{id} [put]
func (h *Handler) updateComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input shopper.UpdateCommentInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateComment(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status{
		Ok: true,
	})
}

// @Summary DeleteComment
// @Security ApiKeyAuth
// @Tags comments
// @Description Delete comment endpoint
// @ID deleteComment
// @Accept  json
// @Param id path int true  "Comment id to be deleted"
// @Produce  json
// @Success 200 {object} status
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/comments/{id} [delete]
func (h *Handler) deleteComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := h.services.DeleteComment(userId, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, status{
		Ok: true,
	})
}

// @Summary GetCommentsByItem
// @Security ApiKeyAuth
// @Tags comments
// @Description Endpoint to get all item`s comments
// @ID getCommentsByItem
// @Accept  json
// @Param item_id path int true "Comment id to be returned"
// @Produce  json
// @Success 200 {object} []shopper.
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/items/{item_id}/comments [get]
func (h *Handler) getCommentsByItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	comments, err := h.services.GetCommentsByItem(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}
