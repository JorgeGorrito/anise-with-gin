package anise

import (
	"errors"

	dependencies "github.com/JorgeGorrito/anise-dependency-injection/andi/port/in"
	"github.com/JorgeGorrito/anise-with-gin/anise/commands"
	"github.com/JorgeGorrito/anise-with-gin/anise/config"
	ea "github.com/JorgeGorrito/anise-with-gin/anise/errors"
	"github.com/JorgeGorrito/anise-with-gin/anise/routing"
	"github.com/gin-gonic/gin"
)

type WebApplication struct {
	engine              *gin.Engine
	configManager       config.Manager
	routesManager       routing.Manager
	dependenciesManager dependencies.Manager
	commandsManager     commands.Manager
	commandsFactory     *commands.Factory
	commandsListener    commands.Listener
}

func NewWebApplication(
	engine *gin.Engine,
	configManager config.Manager,
	routesManager routing.Manager,
	dependenciesManager dependencies.Manager,
	commandsManager commands.Manager,
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

	var commandsFactory *commands.Factory = nil
	var commandsListener commands.Listener = nil
	if commandsManager != nil {
		commandsFactory = commands.NewFactory()
		commandsListener = commands.NewDefaultListener(commandsFactory)
	}
	return &WebApplication{
		engine:              engine,
		configManager:       configManager,
		routesManager:       routesManager,
		dependenciesManager: dependenciesManager,
		commandsManager:     commandsManager,
		commandsFactory:     commandsFactory,
		commandsListener:    commandsListener,
	}
}

func (app *WebApplication) ConfigureApplication() error {
	return app.configManager.ConfigureApplication()
}

func (app *WebApplication) configureEngine() error {
	return app.configManager.ConfigureEngine(app.engine)
}

func (app *WebApplication) registerDependencies() error {
	return app.dependenciesManager.RegisterDependencies()
}

func (app *WebApplication) registerRoutes() error {
	return app.routesManager.RegisterRoutes(app.engine)
}

func (app *WebApplication) RegisterCommands() error {
	return app.commandsManager.RegisterCommands(app.commandsFactory)
}

func (app *WebApplication) Run(addr ...string) {
	var errorList error

	if err := app.ConfigureApplication(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.registerDependencies(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.configureEngine(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if err := app.registerRoutes(); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if app.commandsManager != nil {
		if err := app.RegisterCommands(); err != nil {
			errorList = errors.Join(errorList, err)
		} else {
			go app.commandsListener.Listen(commands.DEFAULT_COMMAND_LISTENER_PORT)
		}
	}

	if err := app.engine.Run(addr...); err != nil {
		errorList = errors.Join(errorList, err)
	}

	if errorList != nil {
		panic(errorList)
	}
}
