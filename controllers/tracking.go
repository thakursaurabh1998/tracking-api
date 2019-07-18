package controllers

import (
	"database/sql"
	"net/http"
	"tracking-api/database"

	"github.com/labstack/echo"
)

// HomeController contains controllers for Home Route
type HomeController struct {
	Db *sql.DB
}

// Init initializes the home controller
func (hc HomeController) Init(g *echo.Group) {
	g.GET("", hc.GetHome)
}

// GetHome gets the home
func (hc HomeController) GetHome(c echo.Context) error {
	d := database.Data{155, 345.34, 2134.54}
	_, err := database.DBQuery(d, "confirm_place_order")

	type Response struct {
		Success bool `json:"success"`
	}

	// var toReturn Response
	var toReturn string
	var statusCode int
	if err == nil {
		statusCode = http.StatusOK
		toReturn = "Success"
	} else {
		statusCode = http.StatusInternalServerError
		toReturn = "Failed!"
	}
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Content-Type", "text/plain")
	return c.String(statusCode, toReturn)
}
