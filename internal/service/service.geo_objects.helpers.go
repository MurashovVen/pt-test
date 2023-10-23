package service

import (
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types/pgeo"

	"pt-test/internal/entity"
	"pt-test/internal/repo/models"
)

func geoObjectFromEntityToDBModel(object *entity.GeoObject) *models.GeoObject {
	return &models.GeoObject{
		ID:        object.ID,
		Geo:       pgeo.NewPoint(object.Latitude, object.Longitude),
		CreatedAt: null.NewTime(object.CreatedAt, !object.CreatedAt.IsZero()),
		UpdatedAt: null.NewTime(object.UpdatedAt, !object.UpdatedAt.IsZero()),
	}
}

func geoObjectFromDBModelToEntity(object *models.GeoObject) *entity.GeoObject {
	return &entity.GeoObject{
		ID:        object.ID,
		Latitude:  object.Geo.X,
		Longitude: object.Geo.Y,
		CreatedAt: object.CreatedAt.Time,
		UpdatedAt: object.UpdatedAt.Time,
	}
}
