package anise

import (
	"errors"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
	"github.com/JorgeGorrito/anise-with-gin/anise/config"
	ea "github.com/JorgeGorrito/anise-with-gin/anise/errors"
	"github.com/JorgeGorrito/anise-with-gin/anise/routing"
	"github.com/gin-gonic/gin"
)

type WebApplication struct {
	engine           *gin.Engine
	configManager    config.Manager
	routesManager    routing.Manager
	commandsManager  command.Manager
	commandsListener command.Listener
}

func NewWebApplication() *WebApplication {
	return &WebApplication{}
}

func (app *WebApplication) SetEngine(engine *gin.Engine) *WebApplication {
	app.engine = engine
	return app
}

func (app *WebApplication) SetConfigManager(configManager config.Manager) *WebApplication {
	app.configManager = configManager
	return app
}

func (app *WebApplication) SetRoutesManager(routesManager routing.Manager) *WebApplication {
	app.routesManager = routesManager
	return app
}

func (app *WebApplication) SetCommandsManager(commandsManager command.Manager) *WebApplication {
	app.commandsManager = commandsManager
	return app
}

func (app *WebApplication) ConfigureApplication() error {
	return app.configManager.ConfigureApplication()
}

func (app *WebApplication) configureEngine() error {
	return app.configManager.ConfigureEngine(app.engine)
}

func (app *WebApplication) registerRoutes() error {
	return app.routesManager.RegisterRoutes(app.engine)
}

func (app *WebApplication) RegisterCommands() error {
	return app.commandsManager.RegisterCommands(app.commandsManager.GetRegistryCommand())
}

func (app *WebApplication) Run(addr ...string) {
	var errorList error

	if app.engine == nil {
		errorList = errors.Join(errorList, ea.ErrEngineIsNil)
	}
	if app.configManager == nil {
		errorList = errors.Join(errorList, ea.ErrConfigManagerIsNil)
	}
	if app.routesManager == nil {
		errorList = errors.Join(errorList, ea.ErrRoutesManagerIsNil)
	}

	if errorList != nil {
		panic(errorList)
	}

	if app.commandsManager != nil {
		app.commandsListener = command.NewDefaultListener(app.commandsManager.GetRetrieverCommand())
	}

	if err := app.ConfigureApplication(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if app.commandsManager != nil {
		if err := app.RegisterCommands(); err != nil {
			errorList = errors.Join(errorList, err)
		} else {
			go app.commandsListener.Listen(command.DEFAULT_COMMAND_LISTENER_PORT)
		}
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

	if errorList != nil {
		panic(errorList)
	}
}

func (app *WebApplication) RunTLS(addr, certFile, keyFile string) {
	var errorList error

	if app.engine == nil {
		errorList = errors.Join(errorList, ea.ErrEngineIsNil)
	}
	if app.configManager == nil {
		errorList = errors.Join(errorList, ea.ErrConfigManagerIsNil)
	}
	if app.routesManager == nil {
		errorList = errors.Join(errorList, ea.ErrRoutesManagerIsNil)
	}

	if errorList != nil {
		panic(errorList)
	}

	if app.commandsManager != nil {
		app.commandsListener = command.NewDefaultListener(app.commandsManager.GetRetrieverCommand())
	}

	if err := app.ConfigureApplication(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if app.commandsManager != nil {
		if err := app.RegisterCommands(); err != nil {
			errorList = errors.Join(errorList, err)
		} else {
			go app.commandsListener.Listen(command.DEFAULT_COMMAND_LISTENER_PORT)
		}
	}

	if err := app.configureEngine(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.registerRoutes(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.engine.RunTLS(addr, certFile, keyFile); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if errorList != nil {
		panic(errorList)
	}
}
