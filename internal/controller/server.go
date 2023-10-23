package controller

//nolint:revive
import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "pt-test/api"
	v1 "pt-test/internal/controller/v1"
	"pt-test/pkg/app/configuration"
	"pt-test/pkg/app/logger"
)

type (
	Server struct {
		env configuration.Environment

		engine *gin.Engine
		server *http.Server

		logger *logger.Logger
	}
)

// New creates a service
// Swagger spec:
// @title       PT-API Service
// @version     0.1.0
// @host        localhost:8080
// @BasePath  	/v1
func New(router *v1.Router, opts ...Option) *Server {
	var (
		engine = defaultEngine()

		srv = &Server{
			env:    configuration.DevEnv,
			engine: engine,
			server: &http.Server{
				Addr:              ":8080",
				Handler:           engine.Handler(),
				ReadTimeout:       5 * time.Second,
				ReadHeaderTimeout: 5 * time.Second,
				WriteTimeout:      5 * time.Second,
				IdleTimeout:       5 * time.Second,
			},
			logger: logger.NewNop(),
		}
	)

	for _, opt := range opts {
		opt(srv)
	}

	api := srv.engine.Group("/v1")
	{
		{
			api.POST("/geo-object", router.GeoObjectPost)
			api.GET("/geo-object/:id/distance", router.GeoObjectDistanceGet)
		}
	}

	if srv.env == configuration.DevEnv {
		srv.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	srv.engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return srv
}

func defaultEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())

	return engine
}
