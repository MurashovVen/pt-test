package service

import (
	"context"

	"github.com/golang/geo/s2"

	"pt-test/internal/entity"
)

func (s *Service) GeoObjectCreate(ctx context.Context, object *entity.GeoObject) (*entity.GeoObject, error) {
	res, err := s.repo.GeoObjectCreate(ctx, geoObjectFromEntityToDBModel(object))
	if err != nil {
		return nil, err
	}

	return geoObjectFromDBModelToEntity(res), nil
}

func (s *Service) GeoObjectDistanceGet(
	ctx context.Context, fromID string, toLatitude, toLongitude float64,
) (float64, error) {
	geoObject, err := s.repo.GeoObjectGet(ctx, fromID)
	if err != nil {
		return 0, err
	}

	var (
		p1 = s2.PointFromLatLng(
			s2.LatLngFromDegrees(geoObject.Geo.X, geoObject.Geo.Y),
		)
		p2 = s2.PointFromLatLng(
			s2.LatLngFromDegrees(toLatitude, toLongitude),
		)

		distance = p1.Distance(p2).Radians()
	)

	// 6371.0 - earthRadius in km
	return distance * 6371.0, nil
}
