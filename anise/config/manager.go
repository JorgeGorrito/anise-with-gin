package config

import "github.com/gin-gonic/gin"

type Manager interface {
	ConfigureEngine(engine *gin.Engine) error
	ConfigureApplication() error
}
