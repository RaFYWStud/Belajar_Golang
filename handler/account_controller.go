package handler

import (
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	service contract.AccountService
}

func (c *AccountController) getPrefix() string {
	return "/account"
}

func (c *AccountController) initService(service *contract.Service) {
	c.service = service.Account
}

func (c *AccountController) initRoute(app *gin.RouterGroup) {
	app.GET("/:id", c.GetAccount)
	app.POST("/register", c.CreateAccount)
	app.POST("/login", c.Login)
}

func (c *AccountController) GetAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	response, err := c.service.GetAccount(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *AccountController) CreateAccount(ctx *gin.Context) {
	var payload dto.AccountRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.CreateAccount(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *AccountController) Login(ctx *gin.Context) {
	var payload dto.AccountRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.Login(&payload)
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
