package v1

import (
	"context"

	"pt-test/internal/entity"
	"pt-test/pkg/app/logger"
)

type (
	Router struct {
		geoObjectService geoObjectService

		logger *logger.Logger
	}

	geoObjectService interface {
		GeoObjectCreate(ctx context.Context, object *entity.GeoObject) (*entity.GeoObject, error)
		GeoObjectDistanceGet(ctx context.Context, fromID string, toLatitude float64, toLongitude float64) (float64, error)
	}
)

func NewRouter(geoObjectService geoObjectService, log *logger.Logger) *Router {
	return &Router{
		geoObjectService: geoObjectService,
		logger:           log.Named("Router"),
	}
}
