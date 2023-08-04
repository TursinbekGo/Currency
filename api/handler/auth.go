package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Create Register
// @Description Register
// @Tags Register
// @Accept json
// @Produce json
// @Param user body models.Register true "CreateRegisterRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) Register(c *gin.Context) {
	var registerUser models.Register

	err := c.ShouldBindJSON(&registerUser) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	if len(registerUser.UserName) < 6 && len(registerUser.Password) < 6 {
		h.handlerResponse(c, "incorrect", http.StatusBadRequest, err.Error())
	}

	_, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{UserName: registerUser.UserName})
	if err == nil {
		h.handlerResponse(c, "storage.user.create", http.StatusBadRequest, "already exists")
		return
	} else {

		if err.Error() == "no rows in result set" {

		} else {
			h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
			return
		}
	}

	id, err := h.strg.User().Create(context.Background(), &models.CreateUser{
		UserName: registerUser.UserName,
		Password: registerUser.Password,
	})
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{
		Id: id,
	})

	if err != nil {
		h.handlerResponse(c, "storage.user.getbyid", http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, user)

}

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Create Login
// @Description Login
// @Tags Login
// @Accept json
// @Produce json
// @Param user body models.Login true "CreateLoginRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) LOGIN(c *gin.Context) {
	var loginUser models.Login

	err := c.ShouldBindJSON(&loginUser) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	if len(loginUser.UserName) < 6 && len(loginUser.Password) < 6 {
		h.handlerResponse(c, "incorrect", http.StatusBadRequest, err.Error())
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{UserName: loginUser.UserName})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, "user not exists")
			return
		}
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

	if loginUser.Password != resp.Password {
		if err != nil {
			h.handlerResponse(c, "storage.user.create", http.StatusBadRequest, "incorrect password")
			return
		}
	}
	data := make(map[string]interface{})
	data["user_id"] = resp.Id

	tkn, err := helper.GenerateJWT(data, time.Hour*100, h.cfg.SecretKey)
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusBadRequest, "token errors")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": tkn})

}
