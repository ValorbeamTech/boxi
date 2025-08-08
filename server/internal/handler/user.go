package handler

import (
	"net/http"
	"strconv"

	"github.com/francis/projectx-api/internal/model"
	"github.com/francis/projectx-api/internal/service"
	"github.com/francis/projectx-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    userService *service.UserService
    logger      logger.Logger
}

func NewUserHandler(userService *service.UserService, logger logger.Logger) *UserHandler {
    return &UserHandler{
        userService: userService,
        logger:      logger,
    }
}

func (h *UserHandler) GetProfile(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, model.ErrorResponse("User not authenticated"))
        return
    }

    user, err := h.userService.GetByID(c.Request.Context(), userID.(int))
    if err != nil {
        h.logger.Error("Failed to get user profile", err)
        c.JSON(http.StatusNotFound, model.ErrorResponse("User not found"))
        return
    }

    c.JSON(http.StatusOK, model.SuccessResponse(user, "Profile retrieved successfully"))
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
    userID, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, model.ErrorResponse("User not authenticated"))
        return
    }

    var updateReq model.User
    if err := c.ShouldBindJSON(&updateReq); err != nil {
        c.JSON(http.StatusBadRequest, model.ErrorResponse("Invalid request body"))
        return
    }

    updateReq.ID = userID.(int)
    if err := h.userService.Update(c.Request.Context(), &updateReq); err != nil {
        h.logger.Error("Failed to update user profile", err)
        c.JSON(http.StatusInternalServerError, model.ErrorResponse("Failed to update profile"))
        return
    }

    c.JSON(http.StatusOK, model.SuccessResponse(updateReq, "Profile updated successfully"))
}

func (h *UserHandler) GetUsers(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    if page < 1 {
        page = 1
    }
    if limit < 1 || limit > 100 {
        limit = 10
    }

    users, err := h.userService.GetUsers(c.Request.Context(), page, limit)
    if err != nil {
        h.logger.Error("Failed to get users", err)
        c.JSON(http.StatusInternalServerError, model.ErrorResponse("Failed to get users"))
        return
    }

    c.JSON(http.StatusOK, model.SuccessResponse(users, "Users retrieved successfully"))
}