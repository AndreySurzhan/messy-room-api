package app

import (
	api "github.com/AndreySurzhan/messy-room-api/gen"
	"github.com/AndreySurzhan/messy-room-api/internal/config"
	"github.com/AndreySurzhan/messy-room-api/internal/controller"
	"github.com/AndreySurzhan/messy-room-api/internal/pkg/logger"
	"github.com/AndreySurzhan/messy-room-api/internal/service"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	apiFile = "./api/api.yaml"
	apiPath = "/docs"
)

// App ...
type App struct {
	service    *service.Service
	controller *controller.Controller
	logger     *logger.Logger
	cfg        *config.Config
}

// New creates new app
func New(cfg *config.Config) (*App, error) {
	a := &App{
		cfg: cfg,
	}

	err := a.initDeps()

	return a, err
}

// initDeps initialize dependencies
func (a *App) initDeps() error {
	inits := []func() error{
		a.initService,
		a.initController,
		a.initLogger,
	}

	for _, fn := range inits {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initService() error {
	a.service = service.New()

	return nil
}

// initImpl initialize API impl
func (a *App) initController() error {
	a.controller = controller.New(
		a.service,
		a.cfg,
	)
	return nil
}

// initLogger initialize logger
func (a *App) initLogger() error {
	a.logger = logger.New(a.cfg)

	return nil
}

// Run runs the app
func (a *App) Run() error {
	r := gin.Default()
	swagger, err := api.GetSwagger()
	if err != nil {
		return err
	}

	a.logger.Use(r)

	registerCustomHandlers(r)
	registerSwagger(r)
	api.RegisterHandlers(r, a.controller)

	r.Use(middleware.OapiRequestValidator(swagger))

	err = r.Run(":" + a.cfg.GetString(config.Port))
	if err != nil {
		return err
	}

	return nil
}

// registerCustomHandlers registers custom handlers
func registerCustomHandlers(router *gin.Engine) *gin.Engine {
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	return router
}

func registerSwagger(router *gin.Engine) *gin.Engine {
	router.StaticFile(apiPath, apiFile)

	router.GET(apiPath+"/*any", ginSwagger.WrapHandler(
		swaggerFiles.NewHandler(), ginSwagger.URL(apiPath),
	))

	return router
}
