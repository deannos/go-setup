package http

import (
	"net/http"
	"strconv"

	"go-setup/internal/entity"
	"go-setup/internal/usecase"
	"go-setup/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUC *usecase.UserUsecase
}

func NewUserHandler(userUC *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request")
		return
	}

	ctx := c.Request.Context()
	if err := h.userUC.CreateUser(ctx, &user); err != nil {
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User created", user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	
	ctx := c.Request.Context()
	user, err := h.userUC.GetUserByID(ctx, id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	response.Success(c, http.StatusOK, "User found", user)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.userUC.ListUsers(ctx)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to list users")
		return
	}

	response.Success(c, http.StatusOK, "Users listed", users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request")
		return
	}
	user.ID = id
	
	ctx := c.Request.Context()
	if err := h.userUC.UpdateUser(ctx, &user); err != nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	response.Success(c, http.StatusOK, "User updated", user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	
	ctx := c.Request.Context()
	if err := h.userUC.DeleteUser(ctx, id); err != nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	response.Success(c, http.StatusOK, "User deleted", nil)
}