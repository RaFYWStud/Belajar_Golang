package handler

import (
    "golang-tutorial/contract"
    "golang-tutorial/dto"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type todoController struct {
    service contract.ToDoService
}

func (c *todoController) getPrefix() string {
    return "/todo"
}

func (c *todoController) initService(service *contract.Service) {
    c.service = service.ToDo
}

func (c *todoController) initRoute(app *gin.RouterGroup) {
    app.POST("/create", c.CreateToDo)
    app.GET("/", c.GetToDos)
    app.GET("/:id", c.GetToDoByID)
    app.PUT("/:id", c.ReplaceToDo)
    app.PATCH("/:id", c.UpdateToDo)
    app.DELETE("/:id", c.DeleteToDo)
}

func (c *todoController) CreateToDo(ctx *gin.Context) {
    var payload dto.ToDoRequest
    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response, err := c.service.CreateToDo(&payload)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, response)
}

func (c *todoController) GetToDos(ctx *gin.Context) {
    response, err := c.service.GetToDos()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, response)
}

func (c *todoController) GetToDoByID(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    response, err := c.service.GetToDoByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, response)
}

func (c *todoController) ReplaceToDo(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var payload dto.ToDoRequest
    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response, err := c.service.ReplaceToDo(id, &payload)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, response)
}

func (c *todoController) UpdateToDo(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var payload dto.ToDoRequest
    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response, err := c.service.UpdateToDo(id, &payload)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, response)
}

func (c *todoController) DeleteToDo(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    err = c.service.DeleteToDo(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Berhasil menghapus data To-Do"})
}