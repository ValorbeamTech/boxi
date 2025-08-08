package handler

import (
	"net/http"

	"github.com/francis/projectx-api/pkg/logger"
	"github.com/francis/projectx-api/pkg/validator"

	"github.com/francis/projectx-api/internal/model"
	"github.com/francis/projectx-api/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    authService *service.AuthService
    logger      logger.Logger
}

func NewAuthHandler(authService *service.AuthService, logger logger.Logger) *AuthHandler {
    return &AuthHandler{
        authService: authService,
        logger:      logger,
    }
}

func (h *AuthHandler) Register(c *gin.Context) {
    var req model.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, model.ErrorResponse("Invalid request body"))
        return
    }

    if err := validator.Validate(&req); err != nil {
        c.JSON(http.StatusBadRequest, model.ErrorResponse(err.Error()))
        return
    }

    user, err := h.authService.Register(c.Request.Context(), &req)
    if err != nil {
        h.logger.Error("Failed to register user", err)
        c.JSON(http.StatusBadRequest, model.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, model.SuccessResponse(user, "User registered successfully"))
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req model.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, model.ErrorResponse("Invalid request body"))
        return
    }

    if err := validator.Validate(&req); err != nil {
        c.JSON(http.StatusBadRequest, model.ErrorResponse(err.Error()))
        return
    }

    response, err := h.authService.Login(c.Request.Context(), &req)
    if err != nil {
        h.logger.Error("Failed to login user", err)
        c.JSON(http.StatusUnauthorized, model.ErrorResponse(err.Error()))
        return
    }

    c.JSON(http.StatusOK, model.SuccessResponse(response, "Login successful"))
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
    // Implementation for refresh token
    c.JSON(http.StatusOK, model.SuccessResponse(nil, "Token refreshed"))
}