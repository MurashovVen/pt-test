package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"pt-test/internal/entity"
)

type (
	GeoObjectPostRequest struct {
		Latitude  float64 `json:"latitude" binding:"required"`
		Longitude float64 `json:"longitude" binding:"required"`
	}

	GeoObject struct {
		ID string `json:"id"`
	}
)

// GeoObjectPost
//
// @Summary     Create geo object
// @Description Method for geo object creation
// @ID          GeoObjectPost
// @Tags  	    geo_object_api
// @Accept      json
// @Produce     json
// @Param       request body GeoObjectPostRequest true "GeoObject params"
// @Success     200 {object} GeoObject
// @Failure     400 {object} Error
// @Failure     500 {object} Error
// @Router      /geo-object [post]
func (c *Router) GeoObjectPost(ctx *gin.Context) {
	var req GeoObjectPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.Error(ctx, err, "invalid object", http.StatusBadRequest)

		return
	}

	geoObjectCreated, err := c.geoObjectService.GeoObjectCreate(
		ctx,
		&entity.GeoObject{
			Latitude:  req.Latitude,
			Longitude: req.Longitude,
		},
	)
	if err != nil {
		c.Error(ctx, err, "failed to create object", http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, &GeoObject{ID: geoObjectCreated.ID})
}

type (
	Distance struct {
		Kilometers float64 `json:"kilometers"`
	}
)

// GeoObjectDistanceGet
//
// @Summary     Get distance
// @Description Get distance between stored GeoObject and provided coordinates
// @ID          GeoObjectDistanceGet
// @Tags  	    geo_object_api
// @Param       id path string true "GeoObject ID"
// @Param       latitude query string true "latitude"
// @Param       longitude query string true "longitude"
// @Produce     json
// @Success     200 {object} Distance
// @Failure     400 {object} Error
// @Failure     500 {object} Error
// @Router      /geo-object/{id}/distance [get]
func (c *Router) GeoObjectDistanceGet(ctx *gin.Context) {
	lat, err := strconv.ParseFloat(ctx.Query("latitude"), 64)
	if err != nil {
		c.Error(ctx, err, "invalid latitude", http.StatusBadRequest)

		return
	}

	lng, err := strconv.ParseFloat(ctx.Query("longitude"), 64)
	if err != nil {
		c.Error(ctx, err, "invalid longitude", http.StatusBadRequest)

		return
	}

	distance, err := c.geoObjectService.GeoObjectDistanceGet(
		ctx,
		ctx.Param("id"),
		lat,
		lng,
	)
	if err != nil {
		c.Error(ctx, err, "computing distance", http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, &Distance{Kilometers: distance})
}
