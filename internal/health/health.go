package health

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h.db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "error",
				"service": "database",
				"message": err.Error(),
			})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "go-setup-api",
			"version": "1.0.0",
		})
	}
}
