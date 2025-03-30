package routing

import (
	"github.com/gin-gonic/gin"
)

type Manager interface {
	RegisterRoutes(engine *gin.Engine) error
}
