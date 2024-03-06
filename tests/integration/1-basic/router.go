package petstore

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	a := A{}
	r := gin.Default()

	r.GET("/pets", getAllPets)
	r.GET("/pets/:id", getPetByID)
	r.POST("/pets", createPet)
	r.DELETE("/pets/:id", deletePet)
	r.GET("/demo", a.Test)

	return r
}

type A struct{}

func (a *A) Test(c *gin.Context) {
	c.JSON(http.StatusOK, []int64{55324})
}
