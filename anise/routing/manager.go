package routing

import (
	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies"
	"github.com/gin-gonic/gin"
)

type Manager interface {
	RegisterRoutes(engine *gin.Engine, depedencyResolver dependencies.Resolver) error
}
