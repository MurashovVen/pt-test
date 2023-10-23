package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"pt-test/internal/repo/models"
)

func (r *Repository) GeoObjectCreate(ctx context.Context, geoObj *models.GeoObject) (*models.GeoObject, error) {
	var (
		columnsNames = models.GeoObjectColumns

		columns = boil.Blacklist(
			columnsNames.ID,
			columnsNames.CreatedAt,
			columnsNames.UpdatedAt,
		)
	)

	return geoObj, geoObj.Insert(ctx, r.db, columns)
}

func (r *Repository) GeoObjectGet(ctx context.Context, id string) (*models.GeoObject, error) {
	return models.FindGeoObject(ctx, r.db, id)
}
