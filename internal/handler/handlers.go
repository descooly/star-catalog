package handler

import (
	"net/http"

	"github.com/descooly/star-catalog/internal/service"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
	repo service.Repository
}

func NewHandler(repo service.Repository) *Handlers {
	return &Handlers{repo: repo}
}

func (h *Handlers) GetStars(c echo.Context) error {
	var response []service.ResultStar

	sn := c.QueryParam("starName")
	cl := c.QueryParam("constellation")
	mind := c.QueryParam("minDistance")
	maxd := c.QueryParam("maxDistance")
	minm := c.QueryParam("minMass")
	maxm := c.QueryParam("maxMass")

	minDist, maxDist, minMass, maxMass := service.ParseQuery(mind, maxd, minm, maxm)

	results, err := h.repo.GetDB(sn, cl, minDist, maxDist, minMass, maxMass)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}

	// Преобразуем в ResultStar
	for _, r := range results {
		response = append(response, service.ResultStar{
			Name:   r.Name,
			Constl: r.Constellation,
			Dist:   r.Distance,
			Mass:   r.Mass,
		})
	}
	if len(response) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "there is no star with these parameters"})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handlers) PostStars(c echo.Context) error {
	newStar := new(service.ResultStar)
	if err := c.Bind(&newStar); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
	}

	h.repo.PostDB(newStar)

	return c.JSON(http.StatusCreated, newStar)
}
