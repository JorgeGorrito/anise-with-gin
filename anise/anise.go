package anise

import (
	"errors"

	"github.com/JorgeGorrito/anise-with-gin/anise/config"
	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies"
	ea "github.com/JorgeGorrito/anise-with-gin/anise/errors"
	"github.com/JorgeGorrito/anise-with-gin/anise/routing"
	"github.com/gin-gonic/gin"
)

type WebApplication struct {
	engine                *gin.Engine
	configManager         config.Manager
	routesManager         routing.Manager
	dependenciesManager   dependencies.Manager
	dependenciesContainer *dependencies.Container
}

func NewWebApplication(
	engine *gin.Engine,
	configManager config.Manager,
	routesManager routing.Manager,
	dependenciesManager dependencies.Manager,
) *WebApplication {
	var errorList error
	if engine == nil {
		errorList = errors.Join(errorList, ea.ErrEngineIsNil)
	}
	if configManager == nil {
		errorList = errors.Join(errorList, ea.ErrConfigManagerIsNil)
	}
	if routesManager == nil {
		errorList = errors.Join(errorList, ea.ErrRoutesManagerIsNil)
	}
	if dependenciesManager == nil {
		errorList = errors.Join(errorList, ea.ErrDependenciesManagerIsNil)
	}

	if errorList != nil {
		panic(errorList)
	}
	return &WebApplication{
		engine:                engine,
		configManager:         configManager,
		routesManager:         routesManager,
		dependenciesManager:   dependenciesManager,
		dependenciesContainer: dependencies.NewContainer(),
	}
}

func (app *WebApplication) ConfigureApplication() error {
	return app.configManager.ConfigureApplication()
}

func (app *WebApplication) configureEngine() error {
	return app.configManager.ConfigureEngine(app.engine)
}

func (app *WebApplication) registerDependencies() error {
	return app.dependenciesManager.RegisterDependencies(app.dependenciesContainer)
}

func (app *WebApplication) registerRoutes() error {
	return app.routesManager.RegisterRoutes(app.engine, app.dependenciesContainer)
}

func (app *WebApplication) Run(addr ...string) {
	var errorList error

	if err := app.registerDependencies(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.ConfigureApplication(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.configureEngine(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.registerRoutes(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.engine.Run(addr...); err != nil {
		errorList = errors.Join(errorList, err)
	}

	panic(errorList)
}
