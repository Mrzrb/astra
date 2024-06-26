package petstore

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Mrzrb/astra/tests/petstore"
)

func getCatByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet, err := petstore.PetByID(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Cat{
		Pet:           *pet,
		Breed:         "Persian",
		IsIndependent: false,
	})
}

func getDogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet, err := petstore.PetByID(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Dog{
		Pet:       *pet,
		Breed:     "Labrador",
		IsTrained: true,
	})
}
