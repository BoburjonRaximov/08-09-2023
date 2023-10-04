package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"playground/cpp-bootcamp/api/response"
	"playground/cpp-bootcamp/models"
	"playground/cpp-bootcamp/pkg/helper"
	"playground/cpp-bootcamp/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

// create person handler
// @Router       /user [post]
// @Summary      create user
// @Description  api for create users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user    body     models.CreateUser  true  "data of user"
// @Success      200  {object}  response.CreateResponse
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateUser(c *gin.Context) {
	var user models.CreateUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "BAD REQUEST", Message: "invalid fields in body"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	bytes, _ := io.ReadAll(c.Request.Body)
	fmt.Println(string(bytes))
	fmt.Println(user)
	resp, err := h.storage.User().Create(user)
	if err != nil {
		fmt.Println("error user Create:", err.Error())
		res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})
}

// @Router       /user/{id} [put]
// @Summary      update user
// @Description  api for update users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of user"
// @Param        user    body     models.CreateUser  true  "data of user"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	user.Id = c.Param("id")
	resp, err := h.storage.User().Update(user)
	if err != nil {
		fmt.Println("error user Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /user/{id} [get]
// @Summary      List users
// @Description  get users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of user"  Format(uuid)
// @Success      200  {object}   models.User
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetUser(c *gin.Context) {
	fmt.Println("MethodGet")

	id := c.Param("id")

	resp, err := h.storage.User().Get(models.RequestByID{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Person Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Security ApiKeyAuth
// @Router       /user [get]
// @Summary      List users
// @Description  get users
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success      200  {array}   models.User
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllUsers(c *gin.Context) {
	h.log.Info("request GetAllUsers")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	age, err := strconv.Atoi(c.DefaultQuery("age", "0"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	resp, err := h.storage.User().GetAll(models.GetAllUsersRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
		Age:    age,
	})
	if err != nil {
		h.log.Error("error User GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllUsers")
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error User GetAll:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.storage.User().Delete(models.RequestByID{ID: id})
	if err != nil {
		h.log.Error("error User GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /change/{id} [put]
// @Summary      change password
// @Description  api for update password
// @Tags         changePassword
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of user"
// @Param        changePassword    body     models.ChangePassword  true  "data of users"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) ChangePassword(c *gin.Context) {
	var user models.ChangePassword
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	user.Id = c.Param("id")
	resp, err := h.storage.User().ChangePassword(user)
	if err != nil {
		fmt.Println("error user changrepassword:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}
