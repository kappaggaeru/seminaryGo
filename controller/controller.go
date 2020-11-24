package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kappaggaeru/seminaryGo/model"
	"github.com/kappaggaeru/seminaryGo/service"
)

// HTTPController ...
type HTTPController struct {
	router  *gin.Engine
	service *service.Movie
}

// NewHTTPController ...
func NewHTTPController(s *service.Movie) HTTPController {
	r := gin.Default()
	makeEndpoints(r, s) // endpoints register
	return HTTPController{r, s}
}

// Run ...
func (c *HTTPController) Run() {
	c.router.Run()
}

func makeEndpoints(r *gin.Engine, s *service.Movie) {
	r.GET("/movies/:id", makeFindHandler(s))
	r.POST("/movies/add", makeAddHandler(s))
	r.DELETE("/movies/delete/:id", makeDeleteHandler(s))
	r.PUT("/movies/update/:id", makeUpdateHandler(s))
}

// Finder
func makeFindHandler(s *service.Movie) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, _ := strconv.Atoi(c.Param("id"))

		m := (*s).FindByID(uint(i))

		c.JSON(http.StatusOK, gin.H{
			"movie": m,
		})
	}
}
// Adder
func makeAddHandler(s *service.Movie) gin.HandlerFunc {
	return func(c *gin.Context) {
		var optionalMovie *model.Movie

		c.BindJSON(&optionalMovie)

		(*s).Add(optionalMovie)

		c.JSON(http.StatusOK, gin.H{
			"movie": &optionalMovie,
		})
	}
}
// Deleter
func makeDeleteHandler(s *service.Movie) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, _ := strconv.Atoi(c.Param("id"))

		c.BindJSON(i)

		(*s).Remove(uint(i))

		c.JSON(http.StatusNoContent, gin.H{})
	}
}
// Updater
func makeUpdateHandler(s *service.Movie) gin.HandlerFunc {
	return func(c *gin.Context) {
		var optionalMovie *model.Movie

		c.BindJSON(&optionalMovie)

		(*s).Update(optionalMovie)

		c.JSON(http.StatusOK, gin.H{
			"movie": &optionalMovie,
		})
	}
}