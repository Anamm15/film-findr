package middleware

import (
	"net/http"
	"slices"

	"FilmFindr/dto"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
)

func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleInterface, exists := c.Get("role")
		if !exists {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, dto.MESSAGE_FAILED_REQUIRED_FIELD, nil)
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		role := roleInterface.(string)
		if slices.Contains(allowedRoles, role) {
			c.Next()
			return
		}

		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UNAUTHORIZED, dto.MESSAGE_FAILED_UNAUTHORIZED, nil)
		c.JSON(http.StatusForbidden, res)
		c.Abort()
	}
}
