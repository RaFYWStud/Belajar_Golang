package handler

import (
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type introController struct {
	service contract.IntroService
}

func (c *introController) getPrefix() string {
	return "/intro"
}

func (c *introController) initService(service *contract.Service) {
	c.service = service.Intro
}

func (c *introController) initRoute(app *gin.RouterGroup) {
	app.GET("/:id", c.GetIntro)
	app.POST("/create", c.CreateIntro)
	app.PUT("/:id", c.UpdateIntro)
	app.PATCH("/:id", c.UpdateIntro)
	app.DELETE("/:id", c.DeleteIntro)
}

func (c *introController) GetIntro(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	response, err := c.service.GetIntro(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *introController) CreateIntro(ctx *gin.Context) {
	var payload dto.IntroRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.CreateIntro(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *introController) UpdateIntro(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var payload dto.IntroRequest
	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.UpdateIntro(intID, &payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (c *introController) DeleteIntro(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	response, err := c.service.DeleteIntro(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
