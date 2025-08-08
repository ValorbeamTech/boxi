package handler

import (
	"database/sql"
	"net/http"

	"github.com/francis/projectx-api/internal/model"
	"github.com/francis/projectx-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
    db     *sql.DB
    logger logger.Logger
}

func NewHealthHandler(db *sql.DB, logger logger.Logger) *HealthHandler {
    return &HealthHandler{
        db:     db,
        logger: logger,
    }
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
    // Check database connection
    if err := h.db.Ping(); err != nil {
        h.logger.Error("Database health check failed", err)
        c.JSON(http.StatusServiceUnavailable, model.ErrorResponse("Database connection failed"))
        return
    }

    response := map[string]interface{}{
        "status":   "healthy",
        "database": "connected",
        "version":  "1.0.0",
    }

    c.JSON(http.StatusOK, model.SuccessResponse(response, "Service is healthy"))
}