package middleware

import (
	"net/http"
	"slices"

	"ReviewPiLem/dto"
	"ReviewPiLem/utils"

	"github.com/gin-gonic/gin"
)

func AuthorizeRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleInterface, exists := c.Get("role")
		if !exists {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, dto.MESSAGE_FAILED_DENIED_ACCESS, nil)
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		role := roleInterface.(string)
		if slices.Contains(allowedRoles, role) {
			c.Next()
			return
		}

		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, dto.MESSAGE_FAILED_DENIED_ACCESS, nil)
		c.JSON(http.StatusForbidden, res)
		c.Abort()
	}
}
