package app

import (
	"github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zsais/go-gin-prometheus"
	"gitlab.stripchat.dev/myclub/go-service/internal/app/api"
	"gitlab.stripchat.dev/myclub/go-service/internal/app/service"
	"gitlab.stripchat.dev/myclub/go-service/internal/config"
	"gitlab.stripchat.dev/myclub/go-service/internal/pkg/logger"
	"gitlab.stripchat.dev/myclub/go-service/internal/pkg/prompts"
)

const (
	apiFile = "./api/api.yaml"
	apiPath = "/docs"
)

// App ...
type App struct {
	impl    *service.Service
	logger  *logrus.Logger
	prompts *ginprometheus.Prometheus
	cfg     *viper.Viper
}

// New creates new app
func New(cfg *viper.Viper) (*App, error) {
	a := &App{}

	a.cfg = cfg

	err := a.initDeps()

	return a, err
}

// initDeps initialize dependencies
func (a *App) initDeps() error {
	inits := []func() error{
		a.initImpl,
		a.initLogger,
		a.initPrompts,
	}

	for _, fn := range inits {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

// initImpl initialize API impl
func (a *App) initImpl() error {
	resourceService := service.NewService()
	a.impl = resourceService

	return nil
}

// initLogger initialize logger
func (a *App) initLogger() error {
	a.logger = logger.New(a.cfg)

	return nil
}

func (a *App) initPrompts() error {
	a.prompts = prompts.New(a.cfg)

	return nil
}

// Run runs the app
func (a *App) Run() error {
	router := gin.Default()
	swagger, err := api.GetSwagger()
	if err != nil {
		return err
	}

	a.prompts.Use(router)

	router = a.registerCustomHandlers(router)
	router = a.registerSwagger(router)
	router = api.RegisterHandlers(router, a.impl)

	router.Use(middleware.OapiRequestValidator(swagger))
	router.Use(logger.Logger(a.logger), gin.Recovery())

	err = router.Run(":" + a.cfg.GetString(config.Port))
	if err != nil {
		return err
	}

	return nil
}

// registerCustomHandlers registers custom handlers
func (a *App) registerCustomHandlers(router *gin.Engine) *gin.Engine {
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	return router
}

func (a *App) registerSwagger(router *gin.Engine) *gin.Engine {
	router.StaticFile(apiPath, apiFile)

	router.GET(apiPath+"/*any", ginSwagger.WrapHandler(
		swaggerFiles.NewHandler(), ginSwagger.URL(apiPath),
	))

	return router
}
