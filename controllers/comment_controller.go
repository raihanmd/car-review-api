package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raihanmd/fp-superbootcamp-go/exceptions"
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	_ "github.com/raihanmd/fp-superbootcamp-go/model/web"
	"github.com/raihanmd/fp-superbootcamp-go/model/web/request"
	_ "github.com/raihanmd/fp-superbootcamp-go/model/web/response"
	"github.com/raihanmd/fp-superbootcamp-go/services"
	"github.com/raihanmd/fp-superbootcamp-go/utils"
)

type CommentController interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type commentControllerImpl struct {
	services.CommentService
}

func NewCommentController(commentService services.CommentService) CommentController {
	return &commentControllerImpl{commentService}
}

// Create comment godoc
// @Summary Create comment.
// @Description Create a comment.
// @Tags Comments
// @Param Body body request.CommentCreateRequest true "the body to create a comment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} web.WebSuccess[response.CommentResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/comments [post]
func (controller *commentControllerImpl) Create(c *gin.Context) {
	var commentCreateReq request.CommentCreateRequest

	err := c.ShouldBindJSON(&commentCreateReq)
	helper.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	commentRes, err := controller.CommentService.Create(c, &commentCreateReq, userID)
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusCreated, commentRes, nil)
}

// Update comment godoc
// @Summary Update comment.
// @Description Update a comment.
// @Tags Comments
// @Param id path int true "Comment ID"
// @Param Body body request.CommentUpdateRequest true "the body to update a comment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} web.WebSuccess[response.CommentResponse]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/comments/{id} [patch]
func (controller *commentControllerImpl) Update(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	var commentUpdateReq request.CommentUpdateRequest

	err = c.ShouldBindJSON(&commentUpdateReq)
	helper.PanicIfError(err)

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	commentRes, err := controller.CommentService.Update(c, &commentUpdateReq, userID, uint(commentID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, commentRes, nil)
}

// Delete comment godoc
// @Summary Delete comment.
// @Description Delete a comment.
// @Tags Comments
// @Param id path int true "Comment ID"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 201 {object} web.WebSuccess[string]
// @Failure 400 {object} web.WebBadRequestError
// @Failure 404 {object} web.WebNotFoundError
// @Failure 500 {object} web.WebInternalServerError
// @Router /api/comments/{id} [delete]
func (controller *commentControllerImpl) Delete(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		panic(exceptions.NewCustomError(http.StatusBadRequest, "id must be an integer"))
	}

	userID, _, err := utils.ExtractTokenClaims(c)
	helper.PanicIfError(err)

	err = controller.CommentService.Delete(c, userID, uint(commentID))
	helper.PanicIfError(err)

	helper.ToResponseJSON(c, http.StatusOK, "comment deleted", nil)
}
