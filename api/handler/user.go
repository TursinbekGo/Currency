package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create User godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "CreateUserRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) CreateUser(c *gin.Context) {

	var createUser models.CreateUser
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "error user should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.User().Create(c.Request.Context(), &createUser)
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create user resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID user godoc
// @ID get_by_id_user
// @Router /user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdUser(c *gin.Context) {
	info, exists := c.Get("user_id")
	if !exists {
		h.handlerResponse(c, "create user", http.StatusUnauthorized, "login first")
		return
	}

	userData := info.(helper.TokenInfo)
	id := ""
	if len(userData.UserID) > 0 {
		id = userData.UserID
	} else {
		id = c.Param("id")
	}
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id User resposne", http.StatusOK, resp)
}

// GetList user godoc
// @ID get_list_user
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListUser(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list User offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list User limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.User().GetList(c.Request.Context(), &models.UserGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.User.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list User resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update user godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param User body models.UpdateUser true "UpdateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateUser(c *gin.Context) {

	var (
		id         string = c.Param("id")
		updateUser models.UpdateUser
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		h.handlerResponse(c, "error User should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateUser.Id = id
	rowsAffected, err := h.strg.User().Update(c.Request.Context(), &updateUser)
	if err != nil {
		h.handlerResponse(c, "storage.User.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.User.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.User().GetByID(c.Request.Context(), &models.UserPrimaryKey{Id: updateUser.Id})
	if err != nil {
		h.handlerResponse(c, "storage.User.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create User resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete user godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteUser(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.User().Delete(c.Request.Context(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create User resposne", http.StatusNoContent, nil)
}
