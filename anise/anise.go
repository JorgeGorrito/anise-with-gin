package anise

import (
	"github.com/JorgeGorrito/anise-with-gin/anise/config"
	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies"
	"github.com/JorgeGorrito/anise-with-gin/anise/errors"
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
) (*WebApplication, error) {
	if engine == nil {
		return nil, errors.ErrEngineIsNil
	}
	if configManager == nil {
		return nil, errors.ErrConfigManagerIsNil
	}
	if routesManager == nil {
		return nil, errors.ErrRoutesManagerIsNil
	}
	if dependenciesManager == nil {
		return nil, errors.ErrDependenciesManagerIsNil
	}
	return &WebApplication{
			engine:                engine,
			configManager:         configManager,
			routesManager:         routesManager,
			dependenciesManager:   dependenciesManager,
			dependenciesContainer: dependencies.NewContainer(),
		},
		nil
}

func (app *WebApplication) configureEngine() error {
	return app.configManager.ConfigureEngine(app.engine)
}

func (app *WebApplication) Run(addr ...string) error {
	if err := app.configureEngine(); err != nil {
		return err
	}

	if err := app.dependenciesManager.RegisterDependencies(app.dependenciesContainer); err != nil {
		return err
	}

	if err := app.routesManager.RegisterRoutes(app.engine, app.dependenciesContainer); err != nil {
		return err
	}

	if err := app.engine.Run(addr...); err != nil {
		return err
	}

	return nil
}
